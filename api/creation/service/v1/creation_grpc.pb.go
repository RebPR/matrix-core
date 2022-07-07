// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: creation/service/v1/creation.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CreationClient is the client API for Creation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreationClient interface {
	GetLeaderBoard(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLeaderBoardReply, error)
	GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListReply, error)
	GetArticleListHot(ctx context.Context, in *GetArticleListHotReq, opts ...grpc.CallOption) (*GetArticleListHotReply, error)
	GetArticleStatistic(ctx context.Context, in *GetArticleStatisticReq, opts ...grpc.CallOption) (*GetArticleStatisticReply, error)
	GetArticleListStatistic(ctx context.Context, in *GetArticleListStatisticReq, opts ...grpc.CallOption) (*GetArticleListStatisticReply, error)
	GetLastArticleDraft(ctx context.Context, in *GetLastArticleDraftReq, opts ...grpc.CallOption) (*GetLastArticleDraftReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateArticleCacheAndSearch(ctx context.Context, in *CreateArticleCacheAndSearchReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateArticleDraft(ctx context.Context, in *CreateArticleDraftReq, opts ...grpc.CallOption) (*CreateArticleDraftReply, error)
	ArticleDraftMark(ctx context.Context, in *ArticleDraftMarkReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetArticleDraftList(ctx context.Context, in *GetArticleDraftListReq, opts ...grpc.CallOption) (*GetArticleDraftListReply, error)
	SendArticle(ctx context.Context, in *SendArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetArticleAgree(ctx context.Context, in *SetArticleAgreeReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetArticleView(ctx context.Context, in *SetArticleViewReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type creationClient struct {
	cc grpc.ClientConnInterface
}

func NewCreationClient(cc grpc.ClientConnInterface) CreationClient {
	return &creationClient{cc}
}

func (c *creationClient) GetLeaderBoard(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLeaderBoardReply, error) {
	out := new(GetLeaderBoardReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetLeaderBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetArticleList(ctx context.Context, in *GetArticleListReq, opts ...grpc.CallOption) (*GetArticleListReply, error) {
	out := new(GetArticleListReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetArticleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetArticleListHot(ctx context.Context, in *GetArticleListHotReq, opts ...grpc.CallOption) (*GetArticleListHotReply, error) {
	out := new(GetArticleListHotReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetArticleListHot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetArticleStatistic(ctx context.Context, in *GetArticleStatisticReq, opts ...grpc.CallOption) (*GetArticleStatisticReply, error) {
	out := new(GetArticleStatisticReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetArticleStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetArticleListStatistic(ctx context.Context, in *GetArticleListStatisticReq, opts ...grpc.CallOption) (*GetArticleListStatisticReply, error) {
	out := new(GetArticleListStatisticReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetArticleListStatistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetLastArticleDraft(ctx context.Context, in *GetLastArticleDraftReq, opts ...grpc.CallOption) (*GetLastArticleDraftReply, error) {
	out := new(GetLastArticleDraftReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetLastArticleDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) CreateArticle(ctx context.Context, in *CreateArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) CreateArticleCacheAndSearch(ctx context.Context, in *CreateArticleCacheAndSearchReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/CreateArticleCacheAndSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) CreateArticleDraft(ctx context.Context, in *CreateArticleDraftReq, opts ...grpc.CallOption) (*CreateArticleDraftReply, error) {
	out := new(CreateArticleDraftReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/CreateArticleDraft", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) ArticleDraftMark(ctx context.Context, in *ArticleDraftMarkReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/ArticleDraftMark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) GetArticleDraftList(ctx context.Context, in *GetArticleDraftListReq, opts ...grpc.CallOption) (*GetArticleDraftListReply, error) {
	out := new(GetArticleDraftListReply)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/GetArticleDraftList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) SendArticle(ctx context.Context, in *SendArticleReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/SendArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) SetArticleAgree(ctx context.Context, in *SetArticleAgreeReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/SetArticleAgree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creationClient) SetArticleView(ctx context.Context, in *SetArticleViewReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/creation.v1.Creation/SetArticleView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreationServer is the server API for Creation service.
// All implementations must embed UnimplementedCreationServer
// for forward compatibility
type CreationServer interface {
	GetLeaderBoard(context.Context, *emptypb.Empty) (*GetLeaderBoardReply, error)
	GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListReply, error)
	GetArticleListHot(context.Context, *GetArticleListHotReq) (*GetArticleListHotReply, error)
	GetArticleStatistic(context.Context, *GetArticleStatisticReq) (*GetArticleStatisticReply, error)
	GetArticleListStatistic(context.Context, *GetArticleListStatisticReq) (*GetArticleListStatisticReply, error)
	GetLastArticleDraft(context.Context, *GetLastArticleDraftReq) (*GetLastArticleDraftReply, error)
	CreateArticle(context.Context, *CreateArticleReq) (*emptypb.Empty, error)
	CreateArticleCacheAndSearch(context.Context, *CreateArticleCacheAndSearchReq) (*emptypb.Empty, error)
	CreateArticleDraft(context.Context, *CreateArticleDraftReq) (*CreateArticleDraftReply, error)
	ArticleDraftMark(context.Context, *ArticleDraftMarkReq) (*emptypb.Empty, error)
	GetArticleDraftList(context.Context, *GetArticleDraftListReq) (*GetArticleDraftListReply, error)
	SendArticle(context.Context, *SendArticleReq) (*emptypb.Empty, error)
	SetArticleAgree(context.Context, *SetArticleAgreeReq) (*emptypb.Empty, error)
	SetArticleView(context.Context, *SetArticleViewReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedCreationServer()
}

// UnimplementedCreationServer must be embedded to have forward compatible implementations.
type UnimplementedCreationServer struct {
}

func (UnimplementedCreationServer) GetLeaderBoard(context.Context, *emptypb.Empty) (*GetLeaderBoardReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderBoard not implemented")
}
func (UnimplementedCreationServer) GetArticleList(context.Context, *GetArticleListReq) (*GetArticleListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleList not implemented")
}
func (UnimplementedCreationServer) GetArticleListHot(context.Context, *GetArticleListHotReq) (*GetArticleListHotReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleListHot not implemented")
}
func (UnimplementedCreationServer) GetArticleStatistic(context.Context, *GetArticleStatisticReq) (*GetArticleStatisticReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleStatistic not implemented")
}
func (UnimplementedCreationServer) GetArticleListStatistic(context.Context, *GetArticleListStatisticReq) (*GetArticleListStatisticReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleListStatistic not implemented")
}
func (UnimplementedCreationServer) GetLastArticleDraft(context.Context, *GetLastArticleDraftReq) (*GetLastArticleDraftReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastArticleDraft not implemented")
}
func (UnimplementedCreationServer) CreateArticle(context.Context, *CreateArticleReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedCreationServer) CreateArticleCacheAndSearch(context.Context, *CreateArticleCacheAndSearchReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleCacheAndSearch not implemented")
}
func (UnimplementedCreationServer) CreateArticleDraft(context.Context, *CreateArticleDraftReq) (*CreateArticleDraftReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleDraft not implemented")
}
func (UnimplementedCreationServer) ArticleDraftMark(context.Context, *ArticleDraftMarkReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ArticleDraftMark not implemented")
}
func (UnimplementedCreationServer) GetArticleDraftList(context.Context, *GetArticleDraftListReq) (*GetArticleDraftListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleDraftList not implemented")
}
func (UnimplementedCreationServer) SendArticle(context.Context, *SendArticleReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendArticle not implemented")
}
func (UnimplementedCreationServer) SetArticleAgree(context.Context, *SetArticleAgreeReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArticleAgree not implemented")
}
func (UnimplementedCreationServer) SetArticleView(context.Context, *SetArticleViewReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArticleView not implemented")
}
func (UnimplementedCreationServer) mustEmbedUnimplementedCreationServer() {}

// UnsafeCreationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreationServer will
// result in compilation errors.
type UnsafeCreationServer interface {
	mustEmbedUnimplementedCreationServer()
}

func RegisterCreationServer(s grpc.ServiceRegistrar, srv CreationServer) {
	s.RegisterService(&Creation_ServiceDesc, srv)
}

func _Creation_GetLeaderBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetLeaderBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetLeaderBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetLeaderBoard(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetArticleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetArticleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetArticleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetArticleList(ctx, req.(*GetArticleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetArticleListHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListHotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetArticleListHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetArticleListHot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetArticleListHot(ctx, req.(*GetArticleListHotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetArticleStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetArticleStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetArticleStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetArticleStatistic(ctx, req.(*GetArticleStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetArticleListStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleListStatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetArticleListStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetArticleListStatistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetArticleListStatistic(ctx, req.(*GetArticleListStatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetLastArticleDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastArticleDraftReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetLastArticleDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetLastArticleDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetLastArticleDraft(ctx, req.(*GetLastArticleDraftReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).CreateArticle(ctx, req.(*CreateArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_CreateArticleCacheAndSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleCacheAndSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).CreateArticleCacheAndSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/CreateArticleCacheAndSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).CreateArticleCacheAndSearch(ctx, req.(*CreateArticleCacheAndSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_CreateArticleDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleDraftReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).CreateArticleDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/CreateArticleDraft",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).CreateArticleDraft(ctx, req.(*CreateArticleDraftReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_ArticleDraftMark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ArticleDraftMarkReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).ArticleDraftMark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/ArticleDraftMark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).ArticleDraftMark(ctx, req.(*ArticleDraftMarkReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_GetArticleDraftList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleDraftListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).GetArticleDraftList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/GetArticleDraftList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).GetArticleDraftList(ctx, req.(*GetArticleDraftListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_SendArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendArticleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).SendArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/SendArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).SendArticle(ctx, req.(*SendArticleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_SetArticleAgree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArticleAgreeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).SetArticleAgree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/SetArticleAgree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).SetArticleAgree(ctx, req.(*SetArticleAgreeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Creation_SetArticleView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArticleViewReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreationServer).SetArticleView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/creation.v1.Creation/SetArticleView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreationServer).SetArticleView(ctx, req.(*SetArticleViewReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Creation_ServiceDesc is the grpc.ServiceDesc for Creation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Creation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "creation.v1.Creation",
	HandlerType: (*CreationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaderBoard",
			Handler:    _Creation_GetLeaderBoard_Handler,
		},
		{
			MethodName: "GetArticleList",
			Handler:    _Creation_GetArticleList_Handler,
		},
		{
			MethodName: "GetArticleListHot",
			Handler:    _Creation_GetArticleListHot_Handler,
		},
		{
			MethodName: "GetArticleStatistic",
			Handler:    _Creation_GetArticleStatistic_Handler,
		},
		{
			MethodName: "GetArticleListStatistic",
			Handler:    _Creation_GetArticleListStatistic_Handler,
		},
		{
			MethodName: "GetLastArticleDraft",
			Handler:    _Creation_GetLastArticleDraft_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Creation_CreateArticle_Handler,
		},
		{
			MethodName: "CreateArticleCacheAndSearch",
			Handler:    _Creation_CreateArticleCacheAndSearch_Handler,
		},
		{
			MethodName: "CreateArticleDraft",
			Handler:    _Creation_CreateArticleDraft_Handler,
		},
		{
			MethodName: "ArticleDraftMark",
			Handler:    _Creation_ArticleDraftMark_Handler,
		},
		{
			MethodName: "GetArticleDraftList",
			Handler:    _Creation_GetArticleDraftList_Handler,
		},
		{
			MethodName: "SendArticle",
			Handler:    _Creation_SendArticle_Handler,
		},
		{
			MethodName: "SetArticleAgree",
			Handler:    _Creation_SetArticleAgree_Handler,
		},
		{
			MethodName: "SetArticleView",
			Handler:    _Creation_SetArticleView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creation/service/v1/creation.proto",
}
