package backends

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/grpc/prpc"

	"go.skia.org/infra/go/buildbucket"
	"go.skia.org/infra/go/skerr"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	bpb "go.chromium.org/luci/buildbucket/proto"
	swarmingpb "go.chromium.org/luci/common/api/swarming/swarming/v1"
	spb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// RBE CAS isolates expire after 32 days. We use 30 out of caution.
	CasExpiration = 30
	// ChromeProject refers to the "chrome" project.
	ChromeProject = "chrome"
	// ChromiumGitilesURL is the default Gitiles URL for chromium/src.
	ChromiumGitilesURL = "https://chromium.googlesource.com/chromium/src"
	// ChromiumGitilesHost is the default Gitiles host for chromium/src.
	ChromiumGitilesHost = "chromium.googlesource.com"
	// ChromiumGitilesProject is the default project name for chromium/src.
	ChromiumGitilesProject = "chromium/src"
	// ChromiumGitilesRefAtHead is the default ref used for Chromium builds.
	ChromiumGitilesRefAtHead = "refs/heads/main"
	// DefaultBucket is the Pinpoint bucket, equivalent to the "try" builds in Buildbucket.
	DefaultBucket = "try"
	// DefaultBuildsetKey is key tagged on builds for how commit information is tracked in Waterfall (CI) and Pinpoint.
	DefaultBuildsetKey = "buildset"
	// DefaultTagValue is the value format for the key above.
	DefaultBuildsetValue = "commit/gitiles/chromium.googlesource.com/chromium/src/+/%s"
	// DefaultCASInstance is the default CAS instance used by Pinpoint builds.
	//
	// TODO(b/315215756): Support other swarming instances. There are three known
	// swarming instances Pinpoint supports. The majority of Pinpoint builds are
	// this defaultInstance. Buildbucket API does not report the swarming instance
	// so our options are to:
	// - include the expected instance in the build tags
	// - try all 3 known swarming instances and brute force it
	DefaultCASInstance = "projects/chrome-swarming/instances/default_instance"
	// DefaultPerRPCTimeout defines the default time permitted for each RPC call.
	DefaultPerRPCTimeout = 90 * time.Second
	// DefaultRetries is the default number of retries for Backoff logic to Buildbucket.
	DefaultRetries = 10
	// SwarmingHashRefKey is the key used to find CAS hashes from successful Pinpoint Buildbucket builds.
	SwarmingHashRefKey = "swarm_hashes_refs"
	// WaterfallBucket is equivalent to the "ci" bucket in Buildbucket.
	WaterfallBucket = "ci"
)

type Buildbucket interface {
	// CancelBuild sends a cancellation request to Buildbucket. It's expected that
	// Buildbucket will cancel the build, whether that's graceful termination or
	// forced cancellation, as long as the request is received.
	CancelBuild(ctx context.Context, buildID int64, summary string) error

	// GetBuilds calls Buildbucket to find existing builds for the given
	// builder and Chromium revision.
	GetBuilds(ctx context.Context, builderName, bucket, commit string, patches []*bpb.GerritChange) ([]*bpb.Build, error)

	// GetBuildWithPatches calls Buildbucket to find existing builds for the
	// given builder, Chromium revision and DEPS overrides combination.
	//
	// TODO(b/315215756): The current mechanism can be updated to utilize
	// tags, so that we aren't operating on O(len(builds) * len(deps_overrides))
	// to find the exact builds. This will require tagging scheduled builds with
	// new tags before it can be utilized.
	GetBuildWithPatches(ctx context.Context, builderName, bucket, commit string, patches []*bpb.GerritChange) (*bpb.Build, error)

	// GetBuildFromWaterfall searches for an existing build using its waterfall
	// (CI) counterpart.
	GetBuildFromWaterfall(ctx context.Context, builderName, commit string) (*bpb.Build, error)

	// GetBuildStatus returns the build status given the ID.
	GetBuildStatus(ctx context.Context, buildID int64) (bpb.Status, error)

	// GetCASReference returns a CAS reference to the output artifacts of a successful build.
	GetCASReference(ctx context.Context, buildID int64, target string) (*swarmingpb.SwarmingRpcsCASReference, error)

	// StartChromeBuild triggers a Chrome build.
	StartChromeBuild(ctx context.Context, pinpointJobID, requestID, builderName, commitHash string, patches []*bpb.GerritChange) (*bpb.Build, error)
}

// BuildbucketClient is an object used to interact with a single Buildbucket instance.
// This extends Skia's Buildbucket wrapper as our single use-case is to create
// builds at specific commits.
type BuildbucketClient struct {
	client bpb.BuildsClient
}

func NewBuildbucketClient(bc bpb.BuildsClient) *BuildbucketClient {
	return &BuildbucketClient{
		client: bc,
	}
}

// createSearchBuildRequest generates a SearchBuildsRequest.
func (b BuildbucketClient) createSearchBuildRequest(builderName, bucket, commit string, patches []*bpb.GerritChange) *bpb.SearchBuildsRequest {
	tags := []*bpb.StringPair{
		{
			Key:   DefaultBuildsetKey,
			Value: fmt.Sprintf(DefaultBuildsetValue, commit),
		},
	}

	// PageSize defaults to 100, with a maximum of 1000 builds.
	req := &bpb.SearchBuildsRequest{
		Predicate: &bpb.BuildPredicate{
			Builder: &bpb.BuilderID{
				Project: ChromeProject,
				Bucket:  bucket,
				Builder: builderName,
			},
			Tags:          tags,
			GerritChanges: patches,
		},
	}

	return req
}

// CancelBuild sends a request to Buildbucket to cancel a build.
func (b BuildbucketClient) CancelBuild(ctx context.Context, buildID int64, summary string) error {
	req := &bpb.CancelBuildRequest{
		Id:              buildID,
		SummaryMarkdown: summary,
	}

	_, err := b.client.CancelBuild(ctx, req)
	if err != nil {
		return skerr.Wrapf(err, "Failed to cancel build %d.", buildID)
	}

	return nil
}

// isBuildTooOld checks whether a terminated build is too old and no longer worth checking.
// Incomplete builds have default endtime of 1970-01-01 00:00 UTC.
func (b BuildbucketClient) isBuildTooOld(build *bpb.Build) bool {
	return (build.Status.Number() > bpb.Status_ENDED_MASK.Number() &&
		time.Now().Sub(build.EndTime.AsTime()).Hours()/24 > float64(CasExpiration))
}

// findMatchingBuild searches the list of builds to find a build in good status (Success, Started, Scheduled)
// with the correct number of patchsets.
func (b BuildbucketClient) findMatchingBuild(builds []*bpb.Build, patches []*bpb.GerritChange) *bpb.Build {
	statusOK := []bpb.Status{
		bpb.Status_SUCCESS,
		bpb.Status_STARTED,
		bpb.Status_SCHEDULED,
	}

	// SearchBuilds returns all builds that contain the GerritChange instead of an exact match,
	// so this logic loops through to ensure we have an identical match.
	// Because of the sorted response (latest -> oldest), this returns the latest matched entry.
	for _, build := range builds {
		// If a completed build is past the expiration point, then all remaining
		// builds are too old, since builds are returned ordered by build number
		// and thus, newest to oldest.
		if b.isBuildTooOld(build) {
			return nil
		}

		if slices.Contains(statusOK, build.GetStatus()) && len(patches) == len(build.GetInput().GetGerritChanges()) {
			return build
		}
	}

	return nil
}

// GetBuilds calls Buildbucket's SearchBuilds.
func (b BuildbucketClient) GetBuilds(ctx context.Context, builderName, bucket, commit string, patches []*bpb.GerritChange) ([]*bpb.Build, error) {
	req := b.createSearchBuildRequest(builderName, bucket, commit, patches)
	resp, err := b.client.SearchBuilds(ctx, req)
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to call Buildbucket with request %v: ", req)
	}

	// Note: This assumes that the result we're looking for is within the first
	// 100 builds, since it's ordered from newest to oldest. Utilize NextPageToken
	// from the response to fetch more responses, or increase the PageSize up to
	// 1000.
	return resp.Builds, nil
}

// GetBuildWithPatches utilizes GetBuilds() and filters to find an exactly matching build, meaning
// that the GerritChanges and base Chromium build commit hash are the same.
func (b BuildbucketClient) GetBuildWithPatches(ctx context.Context, builderName, bucket, commit string, patches []*bpb.GerritChange) (*bpb.Build, error) {
	builds, err := b.GetBuilds(ctx, builderName, bucket, commit, patches)
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to call Buildbucket to find a single matching build.")
	}

	return b.findMatchingBuild(builds, patches), nil
}

// GetBuildFromWaterfall searches for an exactly matching Buildbucket build using information
// from the builderName's CI counterpart.
func (b BuildbucketClient) GetBuildFromWaterfall(ctx context.Context, builderName, commit string) (*bpb.Build, error) {
	mirror, ok := PinpointWaterfall[builderName]
	if !ok {
		return nil, skerr.Fmt("%s has no supported CI waterfall builder.", builderName)
	}

	builds, err := b.GetBuilds(ctx, mirror, WaterfallBucket, commit, nil)
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to find build with using CI counterpart for %s.", builderName)
	}

	// We pass an empty list of len == 0 so that Builds with GerritChanges specified
	// are ignored.
	return b.findMatchingBuild(builds, make([]*bpb.GerritChange, 0)), nil
}

// GetBuildStatus fetches the build status for a given build.
func (b BuildbucketClient) GetBuildStatus(ctx context.Context, buildID int64) (bpb.Status, error) {
	req := &bpb.GetBuildStatusRequest{
		Id: buildID,
	}

	build, err := b.client.GetBuildStatus(ctx, req)
	if err != nil {
		return bpb.Status_STATUS_UNSPECIFIED, err
	}

	return build.Status, nil
}

// createCASReferenceRequest creates a GetBuildRequest that focuses on just the output properties.
func (b BuildbucketClient) createCASReferenceRequest(buildID int64) *bpb.GetBuildRequest {
	return &bpb.GetBuildRequest{
		Id: buildID,
		// To fetch the CAS reference, we just need to focus on output properties.
		Mask: &bpb.BuildMask{
			Fields: &fieldmaskpb.FieldMask{
				Paths: []string{"output.properties"},
			},
		},
	}
}

// GetCASReference parses output.properties of a successful build for a CAS hash.
func (b BuildbucketClient) GetCASReference(ctx context.Context, buildID int64, target string) (*swarmingpb.SwarmingRpcsCASReference, error) {
	req := b.createCASReferenceRequest(buildID)
	build, err := b.client.GetBuild(ctx, req)
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to fetch CAS reference.")
	}
	if build.Status != bpb.Status_SUCCESS {
		return nil, skerr.Fmt("Cannot fetch CAS information from build %d with status %v", buildID, build.Status)
	}
	for k, v := range build.GetOutput().GetProperties().GetFields() {
		if strings.Contains(k, SwarmingHashRefKey) {
			mv := v.GetStructValue().AsMap()
			cas, ok := mv[target].(string)
			if !ok {
				return nil, skerr.Fmt("The target %s cannot be found in the output properties.", target)
			}
			// cas hash is split into two parts: {hash}/{bytes}
			parts := strings.Split(cas, "/")
			if len(parts) != 2 {
				return nil, skerr.Fmt("CAS hash %s has been changed to an unparsable format.", cas)
			}
			// base 10, 64 bit size
			bytes, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, err
			}
			return &swarmingpb.SwarmingRpcsCASReference{
				CasInstance: DefaultCASInstance,
				Digest: &swarmingpb.SwarmingRpcsDigest{
					Hash:      parts[0],
					SizeBytes: bytes,
				},
			}, nil
		}
	}

	return nil, nil
}

// createChromeBuildRequest creates a Chrome Buildbucket build request.
func (b BuildbucketClient) createChromeBuildRequest(pinpointJobID, requestID, builderName, commit string, patches []*bpb.GerritChange) *bpb.ScheduleBuildRequest {
	builder := &bpb.BuilderID{
		Project: ChromeProject,
		Bucket:  DefaultBucket,
		Builder: builderName,
	}

	properties := &spb.Struct{
		Fields: map[string]*spb.Value{
			"clobber": {
				Kind: &spb.Value_BoolValue{
					BoolValue: false,
				},
			},
			"git_repo": {
				Kind: &spb.Value_StringValue{
					StringValue: ChromiumGitilesURL,
				},
			},
			"revision": {
				Kind: &spb.Value_StringValue{
					StringValue: commit,
				},
			},
			"staging": {
				Kind: &spb.Value_BoolValue{
					BoolValue: false,
				},
			},
		},
	}

	gitilesCommit := &bpb.GitilesCommit{
		Host:    ChromiumGitilesHost,
		Project: ChromiumGitilesProject,
		Id:      commit,
		Ref:     ChromiumGitilesRefAtHead,
	}

	// TODO(b/315215756): Implement createTags function to generalize across different job types
	tags := []*bpb.StringPair{
		{
			Key:   "pinpoint_job_id",
			Value: pinpointJobID,
		},
		{
			Key:   "skia_pinpoint_bisect",
			Value: "true",
		},
		{
			Key:   DefaultBuildsetKey,
			Value: fmt.Sprintf(DefaultBuildsetValue, commit),
		},
	}

	return &bpb.ScheduleBuildRequest{
		RequestId:     requestID,
		Builder:       builder,
		Properties:    properties,
		GitilesCommit: gitilesCommit,
		GerritChanges: patches,
		Tags:          tags,
	}
}

// StartChromeBuild schedules a Chrome build through Buildbucket.
func (b BuildbucketClient) StartChromeBuild(ctx context.Context, pinpointJobID, requestID, builderName, commitHash string, patches []*bpb.GerritChange) (*bpb.Build, error) {
	if pinpointJobID == "" {
		pinpointJobID = uuid.New().String()
	}
	if requestID == "" {
		requestID = uuid.New().String()
	}

	req := b.createChromeBuildRequest(pinpointJobID, requestID, builderName, commitHash, patches)
	build, err := b.client.ScheduleBuild(ctx, req)
	if err != nil {
		return nil, skerr.Wrapf(err, "Failed to schedule build with Buildbucket")
	}
	return build, nil
}

// BuildbucketClientConfig represents options for the behavior of the Buildbucket client.
//
// Example:
// bc := DefaultClientConfig().WithClient(c)
// bc.GetBuilds(...)
type BuildbucketClientConfig struct {
	// The buildbucket host to target. See "go.skia.org/infra/go/buildbucket"
	// for the default value.
	Host string

	// Retries, if >= 0, is the number of remaining retries. If <0, no retry
	// count will be applied.
	Retries int

	// Delay is the next generated delay.
	Delay time.Duration

	// MaxDelay is the maximum duration. If <= zero, no maximum will be enforced.
	MaxDelay time.Duration

	// PerRPCTimeout, if > 0, is a timeout that is applied to each call attempt.
	PerRPCTimeout time.Duration
}

// DefaultClientConfig returns a BuildbucketClientConfig with defaults:
//   - Host: cr-buildbucket.appspot.com
//   - Exponential backoff with 10 retries
//   - PerRPCTimeout of 90 seconds. Swarming servers have an internal 60-second
//     deadline to respond to requests.
func DefaultClientConfig() BuildbucketClientConfig {
	return BuildbucketClientConfig{
		Host:          buildbucket.DEFAULT_HOST,
		Retries:       DefaultRetries,
		Delay:         time.Second,
		MaxDelay:      time.Minute,
		PerRPCTimeout: DefaultPerRPCTimeout,
	}
}

// WithClient returns a BuildbucketClient as configured by the ClientConfig
func (bc BuildbucketClientConfig) WithClient(c *http.Client) *BuildbucketClient {
	return &BuildbucketClient{
		client: bpb.NewBuildsPRPCClient(
			&prpc.Client{
				C:    c,
				Host: bc.Host,
				Options: &prpc.Options{
					Retry: func() retry.Iterator {
						return &retry.ExponentialBackoff{
							MaxDelay: bc.MaxDelay,
							Limited: retry.Limited{
								Delay:   bc.Delay,
								Retries: bc.Retries,
							},
						}
					},
					PerRPCTimeout: 90 * time.Second,
				},
			},
		),
	}
}
