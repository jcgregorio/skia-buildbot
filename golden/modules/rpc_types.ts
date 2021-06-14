// DO NOT EDIT. This file is automatically generated.

export interface Changelist {
	system: string;
	id: string;
	owner: string;
	status: string;
	subject: string;
	updated: string;
	url: string;
}

export interface TryJob {
	id: string;
	name: string;
	updated: string;
	system: string;
	url: string;
}

export interface Patchset {
	id: string;
	order: number;
	try_jobs: TryJob[];
}

export interface ChangelistSummaryResponse {
	cl: Changelist;
	patch_sets: Patchset[];
	num_total_patch_sets: number;
}

export interface TriageHistory {
	user: string;
	ts: string;
}

export interface Trace {
	label: TraceID;
	data: number[] | null;
	params: { [key: string]: string } | null;
	comment_indices: number[] | null;
}

export interface DigestStatus {
	digest: Digest;
	status: Label;
}

export interface TraceGroup {
	traces: Trace[] | null;
	digests: DigestStatus[] | null;
	total_digests: number;
}

export interface SRDiffDigest {
	numDiffPixels: number;
	combinedMetric: number;
	pixelDiffPercent: number;
	maxRGBADiffs: number[];
	dimDiffer: boolean;
	digest: Digest;
	status: Label;
	paramset: ParamSet;
}

export interface SearchResult {
	digest: Digest;
	test: TestName;
	status: Label;
	triage_history: TriageHistory[] | null;
	paramset: ParamSet;
	traces: TraceGroup;
	refDiffs: { [key: string]: SRDiffDigest | null } | null;
	closestRef: RefClosest;
}

export interface Commit {
	commit_time: number;
	id: string;
	hash: string;
	author: string;
	message: string;
	cl_url: string;
}

export interface SearchResponse {
	digests: SearchResult[];
	offset: number;
	size: number;
	commits: Commit[] | null;
	bulk_triage_data: TriageRequestData;
}

export interface TriageRequest {
	testDigestStatus: TriageRequestData;
	changelist_id: string;
	crs: string;
	imageMatchingAlgorithm?: string;
}

export interface GUICorpusStatus {
	name: string;
	ok: boolean;
	minCommitHash: string;
	untriagedCount: number;
	negativeCount: number;
}

export interface StatusResponse {
	ok: boolean;
	firstCommit: Commit;
	lastCommit: Commit;
	totalCommits: number;
	filledCommits: number;
	corpStatus: GUICorpusStatus[];
}

export interface TestRollup {
	test: TestName;
	num: number;
	sample_digest: Digest;
}

export interface ByBlameEntry {
	groupID: string;
	nDigests: number;
	nTests: number;
	affectedTests: TestRollup[] | null;
	commits: Commit[] | null;
}

export interface ByBlameResponse {
	data: ByBlameEntry[] | null;
}

export interface TriageDelta {
	test_name: TestName;
	digest: Digest;
	label: Label;
}

export interface TriageLogEntry {
	id: string;
	name: string;
	ts: number;
	changeCount: number;
	details: TriageDelta[] | null;
}

export interface TriageLogResponse {
	entries: TriageLogEntry[] | null;
	offset: number;
	size: number;
	total: number;
}

export interface ChangelistsResponse {
	changelists: Changelist[] | null;
	offset: number;
	size: number;
	total: number;
}

export interface IgnoreRuleBody {
	duration: string;
	filter: string;
	note: string;
}

export interface IgnoreRule {
	id: string;
	name: string;
	updatedBy: string;
	expires: string;
	query: string;
	note: string;
	countAll: number;
	exclusiveCountAll: number;
	count: number;
	exclusiveCount: number;
}

export interface IgnoresResponse {
	rules: IgnoreRule[] | null;
}

export interface TestSummary {
	name: TestName;
	positive_digests: number;
	negative_digests: number;
	untriaged_digests: number;
	total_digests: number;
}

export interface ListTestsResponse {
	tests: TestSummary[] | null;
}

export interface DigestComparison {
	left: SearchResult;
	right: SRDiffDigest | null;
}

export interface DigestDetails {
	digest: SearchResult;
	commits: Commit[] | null;
}

export interface ClusterDiffNode {
	name: Digest;
	status: Label;
}

export interface ClusterDiffLink {
	source: number;
	target: number;
	value: number;
}

export interface ClusterDiffResult {
	nodes: ClusterDiffNode[] | null;
	links: ClusterDiffLink[] | null;
	test: TestName;
	paramsetByDigest: { [key: string]: ParamSet };
	paramsetsUnion: ParamSet;
}

export type ParamSet = { [key: string]: string[] };

export type ParamSetResponse = { [key: string]: string[] | null } | null;

export type Digest = string;

export type TestName = string;

export type Label = "untriaged" | "positive" | "negative";

export type TraceID = string;

export type RefClosest = "pos" | "neg" | "";

export type TriageRequestData = { [key: string]: { [key: string]: Label } };
