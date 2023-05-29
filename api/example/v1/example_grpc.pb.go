// Code generated by entproto. DO NOT EDIT.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc1
// source: example/v1/example.proto

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExampleService_GetExampleList_FullMethodName = "/example.v1.ExampleService/GetExampleList"
	ExampleService_GetExample_FullMethodName     = "/example.v1.ExampleService/GetExample"
	ExampleService_CreateExample_FullMethodName  = "/example.v1.ExampleService/CreateExample"
	ExampleService_UpdateExample_FullMethodName  = "/example.v1.ExampleService/UpdateExample"
	ExampleService_DeleteExample_FullMethodName  = "/example.v1.ExampleService/DeleteExample"
	ExampleService_RecoverExample_FullMethodName = "/example.v1.ExampleService/RecoverExample"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// 列表
	GetExampleList(ctx context.Context, in *GetExampleListReq, opts ...grpc.CallOption) (*GetExampleListPageRes, error)
	// 详情
	GetExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*Example, error)
	// 创建
	CreateExample(ctx context.Context, in *CreateExampleReq, opts ...grpc.CallOption) (*Example, error)
	// 更新
	UpdateExample(ctx context.Context, in *UpdateExampleReq, opts ...grpc.CallOption) (*CheckResponse, error)
	// 删除
	DeleteExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*CheckResponse, error)
	// 恢复
	RecoverExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*CheckResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) GetExampleList(ctx context.Context, in *GetExampleListReq, opts ...grpc.CallOption) (*GetExampleListPageRes, error) {
	out := new(GetExampleListPageRes)
	err := c.cc.Invoke(ctx, ExampleService_GetExampleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, ExampleService_GetExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) CreateExample(ctx context.Context, in *CreateExampleReq, opts ...grpc.CallOption) (*Example, error) {
	out := new(Example)
	err := c.cc.Invoke(ctx, ExampleService_CreateExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) UpdateExample(ctx context.Context, in *UpdateExampleReq, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, ExampleService_UpdateExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) DeleteExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, ExampleService_DeleteExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) RecoverExample(ctx context.Context, in *ExampleIdReq, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, ExampleService_RecoverExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	// 列表
	GetExampleList(context.Context, *GetExampleListReq) (*GetExampleListPageRes, error)
	// 详情
	GetExample(context.Context, *ExampleIdReq) (*Example, error)
	// 创建
	CreateExample(context.Context, *CreateExampleReq) (*Example, error)
	// 更新
	UpdateExample(context.Context, *UpdateExampleReq) (*CheckResponse, error)
	// 删除
	DeleteExample(context.Context, *ExampleIdReq) (*CheckResponse, error)
	// 恢复
	RecoverExample(context.Context, *ExampleIdReq) (*CheckResponse, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) GetExampleList(context.Context, *GetExampleListReq) (*GetExampleListPageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleList not implemented")
}
func (UnimplementedExampleServiceServer) GetExample(context.Context, *ExampleIdReq) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExample not implemented")
}
func (UnimplementedExampleServiceServer) CreateExample(context.Context, *CreateExampleReq) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExample not implemented")
}
func (UnimplementedExampleServiceServer) UpdateExample(context.Context, *UpdateExampleReq) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExample not implemented")
}
func (UnimplementedExampleServiceServer) DeleteExample(context.Context, *ExampleIdReq) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExample not implemented")
}
func (UnimplementedExampleServiceServer) RecoverExample(context.Context, *ExampleIdReq) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoverExample not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_GetExampleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExampleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetExampleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetExampleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetExampleList(ctx, req.(*GetExampleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_GetExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetExample(ctx, req.(*ExampleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_CreateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExampleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).CreateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_CreateExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).CreateExample(ctx, req.(*CreateExampleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_UpdateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExampleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).UpdateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_UpdateExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).UpdateExample(ctx, req.(*UpdateExampleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_DeleteExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).DeleteExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_DeleteExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).DeleteExample(ctx, req.(*ExampleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_RecoverExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).RecoverExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_RecoverExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).RecoverExample(ctx, req.(*ExampleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.v1.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExampleList",
			Handler:    _ExampleService_GetExampleList_Handler,
		},
		{
			MethodName: "GetExample",
			Handler:    _ExampleService_GetExample_Handler,
		},
		{
			MethodName: "CreateExample",
			Handler:    _ExampleService_CreateExample_Handler,
		},
		{
			MethodName: "UpdateExample",
			Handler:    _ExampleService_UpdateExample_Handler,
		},
		{
			MethodName: "DeleteExample",
			Handler:    _ExampleService_DeleteExample_Handler,
		},
		{
			MethodName: "RecoverExample",
			Handler:    _ExampleService_RecoverExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/v1/example.proto",
}
