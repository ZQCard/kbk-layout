package requestInfoMiddleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
)

const domainKey = "x-md-global-domain"

// setRequestInfo 设置Request信息
func setRequestInfo() middleware.Middleware {
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

				// 获取token值
				auths := strings.SplitN(tr.RequestHeader().Get("Authorization"), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
					return nil, errors.BadRequest("BAD REQUEST", "登录已过期")
				}
				jwtToken := auths[1]
				// 设置 Authorization
				tr.RequestHeader().Set("Authorization", "Bearer "+jwtToken)
			}
			return handler(ctx, req)
		}
	}
}
