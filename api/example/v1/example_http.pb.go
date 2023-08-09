// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.2
// - protoc             v4.23.0--rc1
// source: example/v1/example.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationExampleServiceCreateExample = "/example.v1.ExampleService/CreateExample"
const OperationExampleServiceDeleteExample = "/example.v1.ExampleService/DeleteExample"
const OperationExampleServiceGetExample = "/example.v1.ExampleService/GetExample"
const OperationExampleServiceGetExampleList = "/example.v1.ExampleService/GetExampleList"
const OperationExampleServiceRecoverExample = "/example.v1.ExampleService/RecoverExample"
const OperationExampleServiceUpdateExample = "/example.v1.ExampleService/UpdateExample"

type ExampleServiceHTTPServer interface {
	CreateExample(context.Context, *CreateExampleReq) (*Example, error)
	DeleteExample(context.Context, *ExampleIdReq) (*emptypb.Empty, error)
	GetExample(context.Context, *ExampleIdReq) (*Example, error)
	GetExampleList(context.Context, *GetExampleListReq) (*GetExampleListPageRes, error)
	RecoverExample(context.Context, *ExampleIdReq) (*emptypb.Empty, error)
	UpdateExample(context.Context, *UpdateExampleReq) (*emptypb.Empty, error)
}

func RegisterExampleServiceHTTPServer(s *http.Server, srv ExampleServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/example", _ExampleService_GetExampleList0_HTTP_Handler(srv))
	r.GET("/example/{id}", _ExampleService_GetExample0_HTTP_Handler(srv))
	r.POST("/example", _ExampleService_CreateExample0_HTTP_Handler(srv))
	r.PUT("/example", _ExampleService_UpdateExample0_HTTP_Handler(srv))
	r.DELETE("/example", _ExampleService_DeleteExample0_HTTP_Handler(srv))
	r.PATCH("/example", _ExampleService_RecoverExample0_HTTP_Handler(srv))
}

func _ExampleService_GetExampleList0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetExampleListReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceGetExampleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetExampleList(ctx, req.(*GetExampleListReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetExampleListPageRes)
		return ctx.Result(200, reply)
	}
}

func _ExampleService_GetExample0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExampleIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceGetExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetExample(ctx, req.(*ExampleIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Example)
		return ctx.Result(200, reply)
	}
}

func _ExampleService_CreateExample0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateExampleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceCreateExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateExample(ctx, req.(*CreateExampleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Example)
		return ctx.Result(200, reply)
	}
}

func _ExampleService_UpdateExample0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateExampleReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceUpdateExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateExample(ctx, req.(*UpdateExampleReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _ExampleService_DeleteExample0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExampleIdReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceDeleteExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteExample(ctx, req.(*ExampleIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _ExampleService_RecoverExample0_HTTP_Handler(srv ExampleServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExampleIdReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExampleServiceRecoverExample)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RecoverExample(ctx, req.(*ExampleIdReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type ExampleServiceHTTPClient interface {
	CreateExample(ctx context.Context, req *CreateExampleReq, opts ...http.CallOption) (rsp *Example, err error)
	DeleteExample(ctx context.Context, req *ExampleIdReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetExample(ctx context.Context, req *ExampleIdReq, opts ...http.CallOption) (rsp *Example, err error)
	GetExampleList(ctx context.Context, req *GetExampleListReq, opts ...http.CallOption) (rsp *GetExampleListPageRes, err error)
	RecoverExample(ctx context.Context, req *ExampleIdReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateExample(ctx context.Context, req *UpdateExampleReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ExampleServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewExampleServiceHTTPClient(client *http.Client) ExampleServiceHTTPClient {
	return &ExampleServiceHTTPClientImpl{client}
}

func (c *ExampleServiceHTTPClientImpl) CreateExample(ctx context.Context, in *CreateExampleReq, opts ...http.CallOption) (*Example, error) {
	var out Example
	pattern := "/example"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExampleServiceCreateExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExampleServiceHTTPClientImpl) DeleteExample(ctx context.Context, in *ExampleIdReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/example"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationExampleServiceDeleteExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExampleServiceHTTPClientImpl) GetExample(ctx context.Context, in *ExampleIdReq, opts ...http.CallOption) (*Example, error) {
	var out Example
	pattern := "/example/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationExampleServiceGetExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExampleServiceHTTPClientImpl) GetExampleList(ctx context.Context, in *GetExampleListReq, opts ...http.CallOption) (*GetExampleListPageRes, error) {
	var out GetExampleListPageRes
	pattern := "/example"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationExampleServiceGetExampleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExampleServiceHTTPClientImpl) RecoverExample(ctx context.Context, in *ExampleIdReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/example"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExampleServiceRecoverExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExampleServiceHTTPClientImpl) UpdateExample(ctx context.Context, in *UpdateExampleReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/example"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExampleServiceUpdateExample))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
