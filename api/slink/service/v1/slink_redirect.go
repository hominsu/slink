package v1

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RedirectRequest struct {
	Key string `json:"key,omitempty"`
}

type ShortLinkRedirectServiceServer interface {
	Redirect(context.Context, *RedirectRequest) (http.Redirector, error)
	mustEmbedUnimplementedShortLinkRedirectServiceServer()
}

// UnimplementedShortLinkRedirectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortLinkRedirectServiceServer struct{}

func (UnimplementedShortLinkRedirectServiceServer) Redirect(context.Context, *RedirectRequest) (http.Redirector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redirect not implemented")
}

func (UnimplementedShortLinkRedirectServiceServer) mustEmbedUnimplementedShortLinkRedirectServiceServer() {
}

// UnsafeShortLinkRedirectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortLinkServiceServer will
// result in compilation errors.
type UnsafeShortLinkRedirectServiceServer interface {
	mustEmbedUnimplementedShortLinkRedirectServiceServer()
}
