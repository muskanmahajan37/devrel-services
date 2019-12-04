// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue_service.proto

package drghs_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
	OrderBy              string   `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	Comments             bool     `protobuf:"varint,6,opt,name=comments,proto3" json:"comments,omitempty"`
	Reviews              bool     `protobuf:"varint,7,opt,name=reviews,proto3" json:"reviews,omitempty"`
	PullRequest          bool     `protobuf:"varint,8,opt,name=pull_request,json=pullRequest,proto3" json:"pull_request,omitempty"` // Deprecated: Do not use.
	Closed               bool     `protobuf:"varint,9,opt,name=closed,proto3" json:"closed,omitempty"`                              // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

// Deprecated: Do not use.
func (m *ListIssuesRequest) GetPullRequest() bool {
	if m != nil {
		return m.PullRequest
	}
	return false
}

// Deprecated: Do not use.
func (m *ListIssuesRequest) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
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
	Reviews              bool     `protobuf:"varint,3,opt,name=reviews,proto3" json:"reviews,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

// Request message for [DevRelGitHubService.ListGitHubComments][].
type ListGitHubCommentsRequest struct {
	// Required. The fully qualified name of the [Issue][], in the format
	// `owners/*/repositories/*/issues/*`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Limit the number of [GitHubComment]s to include in the
	// response. Fewer GitHubComments than requested might be returned.
	//
	// The maximum page size is `100`. If unspecified, the page size will be the
	// maximum. Further [GitHubComment][]s can subsequently be obtained
	// by including the [ListGitHubCommentsResponse.next_page_token][] in a
	// subsequent request.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. To request the first page of results, `page_token` must be empty.
	// To request the next page of results, page_token must be the value of
	// [ListGitHubCommentsResponse.next_page_token][] returned by a previous call to
	// [Issueservice.ListGitHubComments][].
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
	// Valid filter fields are: `comment.body`, `comment.created_at`,
	// `comment.user`, and `comment.updated_at`.
	//
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	// Optional. Specify how the results should be sorted. The fields supported
	// for sorting are `name` and `size`.
	// The default ordering is by `name`. Prefix with `-` to specify
	// descending order, e.g. `-name`.
	OrderBy              string   `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGitHubCommentsRequest) Reset()         { *m = ListGitHubCommentsRequest{} }
func (m *ListGitHubCommentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListGitHubCommentsRequest) ProtoMessage()    {}
func (*ListGitHubCommentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{4}
}

func (m *ListGitHubCommentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGitHubCommentsRequest.Unmarshal(m, b)
}
func (m *ListGitHubCommentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGitHubCommentsRequest.Marshal(b, m, deterministic)
}
func (m *ListGitHubCommentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGitHubCommentsRequest.Merge(m, src)
}
func (m *ListGitHubCommentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListGitHubCommentsRequest.Size(m)
}
func (m *ListGitHubCommentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGitHubCommentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGitHubCommentsRequest proto.InternalMessageInfo

func (m *ListGitHubCommentsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGitHubCommentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListGitHubCommentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListGitHubCommentsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListGitHubCommentsRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

// Response message for [DevRelGitHubService.ListGitHubComments][].
type ListGitHubCommentsResponse struct {
	// The list of [GitHubComments][].
	Comments []*GitHubComment `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	// A token to retrieve the next page of results, or empty if there are no
	// more results in the list. Pass this value in
	// [ListGitHubCommentsRequest.page_token][] to retrieve the next page of
	// results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The total number of [GitHubComments][] that matched the query.
	Total                int32    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGitHubCommentsResponse) Reset()         { *m = ListGitHubCommentsResponse{} }
func (m *ListGitHubCommentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListGitHubCommentsResponse) ProtoMessage()    {}
func (*ListGitHubCommentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4667488ad260932, []int{5}
}

func (m *ListGitHubCommentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGitHubCommentsResponse.Unmarshal(m, b)
}
func (m *ListGitHubCommentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGitHubCommentsResponse.Marshal(b, m, deterministic)
}
func (m *ListGitHubCommentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGitHubCommentsResponse.Merge(m, src)
}
func (m *ListGitHubCommentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListGitHubCommentsResponse.Size(m)
}
func (m *ListGitHubCommentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGitHubCommentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGitHubCommentsResponse proto.InternalMessageInfo

func (m *ListGitHubCommentsResponse) GetComments() []*GitHubComment {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *ListGitHubCommentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func (m *ListGitHubCommentsResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*ListIssuesRequest)(nil), "drghs.v1.ListIssuesRequest")
	proto.RegisterType((*ListIssuesResponse)(nil), "drghs.v1.ListIssuesResponse")
	proto.RegisterType((*GetIssueRequest)(nil), "drghs.v1.GetIssueRequest")
	proto.RegisterType((*GetIssueResponse)(nil), "drghs.v1.GetIssueResponse")
	proto.RegisterType((*ListGitHubCommentsRequest)(nil), "drghs.v1.ListGitHubCommentsRequest")
	proto.RegisterType((*ListGitHubCommentsResponse)(nil), "drghs.v1.ListGitHubCommentsResponse")
}

func init() { proto.RegisterFile("issue_service.proto", fileDescriptor_d4667488ad260932) }

var fileDescriptor_d4667488ad260932 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x4f, 0x14, 0x4f,
	0x10, 0xcd, 0xec, 0xb2, 0xcb, 0x50, 0xf0, 0xcb, 0xf2, 0x6b, 0x8d, 0x34, 0x03, 0x24, 0xeb, 0x20,
	0xb8, 0xe1, 0xb0, 0x13, 0xe0, 0xe4, 0xc1, 0x0b, 0x1e, 0xd0, 0xc4, 0x83, 0x19, 0xbc, 0x79, 0x98,
	0x0c, 0x4b, 0xb1, 0x76, 0x1c, 0xa6, 0xc7, 0xee, 0xde, 0x45, 0x20, 0x5c, 0x8c, 0x17, 0x4f, 0x1e,
	0xb8, 0xfb, 0xa5, 0xfc, 0x0a, 0x7e, 0x07, 0xaf, 0xa6, 0xff, 0x0c, 0x3b, 0x43, 0x96, 0x4d, 0x8c,
	0x89, 0xb7, 0xa9, 0x7a, 0x35, 0xfd, 0xea, 0xbd, 0xaa, 0x6e, 0x78, 0xc0, 0xa4, 0x1c, 0x61, 0x22,
	0x51, 0x8c, 0xd9, 0x00, 0xfb, 0x85, 0xe0, 0x8a, 0x13, 0xff, 0x44, 0x0c, 0xdf, 0xcb, 0xfe, 0x78,
	0x37, 0x58, 0x1f, 0x72, 0x3e, 0xcc, 0x30, 0x4a, 0x0b, 0x16, 0xa5, 0x79, 0xce, 0x55, 0xaa, 0x18,
	0xcf, 0xa5, 0xad, 0x0b, 0x3a, 0x02, 0x25, 0x1f, 0x89, 0x01, 0x96, 0x89, 0x15, 0x77, 0x4e, 0x72,
	0x07, 0x08, 0x6f, 0x1a, 0xf0, 0xff, 0x6b, 0x26, 0xd5, 0x2b, 0xcd, 0x26, 0x63, 0xfc, 0x38, 0x42,
	0xa9, 0xc8, 0x23, 0x68, 0x17, 0xa9, 0xc0, 0x5c, 0x51, 0xaf, 0xeb, 0xf5, 0x16, 0x62, 0x17, 0x91,
	0x35, 0x58, 0x28, 0xd2, 0x21, 0x26, 0x92, 0x5d, 0x22, 0x6d, 0x74, 0xbd, 0x5e, 0x2b, 0xf6, 0x75,
	0xe2, 0x88, 0x5d, 0x22, 0xd9, 0x00, 0x30, 0xa0, 0xe2, 0x1f, 0x30, 0xa7, 0x4d, 0xf3, 0xa3, 0x29,
	0x7f, 0xab, 0x13, 0xfa, 0xcc, 0x53, 0x96, 0x29, 0x14, 0x74, 0xce, 0x9e, 0x69, 0x23, 0xb2, 0x0a,
	0x3e, 0x17, 0x27, 0x28, 0x92, 0xe3, 0x0b, 0xda, 0x32, 0xc8, 0xbc, 0x89, 0x0f, 0x2e, 0x48, 0x00,
	0xfe, 0x80, 0x9f, 0x9d, 0x61, 0xae, 0x24, 0x6d, 0x77, 0xbd, 0x9e, 0x1f, 0xdf, 0xc6, 0x84, 0xc2,
	0xbc, 0xc0, 0x31, 0xc3, 0x73, 0x49, 0xe7, 0x0d, 0x54, 0x86, 0x64, 0x0b, 0x96, 0x8a, 0x51, 0x96,
	0x25, 0xc2, 0x8a, 0xa1, 0xbe, 0x86, 0x0f, 0x1a, 0xd4, 0x8b, 0x17, 0x75, 0xbe, 0xd4, 0x18, 0x40,
	0x7b, 0x90, 0x71, 0x89, 0x27, 0x74, 0xe1, 0xb6, 0xc0, 0x65, 0xc2, 0x2b, 0x20, 0x55, 0x53, 0x64,
	0xc1, 0x73, 0x89, 0xe4, 0x29, 0xb4, 0xcd, 0x50, 0x24, 0xf5, 0xba, 0xcd, 0xde, 0xe2, 0x5e, 0xa7,
	0x5f, 0x8e, 0xa3, 0x6f, 0x2a, 0x63, 0x07, 0x93, 0x6d, 0xe8, 0xe4, 0xf8, 0x49, 0x25, 0x15, 0x3b,
	0x1a, 0x46, 0xd9, 0x7f, 0x3a, 0xfd, 0xe6, 0xd6, 0x92, 0x87, 0xd0, 0x52, 0x5c, 0xa5, 0x99, 0x31,
	0xab, 0x15, 0xdb, 0x20, 0x7c, 0x07, 0x9d, 0x43, 0xb4, 0xdc, 0x65, 0xaf, 0x04, 0xe6, 0xf2, 0xf4,
	0x0c, 0xdd, 0x34, 0xcc, 0x77, 0xcd, 0x9c, 0xc6, 0xfd, 0xe6, 0x34, 0x6b, 0xe6, 0x84, 0xcf, 0x60,
	0x79, 0x72, 0xb8, 0xd3, 0xb5, 0x05, 0x2d, 0xd3, 0xb8, 0x39, 0x7e, 0x8a, 0x2c, 0x8b, 0x86, 0xdf,
	0x3d, 0x58, 0xd5, 0xae, 0x1c, 0x32, 0xf5, 0x72, 0x74, 0xfc, 0xc2, 0x71, 0xcd, 0x6a, 0xf1, 0xdf,
	0xae, 0x4b, 0xf8, 0xcd, 0x83, 0x60, 0x5a, 0x83, 0x4e, 0xe6, 0x7e, 0xc5, 0x30, 0x3b, 0xc0, 0x95,
	0x89, 0xd2, 0xda, 0x3f, 0x15, 0x27, 0xff, 0x6a, 0x94, 0x7b, 0xbf, 0x9a, 0xb0, 0x64, 0x3c, 0x3c,
	0xb2, 0xd7, 0x8f, 0x7c, 0xf5, 0x60, 0x59, 0xb7, 0x18, 0x63, 0xc1, 0x25, 0x53, 0x5c, 0x30, 0x94,
	0xe4, 0xf1, 0xa4, 0x8d, 0xbb, 0x98, 0x73, 0x37, 0x08, 0x67, 0x95, 0x58, 0x7d, 0x61, 0xff, 0xf3,
	0x8f, 0x9f, 0x37, 0x8d, 0x1e, 0xd9, 0x36, 0x8f, 0xc2, 0x78, 0x37, 0xba, 0xb2, 0xb7, 0xf6, 0x39,
	0x3f, 0xcf, 0x51, 0xc8, 0x68, 0xe7, 0x3a, 0x12, 0x55, 0xda, 0x0c, 0x60, 0xb2, 0xe4, 0x64, 0xad,
	0xce, 0x50, 0x7b, 0x0f, 0x82, 0xf5, 0xe9, 0xa0, 0x23, 0xde, 0x34, 0xc4, 0x1b, 0x64, 0xed, 0x2e,
	0xf1, 0x8e, 0xe6, 0x74, 0x77, 0xe2, 0x14, 0xfc, 0x72, 0xf1, 0xc8, 0x6a, 0xc5, 0xf7, 0xfa, 0xa6,
	0x07, 0xc1, 0x34, 0xe8, 0x5e, 0x1e, 0xbd, 0x65, 0x9a, 0xc5, 0x91, 0x44, 0x3b, 0xd7, 0xe4, 0x8b,
	0x67, 0xef, 0x6e, 0x7d, 0x09, 0xc8, 0x66, 0x5d, 0xc1, 0xd4, 0x1d, 0x0e, 0x9e, 0xcc, 0x2e, 0xfa,
	0x83, 0x36, 0x8e, 0xdb, 0xe6, 0x79, 0xdd, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x8b, 0x42,
	0xdc, 0xc7, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

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
	// Lists [Issues][Issue][Comment].
	ListGitHubComments(ctx context.Context, in *ListGitHubCommentsRequest, opts ...grpc.CallOption) (*ListGitHubCommentsResponse, error)
}

type issueServiceClient struct {
	cc *grpc.ClientConn
}

func NewIssueServiceClient(cc *grpc.ClientConn) IssueServiceClient {
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

func (c *issueServiceClient) ListGitHubComments(ctx context.Context, in *ListGitHubCommentsRequest, opts ...grpc.CallOption) (*ListGitHubCommentsResponse, error) {
	out := new(ListGitHubCommentsResponse)
	err := c.cc.Invoke(ctx, "/drghs.v1.IssueService/ListGitHubComments", in, out, opts...)
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
	// Lists [Issues][Issue][Comment].
	ListGitHubComments(context.Context, *ListGitHubCommentsRequest) (*ListGitHubCommentsResponse, error)
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

func _IssueService_ListGitHubComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGitHubCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueServiceServer).ListGitHubComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drghs.v1.IssueService/ListGitHubComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueServiceServer).ListGitHubComments(ctx, req.(*ListGitHubCommentsRequest))
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
		{
			MethodName: "ListGitHubComments",
			Handler:    _IssueService_ListGitHubComments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issue_service.proto",
}
