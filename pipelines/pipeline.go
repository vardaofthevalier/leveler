package pipelines

import (
	"fmt"
	"sync"
	"errors"
	"context"
	"encoding/json"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
)

const (
	INITIALIZING = iota;
	RUNNING
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
	GetConfig() *PipelineStep
	GetStatus() *PipelineJobStatus
	AddChild(*PipelineJob)
	AddParent(*PipelineJob)
	Init() error
	SyncInputs(context.Context) error 
	SyncOutputs(context.Context) error
	Run(context.Context)
	Watch(*sync.WaitGroup)
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

func (p *Pipeline) Run(ctx context.Context, cancel context.CancelFunc) {
	// IDEA:  ctx will be a context.Context (created using WithCancel method) stored in a map, which can be accessed by ID in order to cancel a pipeline from the server API
	defer cancel()

	// IDEA: to make this work in a distributed system, use central message queues instead of channels
	scheduler := make(chan *PipelineJob)

	var leaves sync.WaitGroup
	
	go func() {
		q := []*PipelineJob{}

		for _, r := range p.RootJobs {
			q = append(q, r)
		}

		for len(q) > 0 {
			current := q[0]
			q = q[1:]

			fmt.Printf("Scheduling job: Id=%+v, Name=%v\n", (*current).GetId(), (*current).GetName())
			scheduler <- current
			for _, c := range (*current).GetChildren() {
				q = append(q, c)
			}
		}

		close(scheduler)
	}()

	for j := range scheduler {
		if len((*j).GetChildren()) == 0 {
			leaves.Add(1)
			go (*j).Watch(&leaves)
		}
		
		go func() {
			fmt.Println("made it")
			if len((*j).GetParents()) > 0 {
				var pg sync.WaitGroup
				pg.Add(len((*j).GetParents()))

				for _, p := range (*j).GetParents() {
					go (*p).Watch(&pg)
				}
		
				pg.Wait()

				for _, p := range (*j).GetParents() {
					if (*p).GetStatus().Status != SUCCEEDED {
						cancel()
						return
					}
				}
			}

			fmt.Println("heyooo")
			go (*j).Run(ctx)
		}()
	}

	leaves.Wait()
}

func (p *Pipeline) Status() ([][]byte, error) {	
	var results [][]byte

	q := queue.New()
	for _, r := range p.RootJobs {
		q.Enqueue(r)
	}

	for q.Len() > 0 {
		node := q.Dequeue()

		switch node.(type) {
		case *PipelineJob:
			for _, c := range node.(PipelineJob).GetChildren() {
				q.Enqueue(c)
			}

			j, err := json.Marshal(node)
			if err != nil {
				return results, err
			}

			fmt.Printf("%s\n", string(j))
			results = append(results, j)
		default:
			return results, errors.New("Malformed pipeline job!") 
		}
		
	}

	return results, nil
}

func (p *Pipeline) Cleanup() error {
	// TODO: BFS and run the job cleanup function for each job
	// close channels?
	return nil
}