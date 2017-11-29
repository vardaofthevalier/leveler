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
	for name, s := range pipelineConfig.Steps {
		inputs := make(map[string]*PipelineInput)
		for _, inp := range s.Inputs {
			inputs[inp] = pipelineConfig.Inputs[inp]
		}

		outputs := make(map[string]*PipelineOutput)
		for _, out := range s.Outputs {
			outputs[out] = pipelineConfig.Outputs[out]
		}

		if _, ok := allJobs[name]; ok {
			return allJobs, p, errors.New("Duplicate job names found in pipeline!")
		} else {
			switch serverConfig.Platform.Name {
			// case "kubernetes":
			// 	j, err := NewKubernetesPipelineJob(serverConfig, name, s, inputs, outputs)
			// 	if err != nil {
			// 		return allJobs, p, err
			// 	}
			// 	allJobs[name] = &j

			// case "docker":
			// 	j, err := NewDockerPipelineJob(serverConfig, name, s, inputs, outputs)
			// 	if err != nil {
			// 		return allJobs, p, err
			// 	}
			// 	allJobs[name] = &j

			case "local":
				j, err := NewLocalPipelineJob(serverConfig, name, s, inputs, outputs)
				if err != nil {
					return allJobs, p, err
				}

				allJobs[name] = &j

			default:
				return allJobs, p, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
			}
			
			if len(s.DependsOn) == 0 {
				ptr := allJobs[name]
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
 
	for name, s := range pipelineConfig.Steps {
		child := allJobs[name]
		var dependencies []string
		for _, i := range s.Inputs {
			if info, ok := pipelineConfig.Outputs[i]; ok {
				dependencies = append(dependencies, info.From)
			}
		}
		for _, d := range dependencies {
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



