package server

import (
	exampleV1 "github.com/ZQCard/kratos-base-layout/api/example/v1"
	"github.com/ZQCard/kratos-base-layout/internal/conf"
	"github.com/ZQCard/kratos-base-layout/internal/service"
	"github.com/ZQCard/kratos-base-layout/pkg/middleware/requestInfo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, service *service.ExampleService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			validate.Validator(),
			recovery.Recovery(),
			// 链路追踪
			tracing.Server(),
			// 访问日志
			logging.Server(logger),
			// 租户
			requestInfo.SetRequestInfo(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	exampleV1.RegisterExampleServiceServer(srv, service)
	return srv
}
