package server

import (
	"context"
	"leveler/data"
	"github.com/golang/protobuf/ptypes/empty"
)

type EndpointServer struct {
	Database data.Database
	SecretStore data.SecretStore
	LogCollector data.LogCollector
	StorageDriver data.StorageDriver
}

type ResourceRequest interface {
	MakeAddRequest(context.Context, *Resource) error
	MakeRemoveRequest(context.Context, *Resource) error 
	MakeCreateRequest(context.Context, *Resource) (*Resource, error)
	MakeGetRequest(context.Context, *Resource) (*Resource, error)
	MakeListRequest(context.Context, *Query) (*ResourceList, error)
	MakeUpdateRequest(context.Context, *Resource) error
	MakePatchRequest(context.Context, *Resource) error
	MakeDeleteRequest(context.Context, *Resource) error
	MakeApplyRequest(context.Context, *Resource) error 
	MakeRunRequest(context.Context, *Resource) (*Resource, error)
	MakeCancelRequest(context.Context, *Resource) error
}

func (s *EndpointServer) Add(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeAddRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}

	return &empty.Empty{}, nil
}

func (s *EndpointServer) Remove(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeRemoveRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}

func (s *EndpointServer) Create(ctx context.Context, resource *Resource) (*Resource, error) {
	r, err := s.MakeCreateRequest(ctx, resource)
	
	if err != nil {
		return r, err 
	}
	return r, nil
}

func (s *EndpointServer) Get(ctx context.Context, resource *Resource) (*Resource, error) {
	r, err := s.MakeGetRequest(ctx, resource)
	
	if err != nil {
		return r, err 
	}
	return r, nil
}

func (s *EndpointServer) List(ctx context.Context, query *Query) (*ResourceList, error) {
	r, err := s.MakeListRequest(ctx, query)
	
	if err != nil {
		return r, err 
	}
	return r, nil
}

func (s *EndpointServer) Update(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeUpdateRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}

func (s *EndpointServer) Patch(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakePatchRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}

func (s *EndpointServer) Delete(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeDeleteRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}

func (s *EndpointServer) Apply(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeApplyRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}

func (s *EndpointServer) Run(ctx context.Context, resource *Resource) (*Resource, error) {
	r, err := s.MakeRunRequest(ctx, resource)
	
	if err != nil {
		return r, err 
	}
	return r, nil
}	

func (s *EndpointServer) Cancel(ctx context.Context, resource *Resource) (*empty.Empty, error) {
	err := s.MakeCancelRequest(ctx, resource)
	
	if err != nil {
		return &empty.Empty{}, err 
	}
	return &empty.Empty{}, nil
}