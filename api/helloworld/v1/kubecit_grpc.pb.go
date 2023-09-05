// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: helloworld/v1/kubecit.proto

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

const (
	Greeter_MostNew_FullMethodName            = "/helloworld.v1.Greeter/MostNew"
	Greeter_GetFirstCategories_FullMethodName = "/helloworld.v1.Greeter/GetFirstCategories"
	Greeter_TagsList_FullMethodName           = "/helloworld.v1.Greeter/TagsList"
	Greeter_SearchCourse_FullMethodName       = "/helloworld.v1.Greeter/SearchCourse"
	Greeter_Category_FullMethodName           = "/helloworld.v1.Greeter/Category"
	Greeter_GetSliders_FullMethodName         = "/helloworld.v1.Greeter/GetSliders"
	Greeter_GetInfo_FullMethodName            = "/helloworld.v1.Greeter/GetInfo"
	Greeter_LoginByJson_FullMethodName        = "/helloworld.v1.Greeter/LoginByJson"
	Greeter_CreateToken_FullMethodName        = "/helloworld.v1.Greeter/CreateToken"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	MostNew(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*MostNewReply, error)
	GetFirstCategories(ctx context.Context, in *GetFirstCategoriesRequest, opts ...grpc.CallOption) (*GetFirstCategoriesReply, error)
	TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error)
	SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error)
	Category(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryResp, error)
	GetSliders(ctx context.Context, in *GetSlidersRequest, opts ...grpc.CallOption) (*GetSlidersReply, error)
	GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error)
	LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error)
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) MostNew(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*MostNewReply, error) {
	out := new(MostNewReply)
	err := c.cc.Invoke(ctx, Greeter_MostNew_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetFirstCategories(ctx context.Context, in *GetFirstCategoriesRequest, opts ...grpc.CallOption) (*GetFirstCategoriesReply, error) {
	out := new(GetFirstCategoriesReply)
	err := c.cc.Invoke(ctx, Greeter_GetFirstCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) TagsList(ctx context.Context, in *TagsListRequest, opts ...grpc.CallOption) (*TagsListReply, error) {
	out := new(TagsListReply)
	err := c.cc.Invoke(ctx, Greeter_TagsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SearchCourse(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*SearchCourseReply, error) {
	out := new(SearchCourseReply)
	err := c.cc.Invoke(ctx, Greeter_SearchCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Category(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryResp, error) {
	out := new(CategoryResp)
	err := c.cc.Invoke(ctx, Greeter_Category_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetSliders(ctx context.Context, in *GetSlidersRequest, opts ...grpc.CallOption) (*GetSlidersReply, error) {
	out := new(GetSlidersReply)
	err := c.cc.Invoke(ctx, Greeter_GetSliders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoReply, error) {
	out := new(GetInfoReply)
	err := c.cc.Invoke(ctx, Greeter_GetInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LoginByJson(ctx context.Context, in *LoginByJsonRequest, opts ...grpc.CallOption) (*LoginByJsonReply, error) {
	out := new(LoginByJsonReply)
	err := c.cc.Invoke(ctx, Greeter_LoginByJson_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenReply, error) {
	out := new(CreateTokenReply)
	err := c.cc.Invoke(ctx, Greeter_CreateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	MostNew(context.Context, *PageRequest) (*MostNewReply, error)
	GetFirstCategories(context.Context, *GetFirstCategoriesRequest) (*GetFirstCategoriesReply, error)
	TagsList(context.Context, *TagsListRequest) (*TagsListReply, error)
	SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error)
	Category(context.Context, *Empty) (*CategoryResp, error)
	GetSliders(context.Context, *GetSlidersRequest) (*GetSlidersReply, error)
	GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error)
	LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error)
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) MostNew(context.Context, *PageRequest) (*MostNewReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostNew not implemented")
}
func (UnimplementedGreeterServer) GetFirstCategories(context.Context, *GetFirstCategoriesRequest) (*GetFirstCategoriesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirstCategories not implemented")
}
func (UnimplementedGreeterServer) TagsList(context.Context, *TagsListRequest) (*TagsListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TagsList not implemented")
}
func (UnimplementedGreeterServer) SearchCourse(context.Context, *SearchCourseRequest) (*SearchCourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourse not implemented")
}
func (UnimplementedGreeterServer) Category(context.Context, *Empty) (*CategoryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Category not implemented")
}
func (UnimplementedGreeterServer) GetSliders(context.Context, *GetSlidersRequest) (*GetSlidersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSliders not implemented")
}
func (UnimplementedGreeterServer) GetInfo(context.Context, *GetInfoRequest) (*GetInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedGreeterServer) LoginByJson(context.Context, *LoginByJsonRequest) (*LoginByJsonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginByJson not implemented")
}
func (UnimplementedGreeterServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_MostNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).MostNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_MostNew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).MostNew(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetFirstCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFirstCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetFirstCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_GetFirstCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetFirstCategories(ctx, req.(*GetFirstCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_TagsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).TagsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_TagsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).TagsList(ctx, req.(*TagsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SearchCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SearchCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SearchCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SearchCourse(ctx, req.(*SearchCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Category_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Category(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_Category_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Category(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetSliders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSlidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetSliders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_GetSliders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetSliders(ctx, req.(*GetSlidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_GetInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetInfo(ctx, req.(*GetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LoginByJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginByJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LoginByJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_LoginByJson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LoginByJson(ctx, req.(*LoginByJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_CreateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MostNew",
			Handler:    _Greeter_MostNew_Handler,
		},
		{
			MethodName: "GetFirstCategories",
			Handler:    _Greeter_GetFirstCategories_Handler,
		},
		{
			MethodName: "TagsList",
			Handler:    _Greeter_TagsList_Handler,
		},
		{
			MethodName: "SearchCourse",
			Handler:    _Greeter_SearchCourse_Handler,
		},
		{
			MethodName: "Category",
			Handler:    _Greeter_Category_Handler,
		},
		{
			MethodName: "GetSliders",
			Handler:    _Greeter_GetSliders_Handler,
		},
		{
			MethodName: "GetInfo",
			Handler:    _Greeter_GetInfo_Handler,
		},
		{
			MethodName: "LoginByJson",
			Handler:    _Greeter_LoginByJson_Handler,
		},
		{
			MethodName: "CreateToken",
			Handler:    _Greeter_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/kubecit.proto",
}
