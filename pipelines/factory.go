package pipelines

import (
	"fmt"
	"errors"
	"leveler/config"
)

// PIPELINE FACTORY

func createJobsMap(serverConfig *config.ServerConfig, pipelineConfig *BasicPipeline) (map[string]PipelineJob, *Pipeline, error) {
	var p = &Pipeline{}
	var allJobs = make(map[string]PipelineJob)

	// process jobs into a map for O(1) lookup later on, and also to verify that no duplicate names are found
	for _, s := range pipelineConfig.Steps {
		if _, ok := allJobs[s.Name]; ok {
			return allJobs, p, errors.New("Duplicate job names found in pipeline!")
		} else {
			switch serverConfig.Platform.Name {
			case "kubernetes":
				j, err := NewKubernetesPipelineJob(serverConfig, s)
				if err != nil {
					return allJobs, p, err
				}
				allJobs[s.Name] = &j

			case "docker":
				j, err := NewDockerPipelineJob(serverConfig, s)
				if err != nil {
					return allJobs, p, err
				}
				allJobs[s.Name] = &j

			case "local":
				j, err := NewLocalPipelineJob(serverConfig, s)
				if err != nil {
					allJobs[s.Name] = &j
				}

			default:
				return allJobs, p, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
			}
			
			if len(s.DependsOn) == 0 {
				ptr := allJobs[s.Name]
				p.RootJobs = append(p.RootJobs, &ptr)
			}
		}
	}

	return allJobs, p, nil
}

func NewBasicPipeline(serverConfig *config.ServerConfig, pipelineConfig *BasicPipeline) (*Pipeline, error) {
	// TODO: validate that integration configurations can be found for the user who submitted this pipeline
	// if not, return an error to the caller

	
	allJobs, p, err := createJobsMap(serverConfig, pipelineConfig)
	if err != nil {
		return p, err
	}

	for _, s := range pipelineConfig.Steps {
		child := allJobs[s.Name]
		for _, d := range s.DependsOn {
			parent := allJobs[d]
			if parent.GetName() == child.GetName() {
				return p, errors.New(fmt.Sprintf("Job '%s' contains a self loop!", child.GetName()))
			}
			
			child.AddParent(&parent)
			parent.AddChild(&child)
		}
	}

	if p.hasCycle() {
		return p, errors.New("Cycle detected in graph!")
	}

	return p, nil
}



