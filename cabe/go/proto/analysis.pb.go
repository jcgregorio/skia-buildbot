// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: cabe/proto/v1/analysis.proto

package proto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AnalysisMetadata defines the metadata of an analysis.
type AnalysisMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The report_id of an analysis
	ReportId    string               `protobuf:"bytes,1,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Diagnostics *AnalysisDiagnostics `protobuf:"bytes,2,opt,name=diagnostics,proto3" json:"diagnostics,omitempty"`
}

func (x *AnalysisMetadata) Reset() {
	*x = AnalysisMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisMetadata) ProtoMessage() {}

func (x *AnalysisMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisMetadata.ProtoReflect.Descriptor instead.
func (*AnalysisMetadata) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *AnalysisMetadata) GetReportId() string {
	if x != nil {
		return x.ReportId
	}
	return ""
}

func (x *AnalysisMetadata) GetDiagnostics() *AnalysisDiagnostics {
	if x != nil {
		return x.Diagnostics
	}
	return nil
}

// AnalysisDiagnostics contains diagnostic messages generated by the Analyzer about
// the replica task pairs and individual tasks during its analysis.
type AnalysisDiagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Things that had to be excluded from the analysis, and why.
	ExcludedSwarmingTasks []*SwarmingTaskDiagnostics `protobuf:"bytes,1,rep,name=excluded_swarming_tasks,json=excludedSwarmingTasks,proto3" json:"excluded_swarming_tasks,omitempty"`
	ExcludedReplicas      []*ReplicaDiagnostics      `protobuf:"bytes,2,rep,name=excluded_replicas,json=excludedReplicas,proto3" json:"excluded_replicas,omitempty"`
	// Things that were included in the analysis as expected.
	IncludedSwarmingTasks []*SwarmingTaskDiagnostics `protobuf:"bytes,3,rep,name=included_swarming_tasks,json=includedSwarmingTasks,proto3" json:"included_swarming_tasks,omitempty"`
	IncludedReplicas      []*ReplicaDiagnostics      `protobuf:"bytes,4,rep,name=included_replicas,json=includedReplicas,proto3" json:"included_replicas,omitempty"`
}

func (x *AnalysisDiagnostics) Reset() {
	*x = AnalysisDiagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisDiagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisDiagnostics) ProtoMessage() {}

func (x *AnalysisDiagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisDiagnostics.ProtoReflect.Descriptor instead.
func (*AnalysisDiagnostics) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *AnalysisDiagnostics) GetExcludedSwarmingTasks() []*SwarmingTaskDiagnostics {
	if x != nil {
		return x.ExcludedSwarmingTasks
	}
	return nil
}

func (x *AnalysisDiagnostics) GetExcludedReplicas() []*ReplicaDiagnostics {
	if x != nil {
		return x.ExcludedReplicas
	}
	return nil
}

func (x *AnalysisDiagnostics) GetIncludedSwarmingTasks() []*SwarmingTaskDiagnostics {
	if x != nil {
		return x.IncludedSwarmingTasks
	}
	return nil
}

func (x *AnalysisDiagnostics) GetIncludedReplicas() []*ReplicaDiagnostics {
	if x != nil {
		return x.IncludedReplicas
	}
	return nil
}

type SwarmingTaskId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId  string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *SwarmingTaskId) Reset() {
	*x = SwarmingTaskId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwarmingTaskId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwarmingTaskId) ProtoMessage() {}

func (x *SwarmingTaskId) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwarmingTaskId.ProtoReflect.Descriptor instead.
func (*SwarmingTaskId) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{2}
}

func (x *SwarmingTaskId) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SwarmingTaskId) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

// SwarmingTaskDiagnostics contains task-specific diagnostic messages
// generated by the Analyzer.
type SwarmingTaskDiagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *SwarmingTaskId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message []string        `protobuf:"bytes,2,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *SwarmingTaskDiagnostics) Reset() {
	*x = SwarmingTaskDiagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwarmingTaskDiagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwarmingTaskDiagnostics) ProtoMessage() {}

func (x *SwarmingTaskDiagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwarmingTaskDiagnostics.ProtoReflect.Descriptor instead.
func (*SwarmingTaskDiagnostics) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{3}
}

func (x *SwarmingTaskDiagnostics) GetId() *SwarmingTaskId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SwarmingTaskDiagnostics) GetMessage() []string {
	if x != nil {
		return x.Message
	}
	return nil
}

// ReplicaDiagnostics contains replica, or task pair-specific diagnostic messages
// generated by the Analyzer.
type ReplicaDiagnostics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaNumber int32           `protobuf:"varint,1,opt,name=replica_number,json=replicaNumber,proto3" json:"replica_number,omitempty"`
	ControlTask   *SwarmingTaskId `protobuf:"bytes,2,opt,name=control_task,json=controlTask,proto3" json:"control_task,omitempty"`
	TreatmentTask *SwarmingTaskId `protobuf:"bytes,3,opt,name=treatment_task,json=treatmentTask,proto3" json:"treatment_task,omitempty"`
	Message       []string        `protobuf:"bytes,4,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *ReplicaDiagnostics) Reset() {
	*x = ReplicaDiagnostics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicaDiagnostics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaDiagnostics) ProtoMessage() {}

func (x *ReplicaDiagnostics) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaDiagnostics.ProtoReflect.Descriptor instead.
func (*ReplicaDiagnostics) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{4}
}

func (x *ReplicaDiagnostics) GetReplicaNumber() int32 {
	if x != nil {
		return x.ReplicaNumber
	}
	return 0
}

func (x *ReplicaDiagnostics) GetControlTask() *SwarmingTaskId {
	if x != nil {
		return x.ControlTask
	}
	return nil
}

func (x *ReplicaDiagnostics) GetTreatmentTask() *SwarmingTaskId {
	if x != nil {
		return x.TreatmentTask
	}
	return nil
}

func (x *ReplicaDiagnostics) GetMessage() []string {
	if x != nil {
		return x.Message
	}
	return nil
}

// AnalysisResult defines the result of an analysis
type AnalysisResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Analysis result id (PK)
	ResultId string `protobuf:"bytes,1,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Analysis experiment spec
	ExperimentSpec *ExperimentSpec `protobuf:"bytes,2,opt,name=experiment_spec,json=experimentSpec,proto3" json:"experiment_spec,omitempty"`
	// The metadata of the analysis
	AnalysisMetadata *AnalysisMetadata `protobuf:"bytes,3,opt,name=analysis_metadata,json=analysisMetadata,proto3" json:"analysis_metadata,omitempty"`
	// The calculated statistic of the analysis
	Statistic *Statistic `protobuf:"bytes,4,opt,name=statistic,proto3" json:"statistic,omitempty"`
}

func (x *AnalysisResult) Reset() {
	*x = AnalysisResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisResult) ProtoMessage() {}

func (x *AnalysisResult) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisResult.ProtoReflect.Descriptor instead.
func (*AnalysisResult) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{5}
}

func (x *AnalysisResult) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

func (x *AnalysisResult) GetExperimentSpec() *ExperimentSpec {
	if x != nil {
		return x.ExperimentSpec
	}
	return nil
}

func (x *AnalysisResult) GetAnalysisMetadata() *AnalysisMetadata {
	if x != nil {
		return x.AnalysisMetadata
	}
	return nil
}

func (x *AnalysisResult) GetStatistic() *Statistic {
	if x != nil {
		return x.Statistic
	}
	return nil
}

// Statistic defines the statistic of an analysis
type Statistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The lower bound of the analysis result
	Lower float64 `protobuf:"fixed64,1,opt,name=lower,proto3" json:"lower,omitempty"`
	// The upper bound of the analysis result
	Upper float64 `protobuf:"fixed64,2,opt,name=upper,proto3" json:"upper,omitempty"`
	// The P value of the analysis result
	PValue float64 `protobuf:"fixed64,3,opt,name=p_value,json=pValue,proto3" json:"p_value,omitempty"`
	// The defined significance level to calculate the lower and upper bound
	SignificanceLevel float64 `protobuf:"fixed64,4,opt,name=significance_level,json=significanceLevel,proto3" json:"significance_level,omitempty"`
	// The point estimate of the analysis result
	PointEstimate float64 `protobuf:"fixed64,6,opt,name=point_estimate,json=pointEstimate,proto3" json:"point_estimate,omitempty"`
	// The median of control arm of the analysis result
	ControlMedian float64 `protobuf:"fixed64,7,opt,name=control_median,json=controlMedian,proto3" json:"control_median,omitempty"`
	// The median of treatment arm of the analysis result
	TreatmentMedian float64 `protobuf:"fixed64,8,opt,name=treatment_median,json=treatmentMedian,proto3" json:"treatment_median,omitempty"`
}

func (x *Statistic) Reset() {
	*x = Statistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cabe_proto_v1_analysis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistic) ProtoMessage() {}

func (x *Statistic) ProtoReflect() protoreflect.Message {
	mi := &file_cabe_proto_v1_analysis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistic.ProtoReflect.Descriptor instead.
func (*Statistic) Descriptor() ([]byte, []int) {
	return file_cabe_proto_v1_analysis_proto_rawDescGZIP(), []int{6}
}

func (x *Statistic) GetLower() float64 {
	if x != nil {
		return x.Lower
	}
	return 0
}

func (x *Statistic) GetUpper() float64 {
	if x != nil {
		return x.Upper
	}
	return 0
}

func (x *Statistic) GetPValue() float64 {
	if x != nil {
		return x.PValue
	}
	return 0
}

func (x *Statistic) GetSignificanceLevel() float64 {
	if x != nil {
		return x.SignificanceLevel
	}
	return 0
}

func (x *Statistic) GetPointEstimate() float64 {
	if x != nil {
		return x.PointEstimate
	}
	return 0
}

func (x *Statistic) GetControlMedian() float64 {
	if x != nil {
		return x.ControlMedian
	}
	return 0
}

func (x *Statistic) GetTreatmentMedian() float64 {
	if x != nil {
		return x.TreatmentMedian
	}
	return 0
}

var File_cabe_proto_v1_analysis_proto protoreflect.FileDescriptor

var file_cabe_proto_v1_analysis_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x63, 0x61, 0x62, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6f, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x22, 0xdd, 0x02, 0x0a, 0x13, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x44,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x58, 0x0a, 0x17, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x61,
	0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x15, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x10, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x58,
	0x0a, 0x17, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x73, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x52, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x53, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x48, 0x0a, 0x11, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x22, 0x43, 0x0a, 0x0e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5c, 0x0a, 0x17, 0x53, 0x77, 0x61, 0x72, 0x6d,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x12, 0x27, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x62, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x3e, 0x0a, 0x0e, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x77, 0x61, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x52, 0x0d, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x0e, 0x41, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x46, 0x0a, 0x11, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x10, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x62, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x22, 0xf8, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x66, 0x69, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0d, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x74, 0x72, 0x65, 0x61, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x6f, 0x2e, 0x73, 0x6b, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x63, 0x61, 0x62, 0x65, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cabe_proto_v1_analysis_proto_rawDescOnce sync.Once
	file_cabe_proto_v1_analysis_proto_rawDescData = file_cabe_proto_v1_analysis_proto_rawDesc
)

func file_cabe_proto_v1_analysis_proto_rawDescGZIP() []byte {
	file_cabe_proto_v1_analysis_proto_rawDescOnce.Do(func() {
		file_cabe_proto_v1_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_cabe_proto_v1_analysis_proto_rawDescData)
	})
	return file_cabe_proto_v1_analysis_proto_rawDescData
}

var file_cabe_proto_v1_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cabe_proto_v1_analysis_proto_goTypes = []interface{}{
	(*AnalysisMetadata)(nil),        // 0: cabe.v1.AnalysisMetadata
	(*AnalysisDiagnostics)(nil),     // 1: cabe.v1.AnalysisDiagnostics
	(*SwarmingTaskId)(nil),          // 2: cabe.v1.SwarmingTaskId
	(*SwarmingTaskDiagnostics)(nil), // 3: cabe.v1.SwarmingTaskDiagnostics
	(*ReplicaDiagnostics)(nil),      // 4: cabe.v1.ReplicaDiagnostics
	(*AnalysisResult)(nil),          // 5: cabe.v1.AnalysisResult
	(*Statistic)(nil),               // 6: cabe.v1.Statistic
	(*ExperimentSpec)(nil),          // 7: cabe.v1.ExperimentSpec
}
var file_cabe_proto_v1_analysis_proto_depIdxs = []int32{
	1,  // 0: cabe.v1.AnalysisMetadata.diagnostics:type_name -> cabe.v1.AnalysisDiagnostics
	3,  // 1: cabe.v1.AnalysisDiagnostics.excluded_swarming_tasks:type_name -> cabe.v1.SwarmingTaskDiagnostics
	4,  // 2: cabe.v1.AnalysisDiagnostics.excluded_replicas:type_name -> cabe.v1.ReplicaDiagnostics
	3,  // 3: cabe.v1.AnalysisDiagnostics.included_swarming_tasks:type_name -> cabe.v1.SwarmingTaskDiagnostics
	4,  // 4: cabe.v1.AnalysisDiagnostics.included_replicas:type_name -> cabe.v1.ReplicaDiagnostics
	2,  // 5: cabe.v1.SwarmingTaskDiagnostics.id:type_name -> cabe.v1.SwarmingTaskId
	2,  // 6: cabe.v1.ReplicaDiagnostics.control_task:type_name -> cabe.v1.SwarmingTaskId
	2,  // 7: cabe.v1.ReplicaDiagnostics.treatment_task:type_name -> cabe.v1.SwarmingTaskId
	7,  // 8: cabe.v1.AnalysisResult.experiment_spec:type_name -> cabe.v1.ExperimentSpec
	0,  // 9: cabe.v1.AnalysisResult.analysis_metadata:type_name -> cabe.v1.AnalysisMetadata
	6,  // 10: cabe.v1.AnalysisResult.statistic:type_name -> cabe.v1.Statistic
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_cabe_proto_v1_analysis_proto_init() }
func file_cabe_proto_v1_analysis_proto_init() {
	if File_cabe_proto_v1_analysis_proto != nil {
		return
	}
	file_cabe_proto_v1_spec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cabe_proto_v1_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisMetadata); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisDiagnostics); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwarmingTaskId); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwarmingTaskDiagnostics); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicaDiagnostics); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisResult); i {
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
		file_cabe_proto_v1_analysis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistic); i {
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
			RawDescriptor: file_cabe_proto_v1_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cabe_proto_v1_analysis_proto_goTypes,
		DependencyIndexes: file_cabe_proto_v1_analysis_proto_depIdxs,
		MessageInfos:      file_cabe_proto_v1_analysis_proto_msgTypes,
	}.Build()
	File_cabe_proto_v1_analysis_proto = out.File
	file_cabe_proto_v1_analysis_proto_rawDesc = nil
	file_cabe_proto_v1_analysis_proto_goTypes = nil
	file_cabe_proto_v1_analysis_proto_depIdxs = nil
}
