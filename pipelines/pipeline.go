package pipelines

import (
	"fmt"
	"sync"
	"errors"
	"reflect"
	"github.com/golang-collections/collections/stack"
	"github.com/golang-collections/collections/queue"
)

const (
	ACKNOWLEDGED = iota;
	INITIALIZING 
	RUNNING
	SUCCEEDED
	FAILED
	CANCELLED
	UNKNOWN
)

type Pipeline struct {
	Id string
	RootJobs []*PipelineJob
}

type PipelineJob interface {
	Quit(status int64, message string)
	GetId() string
	GetName() string
	SetColor(string) 
	GetColor() string
	GetChildren() []*PipelineJob
	GetParents() []*PipelineJob
	GetInputs() map[string]*PipelineInputMapping
	GetOutputs() map[string]*PipelineOutputMapping
	GetConfig() *PipelineStep
	GetStatus() *PipelineJobStatus
	GetJson() (string, error)
	AddChild(*PipelineJob)
	AddParent(*PipelineJob)
	Init() error
	SyncInputs(chan int8) error 
	SyncOutputs(chan int8) error
	Run(chan int8)
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

func (p *Pipeline) Run(quit chan int8) {
	// IDEA:  quit will be a channel stored in a map, which can be accessed by ID in order to cancel a pipeline from the server API

	// IDEA: to make this work in a distributed system, use central message queues instead of channels to coordinate waiting (?)
	scheduler := make(chan *PipelineJob)

	var leaves sync.WaitGroup
	
	var scheduled = make(map[string]string)
	go func() {
		q := []*PipelineJob{}

		for _, r := range p.RootJobs {
			q = append(q, r)
		}

		for len(q) > 0 {
			current := q[0]
			q = q[1:]
			
			if _, ok := scheduled[(*current).GetId()]; !ok {
				scheduler <- current
				scheduled[(*current).GetId()] = ""	

				if len((*current).GetChildren()) == 0 {
					leaves.Add(1)
					go (*current).Watch(&leaves)

				} else {
					for _, c := range (*current).GetChildren() {
						q = append(q, c)
					}
				}		
			}
		}

		close(scheduler)
	}()

	var breakFor bool
	for {
		if breakFor {
			break
		}

		select {
			case <-quit:
				breakFor = true
				break
			case j, ok := <-scheduler:
				if ok {
					go func(j *PipelineJob) {
						if len((*j).GetParents()) > 0 {
							var pg sync.WaitGroup
							pg.Add(len((*j).GetParents()))

							for _, p := range (*j).GetParents() {
								go (*p).Watch(&pg)
							}
							pg.Wait()

							for _, p := range (*j).GetParents() {
								if (*p).GetStatus().Status != SUCCEEDED {
									(*j).Quit(CANCELLED, fmt.Sprintf("Parent job %s failed", ((*p).GetName())))
									return
								}
							}
						}
						(*j).Run(quit)
					}(j)
				} else {
					breakFor = true
					break
				}
		}
	}

	leaves.Wait()
}

func (p *Pipeline) Status() ([]string, error) {	
	var results []string
	var found = make(map[string]string)
	q := queue.New()
	for _, r := range p.RootJobs {
		q.Enqueue(r)
	}

	for q.Len() > 0 {
		node := q.Dequeue()

		switch node.(type) {
		case *PipelineJob:
			ptr := node.(*PipelineJob)

			j, err := (*ptr).GetJson()
			if err != nil {
				return results, err
			}
			
			results = append(results, j)

			for _, c := range (*ptr).GetChildren() {
				if _, ok := found[(*c).GetName()]; !ok {
					found[(*c).GetName()] = ""
					q.Enqueue(c)
				}
			}

		default:
			return results, errors.New(fmt.Sprintf("Expected pipeline job, received %s!", reflect.TypeOf(node))) 
		}
	}

	return results, nil
}

func (p *Pipeline) PrettyPrint() {
	s, err := p.Status() 
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		for _, r := range s {
			fmt.Printf("%v\n", r)
		}
	}
}

func (p *Pipeline) Cleanup() error {
	// TODO: BFS and run the job cleanup function for each job
	// close channels?
	return nil
}