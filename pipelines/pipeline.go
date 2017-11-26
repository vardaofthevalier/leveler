package pipelines

import (
	// "fmt"
	"sync"
	"context"
	"encoding/json"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
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
	Status int64
	Message string
	JobId string
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
				if len((*j).GetParents()) > 0 {
					var pg sync.WaitGroup
					pg.Add(len((*j).GetParents()))

					parentStatuses = make(chan *PipelineJobStatus)
					for _, p := range (*j).GetParents() {
						go (*p).Watch(ctx, parentStatuses, &pg)
					}

					pg.Wait()
				}

				var parentFailures []int64
				for s := range parentStatuses {
					if s.Status != SUCCEEDED {
						parentFailures = append(parentFailures, s.Status)
					}
				}

				close(parentStatuses)

				if len(parentFailures) > 0 {
					cancel()
				} else {
					go (*j).Run(ctx)

					if len((*j).GetChildren()) > 0 {
						for _, c := range (*j).GetChildren() {
							scheduler <- c
						}
					} else {
						leaves.Add(1)
						go (*j).Watch(ctx, leafStatuses, &leaves)
					}
				}
			}()
		case <-ctx.Done():
			break
		}
	}

	leaves.Wait()
}

func (p *Pipeline) Status() ([][]byte, error) {	
	var node *PipelineJobStatus

	q := queue.New()
	for _, r := range p.RootJobs {
		q.Enqueue(r)
	}

	for len(q) > 0 {
		node = q.Dequeue()
		for _, c := range node.GetChildren() {
			q.Enqueue(c)
		}

		j, err := json.Marshal(r)
		if err != nil {
			return results, err
		}

		fmt.Printf("%s\n", string(j))
		results = append(results, j)
	}

	return results, nil
}

func (p *Pipeline) Cleanup() error {
	// TODO: BFS and run the job cleanup function for each job
	// close channels?
	return nil
}