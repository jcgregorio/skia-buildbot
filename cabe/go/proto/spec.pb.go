// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: cabe/proto/v1/spec.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BuildSpec defines what set of executable bits we ship to test machines.
// It should include enough information to tell chrome build infrastructure how
// build the executable from scratch, or how to identify an exact version of a
// pre-built installation (e.g. 3rd party browser other than chrome)
type BuildSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Source checkout (git repo, branch, commit position)
	GitilesCommit *GitilesCommit `protobuf:"bytes,1,opt,name=gitiles_commit,json=gitilesCommit,proto3" json:"gitiles_commit,omitempty"`
	// Applied patches (get repos, branches, commit positions)
	GerritChanges []*GerritChange `protobuf:"bytes,2,rep,name=gerrit_changes,json=gerritChanges,proto3" json:"gerrit_changes,omitempty"`
	// For binaries that use a pre-built installer for CBB experiments.
	InstalledBrowser *InstalledBrowser `protobuf:"bytes,3,opt,name=installed_browser,json=installedBrowser,proto3" json:"installed_browser,omitempty"`
}

func (x *BuildSpec) Reset() {
	*x = BuildSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSpec) ProtoMessage() {}

func (x *BuildSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSpec.ProtoReflect.Descriptor instead.
func (*BuildSpec) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *BuildSpec) GetGitilesCommit() *GitilesCommit {
	if x != nil {
		return x.GitilesCommit
	}
	return nil
}

func (x *BuildSpec) GetGerritChanges() []*GerritChange {
	if x != nil {
		return x.GerritChanges
	}
	return nil
}

func (x *BuildSpec) GetInstalledBrowser() *InstalledBrowser {
	if x != nil {
		return x.InstalledBrowser
	}
	return nil
}

// A Gerrit patchset.
type GerritChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset int64 `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	// Git hash for patchset
	PatchsetHash string `protobuf:"bytes,5,opt,name=patchset_hash,json=patchsetHash,proto3" json:"patchset_hash,omitempty"`
}

func (x *GerritChange) Reset() {
	*x = GerritChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GerritChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GerritChange) ProtoMessage() {}

func (x *GerritChange) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GerritChange.ProtoReflect.Descriptor instead.
func (*GerritChange) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *GerritChange) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GerritChange) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GerritChange) GetChange() int64 {
	if x != nil {
		return x.Change
	}
	return 0
}

func (x *GerritChange) GetPatchset() int64 {
	if x != nil {
		return x.Patchset
	}
	return 0
}

func (x *GerritChange) GetPatchsetHash() string {
	if x != nil {
		return x.PatchsetHash
	}
	return ""
}

// A landed Git commit hosted on Gitiles.
type GitilesCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Gitiles hostname, e.g. "chromium.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Repository name on the host, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Commit HEX SHA1.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Commit ref, e.g. "refs/heads/master".
	// NOT a branch name: if specified, must start with "refs/".
	// If id is set, ref SHOULD also be set, so that git clients can
	// know how to obtain the commit by id.
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Defines a total order of commits on the ref. Requires ref field.
	// Typically 1-based, monotonically increasing, contiguous integer
	// defined by a Gerrit plugin, goto.google.com/git-numberer.
	Position uint32 `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *GitilesCommit) Reset() {
	*x = GitilesCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitilesCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitilesCommit) ProtoMessage() {}

func (x *GitilesCommit) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitilesCommit.ProtoReflect.Descriptor instead.
func (*GitilesCommit) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{2}
}

func (x *GitilesCommit) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *GitilesCommit) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GitilesCommit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GitilesCommit) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *GitilesCommit) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

// Third-party browser builds, not necessarily Chrome.
// These are primarily intended for use by CBB, since it needs to
// compare Chrome to Safari, Edge etc.  These obviously aren't built from source
// but we still need to describe what set of executable bits the benchmark
// exercised.
type InstalledBrowser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. "chrome" or "safari"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// e.g. "104.0.5112.101" or "15.5"
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *InstalledBrowser) Reset() {
	*x = InstalledBrowser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstalledBrowser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstalledBrowser) ProtoMessage() {}

func (x *InstalledBrowser) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstalledBrowser.ProtoReflect.Descriptor instead.
func (*InstalledBrowser) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{3}
}

func (x *InstalledBrowser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InstalledBrowser) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Finch config for Chrome.
type FinchConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. seed hash, seed change list, and seed timestamp.
	SeedHash       string                 `protobuf:"bytes,1,opt,name=seed_hash,json=seedHash,proto3" json:"seed_hash,omitempty"`
	SeedChangelist uint64                 `protobuf:"varint,2,opt,name=seed_changelist,json=seedChangelist,proto3" json:"seed_changelist,omitempty"`
	SeedTimestamp  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=seed_timestamp,json=seedTimestamp,proto3" json:"seed_timestamp,omitempty"`
}

func (x *FinchConfig) Reset() {
	*x = FinchConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinchConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinchConfig) ProtoMessage() {}

func (x *FinchConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinchConfig.ProtoReflect.Descriptor instead.
func (*FinchConfig) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{4}
}

func (x *FinchConfig) GetSeedHash() string {
	if x != nil {
		return x.SeedHash
	}
	return ""
}

func (x *FinchConfig) GetSeedChangelist() uint64 {
	if x != nil {
		return x.SeedChangelist
	}
	return 0
}

func (x *FinchConfig) GetSeedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.SeedTimestamp
	}
	return nil
}

// RunSpec defines where and how to execute the executable bits from a
// BuildSpec. It should include enough information to schedule or locate a set
// of Swarming tasks for a given BuildSpec and AnalysisSpec.
type RunSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OS strings will contain both the OS name and any OS-specific version
	// details.
	Os string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	// Synthetic product names come from Swarming, and uniquely identify specific
	// hardware device configurations.
	SyntheticProductName string `protobuf:"bytes,2,opt,name=synthetic_product_name,json=syntheticProductName,proto3" json:"synthetic_product_name,omitempty"`
	// Finch config (seed hash, change list, and timestamp).
	FinchConfig *FinchConfig `protobuf:"bytes,3,opt,name=finch_config,json=finchConfig,proto3" json:"finch_config,omitempty"`
}

func (x *RunSpec) Reset() {
	*x = RunSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunSpec) ProtoMessage() {}

func (x *RunSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunSpec.ProtoReflect.Descriptor instead.
func (*RunSpec) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{5}
}

func (x *RunSpec) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *RunSpec) GetSyntheticProductName() string {
	if x != nil {
		return x.SyntheticProductName
	}
	return ""
}

func (x *RunSpec) GetFinchConfig() *FinchConfig {
	if x != nil {
		return x.FinchConfig
	}
	return nil
}

// AnalysisSpec defines what benchmarks and measurements we expect to analyze
// from a set of RunSpecs. This type should include all observed potential
// response variables for the experiment.
type AnalysisSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of benchmarks, stories, metrics. CABE ETL will use this as sort of
	// a manifest for results data - it will check to make sure all of these
	// are actually present in the benchmark jobs' collected output.
	Benchmark []*Benchmark `protobuf:"bytes,1,rep,name=benchmark,proto3" json:"benchmark,omitempty"`
}

func (x *AnalysisSpec) Reset() {
	*x = AnalysisSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisSpec) ProtoMessage() {}

func (x *AnalysisSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisSpec.ProtoReflect.Descriptor instead.
func (*AnalysisSpec) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{6}
}

func (x *AnalysisSpec) GetBenchmark() []*Benchmark {
	if x != nil {
		return x.Benchmark
	}
	return nil
}

// Benchmark encapsulates both the coarse grained benchmark suite name
// and all of the more specific workloads (or "stories", to use older
// terminology) that generate measurements.
type Benchmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         // e.g. "Speedometer2"
	Workload []string `protobuf:"bytes,2,rep,name=workload,proto3" json:"workload,omitempty"` // e.g. "React-TodoMVC"
}

func (x *Benchmark) Reset() {
	*x = Benchmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benchmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benchmark) ProtoMessage() {}

func (x *Benchmark) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benchmark.ProtoReflect.Descriptor instead.
func (*Benchmark) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{7}
}

func (x *Benchmark) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Benchmark) GetWorkload() []string {
	if x != nil {
		return x.Workload
	}
	return nil
}

// ArmSpec defines how to build and execute one arm of a performance benchmark
// A/B test. This type should include all observed potential explanatory
// variables for the experiment.
type ArmSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildSpec []*BuildSpec `protobuf:"bytes,1,rep,name=build_spec,json=buildSpec,proto3" json:"build_spec,omitempty"`
	RunSpec   []*RunSpec   `protobuf:"bytes,2,rep,name=run_spec,json=runSpec,proto3" json:"run_spec,omitempty"`
}

func (x *ArmSpec) Reset() {
	*x = ArmSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArmSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArmSpec) ProtoMessage() {}

func (x *ArmSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArmSpec.ProtoReflect.Descriptor instead.
func (*ArmSpec) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{8}
}

func (x *ArmSpec) GetBuildSpec() []*BuildSpec {
	if x != nil {
		return x.BuildSpec
	}
	return nil
}

func (x *ArmSpec) GetRunSpec() []*RunSpec {
	if x != nil {
		return x.RunSpec
	}
	return nil
}

// ExperimentSpec contains all of the necessary information to build, execute
// and analyze a set of benchmark metrics for a controlled experiment.
type ExperimentSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// common contains all of the build/run details that are common to all
	// arms of an experiment. For instance, if you are comparing two different
	// browser build versions across mac, windows and linux, then the mac, windows
	// and linux details would go in the common ArmSpec.  The control and
	// treatment armspecs wouldn't mention mac, windows or linux details since
	// the are implied by the common armspec.
	// Any details specified in both the common ArmSpec and any other arms'
	// ArmSpecs should indicate an unsresolved, invalid ExperimentSpec.
	Common *ArmSpec `protobuf:"bytes,1,opt,name=common,proto3" json:"common,omitempty"`
	// Control and Treatment are somewhat arbitrary distinctions and their meaning
	// is use-case dependent. Values in their ArmSpecs should not conflict
	// with anything in the common ArmSpec.
	Control *ArmSpec `protobuf:"bytes,2,opt,name=control,proto3" json:"control,omitempty"`
	// Treatment may change in the future to be a repeated field to better
	// represent multi-arm trials but for now we'll limit it to a single value.
	Treatment *ArmSpec `protobuf:"bytes,3,opt,name=treatment,proto3" json:"treatment,omitempty"`
	// Analysis describes how we expect CABE to compare the arms of the
	// experiment.
	Analysis *AnalysisSpec `protobuf:"bytes,4,opt,name=analysis,proto3" json:"analysis,omitempty"`
}

func (x *ExperimentSpec) Reset() {
	*x = ExperimentSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_spec_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentSpec) ProtoMessage() {}

func (x *ExperimentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_spec_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentSpec.ProtoReflect.Descriptor instead.
func (*ExperimentSpec) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_spec_proto_rawDescGZIP(), []int{9}
}

func (x *ExperimentSpec) GetCommon() *ArmSpec {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *ExperimentSpec) GetControl() *ArmSpec {
	if x != nil {
		return x.Control
	}
	return nil
}

func (x *ExperimentSpec) GetTreatment() *ArmSpec {
	if x != nil {
		return x.Treatment
	}
	return nil
}

func (x *ExperimentSpec) GetAnalysis() *AnalysisSpec {
	if x != nil {
		return x.Analysis
	}
	return nil
}

var File_cabe_proto_v1_spec_proto protoreflect.FileDescriptor

var file_cabe_proto_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x62, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x3d, 0x0a, 0x0e, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x61, 0x62,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x52, 0x0d, 0x67, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x62, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x0d, 0x67, 0x65, 0x72, 0x72, 0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12,
	0x46, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x5f, 0x62, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x61, 0x62,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64,
	0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x63, 0x68, 0x73, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22,
	0x7b, 0x0a, 0x0d, 0x47, 0x69, 0x74, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x10,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x96,
	0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x65, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x73, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x73, 0x65, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x6f, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63,
	0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x74, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x66, 0x69, 0x6e,
	0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x40, 0x0a, 0x0c, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x30, 0x0a, 0x09, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x09, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x22, 0x3b, 0x0a, 0x09, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x69, 0x0a, 0x07, 0x41, 0x72, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a, 0x0a,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x2b, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x22, 0xc9, 0x01, 0x0a,
	0x0e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x6d, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x62,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x72, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x74, 0x72, 0x65, 0x61,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x6f, 0x2e, 0x73,
	0x6b, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x61,
	0x62, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cabe_proto_v1_spec_proto_rawDescOnce sync.Once
	file_cabe_proto_v1_spec_proto_rawDescData = file_cabe_proto_v1_spec_proto_rawDesc
)

func file_cabe_proto_v1_spec_proto_rawDescGZIP() []byte {
	file_cabe_proto_v1_spec_proto_rawDescOnce.Do(func() {
		file_cabe_proto_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_cabe_proto_v1_spec_proto_rawDescData)
	})
	return file_cabe_proto_v1_spec_proto_rawDescData
}

var file_cabe_proto_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cabe_proto_v1_spec_proto_goTypes = []interface{}{
	(*BuildSpec)(nil),             // 0: cabe.v1.BuildSpec
	(*GerritChange)(nil),          // 1: cabe.v1.GerritChange
	(*GitilesCommit)(nil),         // 2: cabe.v1.GitilesCommit
	(*InstalledBrowser)(nil),      // 3: cabe.v1.InstalledBrowser
	(*FinchConfig)(nil),           // 4: cabe.v1.FinchConfig
	(*RunSpec)(nil),               // 5: cabe.v1.RunSpec
	(*AnalysisSpec)(nil),          // 6: cabe.v1.AnalysisSpec
	(*Benchmark)(nil),             // 7: cabe.v1.Benchmark
	(*ArmSpec)(nil),               // 8: cabe.v1.ArmSpec
	(*ExperimentSpec)(nil),        // 9: cabe.v1.ExperimentSpec
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_cabe_proto_v1_spec_proto_depIdxs = []int32{
	2,  // 0: cabe.v1.BuildSpec.gitiles_commit:type_name -> cabe.v1.GitilesCommit
	1,  // 1: cabe.v1.BuildSpec.gerrit_changes:type_name -> cabe.v1.GerritChange
	3,  // 2: cabe.v1.BuildSpec.installed_browser:type_name -> cabe.v1.InstalledBrowser
	10, // 3: cabe.v1.FinchConfig.seed_timestamp:type_name -> google.protobuf.Timestamp
	4,  // 4: cabe.v1.RunSpec.finch_config:type_name -> cabe.v1.FinchConfig
	7,  // 5: cabe.v1.AnalysisSpec.benchmark:type_name -> cabe.v1.Benchmark
	0,  // 6: cabe.v1.ArmSpec.build_spec:type_name -> cabe.v1.BuildSpec
	5,  // 7: cabe.v1.ArmSpec.run_spec:type_name -> cabe.v1.RunSpec
	8,  // 8: cabe.v1.ExperimentSpec.common:type_name -> cabe.v1.ArmSpec
	8,  // 9: cabe.v1.ExperimentSpec.control:type_name -> cabe.v1.ArmSpec
	8,  // 10: cabe.v1.ExperimentSpec.treatment:type_name -> cabe.v1.ArmSpec
	6,  // 11: cabe.v1.ExperimentSpec.analysis:type_name -> cabe.v1.AnalysisSpec
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_cabe_proto_v1_spec_proto_init() }
func file_cabe_proto_v1_spec_proto_init() {
	if File_cabe_proto_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cabe_proto_v1_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GerritChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitilesCommit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstalledBrowser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinchConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benchmark); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArmSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cabe_proto_v1_spec_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cabe_proto_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cabe_proto_v1_spec_proto_goTypes,
		DependencyIndexes: file_cabe_proto_v1_spec_proto_depIdxs,
		MessageInfos:      file_cabe_proto_v1_spec_proto_msgTypes,
	}.Build()
	File_cabe_proto_v1_spec_proto = out.File
	file_cabe_proto_v1_spec_proto_rawDesc = nil
	file_cabe_proto_v1_spec_proto_goTypes = nil
	file_cabe_proto_v1_spec_proto_depIdxs = nil
}
