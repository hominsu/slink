// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             (unknown)
// source: slink/service/v1/slink.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationShortLinkServiceCreateShortLink = "/slink.service.v1.ShortLinkService/CreateShortLink"

type ShortLinkServiceHTTPServer interface {
	CreateShortLink(context.Context, *CreateShortLinkRequest) (*CreateShortLinkReply, error)
}

func RegisterShortLinkServiceHTTPServer(s *http.Server, srv ShortLinkServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/slinks", _ShortLinkService_CreateShortLink0_HTTP_Handler(srv))
}

func _ShortLinkService_CreateShortLink0_HTTP_Handler(srv ShortLinkServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateShortLinkRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortLinkServiceCreateShortLink)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateShortLink(ctx, req.(*CreateShortLinkRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateShortLinkReply)
		return ctx.Result(200, reply)
	}
}

type ShortLinkServiceHTTPClient interface {
	CreateShortLink(ctx context.Context, req *CreateShortLinkRequest, opts ...http.CallOption) (rsp *CreateShortLinkReply, err error)
}

type ShortLinkServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewShortLinkServiceHTTPClient(client *http.Client) ShortLinkServiceHTTPClient {
	return &ShortLinkServiceHTTPClientImpl{client}
}

func (c *ShortLinkServiceHTTPClientImpl) CreateShortLink(ctx context.Context, in *CreateShortLinkRequest, opts ...http.CallOption) (*CreateShortLinkReply, error) {
	var out CreateShortLinkReply
	pattern := "/v1/slinks"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortLinkServiceCreateShortLink))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
