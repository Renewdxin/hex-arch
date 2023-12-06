// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0--rc2
// source: internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto

package pb

import (
	pb "./internal/adapters/framework/left/grpc/proto/pb"
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
	ArithmeticService_Add_FullMethodName      = "/pb.ArithmeticService/Add"
	ArithmeticService_Subtract_FullMethodName = "/pb.ArithmeticService/Subtract"
	ArithmeticService_Multiply_FullMethodName = "/pb.ArithmeticService/Multiply"
	ArithmeticService_Divide_FullMethodName   = "/pb.ArithmeticService/Divide"
)

// ArithmeticServiceClient is the client API for ArithmeticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArithmeticServiceClient interface {
	Add(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error)
	Subtract(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error)
	Multiply(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error)
	Divide(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error)
}

type arithmeticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArithmeticServiceClient(cc grpc.ClientConnInterface) ArithmeticServiceClient {
	return &arithmeticServiceClient{cc}
}

func (c *arithmeticServiceClient) Add(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error) {
	out := new(pb.Answer)
	err := c.cc.Invoke(ctx, ArithmeticService_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) Subtract(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error) {
	out := new(pb.Answer)
	err := c.cc.Invoke(ctx, ArithmeticService_Subtract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) Multiply(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error) {
	out := new(pb.Answer)
	err := c.cc.Invoke(ctx, ArithmeticService_Multiply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *arithmeticServiceClient) Divide(ctx context.Context, in *pb.OperationParameters, opts ...grpc.CallOption) (*pb.Answer, error) {
	out := new(pb.Answer)
	err := c.cc.Invoke(ctx, ArithmeticService_Divide_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArithmeticServiceServer is the server API for ArithmeticService service.
// All implementations must embed UnimplementedArithmeticServiceServer
// for forward compatibility
type ArithmeticServiceServer interface {
	Add(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	Subtract(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	Multiply(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	Divide(context.Context, *pb.OperationParameters) (*pb.Answer, error)
	mustEmbedUnimplementedArithmeticServiceServer()
}

// UnimplementedArithmeticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedArithmeticServiceServer struct {
}

func (UnimplementedArithmeticServiceServer) Add(context.Context, *pb.OperationParameters) (*pb.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedArithmeticServiceServer) Subtract(context.Context, *pb.OperationParameters) (*pb.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedArithmeticServiceServer) Multiply(context.Context, *pb.OperationParameters) (*pb.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedArithmeticServiceServer) Divide(context.Context, *pb.OperationParameters) (*pb.Answer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedArithmeticServiceServer) mustEmbedUnimplementedArithmeticServiceServer() {}

// UnsafeArithmeticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArithmeticServiceServer will
// result in compilation errors.
type UnsafeArithmeticServiceServer interface {
	mustEmbedUnimplementedArithmeticServiceServer()
}

func RegisterArithmeticServiceServer(s grpc.ServiceRegistrar, srv ArithmeticServiceServer) {
	s.RegisterService(&ArithmeticService_ServiceDesc, srv)
}

func _ArithmeticService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArithmeticService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Add(ctx, req.(*pb.OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArithmeticService_Subtract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Subtract(ctx, req.(*pb.OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArithmeticService_Multiply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Multiply(ctx, req.(*pb.OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArithmeticService_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pb.OperationParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArithmeticServiceServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArithmeticService_Divide_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArithmeticServiceServer).Divide(ctx, req.(*pb.OperationParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// ArithmeticService_ServiceDesc is the grpc.ServiceDesc for ArithmeticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArithmeticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ArithmeticService",
	HandlerType: (*ArithmeticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _ArithmeticService_Add_Handler,
		},
		{
			MethodName: "Subtract",
			Handler:    _ArithmeticService_Subtract_Handler,
		},
		{
			MethodName: "Multiply",
			Handler:    _ArithmeticService_Multiply_Handler,
		},
		{
			MethodName: "Divide",
			Handler:    _ArithmeticService_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto",
}
