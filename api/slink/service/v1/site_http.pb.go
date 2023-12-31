// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.2
// - protoc             (unknown)
// source: slink/service/v1/site.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSiteServiceCount = "/slink.service.v1.SiteService/Count"
const OperationSiteServicePing = "/slink.service.v1.SiteService/Ping"

type SiteServiceHTTPServer interface {
	Count(context.Context, *emptypb.Empty) (*CountReply, error)
	Ping(context.Context, *emptypb.Empty) (*PingReply, error)
}

func RegisterSiteServiceHTTPServer(s *http.Server, srv SiteServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/site/ping", _SiteService_Ping0_HTTP_Handler(srv))
	r.GET("/v1/site/count", _SiteService_Count0_HTTP_Handler(srv))
}

func _SiteService_Ping0_HTTP_Handler(srv SiteServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSiteServicePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PingReply)
		return ctx.Result(200, reply)
	}
}

func _SiteService_Count0_HTTP_Handler(srv SiteServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSiteServiceCount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Count(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CountReply)
		return ctx.Result(200, reply)
	}
}

type SiteServiceHTTPClient interface {
	Count(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CountReply, err error)
	Ping(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *PingReply, err error)
}

type SiteServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSiteServiceHTTPClient(client *http.Client) SiteServiceHTTPClient {
	return &SiteServiceHTTPClientImpl{client}
}

func (c *SiteServiceHTTPClientImpl) Count(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CountReply, error) {
	var out CountReply
	pattern := "/v1/site/count"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSiteServiceCount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SiteServiceHTTPClientImpl) Ping(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*PingReply, error) {
	var out PingReply
	pattern := "/v1/site/ping"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSiteServicePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
