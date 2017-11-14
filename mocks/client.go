package mocks 

import (
	"context"
	"google.golang.org/grpc"
	resources "leveler/resources"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type MockResourceEndpointClient struct {}

func (m *MockResourceEndpointClient) CreateResource(ctx context.Context, obj *resources.Resource, opt ...grpc.CallOption) (*resources.Resource, error) {

	return &resources.Resource{}, nil
}

func (m *MockResourceEndpointClient) GetResource(ctx context.Context, obj *resources.Resource, opt ...grpc.CallOption) (*resources.Resource, error) {

	return &resources.Resource{}, nil
}

func (m *MockResourceEndpointClient) ListResources(ctx context.Context, query *resources.Query, opt ...grpc.CallOption) (*resources.ResourceList, error) {

	return &resources.ResourceList{}, nil
}

func (m *MockResourceEndpointClient) UpdateResource(ctx context.Context, obj *resources.Resource, opt ...grpc.CallOption) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}

func (m *MockResourceEndpointClient) DeleteResource(ctx context.Context, obj *resources.Resource, opt ...grpc.CallOption) (*empty.Empty, error) {
	
	return &empty.Empty{}, nil
}