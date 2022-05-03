package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.skia.org/infra/go/executil"
	"go.skia.org/infra/go/gcs"
	"go.skia.org/infra/go/gcs/mocks"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/now"
	"go.skia.org/infra/go/testutils"
	"go.skia.org/infra/go/testutils/unittest"
)

func setupForTest(t *testing.T, h http.HandlerFunc) (string, *http.Client) {
	client := httputils.DefaultClientConfig().WithoutRetries().With2xxOnly().Client()

	return t.TempDir(), client
}

func TestComputeUploadPathFromTime(t *testing.T) {
	unittest.SmallTest(t)
	var mockTime = time.Unix(1649770315, 12).UTC()
	ctx := context.WithValue(context.Background(), now.ContextKey, mockTime)
	assert.Equal(t, "2022/04/12/13", computeUploadPathFromTime(ctx))
}

var benchmarkScriptArgs = []string{"--output", "results.json"}

func TestRunBenchMarkScript_ScriptReturnsError_ReturnsError(t *testing.T) {
	unittest.MediumTest(t)

	workDir, _ := setupForTest(t, nil)
	ctx := executil.WithFakeTests(context.Background(), "Test_FakeExe_Exec_Fails")

	err := runBenchMarkScript(ctx, benchmarkScriptArgs, workDir)
	require.Error(t, err)
}

func TestRunBenchMarkScript_ScriptSucceeds_DoesNotReturnError(t *testing.T) {
	unittest.MediumTest(t)

	workDir, _ := setupForTest(t, nil)
	ctx := executil.WithFakeTests(context.Background(), "Test_FakeExe_Python_Script_Success")

	err := runBenchMarkScript(ctx, benchmarkScriptArgs, workDir)
	require.NoError(t, err)
}

func Test_FakeExe_Exec_Fails(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}
	os.Exit(1)
}

func Test_FakeExe_Python_Script_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}

	args := executil.OriginalArgs()
	require.Contains(t, args, python)
	require.Contains(t, args, "--output")
	os.Exit(0)
}

const (
	repoURL = "https://example.com/git"
)

var (
	// We need to use real directories for workDir and dest because the exec
	// package checks that Command.Path exists before running the command.
	workDir = os.TempDir()

	dest = os.TempDir()

	directories = []string{"a/b", "c"}

	benchmarkName = "canary"

	benchmarkConfig = benchmark{
		repoURL:       repoURL,
		checkoutPaths: directories,
		scriptName:    "a/b/benchmark.py",
		flags:         []string{"--githash", "abcdef"},
	}
)

func TestNewSparseCheckout_Success(t *testing.T) {
	unittest.MediumTest(t)
	ctx := executil.WithFakeTests(context.Background(),
		"Test_FakeExe_Clone_Success",
		"Test_FakeExe_Sparse_Init_Success",
		"Test_FakeExe_Sparse_Set_Success",
	)

	err := newSparseCheckout(ctx, workDir, repoURL, dest, directories)
	require.NoError(t, err)
}

func TestNewSparseCheckout_GitCommandFails_ReturnsError(t *testing.T) {
	unittest.MediumTest(t)
	ctx := executil.WithFakeTests(context.Background(),
		"Test_FakeExe_Exec_Fails",
	)

	err := newSparseCheckout(ctx, workDir, repoURL, dest, directories)
	require.Error(t, err)
}

func Test_FakeExe_Clone_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}

	args := executil.OriginalArgs()
	expected := []string{"git", "clone", "--depth", "1", "--filter=blob:none", "--sparse", "https://example.com/git"}
	// The checkout dest will always change, so just check all the other arguments.
	require.Equal(t, expected, args[:len(expected)])
	os.Exit(0)
}

func Test_FakeExe_Sparse_Init_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}

	args := executil.OriginalArgs()
	expected := []string{"git", "sparse-checkout", "init", "--cone"}
	require.Equal(t, expected, args)
	os.Exit(0)
}

func Test_FakeExe_Sparse_Set_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}

	args := executil.OriginalArgs()
	expected := []string{"git", "sparse-checkout", "set", "a/b", "c"}
	require.Equal(t, expected, args)
	os.Exit(0)
}

func TestRunSingleBenchmark_HappyPath(t *testing.T) {
	unittest.MediumTest(t)
	ctx := executil.WithFakeTests(context.Background(),
		"Test_FakeExe_Exec_Success",
		"Test_FakeExe_Exec_Success",
		"Test_FakeExe_Exec_Success",
		"Test_FakeExe_Run_Canary_Python_Script_Success",
	)
	workDir := t.TempDir()
	err := os.MkdirAll(filepath.Join(workDir, "git", "canary"), 0755)
	require.NoError(t, err)

	outputFileName, err := runSingleBenchmark(ctx, benchmarkName, benchmarkConfig, "abcdef", workDir)
	require.NoError(t, err)
	require.Contains(t, outputFileName, "canary/results.json")

}

func Test_FakeExe_Exec_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}
	os.Exit(0)
}

func Test_FakeExe_Run_Canary_Python_Script_Success(t *testing.T) {
	unittest.FakeExeTest(t)
	if os.Getenv(executil.OverrideEnvironmentVariable) == "" {
		return
	}

	args := executil.OriginalArgs()
	expected := []string{"python3", "a/b/benchmark.py", "--githash", "abcdef", "--githash", "abcdef", "--output", "canary/results.json"}
	for i, endsWith := range expected {
		require.Contains(t, args[i], endsWith)
	}
	os.Exit(0)
}

// myWriteCloser is a wrapper that turns a bytes.Buffer from an io.Writer to an io.WriteCloser.
type myWriteCloser struct {
	bytes.Buffer
}

func (*myWriteCloser) Close() error {
	return nil
}

var (
	myError = fmt.Errorf("my fake error")

	fileContents = []byte("{}")
)

func setupForUploadTest(t *testing.T) (context.Context, string) {
	// Create a context with a mockTime so that the GCS upload directory is always the same.
	var mockTime = time.Unix(1649770315, 12).UTC()
	ctx := context.WithValue(context.Background(), now.ContextKey, mockTime)

	// Create a results.json file that will be uploaded by uploadResultsFile.
	tempDir := t.TempDir()
	resultsFile := filepath.Join(tempDir, "results.json")
	err := os.WriteFile(resultsFile, fileContents, 0644)
	require.NoError(t, err)

	return ctx, resultsFile
}

func TestUploadResultsFile_HappyPath(t *testing.T) {
	unittest.MediumTest(t)
	ctx, resultsFile := setupForUploadTest(t)

	var b myWriteCloser

	// Mock out the gcs client.
	gcsClient := mocks.NewGCSClient(t)
	gcsClient.On("FileWriter", testutils.AnyContext, fmt.Sprintf("ingest/2022/04/12/13/%s/results.json", benchmarkName), gcs.FileWriteOptions{
		ContentEncoding: "application/json",
	}).Return(&b)

	err := uploadResultsFile(ctx, gcsClient, benchmarkName, resultsFile)
	require.NoError(t, err)
	require.Equal(t, fileContents, b.Bytes())
}

// myWriteCloser is an io.WriteCloser that always fails on Write.
type myFailingWriteCloser struct {
}

func (*myFailingWriteCloser) Write([]byte) (int, error) {
	return 0, myError
}

func (*myFailingWriteCloser) Close() error {
	return nil
}

func TestUploadResultsFile_WriteCloserFailsToWrite_ReturnsError(t *testing.T) {
	unittest.MediumTest(t)
	ctx, resultsFile := setupForUploadTest(t)

	var b myFailingWriteCloser

	// Mock out the gcs client.
	gcsClient := mocks.NewGCSClient(t)
	gcsClient.On("FileWriter", testutils.AnyContext, fmt.Sprintf("ingest/2022/04/12/13/%s/results.json", benchmarkName), gcs.FileWriteOptions{
		ContentEncoding: "application/json",
	}).Return(&b)

	err := uploadResultsFile(ctx, gcsClient, benchmarkName, resultsFile)
	require.Error(t, err)
}