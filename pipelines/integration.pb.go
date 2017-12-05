// Code generated by protoc-gen-go. DO NOT EDIT.
// source: integration.proto

package resources

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Language int32

const (
	Language_java   Language = 0
	Language_golang Language = 1
	Language_python Language = 2
)

var Language_name = map[int32]string{
	0: "java",
	1: "golang",
	2: "python",
}
var Language_value = map[string]int32{
	"java":   0,
	"golang": 1,
	"python": 2,
}

func (x Language) String() string {
	return proto.EnumName(Language_name, int32(x))
}
func (Language) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BuildTool int32

const (
	BuildTool_mvn    BuildTool = 0
	BuildTool_gradle BuildTool = 1
)

var BuildTool_name = map[int32]string{
	0: "mvn",
	1: "gradle",
}
var BuildTool_value = map[string]int32{
	"mvn":    0,
	"gradle": 1,
}

func (x BuildTool) String() string {
	return proto.EnumName(BuildTool_name, int32(x))
}
func (BuildTool) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type Integration struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*Integration_Bitbucket
	//	*Integration_Github
	//	*Integration_Aws
	//	*Integration_Nexus
	//	*Integration_Slack
	Config isIntegration_Config `protobuf_oneof:"config"`
}

func (m *Integration) Reset()                    { *m = Integration{} }
func (m *Integration) String() string            { return proto.CompactTextString(m) }
func (*Integration) ProtoMessage()               {}
func (*Integration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isIntegration_Config interface {
	isIntegration_Config()
}

type Integration_Bitbucket struct {
	Bitbucket *BitbucketAccessConfig `protobuf:"bytes,4,opt,name=bitbucket,oneof"`
}
type Integration_Github struct {
	Github *GithubAccessConfig `protobuf:"bytes,5,opt,name=github,oneof"`
}
type Integration_Aws struct {
	Aws *AWSAccessConfig `protobuf:"bytes,6,opt,name=aws,oneof"`
}
type Integration_Nexus struct {
	Nexus *NexusAccessConfig `protobuf:"bytes,7,opt,name=nexus,oneof"`
}
type Integration_Slack struct {
	Slack *SlackAccessConfig `protobuf:"bytes,8,opt,name=slack,oneof"`
}

func (*Integration_Bitbucket) isIntegration_Config() {}
func (*Integration_Github) isIntegration_Config()    {}
func (*Integration_Aws) isIntegration_Config()       {}
func (*Integration_Nexus) isIntegration_Config()     {}
func (*Integration_Slack) isIntegration_Config()     {}

func (m *Integration) GetConfig() isIntegration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Integration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Integration) GetBitbucket() *BitbucketAccessConfig {
	if x, ok := m.GetConfig().(*Integration_Bitbucket); ok {
		return x.Bitbucket
	}
	return nil
}

func (m *Integration) GetGithub() *GithubAccessConfig {
	if x, ok := m.GetConfig().(*Integration_Github); ok {
		return x.Github
	}
	return nil
}

func (m *Integration) GetAws() *AWSAccessConfig {
	if x, ok := m.GetConfig().(*Integration_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Integration) GetNexus() *NexusAccessConfig {
	if x, ok := m.GetConfig().(*Integration_Nexus); ok {
		return x.Nexus
	}
	return nil
}

func (m *Integration) GetSlack() *SlackAccessConfig {
	if x, ok := m.GetConfig().(*Integration_Slack); ok {
		return x.Slack
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Integration) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Integration_OneofMarshaler, _Integration_OneofUnmarshaler, _Integration_OneofSizer, []interface{}{
		(*Integration_Bitbucket)(nil),
		(*Integration_Github)(nil),
		(*Integration_Aws)(nil),
		(*Integration_Nexus)(nil),
		(*Integration_Slack)(nil),
	}
}

func _Integration_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Integration)
	// config
	switch x := m.Config.(type) {
	case *Integration_Bitbucket:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Bitbucket); err != nil {
			return err
		}
	case *Integration_Github:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Github); err != nil {
			return err
		}
	case *Integration_Aws:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *Integration_Nexus:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nexus); err != nil {
			return err
		}
	case *Integration_Slack:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Slack); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Integration.Config has unexpected type %T", x)
	}
	return nil
}

func _Integration_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Integration)
	switch tag {
	case 4: // config.bitbucket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BitbucketAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &Integration_Bitbucket{msg}
		return true, err
	case 5: // config.github
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GithubAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &Integration_Github{msg}
		return true, err
	case 6: // config.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AWSAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &Integration_Aws{msg}
		return true, err
	case 7: // config.nexus
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NexusAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &Integration_Nexus{msg}
		return true, err
	case 8: // config.slack
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SlackAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &Integration_Slack{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Integration_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Integration)
	// config
	switch x := m.Config.(type) {
	case *Integration_Bitbucket:
		s := proto.Size(x.Bitbucket)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Integration_Github:
		s := proto.Size(x.Github)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Integration_Aws:
		s := proto.Size(x.Aws)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Integration_Nexus:
		s := proto.Size(x.Nexus)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Integration_Slack:
		s := proto.Size(x.Slack)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type IntegrationsList struct {
	Results []*Integration `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *IntegrationsList) Reset()                    { *m = IntegrationsList{} }
func (m *IntegrationsList) String() string            { return proto.CompactTextString(m) }
func (*IntegrationsList) ProtoMessage()               {}
func (*IntegrationsList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *IntegrationsList) GetResults() []*Integration {
	if m != nil {
		return m.Results
	}
	return nil
}

type BitbucketAccessConfig struct {
	Stuff string `protobuf:"bytes,1,opt,name=stuff" json:"stuff,omitempty"`
}

func (m *BitbucketAccessConfig) Reset()                    { *m = BitbucketAccessConfig{} }
func (m *BitbucketAccessConfig) String() string            { return proto.CompactTextString(m) }
func (*BitbucketAccessConfig) ProtoMessage()               {}
func (*BitbucketAccessConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BitbucketAccessConfig) GetStuff() string {
	if m != nil {
		return m.Stuff
	}
	return ""
}

type GithubAccessConfig struct {
	Stuff string `protobuf:"bytes,1,opt,name=stuff" json:"stuff,omitempty"`
}

func (m *GithubAccessConfig) Reset()                    { *m = GithubAccessConfig{} }
func (m *GithubAccessConfig) String() string            { return proto.CompactTextString(m) }
func (*GithubAccessConfig) ProtoMessage()               {}
func (*GithubAccessConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GithubAccessConfig) GetStuff() string {
	if m != nil {
		return m.Stuff
	}
	return ""
}

type AWSAccessConfig struct {
	AwsSecretKey   string `protobuf:"bytes,1,opt,name=aws_secret_key,json=awsSecretKey" json:"aws_secret_key,omitempty"`
	AwsSecretKeyId string `protobuf:"bytes,2,opt,name=aws_secret_key_id,json=awsSecretKeyId" json:"aws_secret_key_id,omitempty"`
}

func (m *AWSAccessConfig) Reset()                    { *m = AWSAccessConfig{} }
func (m *AWSAccessConfig) String() string            { return proto.CompactTextString(m) }
func (*AWSAccessConfig) ProtoMessage()               {}
func (*AWSAccessConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *AWSAccessConfig) GetAwsSecretKey() string {
	if m != nil {
		return m.AwsSecretKey
	}
	return ""
}

func (m *AWSAccessConfig) GetAwsSecretKeyId() string {
	if m != nil {
		return m.AwsSecretKeyId
	}
	return ""
}

type BuildEnvironmentConfig struct {
	Language  Language  `protobuf:"varint,1,opt,name=language,enum=resources.Language" json:"language,omitempty"`
	BuildTool BuildTool `protobuf:"varint,2,opt,name=build_tool,json=buildTool,enum=resources.BuildTool" json:"build_tool,omitempty"`
}

func (m *BuildEnvironmentConfig) Reset()                    { *m = BuildEnvironmentConfig{} }
func (m *BuildEnvironmentConfig) String() string            { return proto.CompactTextString(m) }
func (*BuildEnvironmentConfig) ProtoMessage()               {}
func (*BuildEnvironmentConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BuildEnvironmentConfig) GetLanguage() Language {
	if m != nil {
		return m.Language
	}
	return Language_java
}

func (m *BuildEnvironmentConfig) GetBuildTool() BuildTool {
	if m != nil {
		return m.BuildTool
	}
	return BuildTool_mvn
}

type NexusAccessConfig struct {
	NexusUsername string                  `protobuf:"bytes,1,opt,name=nexus_username,json=nexusUsername" json:"nexus_username,omitempty"`
	NexusPassword string                  `protobuf:"bytes,2,opt,name=nexus_password,json=nexusPassword" json:"nexus_password,omitempty"`
	NexusHost     string                  `protobuf:"bytes,3,opt,name=nexus_host,json=nexusHost" json:"nexus_host,omitempty"`
	NexusPort     string                  `protobuf:"bytes,4,opt,name=nexus_port,json=nexusPort" json:"nexus_port,omitempty"`
	Environment   *BuildEnvironmentConfig `protobuf:"bytes,5,opt,name=environment" json:"environment,omitempty"`
}

func (m *NexusAccessConfig) Reset()                    { *m = NexusAccessConfig{} }
func (m *NexusAccessConfig) String() string            { return proto.CompactTextString(m) }
func (*NexusAccessConfig) ProtoMessage()               {}
func (*NexusAccessConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *NexusAccessConfig) GetNexusUsername() string {
	if m != nil {
		return m.NexusUsername
	}
	return ""
}

func (m *NexusAccessConfig) GetNexusPassword() string {
	if m != nil {
		return m.NexusPassword
	}
	return ""
}

func (m *NexusAccessConfig) GetNexusHost() string {
	if m != nil {
		return m.NexusHost
	}
	return ""
}

func (m *NexusAccessConfig) GetNexusPort() string {
	if m != nil {
		return m.NexusPort
	}
	return ""
}

func (m *NexusAccessConfig) GetEnvironment() *BuildEnvironmentConfig {
	if m != nil {
		return m.Environment
	}
	return nil
}

type SlackAccessConfig struct {
	Stuff string `protobuf:"bytes,1,opt,name=stuff" json:"stuff,omitempty"`
}

func (m *SlackAccessConfig) Reset()                    { *m = SlackAccessConfig{} }
func (m *SlackAccessConfig) String() string            { return proto.CompactTextString(m) }
func (*SlackAccessConfig) ProtoMessage()               {}
func (*SlackAccessConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SlackAccessConfig) GetStuff() string {
	if m != nil {
		return m.Stuff
	}
	return ""
}

func init() {
	proto.RegisterType((*Integration)(nil), "resources.Integration")
	proto.RegisterType((*IntegrationsList)(nil), "resources.IntegrationsList")
	proto.RegisterType((*BitbucketAccessConfig)(nil), "resources.BitbucketAccessConfig")
	proto.RegisterType((*GithubAccessConfig)(nil), "resources.GithubAccessConfig")
	proto.RegisterType((*AWSAccessConfig)(nil), "resources.AWSAccessConfig")
	proto.RegisterType((*BuildEnvironmentConfig)(nil), "resources.BuildEnvironmentConfig")
	proto.RegisterType((*NexusAccessConfig)(nil), "resources.NexusAccessConfig")
	proto.RegisterType((*SlackAccessConfig)(nil), "resources.SlackAccessConfig")
	proto.RegisterEnum("resources.Language", Language_name, Language_value)
	proto.RegisterEnum("resources.BuildTool", BuildTool_name, BuildTool_value)
}

func init() { proto.RegisterFile("integration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xfe, 0xcf, 0x29, 0x94, 0xd6, 0x8c, 0x29, 0x42, 0x4c, 0x2a, 0x11, 0x48, 0x5d,
	0x05, 0x05, 0x75, 0x48, 0xdc, 0xb2, 0x0e, 0xc4, 0x26, 0x26, 0x34, 0xa5, 0x20, 0x2e, 0x2b, 0x27,
	0x75, 0xd3, 0xd0, 0xd4, 0xae, 0x6c, 0xa7, 0xa5, 0x37, 0x3c, 0x1c, 0xcf, 0xc2, 0x83, 0xa0, 0x38,
	0x49, 0x63, 0x92, 0xed, 0xce, 0xf9, 0xce, 0xef, 0xb3, 0xe5, 0xef, 0x1c, 0x07, 0x7a, 0x01, 0x95,
	0xc4, 0xe7, 0x58, 0x06, 0x8c, 0x8e, 0x36, 0x9c, 0x49, 0x86, 0x4c, 0x4e, 0x04, 0x8b, 0xb8, 0x47,
	0x84, 0xfd, 0xa7, 0x02, 0xed, 0xeb, 0x1c, 0x40, 0x08, 0x6a, 0x14, 0xaf, 0x89, 0x65, 0xf4, 0x8d,
	0x81, 0xe9, 0xa8, 0x35, 0xfa, 0x00, 0xa6, 0x1b, 0x48, 0x37, 0xf2, 0x56, 0x44, 0x5a, 0xb5, 0xbe,
	0x31, 0x68, 0x8f, 0xfb, 0xa3, 0xc3, 0x16, 0xa3, 0x49, 0x56, 0xbb, 0xf0, 0x3c, 0x22, 0xc4, 0x25,
	0xa3, 0x8b, 0xc0, 0xbf, 0x3a, 0x72, 0x72, 0x13, 0x7a, 0x0f, 0x0d, 0x3f, 0x90, 0xcb, 0xc8, 0xb5,
	0xea, 0xca, 0x7e, 0xaa, 0xd9, 0x3f, 0xab, 0x42, 0xc1, 0x9b, 0xe2, 0x68, 0x04, 0x55, 0xbc, 0x13,
	0x56, 0x43, 0xb9, 0x9e, 0x6a, 0xae, 0x8b, 0x1f, 0xd3, 0x82, 0x25, 0x06, 0xd1, 0x3b, 0xa8, 0x53,
	0xf2, 0x2b, 0x12, 0x56, 0x53, 0x39, 0x9e, 0x69, 0x8e, 0xaf, 0xb1, 0x5e, 0xf0, 0x24, 0x70, 0xec,
	0x12, 0x21, 0xf6, 0x56, 0x56, 0xab, 0xe4, 0x9a, 0xc6, 0x7a, 0xd1, 0xa5, 0xe0, 0x49, 0x0b, 0x1a,
	0x9e, 0x92, 0xec, 0x8f, 0xd0, 0xd5, 0x32, 0x14, 0x37, 0x81, 0x90, 0xe8, 0x2d, 0x34, 0x39, 0x11,
	0x51, 0x28, 0x85, 0x65, 0xf4, 0xab, 0x83, 0xf6, 0xf8, 0x44, 0xdb, 0x55, 0xa3, 0x9d, 0x0c, 0xb3,
	0x5f, 0xc3, 0x93, 0x3b, 0xa3, 0x44, 0xc7, 0x50, 0x17, 0x32, 0x5a, 0x2c, 0xd2, 0xa6, 0x24, 0x1f,
	0xf6, 0x10, 0x50, 0x39, 0xba, 0x7b, 0x58, 0x17, 0x1e, 0x15, 0x02, 0x43, 0x2f, 0xa0, 0x83, 0x77,
	0x62, 0x26, 0x88, 0xc7, 0x89, 0x9c, 0xad, 0xc8, 0x3e, 0x75, 0x3c, 0xc0, 0x3b, 0x31, 0x55, 0xe2,
	0x17, 0xb2, 0x47, 0x67, 0xd0, 0xfb, 0x9f, 0x9a, 0x05, 0x73, 0xab, 0xa2, 0xc0, 0x8e, 0x0e, 0x5e,
	0xcf, 0xed, 0xdf, 0x70, 0x32, 0x89, 0x82, 0x70, 0xfe, 0x89, 0x6e, 0x03, 0xce, 0xe8, 0x9a, 0x50,
	0x99, 0x1e, 0xf5, 0x06, 0x5a, 0x21, 0xa6, 0x7e, 0x84, 0xfd, 0x64, 0xae, 0x3a, 0xe3, 0xc7, 0x5a,
	0x16, 0x37, 0x69, 0xc9, 0x39, 0x40, 0xe8, 0x1c, 0xc0, 0x8d, 0xb7, 0x9a, 0x49, 0xc6, 0x42, 0x75,
	0x5c, 0x67, 0x7c, 0xac, 0x4f, 0x5c, 0x5c, 0xfc, 0xc6, 0x58, 0xe8, 0x98, 0x6e, 0xb6, 0xb4, 0xff,
	0x1a, 0xd0, 0x2b, 0xf5, 0x18, 0xbd, 0x84, 0x8e, 0xea, 0xf1, 0x2c, 0x12, 0x84, 0x6b, 0x93, 0xfd,
	0x50, 0xa9, 0xdf, 0x53, 0x31, 0xc7, 0x36, 0x58, 0x88, 0x1d, 0xe3, 0xd9, 0x25, 0x13, 0xec, 0x36,
	0x15, 0xd1, 0x29, 0x40, 0x82, 0x2d, 0x99, 0x90, 0x56, 0x55, 0x21, 0xa6, 0x52, 0xae, 0x98, 0x90,
	0x79, 0x79, 0xc3, 0x78, 0xf2, 0x52, 0xb2, 0xf2, 0x2d, 0xe3, 0x12, 0x5d, 0x42, 0x9b, 0xe4, 0xe1,
	0xa4, 0x4f, 0xe1, 0x79, 0xf1, 0x5e, 0xa5, 0xfc, 0x1c, 0xdd, 0x65, 0x9f, 0x41, 0xaf, 0x34, 0x93,
	0x77, 0x77, 0x7d, 0xf8, 0x0a, 0x5a, 0x59, 0xb8, 0xa8, 0x05, 0xb5, 0x9f, 0x78, 0x8b, 0xbb, 0x47,
	0x08, 0xa0, 0xe1, 0xb3, 0x38, 0xea, 0xae, 0x11, 0xaf, 0x37, 0x7b, 0xb9, 0x64, 0xb4, 0x5b, 0x19,
	0xf6, 0xc1, 0x3c, 0xe4, 0x8a, 0x9a, 0x50, 0x5d, 0x6f, 0x69, 0x4a, 0x73, 0x3c, 0x0f, 0x49, 0xd7,
	0x70, 0x1b, 0xea, 0xef, 0x71, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xba, 0x47, 0xaf, 0x39, 0x52,
	0x04, 0x00, 0x00,
}
