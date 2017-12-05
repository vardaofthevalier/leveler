// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

package resources

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogProducer int32

const (
	LogProducer_pipeline LogProducer = 0
	LogProducer_job      LogProducer = 1
)

var LogProducer_name = map[int32]string{
	0: "pipeline",
	1: "job",
}
var LogProducer_value = map[string]int32{
	"pipeline": 0,
	"job":      1,
}

func (x LogProducer) String() string {
	return proto.EnumName(LogProducer_name, int32(x))
}
func (LogProducer) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Query struct {
	Query string `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

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

type Loggable struct {
	Kind LogProducer `protobuf:"varint,1,opt,name=kind,enum=resources.LogProducer" json:"kind,omitempty"`
	Id   string      `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *Loggable) Reset()                    { *m = Loggable{} }
func (m *Loggable) String() string            { return proto.CompactTextString(m) }
func (*Loggable) ProtoMessage()               {}
func (*Loggable) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Loggable) GetKind() LogProducer {
	if m != nil {
		return m.Kind
	}
	return LogProducer_pipeline
}

func (m *Loggable) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Log struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Query)(nil), "resources.Query")
	proto.RegisterType((*Loggable)(nil), "resources.Loggable")
	proto.RegisterType((*Log)(nil), "resources.Log")
	proto.RegisterEnum("resources.LogProducer", LogProducer_name, LogProducer_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Resources service

type ResourcesClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	AddIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*Integration, error)
	ListIntegrations(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegrationsList, error)
	UpdateIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	RemoveIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	AddRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Repository, error)
	ListRepositories(ctx context.Context, in *Query, opts ...grpc.CallOption) (*RepositoriesList, error)
	RemoveRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Pipeline, error)
	ListPipelines(ctx context.Context, in *Query, opts ...grpc.CallOption) (*PipelinesList, error)
	RunPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Pipeline, error)
	CancelPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	ListJobs(ctx context.Context, in *Query, opts ...grpc.CallOption) (*JobsList, error)
	GetLogs(ctx context.Context, in *Loggable, opts ...grpc.CallOption) (Resources_GetLogsClient, error)
}

type resourcesClient struct {
	cc *grpc.ClientConn
}

func NewResourcesClient(cc *grpc.ClientConn) ResourcesClient {
	return &resourcesClient{cc}
}

func (c *resourcesClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/resources.Resources/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) AddIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/AddIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*Integration, error) {
	out := new(Integration)
	err := grpc.Invoke(ctx, "/resources.Resources/GetIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) ListIntegrations(ctx context.Context, in *Query, opts ...grpc.CallOption) (*IntegrationsList, error) {
	out := new(IntegrationsList)
	err := grpc.Invoke(ctx, "/resources.Resources/ListIntegrations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) UpdateIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/UpdateIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) RemoveIntegration(ctx context.Context, in *Integration, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/RemoveIntegration", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) AddRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/AddRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := grpc.Invoke(ctx, "/resources.Resources/GetRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) ListRepositories(ctx context.Context, in *Query, opts ...grpc.CallOption) (*RepositoriesList, error) {
	out := new(RepositoriesList)
	err := grpc.Invoke(ctx, "/resources.Resources/ListRepositories", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) RemoveRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/RemoveRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := grpc.Invoke(ctx, "/resources.Resources/GetPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) ListPipelines(ctx context.Context, in *Query, opts ...grpc.CallOption) (*PipelinesList, error) {
	out := new(PipelinesList)
	err := grpc.Invoke(ctx, "/resources.Resources/ListPipelines", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) RunPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Pipeline, error) {
	out := new(Pipeline)
	err := grpc.Invoke(ctx, "/resources.Resources/RunPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) CancelPipeline(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/resources.Resources/CancelPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/resources.Resources/GetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) ListJobs(ctx context.Context, in *Query, opts ...grpc.CallOption) (*JobsList, error) {
	out := new(JobsList)
	err := grpc.Invoke(ctx, "/resources.Resources/ListJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcesClient) GetLogs(ctx context.Context, in *Loggable, opts ...grpc.CallOption) (Resources_GetLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Resources_serviceDesc.Streams[0], c.cc, "/resources.Resources/GetLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourcesGetLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resources_GetLogsClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type resourcesGetLogsClient struct {
	grpc.ClientStream
}

func (x *resourcesGetLogsClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Resources service

type ResourcesServer interface {
	AddUser(context.Context, *User) (*User, error)
	AddIntegration(context.Context, *Integration) (*google_protobuf1.Empty, error)
	GetIntegration(context.Context, *Integration) (*Integration, error)
	ListIntegrations(context.Context, *Query) (*IntegrationsList, error)
	UpdateIntegration(context.Context, *Integration) (*google_protobuf1.Empty, error)
	RemoveIntegration(context.Context, *Integration) (*google_protobuf1.Empty, error)
	AddRepository(context.Context, *Repository) (*google_protobuf1.Empty, error)
	GetRepository(context.Context, *Repository) (*Repository, error)
	ListRepositories(context.Context, *Query) (*RepositoriesList, error)
	RemoveRepository(context.Context, *Repository) (*google_protobuf1.Empty, error)
	GetPipeline(context.Context, *Pipeline) (*Pipeline, error)
	ListPipelines(context.Context, *Query) (*PipelinesList, error)
	RunPipeline(context.Context, *Pipeline) (*Pipeline, error)
	CancelPipeline(context.Context, *Pipeline) (*google_protobuf1.Empty, error)
	GetJob(context.Context, *Job) (*Job, error)
	ListJobs(context.Context, *Query) (*JobsList, error)
	GetLogs(*Loggable, Resources_GetLogsServer) error
}

func RegisterResourcesServer(s *grpc.Server, srv ResourcesServer) {
	s.RegisterService(&_Resources_serviceDesc, srv)
}

func _Resources_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_AddIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).AddIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/AddIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).AddIntegration(ctx, req.(*Integration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/GetIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetIntegration(ctx, req.(*Integration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/ListIntegrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).ListIntegrations(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_UpdateIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).UpdateIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/UpdateIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).UpdateIntegration(ctx, req.(*Integration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_RemoveIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Integration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).RemoveIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/RemoveIntegration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).RemoveIntegration(ctx, req.(*Integration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_AddRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).AddRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/AddRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).AddRepository(ctx, req.(*Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/GetRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetRepository(ctx, req.(*Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_ListRepositories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).ListRepositories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/ListRepositories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).ListRepositories(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_RemoveRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).RemoveRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/RemoveRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).RemoveRepository(ctx, req.(*Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/GetPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetPipeline(ctx, req.(*Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/ListPipelines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).ListPipelines(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_RunPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).RunPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/RunPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).RunPipeline(ctx, req.(*Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_CancelPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).CancelPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/CancelPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).CancelPipeline(ctx, req.(*Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).GetJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcesServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resources.Resources/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcesServer).ListJobs(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _Resources_GetLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Loggable)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourcesServer).GetLogs(m, &resourcesGetLogsServer{stream})
}

type Resources_GetLogsServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type resourcesGetLogsServer struct {
	grpc.ServerStream
}

func (x *resourcesGetLogsServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

var _Resources_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resources.Resources",
	HandlerType: (*ResourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Resources_AddUser_Handler,
		},
		{
			MethodName: "AddIntegration",
			Handler:    _Resources_AddIntegration_Handler,
		},
		{
			MethodName: "GetIntegration",
			Handler:    _Resources_GetIntegration_Handler,
		},
		{
			MethodName: "ListIntegrations",
			Handler:    _Resources_ListIntegrations_Handler,
		},
		{
			MethodName: "UpdateIntegration",
			Handler:    _Resources_UpdateIntegration_Handler,
		},
		{
			MethodName: "RemoveIntegration",
			Handler:    _Resources_RemoveIntegration_Handler,
		},
		{
			MethodName: "AddRepository",
			Handler:    _Resources_AddRepository_Handler,
		},
		{
			MethodName: "GetRepository",
			Handler:    _Resources_GetRepository_Handler,
		},
		{
			MethodName: "ListRepositories",
			Handler:    _Resources_ListRepositories_Handler,
		},
		{
			MethodName: "RemoveRepository",
			Handler:    _Resources_RemoveRepository_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _Resources_GetPipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _Resources_ListPipelines_Handler,
		},
		{
			MethodName: "RunPipeline",
			Handler:    _Resources_RunPipeline_Handler,
		},
		{
			MethodName: "CancelPipeline",
			Handler:    _Resources_CancelPipeline_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Resources_GetJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Resources_ListJobs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogs",
			Handler:       _Resources_GetLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resources.proto",
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x86, 0x4d, 0x92, 0x06, 0x18, 0x8a, 0x63, 0x26, 0x6d, 0x44, 0x9d, 0x43, 0x2b, 0xab, 0x87,
	0x0a, 0xa9, 0xa4, 0x25, 0xa7, 0x1c, 0xaa, 0x94, 0x22, 0x8a, 0x8a, 0x7c, 0x48, 0xad, 0xe6, 0x01,
	0x6c, 0x76, 0x6a, 0xb9, 0x05, 0xaf, 0xb5, 0xbb, 0xae, 0xc4, 0x43, 0xf5, 0x1d, 0x2b, 0x8c, 0x17,
	0x16, 0xe4, 0x26, 0x55, 0xb9, 0x79, 0xbf, 0x9d, 0xf9, 0x67, 0xfe, 0x9f, 0x15, 0x70, 0x26, 0x48,
	0xf2, 0x5c, 0xcc, 0x48, 0xf6, 0x33, 0xc1, 0x15, 0xc7, 0xe6, 0x06, 0xb8, 0x90, 0x4b, 0x12, 0x6b,
	0xec, 0x76, 0x92, 0x54, 0x51, 0x2c, 0x42, 0x95, 0xf0, 0xb4, 0x44, 0x76, 0x96, 0x64, 0x34, 0x4f,
	0x52, 0x2a, 0xcf, 0x8e, 0xa0, 0x8c, 0xcb, 0x44, 0x71, 0xb1, 0x2c, 0xc9, 0x8b, 0x98, 0xf3, 0x78,
	0x4e, 0x57, 0xc5, 0x29, 0xca, 0xbf, 0x5f, 0x85, 0xa9, 0xbe, 0xba, 0xdc, 0xbf, 0xa2, 0x45, 0xa6,
	0xca, 0x4b, 0xef, 0x3d, 0x3c, 0xf9, 0x9a, 0x93, 0x58, 0xe2, 0xb3, 0xf2, 0xa3, 0x5b, 0x7b, 0x55,
	0x7b, 0xd3, 0x0c, 0x4a, 0x8a, 0x70, 0xf2, 0x6d, 0x99, 0x51, 0xf7, 0xa8, 0x80, 0xc5, 0xb7, 0xf7,
	0x19, 0x1a, 0x3e, 0x8f, 0xe3, 0x30, 0x9a, 0x13, 0xf6, 0xe0, 0xe4, 0x67, 0x92, 0xb2, 0xa2, 0xc9,
	0x1e, 0x5c, 0xf4, 0xb7, 0x16, 0x7d, 0x1e, 0xdf, 0x09, 0xce, 0xf2, 0x19, 0x89, 0xa0, 0xa8, 0x41,
	0x1b, 0x8e, 0x12, 0x56, 0x2a, 0x1d, 0x25, 0xcc, 0x7b, 0x09, 0xc7, 0x3e, 0x8f, 0xb1, 0x0b, 0xf5,
	0x05, 0x49, 0x19, 0xc6, 0x54, 0x8e, 0xd6, 0xc7, 0xde, 0x6b, 0x68, 0x19, 0x2a, 0xf8, 0x14, 0x1a,
	0x3a, 0x06, 0xc7, 0xc2, 0x3a, 0x1c, 0xff, 0xe0, 0x91, 0x53, 0x1b, 0xfc, 0x6e, 0x40, 0x33, 0xd0,
	0x63, 0xf1, 0x2d, 0xd4, 0x87, 0x8c, 0xdd, 0x4b, 0x12, 0x78, 0x66, 0x6c, 0xb3, 0x02, 0xee, 0x3e,
	0xf0, 0x2c, 0xfc, 0x04, 0xf6, 0x90, 0xb1, 0x2f, 0xdb, 0xc0, 0xd1, 0xf4, 0x60, 0x70, 0xf7, 0xa2,
	0xbf, 0x8e, 0xb1, 0xaf, 0x63, 0xec, 0x8f, 0x57, 0x31, 0xae, 0x35, 0x26, 0xa4, 0xfe, 0x4d, 0xa3,
	0x92, 0x7b, 0x16, 0x8e, 0xc0, 0xf1, 0x13, 0x69, 0x8a, 0x48, 0x74, 0x8c, 0xea, 0xe2, 0xd7, 0x70,
	0x2f, 0xab, 0xfb, 0xe5, 0xaa, 0xd5, 0xb3, 0x70, 0x0c, 0x9d, 0xfb, 0x8c, 0x85, 0x8a, 0x0e, 0xf3,
	0x33, 0x86, 0x4e, 0x40, 0x0b, 0xfe, 0xeb, 0x40, 0x99, 0x8f, 0xd0, 0x1e, 0x32, 0x16, 0x6c, 0x1e,
	0x2a, 0x3e, 0x37, 0x24, 0xb6, 0xf8, 0x01, 0x85, 0x5b, 0x68, 0x4f, 0x48, 0x3d, 0xae, 0x50, 0x8d,
	0xb7, 0xa9, 0x6e, 0x58, 0x42, 0x8f, 0xa5, 0x6a, 0x96, 0x96, 0xa9, 0x8e, 0xc0, 0x59, 0xc7, 0x71,
	0x88, 0x95, 0x1b, 0x68, 0x4d, 0x48, 0xdd, 0x95, 0xcf, 0x17, 0xcf, 0x8d, 0x7e, 0x0d, 0xdd, 0x2a,
	0xe8, 0x59, 0xf8, 0x01, 0xda, 0xab, 0x4d, 0x34, 0xa9, 0x72, 0xd0, 0xad, 0xe8, 0xd4, 0xeb, 0xdf,
	0x40, 0x2b, 0xc8, 0xd3, 0xff, 0x9a, 0x7c, 0x0b, 0xf6, 0x28, 0x4c, 0x67, 0x34, 0x7f, 0xb8, 0xfb,
	0xef, 0xae, 0x7b, 0x70, 0x3a, 0x21, 0x35, 0xe5, 0x11, 0xda, 0x46, 0xe3, 0x94, 0x47, 0xee, 0xde,
	0xd9, 0xb3, 0xf0, 0x1a, 0x1a, 0xab, 0x8d, 0xa7, 0x3c, 0xaa, 0x72, 0x78, 0xbe, 0x5b, 0xaf, 0xcd,
	0x0d, 0xa0, 0x3e, 0x21, 0xe5, 0xf3, 0x58, 0xee, 0xac, 0xa6, 0xff, 0x9e, 0x76, 0xc6, 0xf8, 0x3c,
	0xf6, 0xac, 0x77, 0xb5, 0xe8, 0xb4, 0x58, 0xf3, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0x88, 0x0d, 0x76, 0x8f, 0x05, 0x00, 0x00,
}
