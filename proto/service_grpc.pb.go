// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// TODO: how do we secure room service?
	// should be accessible to only internal servers, not external
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
	JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/livekit.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) JoinRoom(ctx context.Context, in *JoinRoomRequest, opts ...grpc.CallOption) (*JoinRoomResponse, error) {
	out := new(JoinRoomResponse)
	err := c.cc.Invoke(ctx, "/livekit.RoomService/JoinRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...grpc.CallOption) (*DeleteRoomResponse, error) {
	out := new(DeleteRoomResponse)
	err := c.cc.Invoke(ctx, "/livekit.RoomService/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	// TODO: how do we secure room service?
	// should be accessible to only internal servers, not external
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
	JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error)
	DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) JoinRoom(context.Context, *JoinRoomRequest) (*JoinRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedRoomServiceServer) DeleteRoom(context.Context, *DeleteRoomRequest) (*DeleteRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livekit.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_JoinRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).JoinRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livekit.RoomService/JoinRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).JoinRoom(ctx, req.(*JoinRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livekit.RoomService/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).DeleteRoom(ctx, req.(*DeleteRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "livekit.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "JoinRoom",
			Handler:    _RoomService_JoinRoom_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _RoomService_DeleteRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

// RTCServiceClient is the client API for RTCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RTCServiceClient interface {
	Offer(ctx context.Context, in *SessionDescription, opts ...grpc.CallOption) (*SessionDescription, error)
	// push channel to allow server to push commands to client
	Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (RTCService_SignalClient, error)
	Trickle(ctx context.Context, in *TrickleRequest, opts ...grpc.CallOption) (*TrickleResponse, error)
}

type rTCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRTCServiceClient(cc grpc.ClientConnInterface) RTCServiceClient {
	return &rTCServiceClient{cc}
}

func (c *rTCServiceClient) Offer(ctx context.Context, in *SessionDescription, opts ...grpc.CallOption) (*SessionDescription, error) {
	out := new(SessionDescription)
	err := c.cc.Invoke(ctx, "/livekit.RTCService/Offer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rTCServiceClient) Signal(ctx context.Context, in *SignalRequest, opts ...grpc.CallOption) (RTCService_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RTCService_serviceDesc.Streams[0], "/livekit.RTCService/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &rTCServiceSignalClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RTCService_SignalClient interface {
	Recv() (*SignalResponse, error)
	grpc.ClientStream
}

type rTCServiceSignalClient struct {
	grpc.ClientStream
}

func (x *rTCServiceSignalClient) Recv() (*SignalResponse, error) {
	m := new(SignalResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rTCServiceClient) Trickle(ctx context.Context, in *TrickleRequest, opts ...grpc.CallOption) (*TrickleResponse, error) {
	out := new(TrickleResponse)
	err := c.cc.Invoke(ctx, "/livekit.RTCService/Trickle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RTCServiceServer is the server API for RTCService service.
// All implementations must embed UnimplementedRTCServiceServer
// for forward compatibility
type RTCServiceServer interface {
	Offer(context.Context, *SessionDescription) (*SessionDescription, error)
	// push channel to allow server to push commands to client
	Signal(*SignalRequest, RTCService_SignalServer) error
	Trickle(context.Context, *TrickleRequest) (*TrickleResponse, error)
	mustEmbedUnimplementedRTCServiceServer()
}

// UnimplementedRTCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRTCServiceServer struct {
}

func (UnimplementedRTCServiceServer) Offer(context.Context, *SessionDescription) (*SessionDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offer not implemented")
}
func (UnimplementedRTCServiceServer) Signal(*SignalRequest, RTCService_SignalServer) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (UnimplementedRTCServiceServer) Trickle(context.Context, *TrickleRequest) (*TrickleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trickle not implemented")
}
func (UnimplementedRTCServiceServer) mustEmbedUnimplementedRTCServiceServer() {}

// UnsafeRTCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RTCServiceServer will
// result in compilation errors.
type UnsafeRTCServiceServer interface {
	mustEmbedUnimplementedRTCServiceServer()
}

func RegisterRTCServiceServer(s *grpc.Server, srv RTCServiceServer) {
	s.RegisterService(&_RTCService_serviceDesc, srv)
}

func _RTCService_Offer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionDescription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RTCServiceServer).Offer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livekit.RTCService/Offer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RTCServiceServer).Offer(ctx, req.(*SessionDescription))
	}
	return interceptor(ctx, in, info, handler)
}

func _RTCService_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SignalRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RTCServiceServer).Signal(m, &rTCServiceSignalServer{stream})
}

type RTCService_SignalServer interface {
	Send(*SignalResponse) error
	grpc.ServerStream
}

type rTCServiceSignalServer struct {
	grpc.ServerStream
}

func (x *rTCServiceSignalServer) Send(m *SignalResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RTCService_Trickle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrickleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RTCServiceServer).Trickle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/livekit.RTCService/Trickle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RTCServiceServer).Trickle(ctx, req.(*TrickleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RTCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "livekit.RTCService",
	HandlerType: (*RTCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Offer",
			Handler:    _RTCService_Offer_Handler,
		},
		{
			MethodName: "Trickle",
			Handler:    _RTCService_Trickle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _RTCService_Signal_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
