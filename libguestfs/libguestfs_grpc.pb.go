// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package libguestfs

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

// VirtSparsifyClient is the client API for VirtSparsify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtSparsifyClient interface {
	Sparsify(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Response, error)
}

type virtSparsifyClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtSparsifyClient(cc grpc.ClientConnInterface) VirtSparsifyClient {
	return &virtSparsifyClient{cc}
}

func (c *virtSparsifyClient) Sparsify(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/alicefr.guestfs.libguestfs.VirtSparsify/Sparsify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtSparsifyServer is the server API for VirtSparsify service.
// All implementations must embed UnimplementedVirtSparsifyServer
// for forward compatibility
type VirtSparsifyServer interface {
	Sparsify(context.Context, *Image) (*Response, error)
	mustEmbedUnimplementedVirtSparsifyServer()
}

// UnimplementedVirtSparsifyServer must be embedded to have forward compatible implementations.
type UnimplementedVirtSparsifyServer struct {
}

func (UnimplementedVirtSparsifyServer) Sparsify(context.Context, *Image) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sparsify not implemented")
}
func (UnimplementedVirtSparsifyServer) mustEmbedUnimplementedVirtSparsifyServer() {}

// UnsafeVirtSparsifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtSparsifyServer will
// result in compilation errors.
type UnsafeVirtSparsifyServer interface {
	mustEmbedUnimplementedVirtSparsifyServer()
}

func RegisterVirtSparsifyServer(s grpc.ServiceRegistrar, srv VirtSparsifyServer) {
	s.RegisterService(&VirtSparsify_ServiceDesc, srv)
}

func _VirtSparsify_Sparsify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtSparsifyServer).Sparsify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alicefr.guestfs.libguestfs.VirtSparsify/Sparsify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtSparsifyServer).Sparsify(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

// VirtSparsify_ServiceDesc is the grpc.ServiceDesc for VirtSparsify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtSparsify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "alicefr.guestfs.libguestfs.VirtSparsify",
	HandlerType: (*VirtSparsifyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sparsify",
			Handler:    _VirtSparsify_Sparsify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libguestfs/libguestfs.proto",
}
