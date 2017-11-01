// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

package leveler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RoleId struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *RoleId) Reset()                    { *m = RoleId{} }
func (m *RoleId) String() string            { return proto.CompactTextString(m) }
func (*RoleId) ProtoMessage()               {}
func (*RoleId) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *RoleId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Role struct {
	Name         string         `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Description  string         `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
	Requirements []*Requirement `protobuf:"bytes,3,rep,name=Requirements" json:"Requirements,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Role) GetRequirements() []*Requirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type RoleList struct {
	Results []*Role `protobuf:"bytes,1,rep,name=Results" json:"Results,omitempty"`
}

func (m *RoleList) Reset()                    { *m = RoleList{} }
func (m *RoleList) String() string            { return proto.CompactTextString(m) }
func (*RoleList) ProtoMessage()               {}
func (*RoleList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *RoleList) GetResults() []*Role {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleId)(nil), "leveler.RoleId")
	proto.RegisterType((*Role)(nil), "leveler.Role")
	proto.RegisterType((*RoleList)(nil), "leveler.RoleList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RoleEndpoint service

type RoleEndpointClient interface {
	GetRole(ctx context.Context, in *RoleId, opts ...grpc.CallOption) (*Role, error)
	ListRoles(ctx context.Context, in *Query, opts ...grpc.CallOption) (*RoleList, error)
	CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	DeleteRole(ctx context.Context, in *RoleId, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type roleEndpointClient struct {
	cc *grpc.ClientConn
}

func NewRoleEndpointClient(cc *grpc.ClientConn) RoleEndpointClient {
	return &roleEndpointClient{cc}
}

func (c *roleEndpointClient) GetRole(ctx context.Context, in *RoleId, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/leveler.RoleEndpoint/GetRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleEndpointClient) ListRoles(ctx context.Context, in *Query, opts ...grpc.CallOption) (*RoleList, error) {
	out := new(RoleList)
	err := grpc.Invoke(ctx, "/leveler.RoleEndpoint/ListRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleEndpointClient) CreateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/leveler.RoleEndpoint/CreateRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleEndpointClient) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := grpc.Invoke(ctx, "/leveler.RoleEndpoint/UpdateRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleEndpointClient) DeleteRole(ctx context.Context, in *RoleId, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/leveler.RoleEndpoint/DeleteRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoleEndpoint service

type RoleEndpointServer interface {
	GetRole(context.Context, *RoleId) (*Role, error)
	ListRoles(context.Context, *Query) (*RoleList, error)
	CreateRole(context.Context, *Role) (*Role, error)
	UpdateRole(context.Context, *Role) (*Role, error)
	DeleteRole(context.Context, *RoleId) (*google_protobuf.Empty, error)
}

func RegisterRoleEndpointServer(s *grpc.Server, srv RoleEndpointServer) {
	s.RegisterService(&_RoleEndpoint_serviceDesc, srv)
}

func _RoleEndpoint_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleEndpointServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leveler.RoleEndpoint/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleEndpointServer).GetRole(ctx, req.(*RoleId))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleEndpoint_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleEndpointServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leveler.RoleEndpoint/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleEndpointServer).ListRoles(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleEndpoint_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleEndpointServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leveler.RoleEndpoint/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleEndpointServer).CreateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleEndpoint_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleEndpointServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leveler.RoleEndpoint/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleEndpointServer).UpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleEndpoint_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleEndpointServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leveler.RoleEndpoint/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleEndpointServer).DeleteRole(ctx, req.(*RoleId))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoleEndpoint_serviceDesc = grpc.ServiceDesc{
	ServiceName: "leveler.RoleEndpoint",
	HandlerType: (*RoleEndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRole",
			Handler:    _RoleEndpoint_GetRole_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _RoleEndpoint_ListRoles_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RoleEndpoint_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleEndpoint_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _RoleEndpoint_DeleteRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}

func init() { proto.RegisterFile("role.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xfb, 0x8f, 0xee, 0xdb, 0x69, 0x5f, 0xa5, 0x41, 0x64, 0x59, 0x3d, 0x2c, 0xb9, 0x58,
	0xb0, 0xa4, 0xd2, 0x1e, 0xf4, 0x6e, 0x8b, 0x14, 0x44, 0x30, 0xe0, 0x07, 0x68, 0xdd, 0xb1, 0x04,
	0xb2, 0x9b, 0x35, 0xc9, 0x16, 0x7a, 0xf5, 0x93, 0x4b, 0xb2, 0x5b, 0xdb, 0xa5, 0x1e, 0xbc, 0xcd,
	0xcc, 0xef, 0xc9, 0x33, 0x13, 0x1e, 0x00, 0xad, 0x24, 0xb2, 0x5c, 0x2b, 0xab, 0x48, 0x20, 0x71,
	0x8b, 0x12, 0x75, 0x34, 0xd4, 0xf8, 0x59, 0x08, 0x8d, 0x29, 0x66, 0xb6, 0x64, 0x11, 0x14, 0x56,
	0xc8, 0xaa, 0xbe, 0xda, 0x28, 0xb5, 0x91, 0x38, 0xf1, 0xdd, 0xba, 0xf8, 0x98, 0x60, 0x9a, 0xdb,
	0x5d, 0x09, 0xe9, 0x35, 0x74, 0xb9, 0x92, 0xb8, 0x4c, 0x08, 0x81, 0xce, 0xcb, 0x2a, 0xc5, 0xb0,
	0x19, 0x37, 0x47, 0x3d, 0xee, 0x6b, 0xba, 0x85, 0x8e, 0xa3, 0xbf, 0x31, 0x12, 0x43, 0x7f, 0x8e,
	0xe6, 0x5d, 0x8b, 0xdc, 0x0a, 0x95, 0x85, 0x2d, 0x8f, 0x8e, 0x47, 0xe4, 0x01, 0x06, 0xfc, 0x70,
	0x99, 0x09, 0xdb, 0x71, 0x7b, 0xd4, 0x9f, 0x5e, 0xb0, 0xea, 0x6e, 0x76, 0x04, 0x79, 0x4d, 0x49,
	0x67, 0xf0, 0xcf, 0xed, 0x7d, 0x16, 0xc6, 0x92, 0x1b, 0x08, 0x38, 0x9a, 0x42, 0x5a, 0x13, 0x36,
	0xbd, 0xc1, 0xff, 0x83, 0x81, 0x92, 0xc8, 0xf7, 0x74, 0xfa, 0xd5, 0x82, 0x81, 0x9b, 0x2c, 0xb2,
	0x24, 0x57, 0x22, 0xb3, 0xe4, 0x16, 0x82, 0x27, 0xb4, 0xfe, 0x03, 0xe7, 0xb5, 0x37, 0xcb, 0x24,
	0xaa, 0x9b, 0xd0, 0x06, 0xb9, 0x83, 0x9e, 0x5b, 0xe7, 0x3a, 0x43, 0xce, 0x7e, 0xe8, 0x6b, 0x81,
	0x7a, 0x17, 0x0d, 0x6b, 0x6a, 0xa7, 0xa3, 0x0d, 0x32, 0x06, 0x78, 0xd4, 0xb8, 0xb2, 0xe8, 0x37,
	0xd4, 0x0d, 0x4f, 0xfd, 0xc7, 0x00, 0x6f, 0x79, 0xf2, 0x57, 0xf5, 0x3d, 0xc0, 0x1c, 0x25, 0x56,
	0xea, 0x93, 0xeb, 0x2f, 0x59, 0x99, 0x29, 0xdb, 0x67, 0xca, 0x16, 0x2e, 0x53, 0xda, 0x58, 0x77,
	0xfd, 0x64, 0xf6, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xe4, 0xed, 0xdd, 0x29, 0x02, 0x00, 0x00,
}