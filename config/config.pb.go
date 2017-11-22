// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	ContainerPlatform
	KubernetesOptions
	Database
	RedisOptions
	SqlOptions
	ServerConfig
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContainerPlatform struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// Types that are valid to be assigned to Opts:
	//	*ContainerPlatform_KubernetesOptions
	Opts isContainerPlatform_Opts `protobuf_oneof:"opts"`
}

func (m *ContainerPlatform) Reset()                    { *m = ContainerPlatform{} }
func (m *ContainerPlatform) String() string            { return proto.CompactTextString(m) }
func (*ContainerPlatform) ProtoMessage()               {}
func (*ContainerPlatform) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isContainerPlatform_Opts interface {
	isContainerPlatform_Opts()
}

type ContainerPlatform_KubernetesOptions struct {
	KubernetesOptions *KubernetesOptions `protobuf:"bytes,4,opt,name=kubernetes_options,json=kubernetesOptions,oneof"`
}

func (*ContainerPlatform_KubernetesOptions) isContainerPlatform_Opts() {}

func (m *ContainerPlatform) GetOpts() isContainerPlatform_Opts {
	if m != nil {
		return m.Opts
	}
	return nil
}

func (m *ContainerPlatform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerPlatform) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ContainerPlatform) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ContainerPlatform) GetKubernetesOptions() *KubernetesOptions {
	if x, ok := m.GetOpts().(*ContainerPlatform_KubernetesOptions); ok {
		return x.KubernetesOptions
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ContainerPlatform) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ContainerPlatform_OneofMarshaler, _ContainerPlatform_OneofUnmarshaler, _ContainerPlatform_OneofSizer, []interface{}{
		(*ContainerPlatform_KubernetesOptions)(nil),
	}
}

func _ContainerPlatform_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ContainerPlatform)
	// opts
	switch x := m.Opts.(type) {
	case *ContainerPlatform_KubernetesOptions:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.KubernetesOptions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ContainerPlatform.Opts has unexpected type %T", x)
	}
	return nil
}

func _ContainerPlatform_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ContainerPlatform)
	switch tag {
	case 4: // opts.kubernetes_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(KubernetesOptions)
		err := b.DecodeMessage(msg)
		m.Opts = &ContainerPlatform_KubernetesOptions{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ContainerPlatform_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ContainerPlatform)
	// opts
	switch x := m.Opts.(type) {
	case *ContainerPlatform_KubernetesOptions:
		s := proto.Size(x.KubernetesOptions)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type KubernetesOptions struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *KubernetesOptions) Reset()                    { *m = KubernetesOptions{} }
func (m *KubernetesOptions) String() string            { return proto.CompactTextString(m) }
func (*KubernetesOptions) ProtoMessage()               {}
func (*KubernetesOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KubernetesOptions) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Database struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	// Types that are valid to be assigned to Opts:
	//	*Database_RedisOptions
	//	*Database_SqlOptions
	Opts     isDatabase_Opts `protobuf_oneof:"opts"`
	Protocol string          `protobuf:"bytes,6,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isDatabase_Opts interface {
	isDatabase_Opts()
}

type Database_RedisOptions struct {
	RedisOptions *RedisOptions `protobuf:"bytes,4,opt,name=redis_options,json=redisOptions,oneof"`
}
type Database_SqlOptions struct {
	SqlOptions *SqlOptions `protobuf:"bytes,5,opt,name=sql_options,json=sqlOptions,oneof"`
}

func (*Database_RedisOptions) isDatabase_Opts() {}
func (*Database_SqlOptions) isDatabase_Opts()   {}

func (m *Database) GetOpts() isDatabase_Opts {
	if m != nil {
		return m.Opts
	}
	return nil
}

func (m *Database) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Database) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Database) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Database) GetRedisOptions() *RedisOptions {
	if x, ok := m.GetOpts().(*Database_RedisOptions); ok {
		return x.RedisOptions
	}
	return nil
}

func (m *Database) GetSqlOptions() *SqlOptions {
	if x, ok := m.GetOpts().(*Database_SqlOptions); ok {
		return x.SqlOptions
	}
	return nil
}

func (m *Database) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Database) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Database_OneofMarshaler, _Database_OneofUnmarshaler, _Database_OneofSizer, []interface{}{
		(*Database_RedisOptions)(nil),
		(*Database_SqlOptions)(nil),
	}
}

func _Database_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Database)
	// opts
	switch x := m.Opts.(type) {
	case *Database_RedisOptions:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RedisOptions); err != nil {
			return err
		}
	case *Database_SqlOptions:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SqlOptions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Database.Opts has unexpected type %T", x)
	}
	return nil
}

func _Database_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Database)
	switch tag {
	case 4: // opts.redis_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RedisOptions)
		err := b.DecodeMessage(msg)
		m.Opts = &Database_RedisOptions{msg}
		return true, err
	case 5: // opts.sql_options
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SqlOptions)
		err := b.DecodeMessage(msg)
		m.Opts = &Database_SqlOptions{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Database_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Database)
	// opts
	switch x := m.Opts.(type) {
	case *Database_RedisOptions:
		s := proto.Size(x.RedisOptions)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Database_SqlOptions:
		s := proto.Size(x.SqlOptions)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RedisOptions struct {
	PoolSize int32 `protobuf:"varint,1,opt,name=poolSize" json:"poolSize,omitempty"`
}

func (m *RedisOptions) Reset()                    { *m = RedisOptions{} }
func (m *RedisOptions) String() string            { return proto.CompactTextString(m) }
func (*RedisOptions) ProtoMessage()               {}
func (*RedisOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RedisOptions) GetPoolSize() int32 {
	if m != nil {
		return m.PoolSize
	}
	return 0
}

type SqlOptions struct {
	Driver   string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	User     string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Database string `protobuf:"bytes,4,opt,name=database" json:"database,omitempty"`
}

func (m *SqlOptions) Reset()                    { *m = SqlOptions{} }
func (m *SqlOptions) String() string            { return proto.CompactTextString(m) }
func (*SqlOptions) ProtoMessage()               {}
func (*SqlOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SqlOptions) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *SqlOptions) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SqlOptions) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *SqlOptions) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type ServerConfig struct {
	Host     string             `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	Port     int32              `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Database *Database          `protobuf:"bytes,3,opt,name=database" json:"database,omitempty"`
	Platform *ContainerPlatform `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServerConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ServerConfig) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerConfig) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ServerConfig) GetPlatform() *ContainerPlatform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func init() {
	proto.RegisterType((*ContainerPlatform)(nil), "config.ContainerPlatform")
	proto.RegisterType((*KubernetesOptions)(nil), "config.KubernetesOptions")
	proto.RegisterType((*Database)(nil), "config.Database")
	proto.RegisterType((*RedisOptions)(nil), "config.RedisOptions")
	proto.RegisterType((*SqlOptions)(nil), "config.SqlOptions")
	proto.RegisterType((*ServerConfig)(nil), "config.ServerConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x3f, 0xf7, 0x4f, 0xd4, 0xde, 0xe6, 0x93, 0xa8, 0x85, 0x50, 0x40, 0x0c, 0x55, 0xa6,
	0x0a, 0xa1, 0x4a, 0x80, 0x3a, 0xb1, 0x51, 0x06, 0x04, 0x03, 0x28, 0x7d, 0x00, 0xe4, 0x36, 0x2e,
	0x44, 0x4d, 0x73, 0x5d, 0xdb, 0x2d, 0x82, 0x57, 0x61, 0xe4, 0xbd, 0x78, 0x16, 0x64, 0x27, 0x71,
	0xa2, 0x96, 0x81, 0xed, 0x9e, 0x6b, 0x5f, 0xfb, 0x9c, 0x9f, 0x0d, 0xfe, 0x1c, 0xb3, 0x45, 0xf2,
	0x32, 0x12, 0x12, 0x35, 0x52, 0x2f, 0x57, 0xe1, 0x17, 0x81, 0xfe, 0x04, 0x33, 0xcd, 0x92, 0x8c,
	0xcb, 0xa7, 0x94, 0xe9, 0x05, 0xca, 0x15, 0xa5, 0xd0, 0xca, 0xd8, 0x8a, 0x07, 0x64, 0x40, 0x86,
	0xdd, 0xc8, 0xd6, 0xa6, 0xf7, 0x8a, 0x4a, 0x07, 0x8d, 0xbc, 0x67, 0x6a, 0xd3, 0x13, 0x28, 0x75,
	0xd0, 0x1c, 0x90, 0x61, 0x3b, 0xb2, 0x35, 0xbd, 0x07, 0xba, 0xdc, 0xcc, 0xb8, 0xcc, 0xb8, 0xe6,
	0xea, 0x19, 0x85, 0x4e, 0x30, 0x53, 0x41, 0x6b, 0x40, 0x86, 0xbd, 0xcb, 0xe3, 0x51, 0x61, 0xe2,
	0xc1, 0xed, 0x78, 0xcc, 0x37, 0xdc, 0xfd, 0x8b, 0xfa, 0xcb, 0xdd, 0xe6, 0x8d, 0x07, 0x2d, 0x14,
	0x5a, 0x85, 0x17, 0xd0, 0xdf, 0x9b, 0xa0, 0xa7, 0xd0, 0x35, 0xc6, 0x94, 0x60, 0xf3, 0xd2, 0x69,
	0xd5, 0x08, 0xbf, 0x09, 0x74, 0x6e, 0x99, 0x66, 0x33, 0xa6, 0xac, 0x77, 0xfd, 0x2e, 0x5c, 0x1e,
	0x53, 0xff, 0x39, 0xcf, 0x35, 0xfc, 0x97, 0x3c, 0x4e, 0x76, 0xa3, 0x1c, 0x96, 0x51, 0x22, 0xb3,
	0x58, 0xa5, 0xf0, 0x65, 0x4d, 0xd3, 0x31, 0xf4, 0xd4, 0x3a, 0x75, 0xa3, 0x6d, 0x3b, 0x4a, 0xcb,
	0xd1, 0xe9, 0x3a, 0xad, 0x06, 0x41, 0x39, 0x45, 0x4f, 0xa0, 0x63, 0x9f, 0x69, 0x8e, 0x69, 0xe0,
	0x59, 0x7f, 0x4e, 0x3b, 0x26, 0x67, 0xe0, 0xd7, 0xaf, 0xb6, 0x33, 0x88, 0xe9, 0x34, 0xf9, 0xc8,
	0x73, 0xb6, 0x23, 0xa7, 0x43, 0x01, 0x50, 0xdd, 0x45, 0x8f, 0xc0, 0x8b, 0x65, 0xb2, 0xe5, 0xb2,
	0xe0, 0x51, 0x28, 0x93, 0x7e, 0xa3, 0xb8, 0x2c, 0x89, 0x98, 0xda, 0x9e, 0xca, 0x94, 0x7a, 0x43,
	0x19, 0x5b, 0x2a, 0xc6, 0x49, 0xa1, 0xcd, 0x5a, 0x5c, 0x10, 0xb6, 0x50, 0xba, 0x91, 0xd3, 0xe1,
	0x27, 0x01, 0x7f, 0xca, 0xe5, 0x96, 0xcb, 0x89, 0xcd, 0xea, 0x70, 0x93, 0x5f, 0x70, 0x37, 0x6a,
	0xb8, 0xcf, 0x6b, 0x87, 0x36, 0x2d, 0xae, 0x83, 0x12, 0x57, 0xf9, 0x9c, 0xd5, 0x35, 0x74, 0x0c,
	0x1d, 0x51, 0x7c, 0xda, 0xdd, 0x2f, 0xb6, 0xf7, 0xab, 0x23, 0xb7, 0x75, 0xe6, 0x59, 0x9a, 0x57,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x92, 0xc2, 0xd5, 0x14, 0x03, 0x00, 0x00,
}
