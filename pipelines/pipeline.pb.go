// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pipeline.proto

/*
Package pipelines is a generated protocol buffer package.

It is generated from these files:
	pipeline.proto
	integration.proto
	data.proto

It has these top-level messages:
	PipelineInput
	PipelineOutput
	PipelineData
	PipelineStep
	BasicPipeline
	CicdPipeline
	Integration
	IntegrationsList
	BitbucketAccessConfig
	GithubAccessConfig
	AWSConfig
	NexusConfig
	SlackConfig
	S3Data
	NexusData
	SCMData
	LocalData
	StreamData
*/
package pipelines

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

type PipelineInput struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Src:
	//	*PipelineInput_External
	//	*PipelineInput_PipelineJob
	Src isPipelineInput_Src `protobuf_oneof:"src"`
}

func (m *PipelineInput) Reset()                    { *m = PipelineInput{} }
func (m *PipelineInput) String() string            { return proto.CompactTextString(m) }
func (*PipelineInput) ProtoMessage()               {}
func (*PipelineInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isPipelineInput_Src interface {
	isPipelineInput_Src()
}

type PipelineInput_External struct {
	External *PipelineData `protobuf:"bytes,2,opt,name=external,oneof"`
}
type PipelineInput_PipelineJob struct {
	PipelineJob string `protobuf:"bytes,3,opt,name=pipelineJob,oneof"`
}

func (*PipelineInput_External) isPipelineInput_Src()    {}
func (*PipelineInput_PipelineJob) isPipelineInput_Src() {}

func (m *PipelineInput) GetSrc() isPipelineInput_Src {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *PipelineInput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineInput) GetExternal() *PipelineData {
	if x, ok := m.GetSrc().(*PipelineInput_External); ok {
		return x.External
	}
	return nil
}

func (m *PipelineInput) GetPipelineJob() string {
	if x, ok := m.GetSrc().(*PipelineInput_PipelineJob); ok {
		return x.PipelineJob
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineInput) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PipelineInput_OneofMarshaler, _PipelineInput_OneofUnmarshaler, _PipelineInput_OneofSizer, []interface{}{
		(*PipelineInput_External)(nil),
		(*PipelineInput_PipelineJob)(nil),
	}
}

func _PipelineInput_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineInput)
	// src
	switch x := m.Src.(type) {
	case *PipelineInput_External:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.External); err != nil {
			return err
		}
	case *PipelineInput_PipelineJob:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.PipelineJob)
	case nil:
	default:
		return fmt.Errorf("PipelineInput.Src has unexpected type %T", x)
	}
	return nil
}

func _PipelineInput_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineInput)
	switch tag {
	case 2: // src.external
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PipelineData)
		err := b.DecodeMessage(msg)
		m.Src = &PipelineInput_External{msg}
		return true, err
	case 3: // src.pipelineJob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Src = &PipelineInput_PipelineJob{x}
		return true, err
	default:
		return false, nil
	}
}

func _PipelineInput_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PipelineInput)
	// src
	switch x := m.Src.(type) {
	case *PipelineInput_External:
		s := proto.Size(x.External)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineInput_PipelineJob:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PipelineJob)))
		n += len(x.PipelineJob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PipelineOutput struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to Dest:
	//	*PipelineOutput_External
	//	*PipelineOutput_PipelineJob
	Dest isPipelineOutput_Dest `protobuf_oneof:"dest"`
}

func (m *PipelineOutput) Reset()                    { *m = PipelineOutput{} }
func (m *PipelineOutput) String() string            { return proto.CompactTextString(m) }
func (*PipelineOutput) ProtoMessage()               {}
func (*PipelineOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isPipelineOutput_Dest interface {
	isPipelineOutput_Dest()
}

type PipelineOutput_External struct {
	External *PipelineData `protobuf:"bytes,2,opt,name=external,oneof"`
}
type PipelineOutput_PipelineJob struct {
	PipelineJob string `protobuf:"bytes,3,opt,name=pipelineJob,oneof"`
}

func (*PipelineOutput_External) isPipelineOutput_Dest()    {}
func (*PipelineOutput_PipelineJob) isPipelineOutput_Dest() {}

func (m *PipelineOutput) GetDest() isPipelineOutput_Dest {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *PipelineOutput) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineOutput) GetExternal() *PipelineData {
	if x, ok := m.GetDest().(*PipelineOutput_External); ok {
		return x.External
	}
	return nil
}

func (m *PipelineOutput) GetPipelineJob() string {
	if x, ok := m.GetDest().(*PipelineOutput_PipelineJob); ok {
		return x.PipelineJob
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineOutput) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PipelineOutput_OneofMarshaler, _PipelineOutput_OneofUnmarshaler, _PipelineOutput_OneofSizer, []interface{}{
		(*PipelineOutput_External)(nil),
		(*PipelineOutput_PipelineJob)(nil),
	}
}

func _PipelineOutput_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineOutput)
	// dest
	switch x := m.Dest.(type) {
	case *PipelineOutput_External:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.External); err != nil {
			return err
		}
	case *PipelineOutput_PipelineJob:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.PipelineJob)
	case nil:
	default:
		return fmt.Errorf("PipelineOutput.Dest has unexpected type %T", x)
	}
	return nil
}

func _PipelineOutput_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineOutput)
	switch tag {
	case 2: // dest.external
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PipelineData)
		err := b.DecodeMessage(msg)
		m.Dest = &PipelineOutput_External{msg}
		return true, err
	case 3: // dest.pipelineJob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Dest = &PipelineOutput_PipelineJob{x}
		return true, err
	default:
		return false, nil
	}
}

func _PipelineOutput_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PipelineOutput)
	// dest
	switch x := m.Dest.(type) {
	case *PipelineOutput_External:
		s := proto.Size(x.External)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineOutput_PipelineJob:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.PipelineJob)))
		n += len(x.PipelineJob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PipelineData struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IsFile      bool   `protobuf:"varint,2,opt,name=isFile" json:"isFile,omitempty"`
	IsDirectory bool   `protobuf:"varint,3,opt,name=isDirectory" json:"isDirectory,omitempty"`
	Shared      bool   `protobuf:"varint,4,opt,name=shared" json:"shared,omitempty"`
	// Types that are valid to be assigned to Datamap:
	//	*PipelineData_S3
	//	*PipelineData_Nexus
	//	*PipelineData_Scm
	//	*PipelineData_Local
	Datamap isPipelineData_Datamap `protobuf_oneof:"datamap"`
}

func (m *PipelineData) Reset()                    { *m = PipelineData{} }
func (m *PipelineData) String() string            { return proto.CompactTextString(m) }
func (*PipelineData) ProtoMessage()               {}
func (*PipelineData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isPipelineData_Datamap interface {
	isPipelineData_Datamap()
}

type PipelineData_S3 struct {
	S3 *S3Data `protobuf:"bytes,5,opt,name=s3,oneof"`
}
type PipelineData_Nexus struct {
	Nexus *NexusData `protobuf:"bytes,6,opt,name=nexus,oneof"`
}
type PipelineData_Scm struct {
	Scm *SCMData `protobuf:"bytes,7,opt,name=scm,oneof"`
}
type PipelineData_Local struct {
	Local *LocalData `protobuf:"bytes,8,opt,name=local,oneof"`
}

func (*PipelineData_S3) isPipelineData_Datamap()    {}
func (*PipelineData_Nexus) isPipelineData_Datamap() {}
func (*PipelineData_Scm) isPipelineData_Datamap()   {}
func (*PipelineData_Local) isPipelineData_Datamap() {}

func (m *PipelineData) GetDatamap() isPipelineData_Datamap {
	if m != nil {
		return m.Datamap
	}
	return nil
}

func (m *PipelineData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineData) GetIsFile() bool {
	if m != nil {
		return m.IsFile
	}
	return false
}

func (m *PipelineData) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *PipelineData) GetShared() bool {
	if m != nil {
		return m.Shared
	}
	return false
}

func (m *PipelineData) GetS3() *S3Data {
	if x, ok := m.GetDatamap().(*PipelineData_S3); ok {
		return x.S3
	}
	return nil
}

func (m *PipelineData) GetNexus() *NexusData {
	if x, ok := m.GetDatamap().(*PipelineData_Nexus); ok {
		return x.Nexus
	}
	return nil
}

func (m *PipelineData) GetScm() *SCMData {
	if x, ok := m.GetDatamap().(*PipelineData_Scm); ok {
		return x.Scm
	}
	return nil
}

func (m *PipelineData) GetLocal() *LocalData {
	if x, ok := m.GetDatamap().(*PipelineData_Local); ok {
		return x.Local
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineData) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PipelineData_OneofMarshaler, _PipelineData_OneofUnmarshaler, _PipelineData_OneofSizer, []interface{}{
		(*PipelineData_S3)(nil),
		(*PipelineData_Nexus)(nil),
		(*PipelineData_Scm)(nil),
		(*PipelineData_Local)(nil),
	}
}

func _PipelineData_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineData)
	// datamap
	switch x := m.Datamap.(type) {
	case *PipelineData_S3:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.S3); err != nil {
			return err
		}
	case *PipelineData_Nexus:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nexus); err != nil {
			return err
		}
	case *PipelineData_Scm:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Scm); err != nil {
			return err
		}
	case *PipelineData_Local:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Local); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PipelineData.Datamap has unexpected type %T", x)
	}
	return nil
}

func _PipelineData_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineData)
	switch tag {
	case 5: // datamap.s3
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(S3Data)
		err := b.DecodeMessage(msg)
		m.Datamap = &PipelineData_S3{msg}
		return true, err
	case 6: // datamap.nexus
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NexusData)
		err := b.DecodeMessage(msg)
		m.Datamap = &PipelineData_Nexus{msg}
		return true, err
	case 7: // datamap.scm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SCMData)
		err := b.DecodeMessage(msg)
		m.Datamap = &PipelineData_Scm{msg}
		return true, err
	case 8: // datamap.local
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LocalData)
		err := b.DecodeMessage(msg)
		m.Datamap = &PipelineData_Local{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PipelineData_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PipelineData)
	// datamap
	switch x := m.Datamap.(type) {
	case *PipelineData_S3:
		s := proto.Size(x.S3)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineData_Nexus:
		s := proto.Size(x.Nexus)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineData_Scm:
		s := proto.Size(x.Scm)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PipelineData_Local:
		s := proto.Size(x.Local)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type PipelineStep struct {
	Name      string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Workdir   string            `protobuf:"bytes,2,opt,name=workdir" json:"workdir,omitempty"`
	Command   string            `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
	Image     string            `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	Inputs    []*PipelineInput  `protobuf:"bytes,5,rep,name=inputs" json:"inputs,omitempty"`
	Outputs   []*PipelineOutput `protobuf:"bytes,6,rep,name=outputs" json:"outputs,omitempty"`
	Timeout   int32             `protobuf:"varint,8,opt,name=timeout" json:"timeout,omitempty"`
	Skip      bool              `protobuf:"varint,9,opt,name=skip" json:"skip,omitempty"`
	DependsOn []string          `protobuf:"bytes,10,rep,name=dependsOn" json:"dependsOn,omitempty"`
}

func (m *PipelineStep) Reset()                    { *m = PipelineStep{} }
func (m *PipelineStep) String() string            { return proto.CompactTextString(m) }
func (*PipelineStep) ProtoMessage()               {}
func (*PipelineStep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PipelineStep) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PipelineStep) GetWorkdir() string {
	if m != nil {
		return m.Workdir
	}
	return ""
}

func (m *PipelineStep) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *PipelineStep) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *PipelineStep) GetInputs() []*PipelineInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PipelineStep) GetOutputs() []*PipelineOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *PipelineStep) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *PipelineStep) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

func (m *PipelineStep) GetDependsOn() []string {
	if m != nil {
		return m.DependsOn
	}
	return nil
}

type BasicPipeline struct {
	Integrations []*Integration  `protobuf:"bytes,1,rep,name=integrations" json:"integrations,omitempty"`
	SharedData   []*PipelineData `protobuf:"bytes,2,rep,name=sharedData" json:"sharedData,omitempty"`
	Steps        []*PipelineStep `protobuf:"bytes,3,rep,name=steps" json:"steps,omitempty"`
}

func (m *BasicPipeline) Reset()                    { *m = BasicPipeline{} }
func (m *BasicPipeline) String() string            { return proto.CompactTextString(m) }
func (*BasicPipeline) ProtoMessage()               {}
func (*BasicPipeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BasicPipeline) GetIntegrations() []*Integration {
	if m != nil {
		return m.Integrations
	}
	return nil
}

func (m *BasicPipeline) GetSharedData() []*PipelineData {
	if m != nil {
		return m.SharedData
	}
	return nil
}

func (m *BasicPipeline) GetSteps() []*PipelineStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

type CicdPipeline struct {
	Integrations []*Integration  `protobuf:"bytes,1,rep,name=integrations" json:"integrations,omitempty"`
	SharedData   []*PipelineData `protobuf:"bytes,2,rep,name=sharedData" json:"sharedData,omitempty"`
	Quality      []*PipelineStep `protobuf:"bytes,3,rep,name=quality" json:"quality,omitempty"`
	Build        []*PipelineStep `protobuf:"bytes,4,rep,name=build" json:"build,omitempty"`
	Integration  []*PipelineStep `protobuf:"bytes,5,rep,name=integration" json:"integration,omitempty"`
	Promote      []*PipelineStep `protobuf:"bytes,6,rep,name=promote" json:"promote,omitempty"`
}

func (m *CicdPipeline) Reset()                    { *m = CicdPipeline{} }
func (m *CicdPipeline) String() string            { return proto.CompactTextString(m) }
func (*CicdPipeline) ProtoMessage()               {}
func (*CicdPipeline) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CicdPipeline) GetIntegrations() []*Integration {
	if m != nil {
		return m.Integrations
	}
	return nil
}

func (m *CicdPipeline) GetSharedData() []*PipelineData {
	if m != nil {
		return m.SharedData
	}
	return nil
}

func (m *CicdPipeline) GetQuality() []*PipelineStep {
	if m != nil {
		return m.Quality
	}
	return nil
}

func (m *CicdPipeline) GetBuild() []*PipelineStep {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *CicdPipeline) GetIntegration() []*PipelineStep {
	if m != nil {
		return m.Integration
	}
	return nil
}

func (m *CicdPipeline) GetPromote() []*PipelineStep {
	if m != nil {
		return m.Promote
	}
	return nil
}

func init() {
	proto.RegisterType((*PipelineInput)(nil), "pipelines.PipelineInput")
	proto.RegisterType((*PipelineOutput)(nil), "pipelines.PipelineOutput")
	proto.RegisterType((*PipelineData)(nil), "pipelines.PipelineData")
	proto.RegisterType((*PipelineStep)(nil), "pipelines.PipelineStep")
	proto.RegisterType((*BasicPipeline)(nil), "pipelines.BasicPipeline")
	proto.RegisterType((*CicdPipeline)(nil), "pipelines.CicdPipeline")
}

func init() { proto.RegisterFile("pipeline.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x9b, 0x64, 0xb3, 0xbb, 0x99, 0xdd, 0x56, 0xaa, 0x55, 0x15, 0x53, 0x71, 0x58, 0x05,
	0x09, 0xed, 0x01, 0x56, 0xd0, 0x08, 0x21, 0x38, 0x6e, 0x2b, 0xd4, 0x22, 0xa0, 0xc8, 0x7d, 0x02,
	0x6f, 0x62, 0x15, 0xab, 0x49, 0x1c, 0x62, 0x47, 0xb4, 0x0f, 0x80, 0x78, 0x0c, 0xae, 0xdc, 0x79,
	0x14, 0x5e, 0x08, 0xd9, 0x8e, 0x5b, 0x23, 0x76, 0xe9, 0xb1, 0xb7, 0xcc, 0xe4, 0x9b, 0xf1, 0x3f,
	0xe3, 0x3f, 0x81, 0x9d, 0x86, 0x37, 0xac, 0xe4, 0x35, 0x5b, 0x34, 0xad, 0x50, 0x02, 0x25, 0x2e,
	0x96, 0x07, 0x50, 0x50, 0x45, 0x6d, 0xfa, 0x60, 0x97, 0xd7, 0x8a, 0x5d, 0xb4, 0x54, 0x71, 0x51,
	0xdb, 0x54, 0xfa, 0x2d, 0x80, 0xed, 0x4f, 0x3d, 0x7c, 0x5a, 0x37, 0x9d, 0x42, 0x08, 0x06, 0x35,
	0xad, 0x18, 0x0e, 0x66, 0xc1, 0x3c, 0x21, 0xe6, 0x19, 0xbd, 0x84, 0x31, 0xbb, 0x52, 0xac, 0xad,
	0x69, 0x89, 0xc3, 0x59, 0x30, 0x9f, 0x1c, 0x3e, 0x58, 0xdc, 0x1c, 0xb1, 0x70, 0xf5, 0xc7, 0x54,
	0xd1, 0x93, 0x2d, 0x72, 0x83, 0xa2, 0x14, 0x26, 0x8e, 0x7a, 0x27, 0x56, 0x38, 0xd2, 0x1d, 0x4f,
	0xb6, 0x88, 0x9f, 0x5c, 0xc6, 0x10, 0xc9, 0x36, 0x4f, 0xbf, 0x07, 0xb0, 0xe3, 0xfa, 0x9c, 0x75,
	0xea, 0x1e, 0x84, 0x0c, 0x61, 0x50, 0x30, 0xa9, 0xd2, 0x9f, 0x21, 0x4c, 0xfd, 0x46, 0x6b, 0x75,
	0xec, 0xc3, 0x90, 0xcb, 0xb7, 0xbc, 0x64, 0x46, 0xc5, 0x98, 0xf4, 0x11, 0x9a, 0xc1, 0x84, 0xcb,
	0x63, 0xde, 0xb2, 0x5c, 0x89, 0xf6, 0xda, 0x1c, 0x34, 0x26, 0x7e, 0x4a, 0x57, 0xca, 0xcf, 0xb4,
	0x65, 0x05, 0x1e, 0xd8, 0x4a, 0x1b, 0xa1, 0xc7, 0x10, 0xca, 0x0c, 0xc7, 0x66, 0xa6, 0x5d, 0x6f,
	0xa6, 0xf3, 0xac, 0x9f, 0x26, 0x94, 0x19, 0x7a, 0x0a, 0x71, 0xcd, 0xae, 0x3a, 0x89, 0x87, 0x86,
	0xdb, 0xf3, 0xb8, 0x8f, 0x3a, 0xdf, 0xa3, 0x16, 0x42, 0x4f, 0x20, 0x92, 0x79, 0x85, 0x47, 0x86,
	0x45, 0x7e, 0xcf, 0xa3, 0x0f, 0x3d, 0xa9, 0x01, 0xdd, 0xb5, 0x14, 0x39, 0x2d, 0xf1, 0xf8, 0x9f,
	0xae, 0xef, 0x75, 0xde, 0x75, 0x35, 0xd0, 0x32, 0x81, 0x91, 0xb6, 0x54, 0x45, 0x9b, 0xf4, 0x87,
	0xb7, 0xaa, 0x73, 0xc5, 0x9a, 0xb5, 0xab, 0xc2, 0x30, 0xfa, 0x2a, 0xda, 0xcb, 0x82, 0xb7, 0x66,
	0x57, 0x09, 0x71, 0xa1, 0x7e, 0x93, 0x8b, 0xaa, 0xa2, 0x75, 0x61, 0x6f, 0x84, 0xb8, 0x10, 0xed,
	0x41, 0xcc, 0x2b, 0x7a, 0xc1, 0xcc, 0x8e, 0x12, 0x62, 0x03, 0xf4, 0x1c, 0x86, 0x5c, 0x5b, 0x54,
	0xe2, 0x78, 0x16, 0xcd, 0x27, 0x87, 0x78, 0xcd, 0xd5, 0x1b, 0x0f, 0x93, 0x9e, 0x43, 0x19, 0x8c,
	0x84, 0x31, 0x93, 0xde, 0x98, 0x2e, 0x79, 0xb8, 0xa6, 0xc4, 0xda, 0x8d, 0x38, 0x52, 0xcb, 0x52,
	0xbc, 0x62, 0xa2, 0x53, 0x66, 0x21, 0x31, 0x71, 0xa1, 0x1e, 0x4f, 0x5e, 0xf2, 0x06, 0x27, 0xe6,
	0xe6, 0xcc, 0x33, 0x7a, 0x04, 0x49, 0xc1, 0x1a, 0x56, 0x17, 0xf2, 0xac, 0xc6, 0x30, 0x8b, 0xe6,
	0x09, 0xb9, 0x4d, 0xa4, 0xbf, 0x02, 0xd8, 0x5e, 0x52, 0xc9, 0x73, 0x77, 0x18, 0x7a, 0x03, 0x53,
	0xef, 0x2b, 0x94, 0x38, 0x30, 0xba, 0xf6, 0x3d, 0x5d, 0xa7, 0xb7, 0xaf, 0xc9, 0x5f, 0x2c, 0x7a,
	0x05, 0x60, 0xdd, 0xa2, 0x6f, 0x04, 0x87, 0xa6, 0x72, 0x93, 0xff, 0x89, 0x87, 0xa2, 0x67, 0x10,
	0x4b, 0xc5, 0x1a, 0x89, 0xa3, 0x8d, 0x35, 0xfa, 0xfe, 0x88, 0xa5, 0xd2, 0xdf, 0x21, 0x4c, 0x8f,
	0x78, 0x5e, 0xdc, 0xaf, 0xe8, 0x17, 0x30, 0xfa, 0xd2, 0xd1, 0x92, 0xab, 0xeb, 0xbb, 0x64, 0x3b,
	0x4e, 0xcf, 0xb9, 0xea, 0x78, 0xa9, 0xbf, 0xad, 0xff, 0xcf, 0x69, 0x28, 0xf4, 0x1a, 0x26, 0x9e,
	0xd4, 0xde, 0x55, 0x1b, 0x8b, 0x7c, 0x56, 0x8b, 0x6b, 0x5a, 0x51, 0x09, 0xc5, 0x7a, 0x67, 0x6d,
	0x16, 0xd7, 0x73, 0xab, 0xa1, 0xf9, 0xe3, 0x66, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x86,
	0x21, 0x01, 0xad, 0x05, 0x00, 0x00,
}
