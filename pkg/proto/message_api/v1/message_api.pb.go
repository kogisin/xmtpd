// Message API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: message_api/v1/message_api.proto

package message_apiv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sort direction
type SortDirection int32

const (
	SortDirection_SORT_DIRECTION_UNSPECIFIED SortDirection = 0
	SortDirection_SORT_DIRECTION_ASCENDING   SortDirection = 1
	SortDirection_SORT_DIRECTION_DESCENDING  SortDirection = 2
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "SORT_DIRECTION_UNSPECIFIED",
		1: "SORT_DIRECTION_ASCENDING",
		2: "SORT_DIRECTION_DESCENDING",
	}
	SortDirection_value = map[string]int32{
		"SORT_DIRECTION_UNSPECIFIED": 0,
		"SORT_DIRECTION_ASCENDING":   1,
		"SORT_DIRECTION_DESCENDING":  2,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_message_api_v1_message_api_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_message_api_v1_message_api_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{0}
}

// This is based off of the go-waku Index type, but with the
// receiverTime and pubsubTopic removed for simplicity.
// Both removed fields are optional
type IndexCursor struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Digest        []byte                 `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
	SenderTimeNs  uint64                 `protobuf:"varint,2,opt,name=sender_time_ns,json=senderTimeNs,proto3" json:"sender_time_ns,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexCursor) Reset() {
	*x = IndexCursor{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexCursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexCursor) ProtoMessage() {}

func (x *IndexCursor) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexCursor.ProtoReflect.Descriptor instead.
func (*IndexCursor) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{0}
}

func (x *IndexCursor) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *IndexCursor) GetSenderTimeNs() uint64 {
	if x != nil {
		return x.SenderTimeNs
	}
	return 0
}

// Wrapper for potentially multiple types of cursor
type Cursor struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Making the cursor a one-of type, as I would like to change the way we
	// handle pagination to use a precomputed sort field.
	// This way we can handle both methods
	//
	// Types that are valid to be assigned to Cursor:
	//
	//	*Cursor_Index
	Cursor        isCursor_Cursor `protobuf_oneof:"cursor"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{1}
}

func (x *Cursor) GetCursor() isCursor_Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *Cursor) GetIndex() *IndexCursor {
	if x != nil {
		if x, ok := x.Cursor.(*Cursor_Index); ok {
			return x.Index
		}
	}
	return nil
}

type isCursor_Cursor interface {
	isCursor_Cursor()
}

type Cursor_Index struct {
	Index *IndexCursor `protobuf:"bytes,1,opt,name=index,proto3,oneof"`
}

func (*Cursor_Index) isCursor_Cursor() {}

// This is based off of the go-waku PagingInfo struct, but with the direction
// changed to our SortDirection enum format
type PagingInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Note: this is a uint32, while go-waku's pageSize is a uint64
	Limit         uint32        `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor        *Cursor       `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Direction     SortDirection `protobuf:"varint,3,opt,name=direction,proto3,enum=xmtp.message_api.v1.SortDirection" json:"direction,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PagingInfo) Reset() {
	*x = PagingInfo{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PagingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagingInfo) ProtoMessage() {}

func (x *PagingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagingInfo.ProtoReflect.Descriptor instead.
func (*PagingInfo) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{2}
}

func (x *PagingInfo) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PagingInfo) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *PagingInfo) GetDirection() SortDirection {
	if x != nil {
		return x.Direction
	}
	return SortDirection_SORT_DIRECTION_UNSPECIFIED
}

// Envelope encapsulates a message while in transit.
type Envelope struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The topic the message belongs to,
	// If the message includes the topic as well
	// it MUST be the same as the topic in the envelope.
	ContentTopic string `protobuf:"bytes,1,opt,name=content_topic,json=contentTopic,proto3" json:"content_topic,omitempty"`
	// Message creation timestamp
	// If the message includes the timestamp as well
	// it MUST be equivalent to the timestamp in the envelope.
	TimestampNs   uint64 `protobuf:"varint,2,opt,name=timestamp_ns,json=timestampNs,proto3" json:"timestamp_ns,omitempty"`
	Message       []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{3}
}

func (x *Envelope) GetContentTopic() string {
	if x != nil {
		return x.ContentTopic
	}
	return ""
}

func (x *Envelope) GetTimestampNs() uint64 {
	if x != nil {
		return x.TimestampNs
	}
	return 0
}

func (x *Envelope) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

// Publish
type PublishRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Envelopes     []*Envelope            `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{4}
}

func (x *PublishRequest) GetEnvelopes() []*Envelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

// Empty message as a response for Publish
type PublishResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{5}
}

// Subscribe
type SubscribeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContentTopics []string               `protobuf:"bytes,1,rep,name=content_topics,json=contentTopics,proto3" json:"content_topics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeRequest) GetContentTopics() []string {
	if x != nil {
		return x.ContentTopics
	}
	return nil
}

// SubscribeAll
type SubscribeAllRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubscribeAllRequest) Reset() {
	*x = SubscribeAllRequest{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeAllRequest) ProtoMessage() {}

func (x *SubscribeAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeAllRequest.ProtoReflect.Descriptor instead.
func (*SubscribeAllRequest) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{7}
}

// Query
type QueryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContentTopics []string               `protobuf:"bytes,1,rep,name=content_topics,json=contentTopics,proto3" json:"content_topics,omitempty"`
	StartTimeNs   uint64                 `protobuf:"varint,2,opt,name=start_time_ns,json=startTimeNs,proto3" json:"start_time_ns,omitempty"`
	EndTimeNs     uint64                 `protobuf:"varint,3,opt,name=end_time_ns,json=endTimeNs,proto3" json:"end_time_ns,omitempty"`
	PagingInfo    *PagingInfo            `protobuf:"bytes,4,opt,name=paging_info,json=pagingInfo,proto3" json:"paging_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{8}
}

func (x *QueryRequest) GetContentTopics() []string {
	if x != nil {
		return x.ContentTopics
	}
	return nil
}

func (x *QueryRequest) GetStartTimeNs() uint64 {
	if x != nil {
		return x.StartTimeNs
	}
	return 0
}

func (x *QueryRequest) GetEndTimeNs() uint64 {
	if x != nil {
		return x.EndTimeNs
	}
	return 0
}

func (x *QueryRequest) GetPagingInfo() *PagingInfo {
	if x != nil {
		return x.PagingInfo
	}
	return nil
}

// The response, containing envelopes, for a query
type QueryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Envelopes     []*Envelope            `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
	PagingInfo    *PagingInfo            `protobuf:"bytes,2,opt,name=paging_info,json=pagingInfo,proto3" json:"paging_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{9}
}

func (x *QueryResponse) GetEnvelopes() []*Envelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

func (x *QueryResponse) GetPagingInfo() *PagingInfo {
	if x != nil {
		return x.PagingInfo
	}
	return nil
}

// BatchQuery
type BatchQueryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Requests      []*QueryRequest        `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchQueryRequest) Reset() {
	*x = BatchQueryRequest{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryRequest) ProtoMessage() {}

func (x *BatchQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryRequest.ProtoReflect.Descriptor instead.
func (*BatchQueryRequest) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{10}
}

func (x *BatchQueryRequest) GetRequests() []*QueryRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

// Response containing a list of QueryResponse messages
type BatchQueryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Responses     []*QueryResponse       `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchQueryResponse) Reset() {
	*x = BatchQueryResponse{}
	mi := &file_message_api_v1_message_api_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchQueryResponse) ProtoMessage() {}

func (x *BatchQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_api_v1_message_api_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchQueryResponse.ProtoReflect.Descriptor instead.
func (*BatchQueryResponse) Descriptor() ([]byte, []int) {
	return file_message_api_v1_message_api_proto_rawDescGZIP(), []int{11}
}

func (x *BatchQueryResponse) GetResponses() []*QueryResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

var File_message_api_v1_message_api_proto protoreflect.FileDescriptor

var file_message_api_v1_message_api_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x4e, 0x73, 0x22, 0x4c, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x48, 0x00, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x22, 0x99, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x09, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x08,
	0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d, 0x0a, 0x0e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x09,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x09,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xbb,
	0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8e, 0x01, 0x0a,
	0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x52, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x52, 0x0a,
	0x11, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x22, 0x56, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2a, 0x6c, 0x0a, 0x0d, 0x53, 0x6f, 0x72,
	0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x4f, 0x52, 0x54,
	0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x32, 0xc6, 0x05, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x12, 0x74, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x23, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x75, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x25, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x32, 0x12, 0x25, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x7f, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x28, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2d, 0x61, 0x6c, 0x6c, 0x30, 0x01, 0x12, 0x6c,
	0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x81, 0x01, 0x0a,
	0x0a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x26, 0x2e, 0x78, 0x6d,
	0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x42, 0xe8, 0x01, 0x92, 0x41, 0x13, 0x12, 0x11, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x70, 0x69, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x78,
	0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x42, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x2f, 0x78, 0x6d, 0x74, 0x70, 0x64, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x58, 0x4d, 0x58, 0xaa, 0x02, 0x12, 0x58, 0x6d, 0x74, 0x70,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x12, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x58, 0x6d, 0x74, 0x70, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x58, 0x6d, 0x74, 0x70, 0x3a, 0x3a, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_message_api_v1_message_api_proto_rawDescOnce sync.Once
	file_message_api_v1_message_api_proto_rawDescData = file_message_api_v1_message_api_proto_rawDesc
)

func file_message_api_v1_message_api_proto_rawDescGZIP() []byte {
	file_message_api_v1_message_api_proto_rawDescOnce.Do(func() {
		file_message_api_v1_message_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_api_v1_message_api_proto_rawDescData)
	})
	return file_message_api_v1_message_api_proto_rawDescData
}

var file_message_api_v1_message_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_api_v1_message_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_message_api_v1_message_api_proto_goTypes = []any{
	(SortDirection)(0),          // 0: xmtp.message_api.v1.SortDirection
	(*IndexCursor)(nil),         // 1: xmtp.message_api.v1.IndexCursor
	(*Cursor)(nil),              // 2: xmtp.message_api.v1.Cursor
	(*PagingInfo)(nil),          // 3: xmtp.message_api.v1.PagingInfo
	(*Envelope)(nil),            // 4: xmtp.message_api.v1.Envelope
	(*PublishRequest)(nil),      // 5: xmtp.message_api.v1.PublishRequest
	(*PublishResponse)(nil),     // 6: xmtp.message_api.v1.PublishResponse
	(*SubscribeRequest)(nil),    // 7: xmtp.message_api.v1.SubscribeRequest
	(*SubscribeAllRequest)(nil), // 8: xmtp.message_api.v1.SubscribeAllRequest
	(*QueryRequest)(nil),        // 9: xmtp.message_api.v1.QueryRequest
	(*QueryResponse)(nil),       // 10: xmtp.message_api.v1.QueryResponse
	(*BatchQueryRequest)(nil),   // 11: xmtp.message_api.v1.BatchQueryRequest
	(*BatchQueryResponse)(nil),  // 12: xmtp.message_api.v1.BatchQueryResponse
}
var file_message_api_v1_message_api_proto_depIdxs = []int32{
	1,  // 0: xmtp.message_api.v1.Cursor.index:type_name -> xmtp.message_api.v1.IndexCursor
	2,  // 1: xmtp.message_api.v1.PagingInfo.cursor:type_name -> xmtp.message_api.v1.Cursor
	0,  // 2: xmtp.message_api.v1.PagingInfo.direction:type_name -> xmtp.message_api.v1.SortDirection
	4,  // 3: xmtp.message_api.v1.PublishRequest.envelopes:type_name -> xmtp.message_api.v1.Envelope
	3,  // 4: xmtp.message_api.v1.QueryRequest.paging_info:type_name -> xmtp.message_api.v1.PagingInfo
	4,  // 5: xmtp.message_api.v1.QueryResponse.envelopes:type_name -> xmtp.message_api.v1.Envelope
	3,  // 6: xmtp.message_api.v1.QueryResponse.paging_info:type_name -> xmtp.message_api.v1.PagingInfo
	9,  // 7: xmtp.message_api.v1.BatchQueryRequest.requests:type_name -> xmtp.message_api.v1.QueryRequest
	10, // 8: xmtp.message_api.v1.BatchQueryResponse.responses:type_name -> xmtp.message_api.v1.QueryResponse
	5,  // 9: xmtp.message_api.v1.MessageApi.Publish:input_type -> xmtp.message_api.v1.PublishRequest
	7,  // 10: xmtp.message_api.v1.MessageApi.Subscribe:input_type -> xmtp.message_api.v1.SubscribeRequest
	7,  // 11: xmtp.message_api.v1.MessageApi.Subscribe2:input_type -> xmtp.message_api.v1.SubscribeRequest
	8,  // 12: xmtp.message_api.v1.MessageApi.SubscribeAll:input_type -> xmtp.message_api.v1.SubscribeAllRequest
	9,  // 13: xmtp.message_api.v1.MessageApi.Query:input_type -> xmtp.message_api.v1.QueryRequest
	11, // 14: xmtp.message_api.v1.MessageApi.BatchQuery:input_type -> xmtp.message_api.v1.BatchQueryRequest
	6,  // 15: xmtp.message_api.v1.MessageApi.Publish:output_type -> xmtp.message_api.v1.PublishResponse
	4,  // 16: xmtp.message_api.v1.MessageApi.Subscribe:output_type -> xmtp.message_api.v1.Envelope
	4,  // 17: xmtp.message_api.v1.MessageApi.Subscribe2:output_type -> xmtp.message_api.v1.Envelope
	4,  // 18: xmtp.message_api.v1.MessageApi.SubscribeAll:output_type -> xmtp.message_api.v1.Envelope
	10, // 19: xmtp.message_api.v1.MessageApi.Query:output_type -> xmtp.message_api.v1.QueryResponse
	12, // 20: xmtp.message_api.v1.MessageApi.BatchQuery:output_type -> xmtp.message_api.v1.BatchQueryResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_message_api_v1_message_api_proto_init() }
func file_message_api_v1_message_api_proto_init() {
	if File_message_api_v1_message_api_proto != nil {
		return
	}
	file_message_api_v1_message_api_proto_msgTypes[1].OneofWrappers = []any{
		(*Cursor_Index)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_api_v1_message_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_api_v1_message_api_proto_goTypes,
		DependencyIndexes: file_message_api_v1_message_api_proto_depIdxs,
		EnumInfos:         file_message_api_v1_message_api_proto_enumTypes,
		MessageInfos:      file_message_api_v1_message_api_proto_msgTypes,
	}.Build()
	File_message_api_v1_message_api_proto = out.File
	file_message_api_v1_message_api_proto_rawDesc = nil
	file_message_api_v1_message_api_proto_goTypes = nil
	file_message_api_v1_message_api_proto_depIdxs = nil
}
