package requestInfo

import (
	"context"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/middleware"
)

const domainKey = "x-md-global-domain"

// setRequestInfo 设置Request信息
func SetRequestInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			if tr, ok := transport.FromServerContext(ctx); ok {
				// 将请求信息放入ctx中
				if ht, ok := tr.(*http.Transport); ok {
					ctx = context.WithValue(ctx, "RemoteAddr", ht.Request().RemoteAddr)
				}

				// 获取请求域
				domain := tr.RequestHeader().Get("domain")
				ctx = metadata.AppendToClientContext(ctx, domainKey, domain)
				ctx = context.WithValue(ctx, "domain", domain)
			}
			return handler(ctx, req)
		}
	}
}
