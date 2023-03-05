package server

import (
	"context"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/internal/conf"
	"github.com/hominsu/slink/app/slink/service/internal/service"
	"github.com/hominsu/slink/app/slink/service/pkgs/middleware"
)

func NewSkipJwtMatcher() selector.MatchFunc {
	skipList := make(map[string]struct{})
	skipList["/slink.service.v1.SiteService/Ping"] = struct{}{}
	skipList["/slink.service.v1.UserService/Login"] = struct{}{}
	skipList["/slink.service.v1.ShortLinkService/CreateShortLink"] = struct{}{}
	skipList["/slink.service.v1.ShortLinkRedirectService/Redirect"] = struct{}{}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipList[operation]; ok {
			return false
		}
		return true
	}
}

func NewAdminMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		return strings.HasPrefix("/slink.service.v1.AdminService/", operation)
	}
}

func NewHTTPServer(
	c *conf.Server,
	ss *service.SiteService,
	us *service.UserService,
	as *service.AdminService,
	sls *service.ShortLinkService,
	slrs *service.ShortLinkRedirectService,
	sc *conf.Secret,
	logger log.Logger,
) *http.Server {
	// server options
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
			selector.Server(
				middleware.JwtAuth(sc.Jwt.Secret),
			).
				Match(NewSkipJwtMatcher()).
				Build(),
			selector.Server(
				middleware.Admin(logger),
			).
				Match(NewAdminMatcher()).
				Build(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	v1.RegisterSiteServiceHTTPServer(srv, ss)
	v1.RegisterUserServiceHTTPServer(srv, us)
	v1.RegisterAdminServiceHTTPServer(srv, as)
	v1.RegisterShortLinkServiceHTTPServer(srv, sls)
	v1.RegisterShortLinkRedirectServiceHTTPServer(srv, slrs)

	return srv
}
