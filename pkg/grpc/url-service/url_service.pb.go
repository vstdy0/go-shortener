// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/grpc/url_service.proto

package urlService

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ShortenURL
type ShortenURLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *ShortenURLReq) Reset() {
	*x = ShortenURLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLReq) ProtoMessage() {}

func (x *ShortenURLReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLReq.ProtoReflect.Descriptor instead.
func (*ShortenURLReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenURLReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenURLResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ShortenURLResp) Reset() {
	*x = ShortenURLResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResp) ProtoMessage() {}

func (x *ShortenURLResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResp.ProtoReflect.Descriptor instead.
func (*ShortenURLResp) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenURLResp) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// ShortenURLsBatch
type ShortenURLsBatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []*ShortenURLsBatchReq_UrlUnit `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
}

func (x *ShortenURLsBatchReq) Reset() {
	*x = ShortenURLsBatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLsBatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLsBatchReq) ProtoMessage() {}

func (x *ShortenURLsBatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLsBatchReq.ProtoReflect.Descriptor instead.
func (*ShortenURLsBatchReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{2}
}

func (x *ShortenURLsBatchReq) GetRequest() []*ShortenURLsBatchReq_UrlUnit {
	if x != nil {
		return x.Request
	}
	return nil
}

type ShortenURLsBatchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*ShortenURLsBatchResp_UrlUnit `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *ShortenURLsBatchResp) Reset() {
	*x = ShortenURLsBatchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLsBatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLsBatchResp) ProtoMessage() {}

func (x *ShortenURLsBatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLsBatchResp.ProtoReflect.Descriptor instead.
func (*ShortenURLsBatchResp) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{3}
}

func (x *ShortenURLsBatchResp) GetResponse() []*ShortenURLsBatchResp_UrlUnit {
	if x != nil {
		return x.Response
	}
	return nil
}

// GetOriginalURL
type GetOrigURLReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrigURLReq) Reset() {
	*x = GetOrigURLReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrigURLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrigURLReq) ProtoMessage() {}

func (x *GetOrigURLReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrigURLReq.ProtoReflect.Descriptor instead.
func (*GetOrigURLReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrigURLReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetUsersURLs
type GetUsersURLsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*GetUsersURLsResp_UrlUnit `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *GetUsersURLsResp) Reset() {
	*x = GetUsersURLsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersURLsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersURLsResp) ProtoMessage() {}

func (x *GetUsersURLsResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersURLsResp.ProtoReflect.Descriptor instead.
func (*GetUsersURLsResp) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetUsersURLsResp) GetResponse() []*GetUsersURLsResp_UrlUnit {
	if x != nil {
		return x.Response
	}
	return nil
}

// DeleteUserURLs
type DelUserURLsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DelUserURLsReq) Reset() {
	*x = DelUserURLsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelUserURLsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelUserURLsReq) ProtoMessage() {}

func (x *DelUserURLsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelUserURLsReq.ProtoReflect.Descriptor instead.
func (*DelUserURLsReq) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{6}
}

func (x *DelUserURLsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ShortenURLsBatchReq_UrlUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	OriginalUrl   string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *ShortenURLsBatchReq_UrlUnit) Reset() {
	*x = ShortenURLsBatchReq_UrlUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLsBatchReq_UrlUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLsBatchReq_UrlUnit) ProtoMessage() {}

func (x *ShortenURLsBatchReq_UrlUnit) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLsBatchReq_UrlUnit.ProtoReflect.Descriptor instead.
func (*ShortenURLsBatchReq_UrlUnit) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ShortenURLsBatchReq_UrlUnit) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenURLsBatchReq_UrlUnit) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type ShortenURLsBatchResp_UrlUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationId string `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ShortUrl      string `protobuf:"bytes,2,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *ShortenURLsBatchResp_UrlUnit) Reset() {
	*x = ShortenURLsBatchResp_UrlUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLsBatchResp_UrlUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLsBatchResp_UrlUnit) ProtoMessage() {}

func (x *ShortenURLsBatchResp_UrlUnit) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLsBatchResp_UrlUnit.ProtoReflect.Descriptor instead.
func (*ShortenURLsBatchResp_UrlUnit) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ShortenURLsBatchResp_UrlUnit) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ShortenURLsBatchResp_UrlUnit) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetUsersURLsResp_UrlUnit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetUsersURLsResp_UrlUnit) Reset() {
	*x = GetUsersURLsResp_UrlUnit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_url_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersURLsResp_UrlUnit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersURLsResp_UrlUnit) ProtoMessage() {}

func (x *GetUsersURLsResp_UrlUnit) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_url_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersURLsResp_UrlUnit.ProtoReflect.Descriptor instead.
func (*GetUsersURLsResp_UrlUnit) Descriptor() ([]byte, []int) {
	return file_api_grpc_url_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetUsersURLsResp_UrlUnit) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *GetUsersURLsResp_UrlUnit) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

var File_api_grpc_url_service_proto protoreflect.FileDescriptor

var file_api_grpc_url_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x72, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x72,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x88, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55,
	0x52, 0x4c, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x41, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x75,
	0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x2e, 0x55, 0x72,
	0x6c, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x5d,
	0x0a, 0x07, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2b, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01,
	0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0xab, 0x01,
	0x0a, 0x14, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c,
	0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x55, 0x72, 0x6c, 0x55, 0x6e,
	0x69, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x4d, 0x0a, 0x07,
	0x55, 0x72, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x1f, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9f, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x2e, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x49, 0x0a, 0x07, 0x55, 0x72, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x22,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x32, 0xf2, 0x03, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x12,
	0x19, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x72, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b,
	0x2f, 0x67, 0x77, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x73,
	0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1f, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x73, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x67, 0x77, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x19, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x67, 0x77, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x67, 0x77, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x1a, 0x2e, 0x75, 0x72, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52,
	0x4c, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x0d, 0x2f, 0x67, 0x77, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x72, 0x6c, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x38, 0x5a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b,
	0x75, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x92, 0x41, 0x14, 0x12, 0x12, 0x0a,
	0x0b, 0x55, 0x72, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_url_service_proto_rawDescOnce sync.Once
	file_api_grpc_url_service_proto_rawDescData = file_api_grpc_url_service_proto_rawDesc
)

func file_api_grpc_url_service_proto_rawDescGZIP() []byte {
	file_api_grpc_url_service_proto_rawDescOnce.Do(func() {
		file_api_grpc_url_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_url_service_proto_rawDescData)
	})
	return file_api_grpc_url_service_proto_rawDescData
}

var file_api_grpc_url_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_grpc_url_service_proto_goTypes = []interface{}{
	(*ShortenURLReq)(nil),                // 0: urlService.ShortenURLReq
	(*ShortenURLResp)(nil),               // 1: urlService.ShortenURLResp
	(*ShortenURLsBatchReq)(nil),          // 2: urlService.ShortenURLsBatchReq
	(*ShortenURLsBatchResp)(nil),         // 3: urlService.ShortenURLsBatchResp
	(*GetOrigURLReq)(nil),                // 4: urlService.GetOrigURLReq
	(*GetUsersURLsResp)(nil),             // 5: urlService.GetUsersURLsResp
	(*DelUserURLsReq)(nil),               // 6: urlService.DelUserURLsReq
	(*ShortenURLsBatchReq_UrlUnit)(nil),  // 7: urlService.ShortenURLsBatchReq.UrlUnit
	(*ShortenURLsBatchResp_UrlUnit)(nil), // 8: urlService.ShortenURLsBatchResp.UrlUnit
	(*GetUsersURLsResp_UrlUnit)(nil),     // 9: urlService.GetUsersURLsResp.UrlUnit
	(*emptypb.Empty)(nil),                // 10: google.protobuf.Empty
}
var file_api_grpc_url_service_proto_depIdxs = []int32{
	7,  // 0: urlService.ShortenURLsBatchReq.request:type_name -> urlService.ShortenURLsBatchReq.UrlUnit
	8,  // 1: urlService.ShortenURLsBatchResp.response:type_name -> urlService.ShortenURLsBatchResp.UrlUnit
	9,  // 2: urlService.GetUsersURLsResp.response:type_name -> urlService.GetUsersURLsResp.UrlUnit
	0,  // 3: urlService.URLService.ShortenURL:input_type -> urlService.ShortenURLReq
	2,  // 4: urlService.URLService.ShortenURLsBatch:input_type -> urlService.ShortenURLsBatchReq
	4,  // 5: urlService.URLService.GetOriginalURL:input_type -> urlService.GetOrigURLReq
	10, // 6: urlService.URLService.GetUsersURLs:input_type -> google.protobuf.Empty
	6,  // 7: urlService.URLService.DeleteUserURLs:input_type -> urlService.DelUserURLsReq
	1,  // 8: urlService.URLService.ShortenURL:output_type -> urlService.ShortenURLResp
	3,  // 9: urlService.URLService.ShortenURLsBatch:output_type -> urlService.ShortenURLsBatchResp
	10, // 10: urlService.URLService.GetOriginalURL:output_type -> google.protobuf.Empty
	5,  // 11: urlService.URLService.GetUsersURLs:output_type -> urlService.GetUsersURLsResp
	10, // 12: urlService.URLService.DeleteUserURLs:output_type -> google.protobuf.Empty
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_grpc_url_service_proto_init() }
func file_api_grpc_url_service_proto_init() {
	if File_api_grpc_url_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_url_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLReq); i {
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
		file_api_grpc_url_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLResp); i {
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
		file_api_grpc_url_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLsBatchReq); i {
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
		file_api_grpc_url_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLsBatchResp); i {
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
		file_api_grpc_url_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrigURLReq); i {
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
		file_api_grpc_url_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersURLsResp); i {
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
		file_api_grpc_url_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelUserURLsReq); i {
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
		file_api_grpc_url_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLsBatchReq_UrlUnit); i {
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
		file_api_grpc_url_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLsBatchResp_UrlUnit); i {
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
		file_api_grpc_url_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersURLsResp_UrlUnit); i {
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
			RawDescriptor: file_api_grpc_url_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_url_service_proto_goTypes,
		DependencyIndexes: file_api_grpc_url_service_proto_depIdxs,
		MessageInfos:      file_api_grpc_url_service_proto_msgTypes,
	}.Build()
	File_api_grpc_url_service_proto = out.File
	file_api_grpc_url_service_proto_rawDesc = nil
	file_api_grpc_url_service_proto_goTypes = nil
	file_api_grpc_url_service_proto_depIdxs = nil
}