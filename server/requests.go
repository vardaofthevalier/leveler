package server

import (
	"fmt"
	"errors"
	"reflect"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func (s *EndpointServer) MakeAddRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeRemoveRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeCreateRequest(ctx context.Context, resource *Resource) (*Resource, error) {
	return &Resource{}, errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeGetRequest(ctx context.Context, resource *Resource) (*Resource, error) {
	return &Resource{}, errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeListRequest(ctx context.Context, query *Query) (*ResourceList, error) {
	return &ResourceList{}, errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeUpdateRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakePatchRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeDeleteRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeApplyRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) MakeRunRequest(ctx context.Context, resource *Resource) (*Resource, error) {
	pb := reflect.ValueOf(proto.MessageType(resource.Type)).Interface()

	switch resource.Type {
	case "PipelineConfig":
		err := ptypes.UnmarshalAny(resource.Message, pb.(*PipelineConfig))
		if err != nil {
			return &Resource{}, err
		}
		p, err := s.RunPipeline(ctx, pb.(*PipelineConfig))
		if err != nil {
			return &Resource{}, err
		}

		a, err := ptypes.MarshalAny(p)
		if err != nil {
			return &Resource{}, err
		} 

		return &Resource{Type: "PipelineConfig", Message: a}, nil
	default:
		return &Resource{}, errors.New(fmt.Sprintf("Unsupported type '%s' for 'run' operation!", resource.Type))
	}
}

func (s *EndpointServer) MakeCancelRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) RunPipeline(ctx context.Context, pipeline *PipelineConfig) (*PipelineConfig, error) {
	return &PipelineConfig{}, nil
}