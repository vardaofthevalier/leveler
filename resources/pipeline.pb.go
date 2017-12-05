// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipeline.proto

package resources

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PipelineIntegration struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Config:
	//	*PipelineIntegration_Bitbucket
	//	*PipelineIntegration_Github
	//	*PipelineIntegration_Aws
	//	*PipelineIntegration_Nexus
	Config isPipelineIntegration_Config `protobuf_oneof:"config"`
}

func (m *PipelineIntegration) Reset()                    { *m = PipelineIntegration{} }
func (m *PipelineIntegration) String() string            { return proto.CompactTextString(m) }
func (*PipelineIntegration) ProtoMessage()               {}
func (*PipelineIntegration) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isPipelineIntegration_Config interface {
	isPipelineIntegration_Config()
}

type PipelineIntegration_Bitbucket struct {
	Bitbucket *BitbucketAccessConfig `protobuf:"bytes,4,opt,name=bitbucket,oneof"`
}
type PipelineIntegration_Github struct {
	Github *GithubAccessConfig `protobuf:"bytes,5,opt,name=github,oneof"`
}
type PipelineIntegration_Aws struct {
	Aws *AWSAccessConfig `protobuf:"bytes,6,opt,name=aws,oneof"`
}
type PipelineIntegration_Nexus struct {
	Nexus *NexusAccessConfig `protobuf:"bytes,7,opt,name=nexus,oneof"`
}

func (*PipelineIntegration_Bitbucket) isPipelineIntegration_Config() {}
func (*PipelineIntegration_Github) isPipelineIntegration_Config()    {}
func (*PipelineIntegration_Aws) isPipelineIntegration_Config()       {}
func (*PipelineIntegration_Nexus) isPipelineIntegration_Config()     {}

func (m *PipelineIntegration) GetConfig() isPipelineIntegration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PipelineIntegration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineIntegration) GetBitbucket() *BitbucketAccessConfig {
	if x, ok := m.GetConfig().(*PipelineIntegration_Bitbucket); ok {
		return x.Bitbucket
	}
	return nil
}

func (m *PipelineIntegration) GetGithub() *GithubAccessConfig {
	if x, ok := m.GetConfig().(*PipelineIntegration_Github); ok {
		return x.Github
	}
	return nil
}

func (m *PipelineIntegration) GetAws() *AWSAccessConfig {
	if x, ok := m.GetConfig().(*PipelineIntegration_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *PipelineIntegration) GetNexus() *NexusAccessConfig {
	if x, ok := m.GetConfig().(*PipelineIntegration_Nexus); ok {
		return x.Nexus
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineIntegration) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PipelineIntegration_OneofMarshaler, _PipelineIntegration_OneofUnmarshaler, _PipelineIntegration_OneofSizer, []interface{}{
		(*PipelineIntegration_Bitbucket)(nil),
		(*PipelineIntegration_Github)(nil),
		(*PipelineIntegration_Aws)(nil),
		(*PipelineIntegration_Nexus)(nil),
	}
}

func _PipelineIntegration_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineIntegration)
	// config
	switch x := m.Config.(type) {
	case *PipelineIntegration_Bitbucket:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Bitbucket); err != nil {
			return err
		}
	case *PipelineIntegration_Github:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Github); err != nil {
			return err
		}
	case *PipelineIntegration_Aws:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *PipelineIntegration_Nexus:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nexus); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PipelineIntegration.Config has unexpected type %T", x)
	}
	return nil
}

func _PipelineIntegration_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineIntegration)
	switch tag {
	case 4: // config.bitbucket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BitbucketAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &PipelineIntegration_Bitbucket{msg}
		return true, err
	case 5: // config.github
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GithubAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &PipelineIntegration_Github{msg}
		return true, err
	case 6: // config.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AWSAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &PipelineIntegration_Aws{msg}
		return true, err
	case 7: // config.nexus
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NexusAccessConfig)
		err := b.DecodeMessage(msg)
		m.Config = &PipelineIntegration_Nexus{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PipelineIntegration_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PipelineIntegration)
	// config
	switch x := m.Config.(type) {
	case *PipelineIntegration_Bitbucket:
		s := proto.Size(x.Bitbucket)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineIntegration_Github:
		s := proto.Size(x.Github)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineIntegration_Aws:
		s := proto.Size(x.Aws)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineIntegration_Nexus:
		s := proto.Size(x.Nexus)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PipelineInput struct {
	From        string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Integration string `protobuf:"bytes,2,opt,name=integration" json:"integration,omitempty"`
	Link        bool   `protobuf:"varint,3,opt,name=link" json:"link,omitempty"`
}

func (m *PipelineInput) Reset()                    { *m = PipelineInput{} }
func (m *PipelineInput) String() string            { return proto.CompactTextString(m) }
func (*PipelineInput) ProtoMessage()               {}
func (*PipelineInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PipelineInput) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *PipelineInput) GetIntegration() string {
	if m != nil {
		return m.Integration
	}
	return ""
}

func (m *PipelineInput) GetLink() bool {
	if m != nil {
		return m.Link
	}
	return false
}

type PipelineOutput struct {
	To          string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
	From        string `protobuf:"bytes,2,opt,name=from" json:"from,omitempty"`
	Integration string `protobuf:"bytes,3,opt,name=integration" json:"integration,omitempty"`
}

func (m *PipelineOutput) Reset()                    { *m = PipelineOutput{} }
func (m *PipelineOutput) String() string            { return proto.CompactTextString(m) }
func (*PipelineOutput) ProtoMessage()               {}
func (*PipelineOutput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PipelineOutput) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *PipelineOutput) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *PipelineOutput) GetIntegration() string {
	if m != nil {
		return m.Integration
	}
	return ""
}

type Job struct {
	Id         string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PipelineId string            `protobuf:"bytes,2,opt,name=pipeline_id,json=pipelineId" json:"pipeline_id,omitempty"`
	Workdir    string            `protobuf:"bytes,3,opt,name=workdir" json:"workdir,omitempty"`
	Command    string            `protobuf:"bytes,4,opt,name=command" json:"command,omitempty"`
	Image      string            `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
	Env        map[string]string `protobuf:"bytes,6,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Inputs     []string          `protobuf:"bytes,7,rep,name=inputs" json:"inputs,omitempty"`
	Outputs    []string          `protobuf:"bytes,8,rep,name=outputs" json:"outputs,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetPipelineId() string {
	if m != nil {
		return m.PipelineId
	}
	return ""
}

func (m *Job) GetWorkdir() string {
	if m != nil {
		return m.Workdir
	}
	return ""
}

func (m *Job) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Job) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Job) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Job) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Job) GetOutputs() []string {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type JobsList struct {
	Results []*Job `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *JobsList) Reset()                    { *m = JobsList{} }
func (m *JobsList) String() string            { return proto.CompactTextString(m) }
func (*JobsList) ProtoMessage()               {}
func (*JobsList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *JobsList) GetResults() []*Job {
	if m != nil {
		return m.Results
	}
	return nil
}

type Pipeline struct {
	Id           string                          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Inputs       map[string]*PipelineInput       `protobuf:"bytes,2,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Outputs      map[string]*PipelineOutput      `protobuf:"bytes,3,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Steps        map[string]*Job                 `protobuf:"bytes,4,rep,name=steps" json:"steps,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GlobalEnv    map[string]string               `protobuf:"bytes,5,rep,name=global_env,json=globalEnv" json:"global_env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Integrations map[string]*PipelineIntegration `protobuf:"bytes,6,rep,name=integrations" json:"integrations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Pipeline) Reset()                    { *m = Pipeline{} }
func (m *Pipeline) String() string            { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()               {}
func (*Pipeline) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *Pipeline) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pipeline) GetInputs() map[string]*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Pipeline) GetOutputs() map[string]*PipelineOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Pipeline) GetSteps() map[string]*Job {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Pipeline) GetGlobalEnv() map[string]string {
	if m != nil {
		return m.GlobalEnv
	}
	return nil
}

func (m *Pipeline) GetIntegrations() map[string]*PipelineIntegration {
	if m != nil {
		return m.Integrations
	}
	return nil
}

type PipelinesList struct {
	Results []*Pipeline `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *PipelinesList) Reset()                    { *m = PipelinesList{} }
func (m *PipelinesList) String() string            { return proto.CompactTextString(m) }
func (*PipelinesList) ProtoMessage()               {}
func (*PipelinesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *PipelinesList) GetResults() []*Pipeline {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*PipelineIntegration)(nil), "resources.PipelineIntegration")
	proto.RegisterType((*PipelineInput)(nil), "resources.PipelineInput")
	proto.RegisterType((*PipelineOutput)(nil), "resources.PipelineOutput")
	proto.RegisterType((*Job)(nil), "resources.Job")
	proto.RegisterType((*JobsList)(nil), "resources.JobsList")
	proto.RegisterType((*Pipeline)(nil), "resources.Pipeline")
	proto.RegisterType((*PipelinesList)(nil), "resources.PipelinesList")
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0xdf, 0x26, 0x6b, 0xd7, 0x9c, 0xee, 0xad, 0x98, 0x87, 0xc0, 0x54, 0xb0, 0x45, 0x11,
	0x48, 0xe5, 0x82, 0x20, 0x8d, 0x89, 0xa1, 0x09, 0x21, 0x36, 0x34, 0xed, 0x8f, 0x10, 0xa0, 0x4c,
	0x80, 0xb8, 0x9a, 0x92, 0xd6, 0x2b, 0x56, 0xdb, 0xb8, 0x8a, 0x9d, 0x8d, 0x7d, 0x0f, 0xf8, 0x8e,
	0x7c, 0x0c, 0x64, 0x27, 0x4e, 0x9c, 0xd6, 0x5c, 0x70, 0x97, 0xe3, 0x73, 0x9e, 0x9f, 0xed, 0xe7,
	0x1c, 0xb7, 0xd0, 0x5f, 0xd0, 0x05, 0x99, 0xd1, 0x94, 0x84, 0x8b, 0x8c, 0x09, 0x86, 0xbc, 0x8c,
	0x70, 0x96, 0x67, 0x23, 0xc2, 0x07, 0x9b, 0x34, 0x15, 0x64, 0x92, 0xc5, 0x82, 0xb2, 0xb4, 0xc8,
	0x06, 0x3f, 0x1d, 0xd8, 0xfa, 0x54, 0x0a, 0xce, 0xea, 0x2c, 0x42, 0xb0, 0x96, 0xc6, 0x73, 0x82,
	0x5b, 0x7e, 0x6b, 0xe8, 0x45, 0xea, 0x1b, 0xbd, 0x05, 0x2f, 0xa1, 0x22, 0xc9, 0x47, 0x53, 0x22,
	0xf0, 0x9a, 0xdf, 0x1a, 0xf6, 0x76, 0xfd, 0xb0, 0xa2, 0x87, 0x47, 0x3a, 0x77, 0x38, 0x1a, 0x11,
	0xce, 0xdf, 0xb1, 0xf4, 0x8a, 0x4e, 0x4e, 0xff, 0x8b, 0x6a, 0x11, 0xda, 0x87, 0xce, 0x84, 0x8a,
	0xef, 0x79, 0x82, 0xdb, 0x4a, 0xfe, 0xc8, 0x90, 0x9f, 0xa8, 0xc4, 0x92, 0xb6, 0x2c, 0x47, 0x21,
	0xb8, 0xf1, 0x0d, 0xc7, 0x1d, 0xa5, 0x1a, 0x18, 0xaa, 0xc3, 0xaf, 0x17, 0x4b, 0x12, 0x59, 0x88,
	0xf6, 0xa0, 0x9d, 0x92, 0x1f, 0x39, 0xc7, 0xeb, 0x4a, 0xf1, 0xd0, 0x50, 0x7c, 0x90, 0xeb, 0x4b,
	0x9a, 0xa2, 0xf8, 0xa8, 0x0b, 0x9d, 0x91, 0x5a, 0x0a, 0xbe, 0xc1, 0xff, 0xb5, 0x2b, 0x8b, 0x5c,
	0x48, 0x3f, 0xae, 0x32, 0x36, 0xd7, 0x7e, 0xc8, 0x6f, 0xe4, 0x43, 0xcf, 0x30, 0x14, 0x3b, 0x2a,
	0x65, 0x2e, 0x49, 0xd5, 0x8c, 0xa6, 0x53, 0xec, 0xfa, 0xad, 0x61, 0x37, 0x52, 0xdf, 0xc1, 0x17,
	0xe8, 0x6b, 0xf4, 0xc7, 0x5c, 0x48, 0x76, 0x1f, 0x1c, 0xc1, 0x4a, 0xb2, 0x23, 0x58, 0xb5, 0x97,
	0xf3, 0xf7, 0xbd, 0xdc, 0x95, 0xbd, 0x82, 0x5f, 0x0e, 0xb8, 0xe7, 0x2c, 0x91, 0x34, 0x3a, 0xd6,
	0x34, 0x3a, 0x46, 0x3b, 0xd0, 0xd3, 0x13, 0x71, 0x49, 0xc7, 0x25, 0x14, 0xf4, 0xd2, 0xd9, 0x18,
	0x61, 0x58, 0xbf, 0x61, 0xd9, 0x74, 0x4c, 0xb3, 0x12, 0xab, 0x43, 0x99, 0x19, 0xb1, 0xf9, 0x3c,
	0x4e, 0xc7, 0xaa, 0xdd, 0x5e, 0xa4, 0x43, 0x74, 0x17, 0xda, 0x74, 0x1e, 0x4f, 0x88, 0xea, 0xa3,
	0x17, 0x15, 0x01, 0x7a, 0x0a, 0x2e, 0x49, 0xaf, 0x71, 0xc7, 0x77, 0x87, 0xbd, 0xdd, 0xfb, 0x86,
	0xe7, 0xe7, 0x2c, 0x09, 0x8f, 0xd3, 0xeb, 0xe3, 0x54, 0x64, 0xb7, 0x91, 0xac, 0x41, 0xf7, 0xa0,
	0x43, 0xa5, 0xb1, 0xb2, 0x43, 0xee, 0xd0, 0x8b, 0xca, 0x48, 0x6e, 0xc9, 0x94, 0x2b, 0x1c, 0x77,
	0x55, 0x42, 0x87, 0x83, 0x97, 0xd0, 0xd5, 0x08, 0x74, 0x07, 0xdc, 0x29, 0xb9, 0x2d, 0x2f, 0x29,
	0x3f, 0xe5, 0x81, 0xae, 0xe3, 0x59, 0x4e, 0xca, 0xfb, 0x15, 0xc1, 0x81, 0xf3, 0xaa, 0x15, 0xec,
	0x41, 0xf7, 0x9c, 0x25, 0xfc, 0x3d, 0xe5, 0x02, 0x0d, 0x61, 0x3d, 0x23, 0x3c, 0x9f, 0x09, 0x8e,
	0x5b, 0xea, 0x90, 0xfd, 0xe6, 0x21, 0x23, 0x9d, 0x0e, 0x7e, 0xb7, 0xa1, 0xab, 0xdb, 0xb4, 0x62,
	0xe9, 0x7e, 0x75, 0x78, 0x47, 0x51, 0x76, 0x0c, 0x8a, 0x16, 0x85, 0x6a, 0x6e, 0x78, 0x71, 0x65,
	0x7d, 0xbb, 0x83, 0xfa, 0x76, 0xae, 0x52, 0xfa, 0x36, 0x65, 0x31, 0x16, 0xa5, 0x54, 0x0b, 0xe4,
	0x48, 0x73, 0x41, 0x16, 0x1c, 0xaf, 0x29, 0xe5, 0xb6, 0x4d, 0x79, 0x21, 0x0b, 0x0a, 0x5d, 0x51,
	0x8c, 0x0e, 0x01, 0x26, 0x33, 0x96, 0xc4, 0xb3, 0x4b, 0xd9, 0x99, 0xb6, 0x92, 0x06, 0x36, 0xe9,
	0x89, 0xaa, 0xaa, 0x9a, 0xe4, 0x4d, 0x74, 0x8c, 0xce, 0x60, 0xc3, 0x98, 0x33, 0x5e, 0xb6, 0xf7,
	0x89, 0xfd, 0xce, 0x75, 0x5d, 0xc1, 0x69, 0x48, 0x07, 0x17, 0xd0, 0x33, 0x6c, 0xb1, 0xb4, 0x31,
	0x34, 0xdb, 0xd8, 0xdb, 0xc5, 0x96, 0x4d, 0x14, 0xc0, 0x68, 0xf0, 0xe0, 0x33, 0x6c, 0x98, 0x8e,
	0x59, 0xa8, 0xcf, 0x9b, 0xd4, 0x07, 0x16, 0x6a, 0x41, 0x30, 0xb1, 0xa7, 0x00, 0xb5, 0x9d, 0x16,
	0xe8, 0xe3, 0x26, 0x74, 0x79, 0x92, 0x0c, 0xd2, 0x6b, 0xe8, 0x37, 0xdd, 0xfd, 0x97, 0xf9, 0x1d,
	0x5c, 0xc2, 0xe6, 0x8a, 0xad, 0x16, 0xc0, 0x5e, 0xf3, 0x38, 0xdb, 0x56, 0xe7, 0x2a, 0x8c, 0xf9,
	0x40, 0xde, 0xd4, 0xbf, 0x75, 0xc5, 0x2b, 0x79, 0xb6, 0xfc, 0x4a, 0xb6, 0x2c, 0xb0, 0xea, 0xa9,
	0x24, 0x1d, 0xf5, 0x4f, 0xf2, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xe2, 0x0a, 0x6a,
	0x79, 0x06, 0x00, 0x00,
}