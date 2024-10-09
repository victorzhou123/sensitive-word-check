// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.17.3
// source: sensitive_word_check.proto

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
	Check_Check_FullMethodName = "/proto.check/check"
)

// CheckClient is the client API for Check service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckClient interface {
	Check(ctx context.Context, in *Word, opts ...grpc.CallOption) (*BoolMsg, error)
}

type checkClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckClient(cc grpc.ClientConnInterface) CheckClient {
	return &checkClient{cc}
}

func (c *checkClient) Check(ctx context.Context, in *Word, opts ...grpc.CallOption) (*BoolMsg, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BoolMsg)
	err := c.cc.Invoke(ctx, Check_Check_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckServer is the server API for Check service.
// All implementations must embed UnimplementedCheckServer
// for forward compatibility.
type CheckServer interface {
	Check(context.Context, *Word) (*BoolMsg, error)
	mustEmbedUnimplementedCheckServer()
}

// UnimplementedCheckServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCheckServer struct{}

func (UnimplementedCheckServer) Check(context.Context, *Word) (*BoolMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedCheckServer) mustEmbedUnimplementedCheckServer() {}
func (UnimplementedCheckServer) testEmbeddedByValue()               {}

// UnsafeCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckServer will
// result in compilation errors.
type UnsafeCheckServer interface {
	mustEmbedUnimplementedCheckServer()
}

func RegisterCheckServer(s grpc.ServiceRegistrar, srv CheckServer) {
	// If the following call pancis, it indicates UnimplementedCheckServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Check_ServiceDesc, srv)
}

func _Check_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Check_Check_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckServer).Check(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

// Check_ServiceDesc is the grpc.ServiceDesc for Check service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Check_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.check",
	HandlerType: (*CheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "check",
			Handler:    _Check_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sensitive_word_check.proto",
}