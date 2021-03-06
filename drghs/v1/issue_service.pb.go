// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_service.proto

package drghs_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for [DevRelGitHubService.ListIssues][].
type ListIssuesRequest struct {
	// Required. The resource name of the repository associated with the
	// [Issues][Issue], in the format `owners/*/repositories/*`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. Limit the number of [Issues][Issue] to include in the
	// response. Fewer Issues than requested might be returned.
	//
	// The maximum page size is `100`. If unspecified, the page size will be the
	// maximum. Further [Issues][Issue] can subsequently be obtained
	// by including the [ListIssuesResponse.next_page_token][] in a
	// subsequent request.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. To request the first page of results, `page_token` must be empty.
	// To request the next page of results, page_token must be the value of
	// [ListIssuesResponse.next_page_token][] returned by a previous call to
	// [Issueservice.ListIssues][].
	//
	// The page token is valid for only 2 hours.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Optional. Filter expression used to only include resources that match the
	// filter in the response. Filter must be in following the format:
	//
	//     field1=123
	//     field2="Foo bar"
	//     field3 IN (value3, value4)
	//     field4 LIKE "%somestring%"
	//
	// Valid filter fields are: `name`, `repo.name`, `repo.owner`, and `size`.
	//
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specify how the results should be sorted. The fields supported
	// for sorting are `name` and `size`.
	// The default ordering is by `name`. Prefix with `-` to specify
	// descending order, e.g. `-name`.
	OrderBy  string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Comments bool   `protobuf:"varint,6,opt,name=comments,proto3" json:"comments,omitempty"`
	Reviews  bool   `protobuf:"varint,7,opt,name=reviews,proto3" json:"reviews,omitempty"`
	// This is a workaround to allow nullable bools
	//
	// Types that are valid to be assigned to PullRequestNullable:
	//	*ListIssuesRequest_PullRequest
	PullRequestNullable isListIssuesRequest_PullRequestNullable `protobuf_oneof:"pull_request_nullable"`
	// This is a workaround to allow nullable bools
	//
	// Types that are valid to be assigned to ClosedNullable:
	//	*ListIssuesRequest_Closed
	ClosedNullable isListIssuesRequest_ClosedNullable `protobuf_oneof:"closed_nullable"`
	// If the FieldMask is NOT set or empty, all fields are returned. If the
	// FieldMask is set, only the specified fields are returned. See
	// https://pkg.go.dev/google.golang.org/genproto/protobuf/field_mask.
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,10,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListIssuesRequest) Reset()         { *m = ListIssuesRequest{} }
func (m *ListIssuesRequest) String() string { return proto.CompactTextString(m) }
func (*ListIssuesRequest) ProtoMessage()    {}
func (*ListIssuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{0}
}

func (m *ListIssuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIssuesRequest.Unmarshal(m, b)
}
func (m *ListIssuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIssuesRequest.Marshal(b, m, deterministic)
}
func (m *ListIssuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIssuesRequest.Merge(m, src)
}
func (m *ListIssuesRequest) XXX_Size() int {
	return xxx_messageInfo_ListIssuesRequest.Size(m)
}
func (m *ListIssuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIssuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIssuesRequest proto.InternalMessageInfo

func (m *ListIssuesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListIssuesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListIssuesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListIssuesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListIssuesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *ListIssuesRequest) GetComments() bool {
	if m != nil {
		return m.Comments
	}
	return false
}

func (m *ListIssuesRequest) GetReviews() bool {
	if m != nil {
		return m.Reviews
	}
	return false
}

type isListIssuesRequest_PullRequestNullable interface {
	isListIssuesRequest_PullRequestNullable()
}

type ListIssuesRequest_PullRequest struct {
	PullRequest bool `protobuf:"varint,8,opt,name=pull_request,json=pullRequest,proto3,oneof"`
}

func (*ListIssuesRequest_PullRequest) isListIssuesRequest_PullRequestNullable() {}

func (m *ListIssuesRequest) GetPullRequestNullable() isListIssuesRequest_PullRequestNullable {
	if m != nil {
		return m.PullRequestNullable
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListIssuesRequest) GetPullRequest() bool {
	if x, ok := m.GetPullRequestNullable().(*ListIssuesRequest_PullRequest); ok {
		return x.PullRequest
	}
	return false
}

type isListIssuesRequest_ClosedNullable interface {
	isListIssuesRequest_ClosedNullable()
}

type ListIssuesRequest_Closed struct {
	Closed bool `protobuf:"varint,9,opt,name=closed,proto3,oneof"`
}

func (*ListIssuesRequest_Closed) isListIssuesRequest_ClosedNullable() {}

func (m *ListIssuesRequest) GetClosedNullable() isListIssuesRequest_ClosedNullable {
	if m != nil {
		return m.ClosedNullable
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListIssuesRequest) GetClosed() bool {
	if x, ok := m.GetClosedNullable().(*ListIssuesRequest_Closed); ok {
		return x.Closed
	}
	return false
}

func (m *ListIssuesRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListIssuesRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListIssuesRequest_PullRequest)(nil),
		(*ListIssuesRequest_Closed)(nil),
	}
}

// Response message for [DevRelGitHubService.ListIssues][].
type ListIssuesResponse struct {
	// The list of [Issues][Issue].
	Issues []*Issue `protobuf:"bytes,1,rep,name=issues,proto3" json:"issues,omitempty"`
	// A token to retrieve the next page of results, or empty if there are no
	// more results in the list. Pass this value in
	// [ListIssuesRequest.page_token][] to retrieve the next page of
	// results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of [Issues][Issue] that matched the query.
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIssuesResponse) Reset()         { *m = ListIssuesResponse{} }
func (m *ListIssuesResponse) String() string { return proto.CompactTextString(m) }
func (*ListIssuesResponse) ProtoMessage()    {}
func (*ListIssuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{1}
}

func (m *ListIssuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIssuesResponse.Unmarshal(m, b)
}
func (m *ListIssuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIssuesResponse.Marshal(b, m, deterministic)
}
func (m *ListIssuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIssuesResponse.Merge(m, src)
}
func (m *ListIssuesResponse) XXX_Size() int {
	return xxx_messageInfo_ListIssuesResponse.Size(m)
}
func (m *ListIssuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIssuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIssuesResponse proto.InternalMessageInfo

func (m *ListIssuesResponse) GetIssues() []*Issue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ListIssuesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListIssuesResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Response message for [DevRelGitHubService.GetIssue][].
type GetIssueRequest struct {
	// Required. The fully qualified name of the [Issue][], in the format
	// `owners/*/repositories/*/issues/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies if to include the [Issue]'s Comments
	Comments bool `protobuf:"varint,2,opt,name=comments,proto3" json:"comments,omitempty"`
	// Specifies if to inlcude the [Issue]'s Reviews
	Reviews bool `protobuf:"varint,3,opt,name=reviews,proto3" json:"reviews,omitempty"`
	// If the FieldMask is NOT set or empty, all fields are returned. If the
	// FieldMask is set, only the specified fields are returned. See
	// https://pkg.go.dev/google.golang.org/genproto/protobuf/field_mask.
	FieldMask            *field_mask.FieldMask `protobuf:"bytes,4,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetIssueRequest) Reset()         { *m = GetIssueRequest{} }
func (m *GetIssueRequest) String() string { return proto.CompactTextString(m) }
func (*GetIssueRequest) ProtoMessage()    {}
func (*GetIssueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{2}
}

func (m *GetIssueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIssueRequest.Unmarshal(m, b)
}
func (m *GetIssueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIssueRequest.Marshal(b, m, deterministic)
}
func (m *GetIssueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIssueRequest.Merge(m, src)
}
func (m *GetIssueRequest) XXX_Size() int {
	return xxx_messageInfo_GetIssueRequest.Size(m)
}
func (m *GetIssueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIssueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIssueRequest proto.InternalMessageInfo

func (m *GetIssueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetIssueRequest) GetComments() bool {
	if m != nil {
		return m.Comments
	}
	return false
}

func (m *GetIssueRequest) GetReviews() bool {
	if m != nil {
		return m.Reviews
	}
	return false
}

func (m *GetIssueRequest) GetFieldMask() *field_mask.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

type GetIssueResponse struct {
	Issue                *Issue   `protobuf:"bytes,1,opt,name=issue,proto3" json:"issue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIssueResponse) Reset()         { *m = GetIssueResponse{} }
func (m *GetIssueResponse) String() string { return proto.CompactTextString(m) }
func (*GetIssueResponse) ProtoMessage()    {}
func (*GetIssueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{3}
}

func (m *GetIssueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIssueResponse.Unmarshal(m, b)
}
func (m *GetIssueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIssueResponse.Marshal(b, m, deterministic)
}
func (m *GetIssueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIssueResponse.Merge(m, src)
}
func (m *GetIssueResponse) XXX_Size() int {
	return xxx_messageInfo_GetIssueResponse.Size(m)
}
func (m *GetIssueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIssueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetIssueResponse proto.InternalMessageInfo

func (m *GetIssueResponse) GetIssue() *Issue {
	if m != nil {
		return m.Issue
	}
	return nil
}

func init() {
	proto.RegisterType((*ListIssuesRequest)(nil), "drghs.v1.ListIssuesRequest")
	proto.RegisterType((*ListIssuesResponse)(nil), "drghs.v1.ListIssuesResponse")
	proto.RegisterType((*GetIssueRequest)(nil), "drghs.v1.GetIssueRequest")
	proto.RegisterType((*GetIssueResponse)(nil), "drghs.v1.GetIssueResponse")
}

func init() {
	proto.RegisterFile("issue_service.proto", fileDescriptor_d4667488ad260932)
}

var fileDescriptor_d4667488ad260932 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6e, 0x93, 0x3a, 0xd3, 0x42, 0xda, 0x2d, 0xb4, 0xae, 0xdb, 0x4a, 0xc1, 0x85, 0x36,
	0xea, 0xc1, 0x56, 0xc3, 0xa9, 0x07, 0x0e, 0xe4, 0xc0, 0x8f, 0x04, 0x12, 0x72, 0xcb, 0xd9, 0x72,
	0xe2, 0x49, 0xb0, 0xe2, 0x78, 0xcd, 0xee, 0x3a, 0xa5, 0xad, 0x7a, 0xe1, 0x86, 0xc4, 0x0d, 0x89,
	0x17, 0xe3, 0x15, 0x78, 0x00, 0x9e, 0x00, 0x21, 0xaf, 0xd7, 0xf9, 0x23, 0xad, 0xc4, 0xcd, 0x33,
	0xdf, 0xec, 0x7c, 0xdf, 0x7e, 0x33, 0x6b, 0xd8, 0x8c, 0x38, 0xcf, 0xd0, 0xe7, 0xc8, 0x46, 0x51,
	0x17, 0x9d, 0x94, 0x51, 0x41, 0x89, 0x11, 0xb2, 0xfe, 0x47, 0xee, 0x8c, 0x4e, 0xac, 0xbd, 0x3e,
	0xa5, 0xfd, 0x18, 0xdd, 0x20, 0x8d, 0xdc, 0x20, 0x49, 0xa8, 0x08, 0x44, 0x44, 0x13, 0x5e, 0xd4,
	0x59, 0x0d, 0x85, 0xca, 0xa8, 0x93, 0xf5, 0xdc, 0x5e, 0x84, 0x71, 0xe8, 0x0f, 0x03, 0x3e, 0x50,
	0x15, 0x75, 0x86, 0x9c, 0x66, 0xac, 0x8b, 0xe5, 0x91, 0x6d, 0xc5, 0xe4, 0xcf, 0x03, 0x9b, 0x41,
	0x38, 0x8c, 0x92, 0x59, 0x21, 0xf6, 0x1f, 0x1d, 0x36, 0xde, 0x46, 0x5c, 0xbc, 0xc9, 0x45, 0x72,
	0x0f, 0x3f, 0x65, 0xc8, 0x05, 0xd9, 0x82, 0x6a, 0x1a, 0x30, 0x4c, 0x84, 0xa9, 0x35, 0xb4, 0x66,
	0xcd, 0x53, 0x11, 0xd9, 0x85, 0x5a, 0x1a, 0xf4, 0xd1, 0xe7, 0xd1, 0x15, 0x9a, 0x7a, 0x43, 0x6b,
	0x56, 0x3c, 0x23, 0x4f, 0x9c, 0x45, 0x57, 0x48, 0xf6, 0x01, 0x24, 0x28, 0xe8, 0x00, 0x13, 0x73,
	0x49, 0x1e, 0x94, 0xe5, 0xe7, 0x79, 0x22, 0xef, 0xd9, 0x8b, 0x62, 0x81, 0xcc, 0x5c, 0x2e, 0x7a,
	0x16, 0x11, 0xd9, 0x01, 0x83, 0xb2, 0x10, 0x99, 0xdf, 0xb9, 0x34, 0x2b, 0x12, 0x59, 0x91, 0x71,
	0xfb, 0x92, 0x58, 0x60, 0x74, 0xe9, 0x70, 0x88, 0x89, 0xe0, 0x66, 0xb5, 0xa1, 0x35, 0x0d, 0x6f,
	0x1c, 0x13, 0x13, 0x56, 0x18, 0x8e, 0x22, 0xbc, 0xe0, 0xe6, 0x8a, 0x84, 0xca, 0x90, 0x1c, 0xc1,
	0x5a, 0x9a, 0xc5, 0xb1, 0xcf, 0x8a, 0xcb, 0x98, 0x46, 0x0e, 0xb7, 0x75, 0x53, 0x7b, 0x7d, 0xcf,
	0x5b, 0xcd, 0x91, 0xf2, 0x96, 0x7b, 0x50, 0xed, 0xc6, 0x94, 0x63, 0x68, 0xd6, 0xc6, 0x25, 0x9a,
	0xa7, 0x72, 0xe4, 0x14, 0x60, 0x62, 0xb6, 0x09, 0x0d, 0xad, 0xb9, 0xda, 0xb2, 0x9c, 0x62, 0x1e,
	0x4e, 0x39, 0x0f, 0xe7, 0x65, 0x5e, 0xf2, 0x2e, 0xe0, 0x03, 0xaf, 0xd6, 0x2b, 0x3f, 0xdb, 0xdb,
	0xf0, 0x68, 0x5a, 0x81, 0x9f, 0x64, 0x71, 0x1c, 0x74, 0x62, 0x6c, 0x6f, 0x40, 0xbd, 0xe8, 0x3e,
	0x4e, 0xd9, 0xd7, 0x40, 0xa6, 0xfd, 0xe7, 0x29, 0x4d, 0x38, 0x92, 0x23, 0xa8, 0xca, 0xb5, 0xe1,
	0xa6, 0xd6, 0x58, 0x6a, 0xae, 0xb6, 0xea, 0x4e, 0xb9, 0x30, 0x8e, 0xac, 0xf4, 0x14, 0x4c, 0x0e,
	0xa1, 0x9e, 0xe0, 0x67, 0xe1, 0x4f, 0x39, 0xaf, 0x4b, 0x13, 0xef, 0xe7, 0xe9, 0xf7, 0x63, 0xf7,
	0x1f, 0x42, 0x45, 0x50, 0x11, 0xc4, 0x72, 0x2e, 0x15, 0xaf, 0x08, 0xec, 0x1f, 0x1a, 0xd4, 0x5f,
	0x61, 0x41, 0x5e, 0xba, 0x42, 0x60, 0x39, 0x09, 0x86, 0xa8, 0x26, 0x2f, 0xbf, 0x67, 0x06, 0xa1,
	0xdf, 0x3e, 0x88, 0xa5, 0xd9, 0x41, 0xcc, 0x3a, 0xb8, 0xfc, 0x1f, 0x0e, 0xda, 0xa7, 0xb0, 0x3e,
	0xd1, 0xa5, 0x3c, 0x79, 0x0a, 0x15, 0x79, 0x69, 0xa9, 0x6c, 0x81, 0x25, 0x05, 0xda, 0xfa, 0xad,
	0xc3, 0x9a, 0x4c, 0x9c, 0x15, 0x8b, 0x4e, 0xbe, 0x6a, 0xb0, 0x9e, 0x5b, 0xec, 0x61, 0x4a, 0x79,
	0x24, 0x28, 0x8b, 0x90, 0x93, 0xc7, 0x93, 0xd3, 0xf3, 0x98, 0x32, 0xc2, 0xb2, 0xef, 0x2a, 0x29,
	0x34, 0xd9, 0xce, 0x97, 0x9f, 0xbf, 0xbe, 0xeb, 0x4d, 0x72, 0x28, 0xdf, 0xef, 0xe8, 0xc4, 0xbd,
	0x2e, 0x5e, 0xca, 0x73, 0x7a, 0x91, 0x20, 0xe3, 0xee, 0xf1, 0x8d, 0xcb, 0xa6, 0x69, 0x63, 0x80,
	0xc9, 0xb4, 0xc9, 0xee, 0x2c, 0xc3, 0xcc, 0x1b, 0xb4, 0xf6, 0x16, 0x83, 0x8a, 0xf8, 0x40, 0x12,
	0xef, 0x93, 0xdd, 0x79, 0xe2, 0xe3, 0x9c, 0x53, 0x2d, 0x47, 0x0f, 0x8c, 0xd2, 0x45, 0xb2, 0x33,
	0x69, 0x37, 0x37, 0x71, 0xcb, 0x5a, 0x04, 0xdd, 0xca, 0x93, 0x2f, 0x44, 0xce, 0xa2, 0x48, 0xdc,
	0xe3, 0x9b, 0xd6, 0x37, 0x0d, 0x36, 0xa6, 0x2d, 0x7f, 0x91, 0xff, 0x68, 0xc8, 0x05, 0x90, 0x0f,
	0x69, 0x18, 0x08, 0x3c, 0x67, 0x41, 0x77, 0x80, 0xa1, 0x34, 0x90, 0x1c, 0x4c, 0xc8, 0xfe, 0x45,
	0x4b, 0x45, 0x4f, 0xee, 0x2e, 0x52, 0xda, 0xb6, 0xa4, 0xb6, 0x75, 0xfb, 0x41, 0xa9, 0x2d, 0x93,
	0xb5, 0x9d, 0xaa, 0xdc, 0xad, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc5, 0x57, 0xfb,
	0x7a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IssueServiceClient is the client API for IssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssueServiceClient interface {
	// Lists [Repositories][Repository].
	ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
	// Lists [Issues][Issue].
	ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error)
	// Gets a [Issue][].
	GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error)
}

type issueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueServiceClient(cc grpc.ClientConnInterface) IssueServiceClient {
	return &issueServiceClient{cc}
}

func (c *issueServiceClient) ListRepositories(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/drghs.v1.IssueService/ListRepositories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) ListIssues(ctx context.Context, in *ListIssuesRequest, opts ...grpc.CallOption) (*ListIssuesResponse, error) {
	out := new(ListIssuesResponse)
	err := c.cc.Invoke(ctx, "/drghs.v1.IssueService/ListIssues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issueServiceClient) GetIssue(ctx context.Context, in *GetIssueRequest, opts ...grpc.CallOption) (*GetIssueResponse, error) {
	out := new(GetIssueResponse)
	err := c.cc.Invoke(ctx, "/drghs.v1.IssueService/GetIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueServiceServer is the server API for IssueService service.
type IssueServiceServer interface {
	// Lists [Repositories][Repository].
	ListRepositories(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
	// Lists [Issues][Issue].
	ListIssues(context.Context, *ListIssuesRequest) (*ListIssuesResponse, error)
	// Gets a [Issue][].
	GetIssue(context.Context, *GetIssueRequest) (*GetIssueResponse, error)
}

// UnimplementedIssueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIssueServiceServer struct {
}

func (*UnimplementedIssueServiceServer) ListRepositories(ctx context.Context, req *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepositories not implemented")
}
func (*UnimplementedIssueServiceServer) ListIssues(ctx context.Context, req *ListIssuesRequest) (*ListIssuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIssues not implemented")
}
func (*UnimplementedIssueServiceServer) GetIssue(ctx context.Context, req *GetIssueRequest) (*GetIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIssue not implemented")
}

func RegisterIssueServiceServer(s *grpc.Server, srv IssueServiceServer) {
	s.RegisterService(&_IssueService_serviceDesc, srv)
}

func _IssueService_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drghs.v1.IssueService/ListRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).ListRepositories(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_ListIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIssuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).ListIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drghs.v1.IssueService/ListIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).ListIssues(ctx, req.(*ListIssuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssueService_GetIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).GetIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drghs.v1.IssueService/GetIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).GetIssue(ctx, req.(*GetIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IssueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drghs.v1.IssueService",
	HandlerType: (*IssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepositories",
			Handler:    _IssueService_ListRepositories_Handler,
		},
		{
			MethodName: "ListIssues",
			Handler:    _IssueService_ListIssues_Handler,
		},
		{
			MethodName: "GetIssue",
			Handler:    _IssueService_GetIssue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issue_service.proto",
}

// IssueServiceAdminClient is the client API for IssueServiceAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssueServiceAdminClient interface {
	UpdateTrackedRepos(ctx context.Context, in *UpdateTrackedReposRequest, opts ...grpc.CallOption) (*UpdateTrackedReposResponse, error)
}

type issueServiceAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewIssueServiceAdminClient(cc grpc.ClientConnInterface) IssueServiceAdminClient {
	return &issueServiceAdminClient{cc}
}

func (c *issueServiceAdminClient) UpdateTrackedRepos(ctx context.Context, in *UpdateTrackedReposRequest, opts ...grpc.CallOption) (*UpdateTrackedReposResponse, error) {
	out := new(UpdateTrackedReposResponse)
	err := c.cc.Invoke(ctx, "/drghs.v1.IssueServiceAdmin/UpdateTrackedRepos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssueServiceAdminServer is the server API for IssueServiceAdmin service.
type IssueServiceAdminServer interface {
	UpdateTrackedRepos(context.Context, *UpdateTrackedReposRequest) (*UpdateTrackedReposResponse, error)
}

// UnimplementedIssueServiceAdminServer can be embedded to have forward compatible implementations.
type UnimplementedIssueServiceAdminServer struct {
}

func (*UnimplementedIssueServiceAdminServer) UpdateTrackedRepos(ctx context.Context, req *UpdateTrackedReposRequest) (*UpdateTrackedReposResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrackedRepos not implemented")
}

func RegisterIssueServiceAdminServer(s *grpc.Server, srv IssueServiceAdminServer) {
	s.RegisterService(&_IssueServiceAdmin_serviceDesc, srv)
}

func _IssueServiceAdmin_UpdateTrackedRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrackedReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceAdminServer).UpdateTrackedRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drghs.v1.IssueServiceAdmin/UpdateTrackedRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceAdminServer).UpdateTrackedRepos(ctx, req.(*UpdateTrackedReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IssueServiceAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drghs.v1.IssueServiceAdmin",
	HandlerType: (*IssueServiceAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateTrackedRepos",
			Handler:    _IssueServiceAdmin_UpdateTrackedRepos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issue_service.proto",
}
