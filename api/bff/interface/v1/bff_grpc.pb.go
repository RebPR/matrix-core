// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: bff/interface/v1/bff.proto

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

// BffClient is the client API for Bff service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BffClient interface {
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterReply, error)
	LoginByPassword(ctx context.Context, in *LoginByPasswordReq, opts ...grpc.CallOption) (*LoginReply, error)
	LoginByCode(ctx context.Context, in *LoginByCodeReq, opts ...grpc.CallOption) (*LoginReply, error)
	LoginByWeChat(ctx context.Context, in *LoginByWeChatReq, opts ...grpc.CallOption) (*LoginReply, error)
	LoginByGithub(ctx context.Context, in *LoginByGithubReq, opts ...grpc.CallOption) (*LoginReply, error)
	LoginPasswordReset(ctx context.Context, in *LoginPasswordResetReq, opts ...grpc.CallOption) (*LoginPasswordResetReply, error)
	SendPhoneCode(ctx context.Context, in *SendPhoneCodeReq, opts ...grpc.CallOption) (*SendPhoneCodeReply, error)
	SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeReply, error)
	GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileReply, error)
}

type bffClient struct {
	cc grpc.ClientConnInterface
}

func NewBffClient(cc grpc.ClientConnInterface) BffClient {
	return &bffClient{cc}
}

func (c *bffClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterReply, error) {
	out := new(UserRegisterReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) LoginByPassword(ctx context.Context, in *LoginByPasswordReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/LoginByPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) LoginByCode(ctx context.Context, in *LoginByCodeReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/LoginByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) LoginByWeChat(ctx context.Context, in *LoginByWeChatReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/LoginByWeChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) LoginByGithub(ctx context.Context, in *LoginByGithubReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/LoginByGithub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) LoginPasswordReset(ctx context.Context, in *LoginPasswordResetReq, opts ...grpc.CallOption) (*LoginPasswordResetReply, error) {
	out := new(LoginPasswordResetReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/LoginPasswordReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) SendPhoneCode(ctx context.Context, in *SendPhoneCodeReq, opts ...grpc.CallOption) (*SendPhoneCodeReply, error) {
	out := new(SendPhoneCodeReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/SendPhoneCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) SendEmailCode(ctx context.Context, in *SendEmailCodeReq, opts ...grpc.CallOption) (*SendEmailCodeReply, error) {
	out := new(SendEmailCodeReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/SendEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bffClient) GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileReply, error) {
	out := new(GetUserProfileReply)
	err := c.cc.Invoke(ctx, "/bff.v1.Bff/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BffServer is the server API for Bff service.
// All implementations must embed UnimplementedBffServer
// for forward compatibility
type BffServer interface {
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterReply, error)
	LoginByPassword(context.Context, *LoginByPasswordReq) (*LoginReply, error)
	LoginByCode(context.Context, *LoginByCodeReq) (*LoginReply, error)
	LoginByWeChat(context.Context, *LoginByWeChatReq) (*LoginReply, error)
	LoginByGithub(context.Context, *LoginByGithubReq) (*LoginReply, error)
	LoginPasswordReset(context.Context, *LoginPasswordResetReq) (*LoginPasswordResetReply, error)
	SendPhoneCode(context.Context, *SendPhoneCodeReq) (*SendPhoneCodeReply, error)
	SendEmailCode(context.Context, *SendEmailCodeReq) (*SendEmailCodeReply, error)
	GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileReply, error)
	mustEmbedUnimplementedBffServer()
}

// UnimplementedBffServer must be embedded to have forward compatible implementations.
type UnimplementedBffServer struct {
}

func (UnimplementedBffServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedBffServer) LoginByPassword(context.Context, *LoginByPasswordReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByPassword not implemented")
}
func (UnimplementedBffServer) LoginByCode(context.Context, *LoginByCodeReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByCode not implemented")
}
func (UnimplementedBffServer) LoginByWeChat(context.Context, *LoginByWeChatReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByWeChat not implemented")
}
func (UnimplementedBffServer) LoginByGithub(context.Context, *LoginByGithubReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByGithub not implemented")
}
func (UnimplementedBffServer) LoginPasswordReset(context.Context, *LoginPasswordResetReq) (*LoginPasswordResetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginPasswordReset not implemented")
}
func (UnimplementedBffServer) SendPhoneCode(context.Context, *SendPhoneCodeReq) (*SendPhoneCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneCode not implemented")
}
func (UnimplementedBffServer) SendEmailCode(context.Context, *SendEmailCodeReq) (*SendEmailCodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailCode not implemented")
}
func (UnimplementedBffServer) GetUserProfile(context.Context, *GetUserProfileReq) (*GetUserProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedBffServer) mustEmbedUnimplementedBffServer() {}

// UnsafeBffServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BffServer will
// result in compilation errors.
type UnsafeBffServer interface {
	mustEmbedUnimplementedBffServer()
}

func RegisterBffServer(s grpc.ServiceRegistrar, srv BffServer) {
	s.RegisterService(&Bff_ServiceDesc, srv)
}

func _Bff_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_LoginByPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).LoginByPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/LoginByPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).LoginByPassword(ctx, req.(*LoginByPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_LoginByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).LoginByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/LoginByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).LoginByCode(ctx, req.(*LoginByCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_LoginByWeChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByWeChatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).LoginByWeChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/LoginByWeChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).LoginByWeChat(ctx, req.(*LoginByWeChatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_LoginByGithub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByGithubReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).LoginByGithub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/LoginByGithub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).LoginByGithub(ctx, req.(*LoginByGithubReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_LoginPasswordReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginPasswordResetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).LoginPasswordReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/LoginPasswordReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).LoginPasswordReset(ctx, req.(*LoginPasswordResetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_SendPhoneCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).SendPhoneCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/SendPhoneCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).SendPhoneCode(ctx, req.(*SendPhoneCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_SendEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).SendEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/SendEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).SendEmailCode(ctx, req.(*SendEmailCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bff_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BffServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bff.v1.Bff/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BffServer).GetUserProfile(ctx, req.(*GetUserProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Bff_ServiceDesc is the grpc.ServiceDesc for Bff service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bff_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bff.v1.Bff",
	HandlerType: (*BffServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _Bff_UserRegister_Handler,
		},
		{
			MethodName: "LoginByPassword",
			Handler:    _Bff_LoginByPassword_Handler,
		},
		{
			MethodName: "LoginByCode",
			Handler:    _Bff_LoginByCode_Handler,
		},
		{
			MethodName: "LoginByWeChat",
			Handler:    _Bff_LoginByWeChat_Handler,
		},
		{
			MethodName: "LoginByGithub",
			Handler:    _Bff_LoginByGithub_Handler,
		},
		{
			MethodName: "LoginPasswordReset",
			Handler:    _Bff_LoginPasswordReset_Handler,
		},
		{
			MethodName: "SendPhoneCode",
			Handler:    _Bff_SendPhoneCode_Handler,
		},
		{
			MethodName: "SendEmailCode",
			Handler:    _Bff_SendEmailCode_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _Bff_GetUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bff/interface/v1/bff.proto",
}
