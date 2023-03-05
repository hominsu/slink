package middleware

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

var (
	ErrNotAdminUser    = errors.Unauthorized(reason, "not admin user")
	ErrMissingUserInfo = errors.Unauthorized(reason, "missing user info")
	ErrInternalServer  = errors.InternalServer(reason, "internal server error")
)

func Admin(logger log.Logger) middleware.Middleware {
	helper := log.NewHelper(log.With(logger, "module", "middleware/admin"))

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			// get user id from context
			v := ctx.Value(string(ContextKeyIsAdmin))
			if v == nil {
				return nil, ErrMissingUserInfo
			}
			isAdmin, ok := v.(string)
			if !ok {
				helper.Warnf("failed to get userid from context")
				return nil, ErrInternalServer
			}

			ok, err := strconv.ParseBool(isAdmin)
			if err != nil {
				helper.Warnf("failed to check user if is admin user, err: %v", err)
				return nil, ErrInternalServer
			}
			if !ok {
				return nil, ErrNotAdminUser
			}

			return handler(ctx, req)
		}
	}
}
