// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: leaf-grpc/v1/snowflake.proto

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

// LeafSnowflakeServiceClient is the client API for LeafSnowflakeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeafSnowflakeServiceClient interface {
	// 雪花ID
	GenSnowflakeId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdReply, error)
	// 解析雪花ID
	DecodeSnowflakeId(ctx context.Context, in *DecodeSnowflakeIdReq, opts ...grpc.CallOption) (*DecodeSnowflakeIdResp, error)
	// 获取当前server的时间戳
	GetServerTimestamp(ctx context.Context, in *GetServerTimestampReq, opts ...grpc.CallOption) (*GetServerTimestampResp, error)
}

type leafSnowflakeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeafSnowflakeServiceClient(cc grpc.ClientConnInterface) LeafSnowflakeServiceClient {
	return &leafSnowflakeServiceClient{cc}
}

func (c *leafSnowflakeServiceClient) GenSnowflakeId(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdReply, error) {
	out := new(IdReply)
	err := c.cc.Invoke(ctx, "/leafgrpc.v1.LeafSnowflakeService/GenSnowflakeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leafSnowflakeServiceClient) DecodeSnowflakeId(ctx context.Context, in *DecodeSnowflakeIdReq, opts ...grpc.CallOption) (*DecodeSnowflakeIdResp, error) {
	out := new(DecodeSnowflakeIdResp)
	err := c.cc.Invoke(ctx, "/leafgrpc.v1.LeafSnowflakeService/DecodeSnowflakeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leafSnowflakeServiceClient) GetServerTimestamp(ctx context.Context, in *GetServerTimestampReq, opts ...grpc.CallOption) (*GetServerTimestampResp, error) {
	out := new(GetServerTimestampResp)
	err := c.cc.Invoke(ctx, "/leafgrpc.v1.LeafSnowflakeService/GetServerTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeafSnowflakeServiceServer is the server API for LeafSnowflakeService service.
// All implementations must embed UnimplementedLeafSnowflakeServiceServer
// for forward compatibility
type LeafSnowflakeServiceServer interface {
	// 雪花ID
	GenSnowflakeId(context.Context, *IdRequest) (*IdReply, error)
	// 解析雪花ID
	DecodeSnowflakeId(context.Context, *DecodeSnowflakeIdReq) (*DecodeSnowflakeIdResp, error)
	// 获取当前server的时间戳
	GetServerTimestamp(context.Context, *GetServerTimestampReq) (*GetServerTimestampResp, error)
	mustEmbedUnimplementedLeafSnowflakeServiceServer()
}

// UnimplementedLeafSnowflakeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLeafSnowflakeServiceServer struct {
}

func (UnimplementedLeafSnowflakeServiceServer) GenSnowflakeId(context.Context, *IdRequest) (*IdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenSnowflakeId not implemented")
}
func (UnimplementedLeafSnowflakeServiceServer) DecodeSnowflakeId(context.Context, *DecodeSnowflakeIdReq) (*DecodeSnowflakeIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeSnowflakeId not implemented")
}
func (UnimplementedLeafSnowflakeServiceServer) GetServerTimestamp(context.Context, *GetServerTimestampReq) (*GetServerTimestampResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTimestamp not implemented")
}
func (UnimplementedLeafSnowflakeServiceServer) mustEmbedUnimplementedLeafSnowflakeServiceServer() {}

// UnsafeLeafSnowflakeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeafSnowflakeServiceServer will
// result in compilation errors.
type UnsafeLeafSnowflakeServiceServer interface {
	mustEmbedUnimplementedLeafSnowflakeServiceServer()
}

func RegisterLeafSnowflakeServiceServer(s grpc.ServiceRegistrar, srv LeafSnowflakeServiceServer) {
	s.RegisterService(&LeafSnowflakeService_ServiceDesc, srv)
}

func _LeafSnowflakeService_GenSnowflakeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeafSnowflakeServiceServer).GenSnowflakeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leafgrpc.v1.LeafSnowflakeService/GenSnowflakeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeafSnowflakeServiceServer).GenSnowflakeId(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeafSnowflakeService_DecodeSnowflakeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeSnowflakeIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeafSnowflakeServiceServer).DecodeSnowflakeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leafgrpc.v1.LeafSnowflakeService/DecodeSnowflakeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeafSnowflakeServiceServer).DecodeSnowflakeId(ctx, req.(*DecodeSnowflakeIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeafSnowflakeService_GetServerTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerTimestampReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeafSnowflakeServiceServer).GetServerTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leafgrpc.v1.LeafSnowflakeService/GetServerTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeafSnowflakeServiceServer).GetServerTimestamp(ctx, req.(*GetServerTimestampReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LeafSnowflakeService_ServiceDesc is the grpc.ServiceDesc for LeafSnowflakeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeafSnowflakeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leafgrpc.v1.LeafSnowflakeService",
	HandlerType: (*LeafSnowflakeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenSnowflakeId",
			Handler:    _LeafSnowflakeService_GenSnowflakeId_Handler,
		},
		{
			MethodName: "DecodeSnowflakeId",
			Handler:    _LeafSnowflakeService_DecodeSnowflakeId_Handler,
		},
		{
			MethodName: "GetServerTimestamp",
			Handler:    _LeafSnowflakeService_GetServerTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaf-grpc/v1/snowflake.proto",
}