package poller

// Initializes and polls the various issue frameworks.

import (
	"context"
	"os/user"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"

	"go.skia.org/infra/bugs-central/go/bugs"
	"go.skia.org/infra/bugs-central/go/bugs/github"
	"go.skia.org/infra/bugs-central/go/bugs/issuetracker"
	"go.skia.org/infra/bugs-central/go/bugs/monorail"
	"go.skia.org/infra/bugs-central/go/db"
	"go.skia.org/infra/bugs-central/go/types"
	"go.skia.org/infra/go/baseapp"
	"go.skia.org/infra/go/cleanup"
	github_lib "go.skia.org/infra/go/github"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sklog"
)

const (
	// All recognized clients.
	AndroidClient       types.RecognizedClient = "Android"
	ChromiumClient      types.RecognizedClient = "Chromium"
	FlutterNativeClient types.RecognizedClient = "Flutter-native"
	FlutterOnWebClient  types.RecognizedClient = "Flutter-on-web"
	SkiaClient          types.RecognizedClient = "Skia"
)

// IssuesPoller will be used to poll the different issue frameworks.
type IssuesPoller struct {
	storageClient            *storage.Client
	pathToGithubToken        string
	pathToServiceAccountFile string

	dbClient *db.FirestoreDB
}

// New returns an instance of IssuesPoller.
func New(ctx context.Context, ts oauth2.TokenSource, pathToServiceAccountFile string, dbClient *db.FirestoreDB) (*IssuesPoller, error) {
	httpClient := httputils.DefaultClientConfig().WithTokenSource(ts).With2xxOnly().Client()
	storageClient, err := storage.NewClient(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, skerr.Wrapf(err, "failed to init storage client")
	}

	pathToGithubToken := filepath.Join(github_lib.GITHUB_TOKEN_SERVER_PATH, github_lib.GITHUB_TOKEN_FILENAME)
	if *baseapp.Local {
		usr, err := user.Current()
		if err != nil {
			return nil, err
		}
		pathToGithubToken = filepath.Join(usr.HomeDir, github_lib.GITHUB_TOKEN_FILENAME)
	}

	return &IssuesPoller{
		storageClient:            storageClient,
		pathToGithubToken:        pathToGithubToken,
		pathToServiceAccountFile: pathToServiceAccountFile,
		dbClient:                 dbClient,
	}, nil
}

// Start polls the different issue frameworks and populates DB and an in-memory object with that data.
// It hardcodes information about Skia's various clients. It may be possible to extract some/all of these into
// flags or YAML config files in the future.
func (p *IssuesPoller) Start(ctx context.Context, pollInterval time.Duration) error {
	// Instantiate the in-memory open issues object that will be passed to the different frameworks to
	// populate.
	openIssues := bugs.InitOpenIssues()

	// Instantiate the bug frameworks with the different client configurations and then poll them.
	bugFrameworks := []bugs.BugFramework{}

	//////////////////// Android - IssueTracker ////////////////////
	androidQueryConfig := &issuetracker.IssueTrackerQueryConfig{
		Query:               "componentid:1346 status:open",
		Client:              AndroidClient,
		UntriagedPriorities: []string{"P4"},
		UntriagedAliases:    []string{"skia-android-triage@google.com"},
	}
	androidIssueTracker, err := issuetracker.New(p.storageClient, openIssues, androidQueryConfig)
	if err != nil {
		return skerr.Wrapf(err, "failed to init issuetracker for android")
	}
	bugFrameworks = append(bugFrameworks, androidIssueTracker)

	//////////////////// Flutter_on_web - Github ////////////////////
	flutterOnWebQueryConfig := &github.GithubQueryConfig{
		Labels:           []string{"e: web_canvaskit"},
		Open:             true,
		PriorityRequired: true,
		Client:           FlutterOnWebClient,
	}
	flutterOnWebGithub, err := github.New(ctx, "flutter", "flutter", p.pathToGithubToken, openIssues, flutterOnWebQueryConfig)
	if err != nil {
		return skerr.Wrapf(err, "failed to init github for flutter-on-web")
	}
	bugFrameworks = append(bugFrameworks, flutterOnWebGithub)

	//////////////////// Flutter_native - Github ////////////////////
	flutterNativeQueryConfig := &github.GithubQueryConfig{
		Labels:           []string{"dependency: skia"},
		ExcludeLabels:    []string{"e: web_canvaskit"}, // These issues are already covered by flutter-on-web
		Open:             true,
		PriorityRequired: false,
		Client:           FlutterNativeClient,
	}
	flutterNativeGithub, err := github.New(ctx, "flutter", "flutter", p.pathToGithubToken, openIssues, flutterNativeQueryConfig)
	if err != nil {
		return skerr.Wrapf(err, "failed to init github for flutter-on-web")
	}
	bugFrameworks = append(bugFrameworks, flutterNativeGithub)

	//////////////////// Chromium:Internals>Skia - Monorail ////////////////////
	crQueryConfig1 := &monorail.MonorailQueryConfig{
		Instance:          "chromium",
		Query:             "is:open component=Internals>Skia",
		Client:            ChromiumClient,
		UntriagedStatuses: []string{"Untriaged", "Unconfirmed"},
	}
	crMonorail1, err := monorail.New(ctx, p.pathToServiceAccountFile, openIssues, crQueryConfig1)
	if err != nil {
		return skerr.Wrapf(err, "failed to init monorail for chromium")
	}
	bugFrameworks = append(bugFrameworks, crMonorail1)

	//////////////////// Chromium:Internals>Skia>Compositing - Monorail ////////////////////
	crQueryConfig2 := &monorail.MonorailQueryConfig{
		Instance:          "chromium",
		Query:             "is:open component=Internals>Skia>Compositing",
		Client:            ChromiumClient,
		UntriagedStatuses: []string{"Untriaged", "Unconfirmed"},
	}
	crMonorail2, err := monorail.New(ctx, p.pathToServiceAccountFile, openIssues, crQueryConfig2)
	if err != nil {
		return skerr.Wrapf(err, "failed to init monorail for chromium")
	}
	bugFrameworks = append(bugFrameworks, crMonorail2)

	//////////////////// Chromium:Internals>Skia>PDF - Monorail ////////////////////
	crQueryConfig3 := &monorail.MonorailQueryConfig{
		Instance:          "chromium",
		Query:             "is:open component=Internals>Skia>PDF",
		Client:            ChromiumClient,
		UntriagedStatuses: []string{"Untriaged", "Unconfirmed"},
	}
	crMonorail3, err := monorail.New(ctx, p.pathToServiceAccountFile, openIssues, crQueryConfig3)
	if err != nil {
		return skerr.Wrapf(err, "failed to init monorail for chromium")
	}
	bugFrameworks = append(bugFrameworks, crMonorail3)

	//////////////////// Skia - Monorail ////////////////////
	skQueryConfig := &monorail.MonorailQueryConfig{
		Instance:          "skia",
		Query:             "is:open",
		Client:            SkiaClient,
		UntriagedStatuses: []string{"New"},
	}
	skMonorail, err := monorail.New(ctx, p.pathToServiceAccountFile, openIssues, skQueryConfig)
	if err != nil {
		return skerr.Wrapf(err, "failed to init monorail for skia")
	}
	bugFrameworks = append(bugFrameworks, skMonorail)

	cleanup.Repeat(pollInterval, func(ctx context.Context) {
		if !*baseapp.Local {
			// Ignore the passed-in context; this allows us to continue running even if the
			// context is canceled due to transient errors.
			ctx = context.Background()
		}

		// Create a runID timestamp to associate all found issues with this poll iteration.
		runId := p.dbClient.GenerateRunId(time.Now())

		// Search all bug frameworks.
		for _, b := range bugFrameworks {
			if err := b.SearchClientAndPersist(ctx, p.dbClient, runId); err != nil {
				sklog.Errorf("Error when searching and saving issues: %s", err)
				return
			}
		}

		// We are done with this iteration. Add the runId timestamp to the DB.
		if err := p.dbClient.StoreRunId(context.Background(), runId); err != nil {
			sklog.Errorf("Could not store runId in DB: %s", err)
			return
		}

		openIssues.PrettyPrintOpenIssues()
	}, nil)

	return nil
}
