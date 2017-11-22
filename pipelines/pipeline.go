package pipelines

import (
	"fmt"
	"sync"
	//"os/exec"
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
	SetVisited() 
	GetVisited() bool
	GetChildren() []*PipelineJob
	GetParents() []*PipelineJob
	AddChild(*PipelineJob)
	AddParent(*PipelineJob)
	Run() error 
	Watch(chan *PipelineJobStatus, *sync.WaitGroup)
	Cleanup() error
}

type PipelineJobStatus struct {
	Status int
	Message string
	JobSpec *PipelineJob
}

func (p *Pipeline) HasCycle() bool {
	s := stack.New()

	for _, r := range p.RootJobs {
		s.Push(r)
	}

	for s.Len() > 0 {
		current := s.Pop()
		value := current.(*PipelineJob)
		fmt.Println((*value).GetId())
		if (*value).GetVisited() {
			return true
		} else {
			(*value).SetVisited()
			fmt.Println((*value).GetChildren())
			for _, c := range (*value).GetChildren() {
				s.Push(c)
			}
		}
	}

	return false
}

func (p *Pipeline) Run(quit chan bool) map[string]PipelineJobStatus {
	// IDEA:  quit will be a channel stored in a map, which can be accessed by ID in order to kill a pipeline from the server API
	// In addition to this, when a cancel command is sent to the server, some initial job killing can occur before the quit message is sent

	if p.HasCycle() {
		// fatal error
	}

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
							go (*p).Watch(parentStatuses, &wg)
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
						close(scheduler)
					}

					go (*j).Run()

					if len((*j).GetChildren()) > 0 {
						for _, c := range (*j).GetChildren() {
							scheduler <- c
						}
					} else {
						leaves.Add(1)
						(*j).Watch(leafStatuses, &leaves)
					}
				}
			}()
		case <-quit:
			break
		}
	}

	leaves.Wait()

	for s := range leafStatuses {
		jobStatuses[(*s.JobSpec).GetId()] = *s
	}
	
	return jobStatuses
}