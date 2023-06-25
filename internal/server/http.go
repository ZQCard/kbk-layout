package server

import (
	exampleV1 "github.com/ZQCard/kbk-layout/api/example/v1"
	"github.com/ZQCard/kbk-layout/internal/conf"
	"github.com/ZQCard/kbk-layout/internal/service"
	"github.com/ZQCard/kbk-layout/pkg/middleware/requestInfo"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(conf *conf.Bootstrap, server *conf.Server, service *service.ExampleService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
			// 链路追踪
			tracing.Server(),
			// 访问日志
			logging.Server(logger),
			// 租户
			requestInfo.SetRequestInfo(),
		),
	}
	if server.Http.Network != "" {
		opts = append(opts, http.Network(server.Http.Network))
	}
	if server.Http.Addr != "" {
		opts = append(opts, http.Address(server.Http.Addr))
	}
	if server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	exampleV1.RegisterExampleServiceHTTPServer(srv, service)
	return srv
}
