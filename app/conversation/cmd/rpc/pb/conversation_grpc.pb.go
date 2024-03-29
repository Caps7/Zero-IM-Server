// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: conversation.proto

package pb

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

// ConversationServiceClient is the client API for ConversationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationServiceClient interface {
	ModifyConversationField(ctx context.Context, in *ModifyConversationFieldReq, opts ...grpc.CallOption) (*ModifyConversationFieldResp, error)
}

type conversationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationServiceClient(cc grpc.ClientConnInterface) ConversationServiceClient {
	return &conversationServiceClient{cc}
}

func (c *conversationServiceClient) ModifyConversationField(ctx context.Context, in *ModifyConversationFieldReq, opts ...grpc.CallOption) (*ModifyConversationFieldResp, error) {
	out := new(ModifyConversationFieldResp)
	err := c.cc.Invoke(ctx, "/pb.conversationService/ModifyConversationField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServiceServer is the server API for ConversationService service.
// All implementations must embed UnimplementedConversationServiceServer
// for forward compatibility
type ConversationServiceServer interface {
	ModifyConversationField(context.Context, *ModifyConversationFieldReq) (*ModifyConversationFieldResp, error)
	mustEmbedUnimplementedConversationServiceServer()
}

// UnimplementedConversationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversationServiceServer struct {
}

func (UnimplementedConversationServiceServer) ModifyConversationField(context.Context, *ModifyConversationFieldReq) (*ModifyConversationFieldResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyConversationField not implemented")
}
func (UnimplementedConversationServiceServer) mustEmbedUnimplementedConversationServiceServer() {}

// UnsafeConversationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversationServiceServer will
// result in compilation errors.
type UnsafeConversationServiceServer interface {
	mustEmbedUnimplementedConversationServiceServer()
}

func RegisterConversationServiceServer(s grpc.ServiceRegistrar, srv ConversationServiceServer) {
	s.RegisterService(&ConversationService_ServiceDesc, srv)
}

func _ConversationService_ModifyConversationField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyConversationFieldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServiceServer).ModifyConversationField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.conversationService/ModifyConversationField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServiceServer).ModifyConversationField(ctx, req.(*ModifyConversationFieldReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversationService_ServiceDesc is the grpc.ServiceDesc for ConversationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.conversationService",
	HandlerType: (*ConversationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ModifyConversationField",
			Handler:    _ConversationService_ModifyConversationField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation.proto",
}
