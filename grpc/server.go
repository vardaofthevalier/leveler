package grpc

import (
	"log"
	"errors"
	"context"
	"leveler/data"
	"leveler/resources"
	"leveler/pipelines"
	"github.com/golang/protobuf/ptypes/empty"
)

type EndpointServer struct {
	Database data.Database
	SecretStore data.SecretStore
	LogCollector data.LogCollector
	StorageDriver data.StorageDriver
}

func (s *EndpointServer) GetPipeline(ctx context.Context, pipeline *resources.Pipeline) *resources.Pipeline {

}

func (s *EndpointServer) ListPipelines(ctx context.Context, query *resources.Query) *resources.PipelinesList {

}

func (s *EndpointServer) RunPipeline(ctx context.Context, pipeline *resources.Pipeline) *resource.Pipeline {

}

func (s *EndpointServer) CancelPipeline(ctx context.Context, pipeline *resources.Pipeline) *empty.Empty {

}