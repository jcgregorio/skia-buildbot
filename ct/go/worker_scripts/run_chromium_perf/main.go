// run_chromium_perf is an application that runs the specified benchmark over CT's
// webpage archives.
package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/skia-dev/glog"

	"strings"

	"go.skia.org/infra/ct/go/adb"
	"go.skia.org/infra/ct/go/util"
	"go.skia.org/infra/ct/go/worker_scripts/worker_common"
	"go.skia.org/infra/go/common"
	skutil "go.skia.org/infra/go/util"
)

const (
	// The number of goroutines that will run in parallel to run benchmarks.
	WORKER_POOL_SIZE = 10
)

var (
	workerNum                 = flag.Int("worker_num", 1, "The number of this CT worker. It will be in the {1..100} range.")
	pagesetType               = flag.String("pageset_type", util.PAGESET_TYPE_MOBILE_10k, "The type of pagesets to create from the Alexa CSV list. Eg: 10k, Mobile10k, All.")
	chromiumBuildNoPatch      = flag.String("chromium_build_nopatch", "", "The chromium build to use for the nopatch run.")
	chromiumBuildWithPatch    = flag.String("chromium_build_withpatch", "", "The chromium build to use for the withpatch run.")
	runID                     = flag.String("run_id", "", "The unique run id (typically requester + timestamp).")
	runIDNoPatch              = flag.String("run_id_nopatch", "", "The unique run id (typically requester + timestamp) for the nopatch run.")
	runIDWithPatch            = flag.String("run_id_withpatch", "", "The unique run id (typically requester + timestamp) for the withpatch run.")
	benchmarkName             = flag.String("benchmark_name", "", "The telemetry benchmark to run on this worker.")
	benchmarkExtraArgs        = flag.String("benchmark_extra_args", "", "The extra arguments that are passed to the specified benchmark.")
	browserExtraArgsNoPatch   = flag.String("browser_extra_args_nopatch", "", "The extra arguments that are passed to the browser while running the benchmark during the nopatch run.")
	browserExtraArgsWithPatch = flag.String("browser_extra_args_withpatch", "", "The extra arguments that are passed to the browser while running the benchmark during the withpatch run.")
	repeatBenchmark           = flag.Int("repeat_benchmark", 3, "The number of times the benchmark should be repeated. For skpicture_printer benchmark this value is always 1.")
	runInParallel             = flag.Bool("run_in_parallel", false, "Run the benchmark by bringing up multiple chrome instances in parallel.")
	targetPlatform            = flag.String("target_platform", util.PLATFORM_ANDROID, "The platform the benchmark will run on (Android / Linux).")
	chromeCleanerTimer        = flag.Duration("cleaner_timer", 15*time.Minute, "How often all chrome processes will be killed on this slave.")
)

func main() {
	defer common.LogPanic()
	worker_common.Init()
	if !*worker_common.Local {
		defer util.CleanTmpDir()
	}
	defer util.TimeTrack(time.Now(), "Running Chromium Perf")
	defer glog.Flush()

	// Validate required arguments.
	if *chromiumBuildNoPatch == "" {
		glog.Error("Must specify --chromium_build_nopatch")
		return
	}
	if *chromiumBuildWithPatch == "" {
		glog.Error("Must specify --chromium_build_withpatch")
		return
	}
	if *runID == "" {
		glog.Error("Must specify --run_id")
		return
	}
	if *runIDNoPatch == "" {
		glog.Error("Must specify --run_id_nopatch")
		return
	}
	if *runIDWithPatch == "" {
		glog.Error("Must specify --run_id_withpatch")
		return
	}
	if *benchmarkName == "" {
		glog.Error("Must specify --benchmark_name")
		return
	}

	// Reset the local chromium checkout.
	if err := util.ResetCheckout(util.ChromiumSrcDir); err != nil {
		glog.Errorf("Could not reset %s: %s", util.ChromiumSrcDir, err)
		return
	}
	// Sync the local chromium checkout.
	if err := util.SyncDir(util.ChromiumSrcDir); err != nil {
		glog.Errorf("Could not gclient sync %s: %s", util.ChromiumSrcDir, err)
		return
	}

	// Create the task file so that the master knows this worker is still busy.
	skutil.LogErr(util.CreateTaskFile(util.ACTIVITY_RUNNING_CHROMIUM_PERF))
	defer util.DeleteTaskFile(util.ACTIVITY_RUNNING_CHROMIUM_PERF)

	if *targetPlatform == util.PLATFORM_ANDROID {
		if err := adb.VerifyLocalDevice(); err != nil {
			// Android device missing or offline.
			glog.Errorf("Could not find Android device: %s", err)
			return
		}
		// Kill adb server to make sure we start from a clean slate.
		skutil.LogErr(util.ExecuteCmd(util.BINARY_ADB, []string{"kill-server"}, []string{},
			util.ADB_ROOT_TIMEOUT, nil, nil))
		// Make sure adb shell is running as root.
		skutil.LogErr(util.ExecuteCmd(util.BINARY_ADB, []string{"root"}, []string{},
			util.ADB_ROOT_TIMEOUT, nil, nil))
	}

	// Instantiate GsUtil object.
	gs, err := util.NewGsUtil(nil)
	if err != nil {
		glog.Error(err)
		return
	}

	// Download the benchmark patch for this run from Google storage.
	benchmarkPatchName := *runID + ".benchmark.patch"
	benchmarkPatchLocalPath := filepath.Join(os.TempDir(), benchmarkPatchName)
	remoteDir := filepath.Join(util.ChromiumPerfRunsDir, *runID)
	benchmarkPatchRemotePath := filepath.Join(remoteDir, benchmarkPatchName)
	respBody, err := gs.GetRemoteFileContents(benchmarkPatchRemotePath)
	if err != nil {
		glog.Errorf("Could not fetch %s: %s", benchmarkPatchRemotePath, err)
		return
	}
	defer skutil.Close(respBody)
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(respBody); err != nil {
		glog.Errorf("Could not read from %s: %s", benchmarkPatchRemotePath, err)
		return
	}
	if err := ioutil.WriteFile(benchmarkPatchLocalPath, buf.Bytes(), 0666); err != nil {
		glog.Errorf("Unable to create file %s: %s", benchmarkPatchLocalPath, err)
		return
	}
	defer skutil.Remove(benchmarkPatchLocalPath)
	// Apply benchmark patch to the local chromium checkout.
	if buf.Len() > 10 {
		if err := util.ApplyPatch(benchmarkPatchLocalPath, util.ChromiumSrcDir); err != nil {
			glog.Errorf("Could not apply Telemetry's patch in %s: %s", util.ChromiumSrcDir, err)
			return
		}
	}

	// Download the specified chromium builds.
	for _, chromiumBuild := range []string{*chromiumBuildNoPatch, *chromiumBuildWithPatch} {
		if err := gs.DownloadChromiumBuild(chromiumBuild); err != nil {
			glog.Error(err)
			return
		}
		//Delete the chromium build to save space when we are done.
		defer skutil.RemoveAll(filepath.Join(util.ChromiumBuildsDir, chromiumBuild))
	}

	chromiumBinaryNoPatch := filepath.Join(util.ChromiumBuildsDir, *chromiumBuildNoPatch, util.BINARY_CHROME)
	chromiumBinaryWithPatch := filepath.Join(util.ChromiumBuildsDir, *chromiumBuildWithPatch, util.BINARY_CHROME)

	// Download pagesets if they do not exist locally.
	if err := gs.DownloadWorkerArtifacts(util.PAGESETS_DIR_NAME, *pagesetType, *workerNum); err != nil {
		glog.Error(err)
		return
	}
	pathToPagesets := filepath.Join(util.PagesetsDir, *pagesetType)

	// Download archives if they do not exist locally.
	if err := gs.DownloadWorkerArtifacts(util.WEB_ARCHIVES_DIR_NAME, *pagesetType, *workerNum); err != nil {
		glog.Error(err)
		return
	}

	// Establish nopatch output paths.
	localOutputDirNoPatch := filepath.Join(util.StorageDir, util.BenchmarkRunsDir, *runIDNoPatch)
	skutil.RemoveAll(localOutputDirNoPatch)
	skutil.MkdirAll(localOutputDirNoPatch, 0700)
	defer skutil.RemoveAll(localOutputDirNoPatch)
	remoteDirNoPatch := filepath.Join(util.BenchmarkRunsDir, *runIDNoPatch)

	// Establish withpatch output paths.
	localOutputDirWithPatch := filepath.Join(util.StorageDir, util.BenchmarkRunsDir, *runIDWithPatch)
	skutil.RemoveAll(localOutputDirWithPatch)
	skutil.MkdirAll(localOutputDirWithPatch, 0700)
	defer skutil.RemoveAll(localOutputDirWithPatch)
	remoteDirWithPatch := filepath.Join(util.BenchmarkRunsDir, *runIDWithPatch)

	// Construct path to the ct_run_benchmark python script.
	_, currentFile, _, _ := runtime.Caller(0)
	pathToPyFiles := filepath.Join(
		filepath.Dir((filepath.Dir(filepath.Dir(filepath.Dir(currentFile))))),
		"py")

	fileInfos, err := ioutil.ReadDir(pathToPagesets)
	if err != nil {
		glog.Errorf("Unable to read the pagesets dir %s: %s", pathToPagesets, err)
		return
	}

	numWorkers := WORKER_POOL_SIZE
	if *targetPlatform == util.PLATFORM_ANDROID || !*runInParallel {
		// Do not run page sets in parallel if the target platform is Android.
		// This is because the nopatch/withpatch APK needs to be installed prior to
		// each run and this will interfere with the parallel runs. Instead of trying
		// to find a complicated solution to this, it makes sense for Android to
		// continue to be serial because it will help guard against
		// crashes/flakiness/inconsistencies which are more prevalent in mobile runs.
		numWorkers = 1
		glog.Infoln("===== Going to run the task serially =====")
	} else {
		glog.Infoln("===== Going to run the task with parallel chrome processes =====")
	}

	// Create channel that contains all pageset file names. This channel will
	// be consumed by the worker pool.
	pagesetRequests := util.GetClosedChannelOfPagesets(fileInfos)

	var wg sync.WaitGroup
	// Use a RWMutex for the chromeProcessesCleaner goroutine to communicate to
	// the workers (acting as "readers") when it wants to be the "writer" and
	// kill all zombie chrome processes.
	var mutex sync.RWMutex

	// Loop through workers in the worker pool.
	for i := 0; i < numWorkers; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Create and run a goroutine closure that captures SKPs.
		go func() {
			// Decrement the WaitGroup counter when the goroutine completes.
			defer wg.Done()

			for pagesetName := range pagesetRequests {

				mutex.RLock()

				if err := runBenchmark(pagesetName, pathToPagesets, pathToPyFiles, localOutputDirNoPatch, *chromiumBuildNoPatch, chromiumBinaryNoPatch, *runIDNoPatch, *browserExtraArgsNoPatch); err != nil {
					glog.Errorf("Error while running nopatch benchmark: %s", err)
					return
				}
				if err := runBenchmark(pagesetName, pathToPagesets, pathToPyFiles, localOutputDirWithPatch, *chromiumBuildWithPatch, chromiumBinaryWithPatch, *runIDWithPatch, *browserExtraArgsWithPatch); err != nil {
					glog.Errorf("Error while running withpatch benchmark: %s", err)
					return
				}
				mutex.RUnlock()
			}
		}()
	}

	if !*worker_common.Local {
		// Start the cleaner.
		go util.ChromeProcessesCleaner(&mutex, *chromeCleanerTimer)
	}

	// Wait for all spawned goroutines to complete.
	wg.Wait()

	// If "--output-format=csv-pivot-table" was specified then merge all CSV files and upload.
	if strings.Contains(*benchmarkExtraArgs, "--output-format=csv-pivot-table") {
		if err := mergeUploadCSVFiles(localOutputDirNoPatch, pathToPyFiles, *runIDNoPatch, remoteDirNoPatch, gs); err != nil {
			glog.Errorf("Error while processing nopatch CSV files: %s", err)
			return
		}
		if err := mergeUploadCSVFiles(localOutputDirWithPatch, pathToPyFiles, *runIDWithPatch, remoteDirWithPatch, gs); err != nil {
			glog.Errorf("Error while processing withpatch CSV files: %s", err)
			return
		}
	}
}

func runBenchmark(fileInfoName, pathToPagesets, pathToPyFiles, localOutputDir, chromiumBuildName, chromiumBinary, runID, browserExtraArgs string) error {
	pagesetBaseName := filepath.Base(fileInfoName)
	if pagesetBaseName == util.TIMESTAMP_FILE_NAME || filepath.Ext(pagesetBaseName) == ".pyc" {
		// Ignore timestamp files and .pyc files.
		return nil
	}

	// Read the pageset.
	pagesetName := strings.TrimSuffix(pagesetBaseName, filepath.Ext(pagesetBaseName))
	pagesetPath := filepath.Join(pathToPagesets, fileInfoName)
	decodedPageset, err := util.ReadPageset(pagesetPath)
	if err != nil {
		return fmt.Errorf("Could not read %s: %s", pagesetPath, err)
	}

	glog.Infof("===== Processing %s for %s =====", pagesetPath, runID)
	pagesetName, present := util.BenchmarksToPagesetName[*benchmarkName]
	if !present {
		// If it is custom benchmark use the entered benchmark name.
		pagesetName = *benchmarkName
	}
	args := []string{
		filepath.Join(util.TelemetryBinariesDir, util.BINARY_RUN_BENCHMARK),
		pagesetName,
		"--also-run-disabled-tests",
		"--user-agent=" + decodedPageset.UserAgent,
		"--urls-list=" + decodedPageset.UrlsList,
		"--archive-data-file=" + decodedPageset.ArchiveDataFile,
	}

	// Need to capture output for all benchmarks.
	outputDirArgValue := filepath.Join(localOutputDir, pagesetName)
	args = append(args, "--output-dir="+outputDirArgValue)
	// Figure out which browser should be used.
	if *targetPlatform == util.PLATFORM_ANDROID {
		if err := util.InstallChromeAPK(chromiumBuildName); err != nil {
			return fmt.Errorf("Error while installing APK: %s", err)
		}
		args = append(args, "--browser=android-chromium")
	} else {
		args = append(args, "--browser=exact", "--browser-executable="+chromiumBinary)
	}
	// Split benchmark args if not empty and append to args.
	if *benchmarkExtraArgs != "" {
		for _, benchmarkArg := range strings.Split(*benchmarkExtraArgs, " ") {
			args = append(args, benchmarkArg)
		}
	}
	// Add the number of times to repeat.
	args = append(args, fmt.Sprintf("--page-repeat=%d", *repeatBenchmark))
	// Add browserArgs if not empty to args.
	if browserExtraArgs != "" {
		args = append(args, "--extra-browser-args="+browserExtraArgs)
	}
	// Set the PYTHONPATH to the pagesets and the telemetry dirs.
	env := []string{
		fmt.Sprintf("PYTHONPATH=%s:%s:%s:$PYTHONPATH", pathToPagesets, util.TelemetryBinariesDir, util.TelemetrySrcDir),
		"DISPLAY=:0",
	}
	timeoutSecs := util.PagesetTypeToInfo[*pagesetType].RunChromiumPerfTimeoutSecs
	if err := util.ExecuteCmd("python", args, env, time.Duration(timeoutSecs)*time.Second, nil, nil); err != nil {
		glog.Errorf("Run benchmark command failed with: %s", err)
	}
	return nil
}

func mergeUploadCSVFiles(localOutputDir, pathToPyFiles, runID, remoteDir string, gs *util.GsUtil) error {
	// Move all results into a single directory.
	fileInfos, err := ioutil.ReadDir(localOutputDir)
	if err != nil {
		return fmt.Errorf("Unable to read %s: %s", localOutputDir, err)
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			continue
		}
		outputFile := filepath.Join(localOutputDir, fileInfo.Name(), "results-pivot-table.csv")
		newFile := filepath.Join(localOutputDir, fmt.Sprintf("%s.csv", fileInfo.Name()))
		if err := os.Rename(outputFile, newFile); err != nil {
			glog.Errorf("Could not rename %s to %s: %s", outputFile, newFile, err)
			continue
		}
		// Add the rank of the page to the CSV file.
		headers, values, err := getRowsFromCSV(newFile)
		if err != nil {
			glog.Errorf("Could not read %s: %s", newFile, err)
			continue
		}
		pageRank := strings.Split(fileInfo.Name(), "_")[1]
		for i := range headers {
			for j := range values {
				if headers[i] == "page" {
					values[j][i] = fmt.Sprintf("%s (#%s)", values[j][i], pageRank)
				}
			}
		}
		if err := writeRowsToCSV(newFile, headers, values); err != nil {
			glog.Errorf("Could not write to %s: %s", newFile, err)
			continue
		}
	}
	// Call csv_pivot_table_merger.py to merge all results into a single results CSV.
	pathToCsvMerger := filepath.Join(pathToPyFiles, "csv_pivot_table_merger.py")
	outputFileName := runID + ".output"
	args := []string{
		pathToCsvMerger,
		"--csv_dir=" + localOutputDir,
		"--output_csv_name=" + filepath.Join(localOutputDir, outputFileName),
	}
	err = util.ExecuteCmd("python", args, []string{}, util.CSV_PIVOT_TABLE_MERGER_TIMEOUT, nil,
		nil)
	if err != nil {
		return fmt.Errorf("Error running csv_pivot_table_merger.py: %s", err)
	}
	// Copy the output file to Google Storage.
	remoteOutputDir := filepath.Join(remoteDir, fmt.Sprintf("slave%d", *workerNum), "outputs")
	if err := gs.UploadFile(outputFileName, localOutputDir, remoteOutputDir); err != nil {
		return fmt.Errorf("Unable to upload %s to %s: %s", outputFileName, remoteOutputDir, err)
	}
	return nil
}

func getRowsFromCSV(csvPath string) ([]string, [][]string, error) {
	csvFile, err := os.Open(csvPath)
	defer skutil.Close(csvFile)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not open %s: %s", csvPath, err)
	}
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("Could not read %s: %s", csvPath, err)
	}
	if len(rawCSVdata) < 2 {
		return nil, nil, fmt.Errorf("No data in %s", csvPath)
	}
	return rawCSVdata[0], rawCSVdata[1:], nil
}

func writeRowsToCSV(csvPath string, headers []string, values [][]string) error {
	csvFile, err := os.OpenFile(csvPath, os.O_WRONLY, 666)
	defer skutil.Close(csvFile)
	if err != nil {
		return fmt.Errorf("Could not open %s: %s", csvPath, err)
	}
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	// Write the headers.
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("Could not write to %s: %s", csvPath, err)
	}
	// Write all values.
	for _, row := range values {
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("Could not write to %s: %s", csvPath, err)
		}
	}
	return nil
}
