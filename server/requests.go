package server

import (
	"fmt"
	"errors"
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func (s *EndpointServer) MakeAddRequest(ctx context.Context, resource *Resource) error {
	switch resource.Kind {
	case "Integration":
		m := &Integration{}
		err := ptypes.UnmarshalAny(resource.Message, m)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Printf("Received integration: %+v\n", m)
		// p, err := s.RunPipeline(ctx, pb.(*PipelineConfig))
		// if err != nil {
		// 	return &Resource{}, err
		// }

		_, err = ptypes.MarshalAny(m)
		if err != nil {
			return err
		} 

	default:
		return errors.New(fmt.Sprintf("Unsupported type '%s' for 'run' operation!", resource.Kind))
	}

	return nil
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
	var a *any.Any
	switch resource.Kind {
	case "PipelineConfig":
		m := &PipelineConfig{}
		err := ptypes.UnmarshalAny(resource.Message, m)
		if err != nil {
			fmt.Println(err)
			return &Resource{}, err
		}

		fmt.Printf("Received pipeline: %+v\n", m)
		// p, err := s.RunPipeline(ctx, pb.(*PipelineConfig))
		// if err != nil {
		// 	return &Resource{}, err
		// }

		a, err = ptypes.MarshalAny(m)
		if err != nil {
			return &Resource{Kind: "PipelineConfig", Message: a,}, err
		} 

	default:
		return &Resource{Kind: resource.Kind, Message: a,}, errors.New(fmt.Sprintf("Unsupported type '%s' for 'run' operation!", resource.Kind))
	}

	return &Resource{Kind: resource.Kind, Message: a}, nil
}

func (s *EndpointServer) MakeCancelRequest(ctx context.Context, resource *Resource) error {
	return errors.New("Not yet implemented!")
}

func (s *EndpointServer) RunPipeline(ctx context.Context, pipeline *PipelineConfig) (*PipelineConfig, error) {
	return &PipelineConfig{}, nil
}