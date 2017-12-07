// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

package server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Resource struct {
	Type    string               `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Message *google_protobuf.Any `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetMessage() *google_protobuf.Any {
	if m != nil {
		return m.Message
	}
	return nil
}

type ResourceList struct {
	Type     string                 `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Messages []*google_protobuf.Any `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *ResourceList) Reset()                    { *m = ResourceList{} }
func (m *ResourceList) String() string            { return proto.CompactTextString(m) }
func (*ResourceList) ProtoMessage()               {}
func (*ResourceList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ResourceList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ResourceList) GetMessages() []*google_protobuf.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Query struct {
	Query string `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Query) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Query) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Resource)(nil), "server.Resource")
	proto.RegisterType((*ResourceList)(nil), "server.ResourceList")
	proto.RegisterType((*Query)(nil), "server.Query")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resources service

type ResourcesClient interface {
	Add(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Remove(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Create(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error)
	Get(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error)
	List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ResourceList, error)
	Update(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Patch(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Delete(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Apply(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Run(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error)
	Cancel(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type resourcesClient struct {
	cc *grpc.ClientConn
}

func NewResourcesClient(cc *grpc.ClientConn) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) Add(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Remove(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Create(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := grpc.Invoke(ctx, "/server.Resources/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Get(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := grpc.Invoke(ctx, "/server.Resources/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*ResourceList, error) {
	out := new(ResourceList)
	err := grpc.Invoke(ctx, "/server.Resources/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Update(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Patch(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Patch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Delete(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Apply(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Run(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*Resource, error) {
	out := new(Resource)
	err := grpc.Invoke(ctx, "/server.Resources/Run", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) Cancel(ctx context.Context, in *Resource, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/server.Resources/Cancel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Resources service

type ResourcesServer interface {
	Add(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Remove(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Create(context.Context, *Resource) (*Resource, error)
	Get(context.Context, *Resource) (*Resource, error)
	List(context.Context, *Query) (*ResourceList, error)
	Update(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Patch(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Delete(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Apply(context.Context, *Resource) (*google_protobuf1.Empty, error)
	Run(context.Context, *Resource) (*Resource, error)
	Cancel(context.Context, *Resource) (*google_protobuf1.Empty, error)
}

func RegisterResourcesServer(s *grpc.Server, srv ResourcesServer) {
	s.RegisterService(&_Resources_serviceDesc, srv)
}

func _Resources_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Add(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Remove(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Create(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Get(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).List(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Update(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Patch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Patch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Patch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Patch(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Delete(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Apply(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Run(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Resource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Resources/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).Cancel(ctx, req.(*Resource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Resources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Resources_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Resources_Remove_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Resources_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Resources_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Resources_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Resources_Update_Handler,
		},
		{
			MethodName: "Patch",
			Handler:    _Resources_Patch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Resources_Delete_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Resources_Apply_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _Resources_Run_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Resources_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resources.proto",
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xdd, 0x4e, 0xf2, 0x40,
	0x10, 0x05, 0x5a, 0xfa, 0xc1, 0x7c, 0x1a, 0xcd, 0x86, 0x18, 0xac, 0x37, 0xa4, 0x57, 0x24, 0xc4,
	0x45, 0x7e, 0x5e, 0x80, 0xa0, 0xf1, 0xc6, 0x18, 0x6d, 0xf0, 0x01, 0x0a, 0x1d, 0xab, 0x49, 0xdb,
	0xdd, 0xec, 0x6e, 0x49, 0xfa, 0x68, 0xbe, 0x9d, 0x69, 0xbb, 0xeb, 0x05, 0x94, 0x84, 0xde, 0x4d,
	0xce, 0x39, 0x73, 0xce, 0xcc, 0x81, 0x2b, 0x81, 0x92, 0x65, 0x62, 0x87, 0x92, 0x72, 0xc1, 0x14,
	0x23, 0x8e, 0x44, 0xb1, 0x47, 0xe1, 0xde, 0x46, 0x8c, 0x45, 0x31, 0x4e, 0x4b, 0x74, 0x9b, 0x7d,
	0x4e, 0x83, 0x34, 0xaf, 0x24, 0xee, 0xdd, 0x21, 0x85, 0x09, 0x57, 0x9a, 0xf4, 0x5e, 0xa1, 0xe7,
	0x6b, 0x4b, 0x42, 0xc0, 0x56, 0x39, 0xc7, 0x61, 0x7b, 0xd4, 0x1e, 0xf7, 0xfd, 0x72, 0x26, 0x14,
	0xfe, 0x25, 0x28, 0x65, 0x10, 0xe1, 0xb0, 0x33, 0x6a, 0x8f, 0xff, 0xcf, 0x07, 0xb4, 0xb2, 0xa3,
	0xc6, 0x8e, 0xae, 0xd2, 0xdc, 0x37, 0x22, 0x6f, 0x03, 0x17, 0xc6, 0xef, 0xe5, 0x5b, 0xaa, 0x5a,
	0xcf, 0x07, 0xe8, 0x69, 0xb9, 0x1c, 0x76, 0x46, 0xd6, 0x49, 0xd3, 0x3f, 0x95, 0x37, 0x83, 0xee,
	0x7b, 0x86, 0x22, 0x27, 0x03, 0x3d, 0x68, 0x3f, 0x8d, 0x12, 0xb0, 0x37, 0x45, 0x48, 0xa7, 0x0a,
	0x29, 0xe6, 0xf9, 0x8f, 0x0d, 0x7d, 0x73, 0x89, 0x24, 0x33, 0xb0, 0x56, 0x61, 0x48, 0xae, 0x69,
	0x55, 0x17, 0x35, 0x8c, 0x7b, 0x73, 0x94, 0xfc, 0x54, 0xb4, 0xe3, 0xb5, 0xc8, 0x12, 0x1c, 0x1f,
	0x13, 0xb6, 0xc7, 0x46, 0x5b, 0x14, 0x9c, 0xb5, 0xc0, 0x40, 0xd5, 0x6d, 0x1d, 0x21, 0x5e, 0x8b,
	0x4c, 0xc0, 0x7a, 0x46, 0x75, 0xa6, 0xf8, 0x1e, 0xec, 0xb2, 0xd4, 0x4b, 0xc3, 0x95, 0xef, 0xbb,
	0x83, 0x43, 0x69, 0x21, 0xaa, 0x3e, 0xf8, 0xe0, 0x61, 0xfd, 0x2d, 0xa7, 0x3f, 0x58, 0x40, 0xf7,
	0x2d, 0x50, 0xbb, 0xaf, 0xa6, 0x65, 0x3d, 0x62, 0x8c, 0xcd, 0xa3, 0x56, 0x9c, 0xc7, 0x79, 0xa3,
	0xa5, 0x09, 0x58, 0x7e, 0x96, 0x9e, 0xd9, 0xd8, 0x12, 0x9c, 0x75, 0x90, 0xee, 0x30, 0x6e, 0x12,
	0xb1, 0x75, 0x4a, 0x64, 0xf1, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xe1, 0xda, 0x7e, 0x6e, 0x03,
	0x00, 0x00,
}
