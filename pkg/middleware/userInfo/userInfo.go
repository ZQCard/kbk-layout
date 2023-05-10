package requestInfo

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

const userIdKey = "x-md-global-userId"
const userNameKey = "x-md-global-username"

// 设置用户信息
func SetUserInfo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			claim, _ := jwt.FromContext(ctx)
			if claim == nil {
				return nil, errors.BadRequest("BAD REQUEST", "登录已过期")
			}
			claimInfo := claim.(jwt2.MapClaims)
			AdministratorId := int64(claimInfo["userId"].(float64))
			ctx = context.WithValue(ctx, userIdKey, AdministratorId)
			ctx = context.WithValue(ctx, userNameKey, claimInfo["username"])
			return handler(ctx, req)
		}
	}
}
