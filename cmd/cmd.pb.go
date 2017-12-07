// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd.proto

/*
Package cmd is a generated protocol buffer package.

It is generated from these files:
	cmd.proto

It has these top-level messages:
	ResourceCmdConfig
	FileSource
	OptionsSource
	CmdConfig
	SubCmdConfig
	Option
*/
package cmd

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

type Operation int32

const (
	Operation_create Operation = 0
	Operation_get    Operation = 1
	Operation_list   Operation = 2
	Operation_update Operation = 3
	Operation_patch  Operation = 4
	Operation_delete Operation = 5
	Operation_apply  Operation = 6
	Operation_add    Operation = 7
	Operation_remove Operation = 8
	Operation_run    Operation = 9
	Operation_cancel Operation = 10
)

var Operation_name = map[int32]string{
	0:  "create",
	1:  "get",
	2:  "list",
	3:  "update",
	4:  "patch",
	5:  "delete",
	6:  "apply",
	7:  "add",
	8:  "remove",
	9:  "run",
	10: "cancel",
}
var Operation_value = map[string]int32{
	"create": 0,
	"get":    1,
	"list":   2,
	"update": 3,
	"patch":  4,
	"delete": 5,
	"apply":  6,
	"add":    7,
	"remove": 8,
	"run":    9,
	"cancel": 10,
}

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}
func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (x *Operation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Operation_value, data, "Operation")
	if err != nil {
		return err
	}
	*x = Operation(value)
	return nil
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ResourceCmdConfig struct {
	Version          *string      `protobuf:"bytes,1,req,name=Version" json:"Version,omitempty"`
	Resources        []*CmdConfig `protobuf:"bytes,2,rep,name=Resources" json:"Resources,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ResourceCmdConfig) Reset()                    { *m = ResourceCmdConfig{} }
func (m *ResourceCmdConfig) String() string            { return proto.CompactTextString(m) }
func (*ResourceCmdConfig) ProtoMessage()               {}
func (*ResourceCmdConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceCmdConfig) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *ResourceCmdConfig) GetResources() []*CmdConfig {
	if m != nil {
		return m.Resources
	}
	return nil
}

type FileSource struct {
	MergeOptions     []*Option `protobuf:"bytes,1,rep,name=MergeOptions" json:"MergeOptions,omitempty"`
	Required         []string  `protobuf:"bytes,2,rep,name=Required" json:"Required,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *FileSource) Reset()                    { *m = FileSource{} }
func (m *FileSource) String() string            { return proto.CompactTextString(m) }
func (*FileSource) ProtoMessage()               {}
func (*FileSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FileSource) GetMergeOptions() []*Option {
	if m != nil {
		return m.MergeOptions
	}
	return nil
}

func (m *FileSource) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

type OptionsSource struct {
	Options          []*Option       `protobuf:"bytes,1,rep,name=Options" json:"Options,omitempty"`
	Subcommands      []*SubCmdConfig `protobuf:"bytes,2,rep,name=Subcommands" json:"Subcommands,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *OptionsSource) Reset()                    { *m = OptionsSource{} }
func (m *OptionsSource) String() string            { return proto.CompactTextString(m) }
func (*OptionsSource) ProtoMessage()               {}
func (*OptionsSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OptionsSource) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *OptionsSource) GetSubcommands() []*SubCmdConfig {
	if m != nil {
		return m.Subcommands
	}
	return nil
}

type CmdConfig struct {
	Name                *string     `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Usage               *string     `protobuf:"bytes,2,req,name=Usage" json:"Usage,omitempty"`
	ShortDescription    *string     `protobuf:"bytes,3,req,name=ShortDescription" json:"ShortDescription,omitempty"`
	LongDescription     *string     `protobuf:"bytes,4,req,name=LongDescription" json:"LongDescription,omitempty"`
	SupportedOperations []Operation `protobuf:"varint,5,rep,name=SupportedOperations,enum=cmd.Operation" json:"SupportedOperations,omitempty"`
	// Types that are valid to be assigned to Source:
	//	*CmdConfig_FromFile
	//	*CmdConfig_FromOptions
	Source           isCmdConfig_Source `protobuf_oneof:"source"`
	ProtobufType     *string            `protobuf:"bytes,9,req,name=ProtobufType" json:"ProtobufType,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CmdConfig) Reset()                    { *m = CmdConfig{} }
func (m *CmdConfig) String() string            { return proto.CompactTextString(m) }
func (*CmdConfig) ProtoMessage()               {}
func (*CmdConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isCmdConfig_Source interface {
	isCmdConfig_Source()
}

type CmdConfig_FromFile struct {
	FromFile *FileSource `protobuf:"bytes,6,opt,name=FromFile,oneof"`
}
type CmdConfig_FromOptions struct {
	FromOptions *OptionsSource `protobuf:"bytes,7,opt,name=FromOptions,oneof"`
}

func (*CmdConfig_FromFile) isCmdConfig_Source()    {}
func (*CmdConfig_FromOptions) isCmdConfig_Source() {}

func (m *CmdConfig) GetSource() isCmdConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *CmdConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CmdConfig) GetUsage() string {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return ""
}

func (m *CmdConfig) GetShortDescription() string {
	if m != nil && m.ShortDescription != nil {
		return *m.ShortDescription
	}
	return ""
}

func (m *CmdConfig) GetLongDescription() string {
	if m != nil && m.LongDescription != nil {
		return *m.LongDescription
	}
	return ""
}

func (m *CmdConfig) GetSupportedOperations() []Operation {
	if m != nil {
		return m.SupportedOperations
	}
	return nil
}

func (m *CmdConfig) GetFromFile() *FileSource {
	if x, ok := m.GetSource().(*CmdConfig_FromFile); ok {
		return x.FromFile
	}
	return nil
}

func (m *CmdConfig) GetFromOptions() *OptionsSource {
	if x, ok := m.GetSource().(*CmdConfig_FromOptions); ok {
		return x.FromOptions
	}
	return nil
}

func (m *CmdConfig) GetProtobufType() string {
	if m != nil && m.ProtobufType != nil {
		return *m.ProtobufType
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CmdConfig) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CmdConfig_OneofMarshaler, _CmdConfig_OneofUnmarshaler, _CmdConfig_OneofSizer, []interface{}{
		(*CmdConfig_FromFile)(nil),
		(*CmdConfig_FromOptions)(nil),
	}
}

func _CmdConfig_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CmdConfig)
	// source
	switch x := m.Source.(type) {
	case *CmdConfig_FromFile:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FromFile); err != nil {
			return err
		}
	case *CmdConfig_FromOptions:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FromOptions); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CmdConfig.Source has unexpected type %T", x)
	}
	return nil
}

func _CmdConfig_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CmdConfig)
	switch tag {
	case 6: // source.FromFile
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileSource)
		err := b.DecodeMessage(msg)
		m.Source = &CmdConfig_FromFile{msg}
		return true, err
	case 7: // source.FromOptions
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OptionsSource)
		err := b.DecodeMessage(msg)
		m.Source = &CmdConfig_FromOptions{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CmdConfig_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CmdConfig)
	// source
	switch x := m.Source.(type) {
	case *CmdConfig_FromFile:
		s := proto.Size(x.FromFile)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CmdConfig_FromOptions:
		s := proto.Size(x.FromOptions)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SubCmdConfig struct {
	Name             *string   `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Usage            *string   `protobuf:"bytes,2,req,name=Usage" json:"Usage,omitempty"`
	ShortDescription *string   `protobuf:"bytes,3,req,name=ShortDescription" json:"ShortDescription,omitempty"`
	LongDescription  *string   `protobuf:"bytes,4,req,name=LongDescription" json:"LongDescription,omitempty"`
	Options          []*Option `protobuf:"bytes,5,rep,name=Options" json:"Options,omitempty"`
	ProtobufType     *string   `protobuf:"bytes,6,req,name=ProtobufType" json:"ProtobufType,omitempty"`
	ParentField      *string   `protobuf:"bytes,7,req,name=ParentField" json:"ParentField,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SubCmdConfig) Reset()                    { *m = SubCmdConfig{} }
func (m *SubCmdConfig) String() string            { return proto.CompactTextString(m) }
func (*SubCmdConfig) ProtoMessage()               {}
func (*SubCmdConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SubCmdConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SubCmdConfig) GetUsage() string {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return ""
}

func (m *SubCmdConfig) GetShortDescription() string {
	if m != nil && m.ShortDescription != nil {
		return *m.ShortDescription
	}
	return ""
}

func (m *SubCmdConfig) GetLongDescription() string {
	if m != nil && m.LongDescription != nil {
		return *m.LongDescription
	}
	return ""
}

func (m *SubCmdConfig) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *SubCmdConfig) GetProtobufType() string {
	if m != nil && m.ProtobufType != nil {
		return *m.ProtobufType
	}
	return ""
}

func (m *SubCmdConfig) GetParentField() string {
	if m != nil && m.ParentField != nil {
		return *m.ParentField
	}
	return ""
}

type Option struct {
	Name             *string  `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Required         []string `protobuf:"bytes,2,rep,name=Required" json:"Required,omitempty"`
	Description      *string  `protobuf:"bytes,3,req,name=Description" json:"Description,omitempty"`
	Type             *string  `protobuf:"bytes,4,req,name=Type" json:"Type,omitempty"`
	Default          *string  `protobuf:"bytes,5,opt,name=Default" json:"Default,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Option) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Option) GetRequired() []string {
	if m != nil {
		return m.Required
	}
	return nil
}

func (m *Option) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Option) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Option) GetDefault() string {
	if m != nil && m.Default != nil {
		return *m.Default
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceCmdConfig)(nil), "cmd.ResourceCmdConfig")
	proto.RegisterType((*FileSource)(nil), "cmd.FileSource")
	proto.RegisterType((*OptionsSource)(nil), "cmd.OptionsSource")
	proto.RegisterType((*CmdConfig)(nil), "cmd.CmdConfig")
	proto.RegisterType((*SubCmdConfig)(nil), "cmd.SubCmdConfig")
	proto.RegisterType((*Option)(nil), "cmd.Option")
	proto.RegisterEnum("cmd.Operation", Operation_name, Operation_value)
}

func init() { proto.RegisterFile("cmd.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x9a, 0x26, 0x69, 0x5e, 0xc6, 0xe6, 0x19, 0x0e, 0x11, 0xa7, 0x28, 0x12, 0x52, 0x35,
	0xc1, 0x90, 0x86, 0xc4, 0x19, 0xb1, 0xa9, 0xe2, 0x00, 0x6c, 0x4a, 0x01, 0x09, 0x71, 0x72, 0xe3,
	0xd7, 0x2e, 0x22, 0x89, 0x8d, 0xe3, 0x20, 0xed, 0xca, 0x89, 0xdf, 0xc9, 0xdf, 0xe0, 0x82, 0xec,
	0xb4, 0x69, 0xb6, 0x4e, 0xbd, 0x72, 0xf3, 0xfb, 0xde, 0xf7, 0x3e, 0xfb, 0x7b, 0xef, 0x19, 0xc2,
	0xbc, 0xe2, 0x67, 0x52, 0x09, 0x2d, 0xa8, 0x9b, 0x57, 0x3c, 0xfd, 0x06, 0x27, 0x19, 0x36, 0xa2,
	0x55, 0x39, 0x5e, 0x54, 0xfc, 0x42, 0xd4, 0xcb, 0x62, 0x45, 0x63, 0x08, 0xbe, 0xa0, 0x6a, 0x0a,
	0x51, 0xc7, 0x4e, 0x32, 0x9a, 0x86, 0xd9, 0x26, 0xa4, 0xcf, 0x21, 0xdc, 0xd0, 0x9b, 0x78, 0x94,
	0xb8, 0xd3, 0xe8, 0xfc, 0xe8, 0xcc, 0x48, 0xf6, 0xc5, 0xd9, 0x96, 0x90, 0x7e, 0x05, 0x98, 0x15,
	0x25, 0xce, 0x6d, 0x48, 0x5f, 0xc2, 0xe1, 0x07, 0x54, 0x2b, 0xbc, 0x92, 0xba, 0x10, 0x75, 0x13,
	0x3b, 0xb6, 0x3c, 0xb2, 0xe5, 0x1d, 0x96, 0xdd, 0x21, 0xd0, 0xa7, 0x30, 0xc9, 0xf0, 0x47, 0x5b,
	0x28, 0xe4, 0xf6, 0xae, 0x30, 0xeb, 0xe3, 0xf4, 0x3b, 0x3c, 0x5a, 0xd3, 0xd6, 0xea, 0xcf, 0x20,
	0xd8, 0x23, 0xbc, 0xc9, 0xd1, 0x57, 0x10, 0xcd, 0xdb, 0x45, 0x2e, 0xaa, 0x8a, 0xd5, 0x7c, 0x63,
	0xe1, 0xc4, 0x52, 0xe7, 0xed, 0x62, 0xeb, 0x62, 0xc8, 0x4a, 0xff, 0x8c, 0x20, 0xdc, 0x76, 0x87,
	0xc2, 0xf8, 0x23, 0xab, 0x70, 0xdd, 0x1a, 0x7b, 0xa6, 0x4f, 0xc0, 0xfb, 0xdc, 0xb0, 0x15, 0xc6,
	0x23, 0x0b, 0x76, 0x01, 0x3d, 0x05, 0x32, 0xbf, 0x11, 0x4a, 0x5f, 0x62, 0x93, 0xab, 0xc2, 0xbe,
	0x20, 0x76, 0x2d, 0x61, 0x07, 0xa7, 0x53, 0x38, 0x7e, 0x2f, 0xea, 0xd5, 0x90, 0x3a, 0xb6, 0xd4,
	0xfb, 0x30, 0x7d, 0x03, 0x8f, 0xe7, 0xad, 0x94, 0x42, 0x69, 0xe4, 0x57, 0x12, 0x15, 0xeb, 0x5c,
	0x7b, 0x89, 0x3b, 0x3d, 0x5a, 0x4f, 0xa3, 0x87, 0xb3, 0x87, 0xa8, 0xf4, 0x05, 0x4c, 0x66, 0x4a,
	0x54, 0x66, 0x36, 0xb1, 0x9f, 0x38, 0xd3, 0xe8, 0xfc, 0xd8, 0x96, 0x6d, 0x87, 0xf5, 0xee, 0x20,
	0xeb, 0x29, 0xf4, 0x35, 0x44, 0xe6, 0xbc, 0x69, 0x6f, 0x60, 0x2b, 0xe8, 0xa0, 0xbd, 0x4d, 0x5f,
	0x34, 0x24, 0xd2, 0x14, 0x0e, 0xaf, 0xcd, 0xa6, 0x2d, 0xda, 0xe5, 0xa7, 0x5b, 0x89, 0x71, 0x68,
	0xfd, 0xdc, 0xc1, 0xde, 0x4e, 0xc0, 0xef, 0xb6, 0x25, 0xfd, 0xeb, 0xc0, 0xe1, 0x70, 0x04, 0xff,
	0xbd, 0xcf, 0x83, 0x8d, 0xf2, 0xf6, 0x6c, 0xd4, 0x7d, 0x97, 0xfe, 0xae, 0x4b, 0x9a, 0x40, 0x74,
	0xcd, 0x14, 0xd6, 0x7a, 0x56, 0x60, 0xc9, 0xe3, 0xc0, 0x52, 0x86, 0x50, 0xfa, 0xdb, 0x01, 0xbf,
	0x53, 0x7c, 0xd0, 0xf7, 0x9e, 0xaf, 0x60, 0xc4, 0x77, 0x8d, 0x0f, 0x21, 0xa3, 0x68, 0x9f, 0xd6,
	0x19, 0xb5, 0x67, 0xf3, 0xc7, 0x2f, 0x71, 0xc9, 0xda, 0x52, 0xc7, 0x5e, 0xe2, 0x98, 0x3f, 0xbe,
	0x0e, 0x4f, 0x7f, 0x39, 0x10, 0xf6, 0xcb, 0x42, 0x01, 0xfc, 0x5c, 0x21, 0xd3, 0x48, 0x0e, 0x68,
	0x00, 0xee, 0x0a, 0x35, 0x71, 0xe8, 0x04, 0xc6, 0x65, 0xd1, 0x68, 0x32, 0x32, 0xe9, 0x56, 0x72,
	0x93, 0x76, 0x69, 0x08, 0x9e, 0x64, 0x3a, 0xbf, 0x21, 0x63, 0x03, 0x73, 0x2c, 0x51, 0x23, 0xf1,
	0x0c, 0xcc, 0xa4, 0x2c, 0x6f, 0x89, 0x6f, 0x04, 0x18, 0xe7, 0x24, 0x30, 0x79, 0x85, 0x95, 0xf8,
	0x89, 0x64, 0x62, 0x40, 0xd5, 0xd6, 0x24, 0xb4, 0x57, 0xb1, 0x3a, 0xc7, 0x92, 0xc0, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x22, 0xa6, 0x92, 0x0c, 0xa8, 0x04, 0x00, 0x00,
}