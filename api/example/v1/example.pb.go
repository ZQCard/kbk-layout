// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.0--rc1
// source: example/v1/example.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求 - 更新
type UpdateExampleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主键id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 状态
	Status bool `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateExampleReq) Reset() {
	*x = UpdateExampleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExampleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExampleReq) ProtoMessage() {}

func (x *UpdateExampleReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExampleReq.ProtoReflect.Descriptor instead.
func (*UpdateExampleReq) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateExampleReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateExampleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateExampleReq) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// 请求 - 创建
type CreateExampleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 名称
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status bool   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateExampleReq) Reset() {
	*x = CreateExampleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExampleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExampleReq) ProtoMessage() {}

func (x *CreateExampleReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExampleReq.ProtoReflect.Descriptor instead.
func (*CreateExampleReq) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExampleReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateExampleReq) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

// 响应 - 示例信息
type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主键id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 状态
	Status bool `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{2}
}

func (x *Example) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Example) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Example) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Example) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Example) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Example) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

// 响应 - 分页
type GetExampleListPageRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// 列表
	List []*Example `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GetExampleListPageRes) Reset() {
	*x = GetExampleListPageRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleListPageRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleListPageRes) ProtoMessage() {}

func (x *GetExampleListPageRes) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleListPageRes.ProtoReflect.Descriptor instead.
func (*GetExampleListPageRes) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{3}
}

func (x *GetExampleListPageRes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetExampleListPageRes) GetList() []*Example {
	if x != nil {
		return x.List
	}
	return nil
}

// 请求 - 分页列表
type GetExampleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主键id
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// 页记录数
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 状态 true正常 false冻结
	Status *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// 是否删除
	IsDeleted *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	// 创建开始时间
	CreatedAtStart string `protobuf:"bytes,6,opt,name=created_at_start,json=createdAtStart,proto3" json:"created_at_start,omitempty"`
	// 创建结束时间
	CreatedAtEnd string `protobuf:"bytes,7,opt,name=created_at_end,json=createdAtEnd,proto3" json:"created_at_end,omitempty"`
}

func (x *GetExampleListReq) Reset() {
	*x = GetExampleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleListReq) ProtoMessage() {}

func (x *GetExampleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleListReq.ProtoReflect.Descriptor instead.
func (*GetExampleListReq) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{4}
}

func (x *GetExampleListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetExampleListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetExampleListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetExampleListReq) GetStatus() *wrapperspb.BoolValue {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *GetExampleListReq) GetIsDeleted() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsDeleted
	}
	return nil
}

func (x *GetExampleListReq) GetCreatedAtStart() string {
	if x != nil {
		return x.CreatedAtStart
	}
	return ""
}

func (x *GetExampleListReq) GetCreatedAtEnd() string {
	if x != nil {
		return x.CreatedAtEnd
	}
	return ""
}

// 请求 - 主键id
type ExampleIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 主键id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExampleIdReq) Reset() {
	*x = ExampleIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_v1_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleIdReq) ProtoMessage() {}

func (x *ExampleIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_v1_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleIdReq.ProtoReflect.Descriptor instead.
func (*ExampleIdReq) Descriptor() ([]byte, []int) {
	return file_example_v1_example_proto_rawDescGZIP(), []int{5}
}

func (x *ExampleIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_example_v1_example_proto protoreflect.FileDescriptor

var file_example_v1_example_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x3c, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x3c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x56, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0xa8, 0x02, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x45, 0x6e, 0x64, 0x22, 0x1e,
	0x0a, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xad,
	0x04, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x57, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x5a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x1a, 0x08, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x53, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x3a, 0x01, 0x2a, 0x32, 0x08, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x62,
	0xba, 0x47, 0x32, 0x12, 0x30, 0x0a, 0x0c, 0xe7, 0xa4, 0xba, 0xe4, 0xbe, 0x8b, 0xe6, 0x9c, 0x8d,
	0xe5, 0x8a, 0xa1, 0x12, 0x0c, 0xe7, 0xa4, 0xba, 0xe4, 0xbe, 0x8b, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a,
	0xa1, 0x22, 0x0b, 0x0a, 0x09, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x05,
	0x30, 0x2e, 0x30, 0x2e, 0x31, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x5a, 0x51, 0x43, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x62, 0x6b, 0x2d, 0x6c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_v1_example_proto_rawDescOnce sync.Once
	file_example_v1_example_proto_rawDescData = file_example_v1_example_proto_rawDesc
)

func file_example_v1_example_proto_rawDescGZIP() []byte {
	file_example_v1_example_proto_rawDescOnce.Do(func() {
		file_example_v1_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_v1_example_proto_rawDescData)
	})
	return file_example_v1_example_proto_rawDescData
}

var file_example_v1_example_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_example_v1_example_proto_goTypes = []interface{}{
	(*UpdateExampleReq)(nil),      // 0: example.v1.UpdateExampleReq
	(*CreateExampleReq)(nil),      // 1: example.v1.CreateExampleReq
	(*Example)(nil),               // 2: example.v1.Example
	(*GetExampleListPageRes)(nil), // 3: example.v1.GetExampleListPageRes
	(*GetExampleListReq)(nil),     // 4: example.v1.GetExampleListReq
	(*ExampleIdReq)(nil),          // 5: example.v1.ExampleIdReq
	(*wrapperspb.BoolValue)(nil),  // 6: google.protobuf.BoolValue
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_example_v1_example_proto_depIdxs = []int32{
	2, // 0: example.v1.GetExampleListPageRes.list:type_name -> example.v1.Example
	6, // 1: example.v1.GetExampleListReq.status:type_name -> google.protobuf.BoolValue
	6, // 2: example.v1.GetExampleListReq.is_deleted:type_name -> google.protobuf.BoolValue
	4, // 3: example.v1.ExampleService.GetExampleList:input_type -> example.v1.GetExampleListReq
	5, // 4: example.v1.ExampleService.GetExample:input_type -> example.v1.ExampleIdReq
	1, // 5: example.v1.ExampleService.CreateExample:input_type -> example.v1.CreateExampleReq
	0, // 6: example.v1.ExampleService.UpdateExample:input_type -> example.v1.UpdateExampleReq
	5, // 7: example.v1.ExampleService.DeleteExample:input_type -> example.v1.ExampleIdReq
	5, // 8: example.v1.ExampleService.RecoverExample:input_type -> example.v1.ExampleIdReq
	3, // 9: example.v1.ExampleService.GetExampleList:output_type -> example.v1.GetExampleListPageRes
	2, // 10: example.v1.ExampleService.GetExample:output_type -> example.v1.Example
	2, // 11: example.v1.ExampleService.CreateExample:output_type -> example.v1.Example
	7, // 12: example.v1.ExampleService.UpdateExample:output_type -> google.protobuf.Empty
	7, // 13: example.v1.ExampleService.DeleteExample:output_type -> google.protobuf.Empty
	7, // 14: example.v1.ExampleService.RecoverExample:output_type -> google.protobuf.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_example_v1_example_proto_init() }
func file_example_v1_example_proto_init() {
	if File_example_v1_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_v1_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateExampleReq); i {
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
		file_example_v1_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExampleReq); i {
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
		file_example_v1_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_example_v1_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExampleListPageRes); i {
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
		file_example_v1_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExampleListReq); i {
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
		file_example_v1_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleIdReq); i {
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
			RawDescriptor: file_example_v1_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_v1_example_proto_goTypes,
		DependencyIndexes: file_example_v1_example_proto_depIdxs,
		MessageInfos:      file_example_v1_example_proto_msgTypes,
	}.Build()
	File_example_v1_example_proto = out.File
	file_example_v1_example_proto_rawDesc = nil
	file_example_v1_example_proto_goTypes = nil
	file_example_v1_example_proto_depIdxs = nil
}
