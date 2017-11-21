package pipeline

import (
	"fmt"
	"sync"
	"os/exec"
)

const (
	RUNNING = iota;
	SUCCEEDED
	FAILED
	CANCELLED
	UNKNOWN
)

type Pipeline struct {
	Config *PipelineConfig
	Root *PipelineJob
	Alerts []*PipelineAlertConfig
}

type PipelineJob interface {
	Run() 
	Watch(chan *PipelineJob, *sync.WaitGroup)
}

type PipelineJobStatus struct {
	Status int
	Message string
	JobSpec *PipelineJob
}

func (p *Pipeline) Run(quit chan bool) map[string]PipelineJobStatus {
	// IDEA:  quit will be a channel stored a map, which can be accessed by ID in order to kill a pipeline from the server API
	// In addition to this, when a cancel command is sent to the server, some initial job killing can occur before the quit message is sent
	jobStatuses := map[string]PipelineJobStatus

	scheduler := make(chan *PipelineJob)
	defer scheduler.Close()
	scheduler <- p.Root

	var leaves sync.Waitgroup
	defer leaves.Done()

	var leafStatuses = make(chan PipelineJobStatus)
	var parentStatuses *chan PipelineJobStatus

	for {
		select {
		case scheduler <- j:
			go func() {
				k, ok := jobStatuses[j.Id]; !ok {
					if len(j.Parents) > 0 {
						var wg sync.Waitgroup
						wg.Add(len(j.Parents))

						parentStatuses = make(chan PipelineJobStatus)
						defer parentStatuses.Close()

						for _, p := range j.Parents {
							go p.Watch(&parentStatuses, &wg)
						}

						wg.Wait()
					}

					var parentFailures []int
					for s := range parentStatuses {
						if s.Status != SUCCEEDED {
							parentFailures = append(parentFailures, s)
						}

						jobStatuses[s.JobSpec.Id] = s
					}

					if len(parentFailures) > 0 {
						close(scheduler)
					}
					// TODO: check the database to make sure the job and/or pipeline hasn't been cancelled by the user
					go j.Run()

					if len(current.Children) > 0 {
						for _, c := range current.Children {
							scheduler <- c
						}
					} else {
						leaves.Add(1)
						j.Watch(&leafStatuses, &wg)
					}
				}
			}
		case <-quit:
			break
		}
	}

	leaves.Wait()

	for s := range leafStatuses {
		jobStatuses[s.JobSpec.Id] = s
	}
	
	return jobStatuses
}