// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: protobufs/mamar/port.proto

package proto

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

// MamarClient is the client API for Mamar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MamarClient interface {
	GetPort(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Port, error)
}

type mamarClient struct {
	cc grpc.ClientConnInterface
}

func NewMamarClient(cc grpc.ClientConnInterface) MamarClient {
	return &mamarClient{cc}
}

func (c *mamarClient) GetPort(ctx context.Context, in *Service, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/mamar.mamar/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MamarServer is the server API for Mamar service.
// All implementations must embed UnimplementedMamarServer
// for forward compatibility
type MamarServer interface {
	GetPort(context.Context, *Service) (*Port, error)
	mustEmbedUnimplementedMamarServer()
}

// UnimplementedMamarServer must be embedded to have forward compatible implementations.
type UnimplementedMamarServer struct {
}

func (UnimplementedMamarServer) GetPort(context.Context, *Service) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (UnimplementedMamarServer) mustEmbedUnimplementedMamarServer() {}

// UnsafeMamarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MamarServer will
// result in compilation errors.
type UnsafeMamarServer interface {
	mustEmbedUnimplementedMamarServer()
}

func RegisterMamarServer(s grpc.ServiceRegistrar, srv MamarServer) {
	s.RegisterService(&Mamar_ServiceDesc, srv)
}

func _Mamar_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MamarServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mamar.mamar/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MamarServer).GetPort(ctx, req.(*Service))
	}
	return interceptor(ctx, in, info, handler)
}

// Mamar_ServiceDesc is the grpc.ServiceDesc for Mamar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mamar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mamar.mamar",
	HandlerType: (*MamarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _Mamar_GetPort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/mamar/port.proto",
}
