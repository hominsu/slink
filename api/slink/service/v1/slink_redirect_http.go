package v1

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var (
	_ = new(context.Context)
	_ = binding.EncodeURL
)

const _ = http.SupportPackageIsVersion1

const OperationShortLinkRedirectServiceRedirect = "/slink.service.v1.ShortLinkRedirectService/Redirect"

type ShortLinkRedirectServiceHTTPServer interface {
	Redirect(context.Context, *RedirectRequest) (http.Redirector, error)
}

func RegisterShortLinkRedirectServiceHTTPServer(s *http.Server, srv ShortLinkRedirectServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/s/{key}", ShortLinkRedirectServiceRedirectHTTPHandler(srv))
}

func ShortLinkRedirectServiceRedirectHTTPHandler(srv ShortLinkRedirectServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RedirectRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortLinkRedirectServiceRedirect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Redirect(ctx, req.(*RedirectRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}
