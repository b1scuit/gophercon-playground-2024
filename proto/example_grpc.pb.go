// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.4
// source: proto/example.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ExampleServer_GetExample_FullMethodName = "/cpuprofile.ExampleServer/GetExample"
)

// ExampleServerClient is the client API for ExampleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServerClient interface {
	GetExample(ctx context.Context, in *Example, opts ...grpc.CallOption) (*Example, error)
}

type exampleServerClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServerClient(cc grpc.ClientConnInterface) ExampleServerClient {
	return &exampleServerClient{cc}
}

func (c *exampleServerClient) GetExample(ctx context.Context, in *Example, opts ...grpc.CallOption) (*Example, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Example)
	err := c.cc.Invoke(ctx, ExampleServer_GetExample_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServerServer is the server API for ExampleServer service.
// All implementations must embed UnimplementedExampleServerServer
// for forward compatibility.
type ExampleServerServer interface {
	GetExample(context.Context, *Example) (*Example, error)
	mustEmbedUnimplementedExampleServerServer()
}

// UnimplementedExampleServerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExampleServerServer struct{}

func (UnimplementedExampleServerServer) GetExample(context.Context, *Example) (*Example, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExample not implemented")
}
func (UnimplementedExampleServerServer) mustEmbedUnimplementedExampleServerServer() {}
func (UnimplementedExampleServerServer) testEmbeddedByValue()                       {}

// UnsafeExampleServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServerServer will
// result in compilation errors.
type UnsafeExampleServerServer interface {
	mustEmbedUnimplementedExampleServerServer()
}

func RegisterExampleServerServer(s grpc.ServiceRegistrar, srv ExampleServerServer) {
	// If the following call pancis, it indicates UnimplementedExampleServerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExampleServer_ServiceDesc, srv)
}

func _ExampleServer_GetExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Example)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServerServer).GetExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleServer_GetExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServerServer).GetExample(ctx, req.(*Example))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleServer_ServiceDesc is the grpc.ServiceDesc for ExampleServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cpuprofile.ExampleServer",
	HandlerType: (*ExampleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExample",
			Handler:    _ExampleServer_GetExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/example.proto",
}
