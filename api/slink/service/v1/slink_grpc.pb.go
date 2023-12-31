// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: slink/service/v1/slink.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ShortLinkService_CreateShortLink_FullMethodName = "/slink.service.v1.ShortLinkService/CreateShortLink"
)

// ShortLinkServiceClient is the client API for ShortLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortLinkServiceClient interface {
	CreateShortLink(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkReply, error)
}

type shortLinkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShortLinkServiceClient(cc grpc.ClientConnInterface) ShortLinkServiceClient {
	return &shortLinkServiceClient{cc}
}

func (c *shortLinkServiceClient) CreateShortLink(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkReply, error) {
	out := new(CreateShortLinkReply)
	err := c.cc.Invoke(ctx, ShortLinkService_CreateShortLink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortLinkServiceServer is the server API for ShortLinkService service.
// All implementations must embed UnimplementedShortLinkServiceServer
// for forward compatibility
type ShortLinkServiceServer interface {
	CreateShortLink(context.Context, *CreateShortLinkRequest) (*CreateShortLinkReply, error)
	mustEmbedUnimplementedShortLinkServiceServer()
}

// UnimplementedShortLinkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShortLinkServiceServer struct {
}

func (UnimplementedShortLinkServiceServer) CreateShortLink(context.Context, *CreateShortLinkRequest) (*CreateShortLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortLink not implemented")
}
func (UnimplementedShortLinkServiceServer) mustEmbedUnimplementedShortLinkServiceServer() {}

// UnsafeShortLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortLinkServiceServer will
// result in compilation errors.
type UnsafeShortLinkServiceServer interface {
	mustEmbedUnimplementedShortLinkServiceServer()
}

func RegisterShortLinkServiceServer(s grpc.ServiceRegistrar, srv ShortLinkServiceServer) {
	s.RegisterService(&ShortLinkService_ServiceDesc, srv)
}

func _ShortLinkService_CreateShortLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShortLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortLinkServiceServer).CreateShortLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortLinkService_CreateShortLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortLinkServiceServer).CreateShortLink(ctx, req.(*CreateShortLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortLinkService_ServiceDesc is the grpc.ServiceDesc for ShortLinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortLinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "slink.service.v1.ShortLinkService",
	HandlerType: (*ShortLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortLink",
			Handler:    _ShortLinkService_CreateShortLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slink/service/v1/slink.proto",
}