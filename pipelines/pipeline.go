package pipelines

import (
	// "fmt"
	"sync"
	"context"
	"github.com/golang-collections/collections/stack"
)

const (
	RUNNING = iota;
	SUCCEEDED
	FAILED
	CANCELLED
	UNKNOWN
)

type Pipeline struct {
	RootJobs []*PipelineJob
}

type PipelineJob interface {
	GetId() string
	GetName() string
	SetColor(string) 
	GetColor() string
	GetChildren() []*PipelineJob
	GetParents() []*PipelineJob
	AddChild(*PipelineJob)
	AddParent(*PipelineJob)
	Init() error
	SyncInputs() error 
	SyncOutputs() error
	Run(chan *PipelineJobStatus) error 
	Watch(chan *PipelineJobStatus, *sync.WaitGroup)
	Cleanup() error
}

type PipelineJobStatus struct {
	Status int
	Message string
	JobSpec *PipelineJob
}

func (p *Pipeline) hasCycle() bool {
	s := stack.New()

	for _, r := range p.RootJobs {
		s.Push(r)
	}

	for s.Len() > 0 {
		current := s.Pop()
		value := current.(*PipelineJob)
		if (*value).GetColor() == "grey" {
			for _, c := range (*value).GetChildren() {
				if (*c).GetColor() != "black" {
					return true
				}
			}
			(*value).SetColor("black")
		} else if (*value).GetColor() == "white" {
			(*value).SetColor("grey")
			s.Push(value)
			for _, c := range (*value).GetChildren() {
				s.Push(c)
			}
		}
	}

	return false
}

func (p *Pipeline) Run(ctx context.Context, cancel context.CancelFunc) map[string]PipelineJobStatus {
	// IDEA:  ctx will be a context.Context (created using WithCancel method) stored in a map, which can be accessed by ID in order to cancel a pipeline from the server API
	defer cancel()

	// IDEA: to make this work in a distributed system, use central message queues instead of channels
	scheduler := make(chan *PipelineJob)
	defer close(scheduler)

	var jobStatuses map[string]PipelineJobStatus
	var parentStatuses chan *PipelineJobStatus

	var leafStatuses = make(chan *PipelineJobStatus)
	defer close(leafStatuses)

	var leaves sync.WaitGroup
	defer leaves.Done()
	
	for _, j := range p.RootJobs {
		scheduler <- j
	}

	var j *PipelineJob
	for {
		select {
		case scheduler <- j:
			go func() {
				if _, ok := jobStatuses[(*j).GetId()]; !ok {
					if len((*j).GetParents()) > 0 {
						var wg sync.WaitGroup
						wg.Add(len((*j).GetParents()))

						parentStatuses := make(chan *PipelineJobStatus)
						defer close(parentStatuses)

						for _, p := range (*j).GetParents() {
							go (*p).Watch(ctx, parentStatuses, &wg)
						}

						wg.Wait()
					}

					var parentFailures []int
					for s := range parentStatuses {
						if s.Status != SUCCEEDED {
							parentFailures = append(parentFailures, s.Status)
						}

						jobStatuses[(*s.JobSpec).GetId()] = *s
					}

					if len(parentFailures) > 0 {
						quit <- true
					}

					if len((*j).GetChildren()) > 0 {
						for _, c := range (*j).GetChildren() {
							scheduler <- c
						}
					} else {
						leaves.Add(1)
						go (*j).Watch(ctx, leafStatuses, &leaves)
					}

					err = (*j).Init()
					if err != nil {
						// TODO:
					}

					err = (*j).Run()
					if err != nil {
						// TODO: create pipeline job status from error
					}
				}
			}()
		case jobQuit <- jq:
			jobStatuses[(*s.JobSpec).GetId()] = *jq
			quit <- true
		case <-ctx.Done():
			close(scheduler)
			break
		}
	}

	leaves.Wait()

	for s := range leafStatuses {
		jobStatuses[(*s.JobSpec).GetId()] = *s
	}
	
	// TODO: run cleanup logic (toggle true/false in config)
	return jobStatuses
}

func (p *Pipeline) Cleanup() error {
	// TODO: BFS and run the job cleanup function for each job
	return nil
}