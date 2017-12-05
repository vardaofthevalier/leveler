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
	/*
		TODO:
		- add worker queue
	*/
}

var pipelineMap = make(map[string]*pipelines.Pipeline)
var pipelineTracker = make(map[string]chan int8)

func (s *EndpointServer) GetPipeline(ctx context.Context, pipeline *resources.Pipeline) *resources.Pipeline {
	// TODO: figure out pipeline storage solution -- map + cache + db?
	var result *resources.Pipeline 
	return result
}

func (s *EndpointServer) ListPipelines(ctx context.Context, query *resources.Query) *resources.PipelinesList {
	var result *resources.PipelinesList
	return result
}

func (s *EndpointServer) RunPipeline(ctx context.Context, pipeline *resources.Pipeline) *resource.Pipeline {
	p, err := pipelines.NewPipeline(pipeline)
	pipeline.Id = p.Id

	pipelineMap[p.Id] = p
	pipelineTracker[p.Id] = make(chan int8)

	go func() {
		// watch the pipeline's done channel for a message -- once it's received, write the pipeline status to permanent storage
	}()

	p.Run(pipelineTracker[p.Id])

	return pipeline
}

func (s *EndpointServer) CancelPipeline(ctx context.Context, pipeline *resources.Pipeline) *empty.Empty {
	return &empty.Empty{}
}