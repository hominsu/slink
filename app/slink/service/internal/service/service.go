package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/internal/biz"
	"github.com/hominsu/slink/app/slink/service/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewSiteService,
	NewUserService,
	NewAdminService,
	NewShortLinkService,
	NewShortLinkRedirectService,
)

type SiteService struct {
	v1.UnimplementedSiteServiceServer

	log     *log.Helper
	version string
}

func NewSiteService(logger log.Logger, version string) *SiteService {
	return &SiteService{
		log:     log.NewHelper(log.With(logger, "module", "service/site")),
		version: version,
	}
}

type UserService struct {
	v1.UnimplementedUserServiceServer

	secret *conf.Secret
	log    *log.Helper
}

func NewUserService(secret *conf.Secret, logger log.Logger) *UserService {
	return &UserService{
		secret: secret,
		log:    log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

type AdminService struct {
	v1.UnimplementedAdminServiceServer

	su  *biz.ShortLinkRepoUsecase
	log *log.Helper
}

func NewAdminService(su *biz.ShortLinkRepoUsecase, logger log.Logger) *AdminService {
	return &AdminService{
		su:  su,
		log: log.NewHelper(log.With(logger, "module", "service/admin")),
	}
}

type ShortLinkService struct {
	v1.UnimplementedShortLinkServiceServer

	su  *biz.ShortLinkRepoUsecase
	log *log.Helper
}

func NewShortLinkService(su *biz.ShortLinkRepoUsecase, logger log.Logger) *ShortLinkService {
	return &ShortLinkService{
		su:  su,
		log: log.NewHelper(log.With(logger, "module", "service/slink")),
	}
}

type ShortLinkRedirectService struct {
	v1.UnimplementedShortLinkRedirectServiceServer

	su  *biz.ShortLinkRepoUsecase
	log *log.Helper
}

func NewShortLinkRedirectService(su *biz.ShortLinkRepoUsecase, logger log.Logger) *ShortLinkRedirectService {
	return &ShortLinkRedirectService{
		su:  su,
		log: log.NewHelper(log.With(logger, "module", "service/slink-redirect")),
	}
}
