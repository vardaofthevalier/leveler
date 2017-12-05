// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository.proto

package resources

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Repository struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Repository) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type RepositoriesList struct {
	Results []*Repository `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *RepositoriesList) Reset()                    { *m = RepositoriesList{} }
func (m *RepositoriesList) String() string            { return proto.CompactTextString(m) }
func (*RepositoriesList) ProtoMessage()               {}
func (*RepositoriesList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *RepositoriesList) GetResults() []*Repository {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Repository)(nil), "resources.Repository")
	proto.RegisterType((*RepositoriesList)(nil), "resources.RepositoriesList")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2d, 0xc8,
	0x2f, 0xce, 0x2c, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x4a,
	0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x56, 0x92, 0xe3, 0xe2, 0x0a, 0x82, 0x4b, 0x0b, 0x09,
	0x70, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0xce,
	0x5c, 0x02, 0x70, 0xf9, 0xcc, 0xd4, 0x62, 0x9f, 0xcc, 0xe2, 0x12, 0x21, 0x7d, 0x2e, 0xf6, 0xa2,
	0xd4, 0xe2, 0xd2, 0x9c, 0x92, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x51, 0x3d, 0xb8,
	0x81, 0x7a, 0x08, 0xd3, 0x82, 0x60, 0xaa, 0x92, 0xd8, 0xc0, 0xd6, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x65, 0x4a, 0x29, 0x8a, 0x00, 0x00, 0x00,
}
