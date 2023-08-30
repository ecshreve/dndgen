// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: entpb/entpb.proto

package entpb

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

// AbilityScoreServiceClient is the client API for AbilityScoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AbilityScoreServiceClient interface {
	Create(ctx context.Context, in *CreateAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error)
	Get(ctx context.Context, in *GetAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error)
	Update(ctx context.Context, in *UpdateAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error)
	Delete(ctx context.Context, in *DeleteAbilityScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ListAbilityScoreRequest, opts ...grpc.CallOption) (*ListAbilityScoreResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateAbilityScoresRequest, opts ...grpc.CallOption) (*BatchCreateAbilityScoresResponse, error)
}

type abilityScoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAbilityScoreServiceClient(cc grpc.ClientConnInterface) AbilityScoreServiceClient {
	return &abilityScoreServiceClient{cc}
}

func (c *abilityScoreServiceClient) Create(ctx context.Context, in *CreateAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error) {
	out := new(AbilityScore)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abilityScoreServiceClient) Get(ctx context.Context, in *GetAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error) {
	out := new(AbilityScore)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abilityScoreServiceClient) Update(ctx context.Context, in *UpdateAbilityScoreRequest, opts ...grpc.CallOption) (*AbilityScore, error) {
	out := new(AbilityScore)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abilityScoreServiceClient) Delete(ctx context.Context, in *DeleteAbilityScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abilityScoreServiceClient) List(ctx context.Context, in *ListAbilityScoreRequest, opts ...grpc.CallOption) (*ListAbilityScoreResponse, error) {
	out := new(ListAbilityScoreResponse)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *abilityScoreServiceClient) BatchCreate(ctx context.Context, in *BatchCreateAbilityScoresRequest, opts ...grpc.CallOption) (*BatchCreateAbilityScoresResponse, error) {
	out := new(BatchCreateAbilityScoresResponse)
	err := c.cc.Invoke(ctx, "/entpb.AbilityScoreService/BatchCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AbilityScoreServiceServer is the server API for AbilityScoreService service.
// All implementations must embed UnimplementedAbilityScoreServiceServer
// for forward compatibility
type AbilityScoreServiceServer interface {
	Create(context.Context, *CreateAbilityScoreRequest) (*AbilityScore, error)
	Get(context.Context, *GetAbilityScoreRequest) (*AbilityScore, error)
	Update(context.Context, *UpdateAbilityScoreRequest) (*AbilityScore, error)
	Delete(context.Context, *DeleteAbilityScoreRequest) (*emptypb.Empty, error)
	List(context.Context, *ListAbilityScoreRequest) (*ListAbilityScoreResponse, error)
	BatchCreate(context.Context, *BatchCreateAbilityScoresRequest) (*BatchCreateAbilityScoresResponse, error)
	mustEmbedUnimplementedAbilityScoreServiceServer()
}

// UnimplementedAbilityScoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAbilityScoreServiceServer struct {
}

func (UnimplementedAbilityScoreServiceServer) Create(context.Context, *CreateAbilityScoreRequest) (*AbilityScore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAbilityScoreServiceServer) Get(context.Context, *GetAbilityScoreRequest) (*AbilityScore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAbilityScoreServiceServer) Update(context.Context, *UpdateAbilityScoreRequest) (*AbilityScore, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAbilityScoreServiceServer) Delete(context.Context, *DeleteAbilityScoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAbilityScoreServiceServer) List(context.Context, *ListAbilityScoreRequest) (*ListAbilityScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAbilityScoreServiceServer) BatchCreate(context.Context, *BatchCreateAbilityScoresRequest) (*BatchCreateAbilityScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedAbilityScoreServiceServer) mustEmbedUnimplementedAbilityScoreServiceServer() {}

// UnsafeAbilityScoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AbilityScoreServiceServer will
// result in compilation errors.
type UnsafeAbilityScoreServiceServer interface {
	mustEmbedUnimplementedAbilityScoreServiceServer()
}

func RegisterAbilityScoreServiceServer(s grpc.ServiceRegistrar, srv AbilityScoreServiceServer) {
	s.RegisterService(&AbilityScoreService_ServiceDesc, srv)
}

func _AbilityScoreService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAbilityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).Create(ctx, req.(*CreateAbilityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbilityScoreService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAbilityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).Get(ctx, req.(*GetAbilityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbilityScoreService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAbilityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).Update(ctx, req.(*UpdateAbilityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbilityScoreService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAbilityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).Delete(ctx, req.(*DeleteAbilityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbilityScoreService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAbilityScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).List(ctx, req.(*ListAbilityScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AbilityScoreService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateAbilityScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AbilityScoreServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.AbilityScoreService/BatchCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AbilityScoreServiceServer).BatchCreate(ctx, req.(*BatchCreateAbilityScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AbilityScoreService_ServiceDesc is the grpc.ServiceDesc for AbilityScoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AbilityScoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.AbilityScoreService",
	HandlerType: (*AbilityScoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AbilityScoreService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AbilityScoreService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AbilityScoreService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AbilityScoreService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AbilityScoreService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _AbilityScoreService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}

// SkillServiceClient is the client API for SkillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkillServiceClient interface {
	Create(ctx context.Context, in *CreateSkillRequest, opts ...grpc.CallOption) (*Skill, error)
	Get(ctx context.Context, in *GetSkillRequest, opts ...grpc.CallOption) (*Skill, error)
	Update(ctx context.Context, in *UpdateSkillRequest, opts ...grpc.CallOption) (*Skill, error)
	Delete(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ListSkillRequest, opts ...grpc.CallOption) (*ListSkillResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateSkillsRequest, opts ...grpc.CallOption) (*BatchCreateSkillsResponse, error)
}

type skillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSkillServiceClient(cc grpc.ClientConnInterface) SkillServiceClient {
	return &skillServiceClient{cc}
}

func (c *skillServiceClient) Create(ctx context.Context, in *CreateSkillRequest, opts ...grpc.CallOption) (*Skill, error) {
	out := new(Skill)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) Get(ctx context.Context, in *GetSkillRequest, opts ...grpc.CallOption) (*Skill, error) {
	out := new(Skill)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) Update(ctx context.Context, in *UpdateSkillRequest, opts ...grpc.CallOption) (*Skill, error) {
	out := new(Skill)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) Delete(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) List(ctx context.Context, in *ListSkillRequest, opts ...grpc.CallOption) (*ListSkillResponse, error) {
	out := new(ListSkillResponse)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *skillServiceClient) BatchCreate(ctx context.Context, in *BatchCreateSkillsRequest, opts ...grpc.CallOption) (*BatchCreateSkillsResponse, error) {
	out := new(BatchCreateSkillsResponse)
	err := c.cc.Invoke(ctx, "/entpb.SkillService/BatchCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkillServiceServer is the server API for SkillService service.
// All implementations must embed UnimplementedSkillServiceServer
// for forward compatibility
type SkillServiceServer interface {
	Create(context.Context, *CreateSkillRequest) (*Skill, error)
	Get(context.Context, *GetSkillRequest) (*Skill, error)
	Update(context.Context, *UpdateSkillRequest) (*Skill, error)
	Delete(context.Context, *DeleteSkillRequest) (*emptypb.Empty, error)
	List(context.Context, *ListSkillRequest) (*ListSkillResponse, error)
	BatchCreate(context.Context, *BatchCreateSkillsRequest) (*BatchCreateSkillsResponse, error)
	mustEmbedUnimplementedSkillServiceServer()
}

// UnimplementedSkillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSkillServiceServer struct {
}

func (UnimplementedSkillServiceServer) Create(context.Context, *CreateSkillRequest) (*Skill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSkillServiceServer) Get(context.Context, *GetSkillRequest) (*Skill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSkillServiceServer) Update(context.Context, *UpdateSkillRequest) (*Skill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSkillServiceServer) Delete(context.Context, *DeleteSkillRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSkillServiceServer) List(context.Context, *ListSkillRequest) (*ListSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSkillServiceServer) BatchCreate(context.Context, *BatchCreateSkillsRequest) (*BatchCreateSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreate not implemented")
}
func (UnimplementedSkillServiceServer) mustEmbedUnimplementedSkillServiceServer() {}

// UnsafeSkillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SkillServiceServer will
// result in compilation errors.
type UnsafeSkillServiceServer interface {
	mustEmbedUnimplementedSkillServiceServer()
}

func RegisterSkillServiceServer(s grpc.ServiceRegistrar, srv SkillServiceServer) {
	s.RegisterService(&SkillService_ServiceDesc, srv)
}

func _SkillService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).Create(ctx, req.(*CreateSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).Get(ctx, req.(*GetSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).Update(ctx, req.(*UpdateSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).Delete(ctx, req.(*DeleteSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).List(ctx, req.(*ListSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SkillService_BatchCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchCreateSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkillServiceServer).BatchCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entpb.SkillService/BatchCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkillServiceServer).BatchCreate(ctx, req.(*BatchCreateSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SkillService_ServiceDesc is the grpc.ServiceDesc for SkillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "entpb.SkillService",
	HandlerType: (*SkillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SkillService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SkillService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SkillService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SkillService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SkillService_List_Handler,
		},
		{
			MethodName: "BatchCreate",
			Handler:    _SkillService_BatchCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entpb/entpb.proto",
}
