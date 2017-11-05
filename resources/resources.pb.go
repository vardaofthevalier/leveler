// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resources.proto

/*
Package leveler is a generated protocol buffer package.

It is generated from these files:
	resources.proto

It has these top-level messages:
	ResourceConfig
	Resource
	Option
*/
package leveler

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

type ResourceConfig struct {
	Resources        []*Resource `protobuf:"bytes,1,rep,name=Resources" json:"Resources,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ResourceConfig) Reset()                    { *m = ResourceConfig{} }
func (m *ResourceConfig) String() string            { return proto.CompactTextString(m) }
func (*ResourceConfig) ProtoMessage()               {}
func (*ResourceConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResourceConfig) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type Resource struct {
	Name             *string   `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Usage            *string   `protobuf:"bytes,2,req,name=Usage" json:"Usage,omitempty"`
	ShortDescription *string   `protobuf:"bytes,3,req,name=ShortDescription" json:"ShortDescription,omitempty"`
	LongDescription  *string   `protobuf:"bytes,4,req,name=LongDescription" json:"LongDescription,omitempty"`
	Options          []*Option `protobuf:"bytes,5,rep,name=Options" json:"Options,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resource) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resource) GetUsage() string {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return ""
}

func (m *Resource) GetShortDescription() string {
	if m != nil && m.ShortDescription != nil {
		return *m.ShortDescription
	}
	return ""
}

func (m *Resource) GetLongDescription() string {
	if m != nil && m.LongDescription != nil {
		return *m.LongDescription
	}
	return ""
}

func (m *Resource) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type Option struct {
	Name             *string `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	ShortName        *string `protobuf:"bytes,2,req,name=ShortName" json:"ShortName,omitempty"`
	Required         *bool   `protobuf:"varint,3,req,name=Required" json:"Required,omitempty"`
	Type             *string `protobuf:"bytes,4,req,name=Type" json:"Type,omitempty"`
	Description      *string `protobuf:"bytes,5,req,name=Description" json:"Description,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Option) Reset()                    { *m = Option{} }
func (m *Option) String() string            { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()               {}
func (*Option) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Option) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Option) GetShortName() string {
	if m != nil && m.ShortName != nil {
		return *m.ShortName
	}
	return ""
}

func (m *Option) GetRequired() bool {
	if m != nil && m.Required != nil {
		return *m.Required
	}
	return false
}

func (m *Option) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Option) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceConfig)(nil), "leveler.ResourceConfig")
	proto.RegisterType((*Resource)(nil), "leveler.Resource")
	proto.RegisterType((*Option)(nil), "leveler.Option")
}

func init() { proto.RegisterFile("resources.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0xc9, 0xcc, 0xd4, 0x99, 0xbe, 0x01, 0xab, 0x0f, 0x17, 0x41, 0x5c, 0x94, 0xae, 0xaa,
	0x8b, 0x0a, 0xde, 0x40, 0x74, 0x29, 0x0a, 0x51, 0x0f, 0x30, 0x8c, 0xcf, 0x1a, 0x18, 0x9b, 0x9a,
	0xb4, 0x82, 0x77, 0xf0, 0x3a, 0xde, 0x4f, 0xfa, 0xd2, 0xb4, 0x45, 0x67, 0xf7, 0xbf, 0xef, 0xff,
	0x09, 0x5f, 0x20, 0xb1, 0xe4, 0x4c, 0x6b, 0xb7, 0xe4, 0x8a, 0xda, 0x9a, 0xc6, 0xe0, 0x72, 0x47,
	0x9f, 0xb4, 0x23, 0x9b, 0x5d, 0xc3, 0xa1, 0xea, 0xbb, 0x1b, 0x53, 0xbd, 0xea, 0x12, 0x2f, 0x21,
	0x0e, 0xc4, 0x49, 0x91, 0xce, 0xf3, 0xf5, 0xd5, 0x71, 0xd1, 0xcf, 0x8b, 0xd0, 0xa8, 0x71, 0x93,
	0xfd, 0x08, 0x58, 0x85, 0x0b, 0x11, 0x16, 0xf7, 0x9b, 0x77, 0x92, 0x22, 0x9d, 0xe5, 0xb1, 0xe2,
	0x8c, 0x27, 0x10, 0x3d, 0xbb, 0x4d, 0x49, 0x72, 0xc6, 0xd0, 0x1f, 0x78, 0x01, 0x47, 0x8f, 0x6f,
	0xc6, 0x36, 0xb7, 0xe4, 0xb6, 0x56, 0xd7, 0x8d, 0x36, 0x95, 0x9c, 0xf3, 0xe0, 0x1f, 0xc7, 0x1c,
	0x92, 0x3b, 0x53, 0x95, 0xd3, 0xe9, 0x82, 0xa7, 0x7f, 0x31, 0x9e, 0xc3, 0xf2, 0x81, 0x93, 0x93,
	0x11, 0xbb, 0x27, 0x83, 0xbb, 0xe7, 0x2a, 0xf4, 0xd9, 0xb7, 0x80, 0x03, 0x9f, 0xf7, 0x5a, 0x9f,
	0x41, 0xcc, 0x1e, 0x5c, 0x78, 0xf3, 0x11, 0xe0, 0x69, 0xf7, 0xe7, 0x8f, 0x56, 0x5b, 0x7a, 0x61,
	0xeb, 0x95, 0x1a, 0xee, 0xee, 0xb5, 0xa7, 0xaf, 0x9a, 0x7a, 0x45, 0xce, 0x98, 0xc2, 0x7a, 0x6a,
	0x1f, 0x71, 0x35, 0x45, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe7, 0xc5, 0x8a, 0xa4, 0x01,
	0x00, 0x00,
}
