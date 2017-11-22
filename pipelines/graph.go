package pipelines

import (
	"fmt"
	"errors"
	"leveler/config"
)

func BuildBasicPipelineGraph(serverConfig *config.ServerConfig, pipelineConfig *BasicPipeline) (*Pipeline, error) {
	// TODO: validate that integration configurations can be found for the user who submitted this pipeline
	// if not, return an error to the caller

	var p = &Pipeline{}
	var allJobs = make(map[string]PipelineJob)

	// process jobs into a map for O(1) lookup later on, and also to verify that no duplicate names are found
	for _, s := range pipelineConfig.Steps {
		if _, ok := allJobs[s.Name]; ok {
			return p, errors.New("Duplicate job names found in pipeline!")
		} else {
			switch serverConfig.Platform.Name {
			case "kubernetes":
				j, err := NewKubernetesPipelineJob(serverConfig, s)
				if err != nil {
					return p, err
				}
				allJobs[s.Name] = &j

			default:
				return p, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
			}
			
			if len(s.DependsOn) == 0 {
				ptr := allJobs[s.Name]
				p.RootJobs = append(p.RootJobs, &ptr)
			}
		}
	}

	for _, s := range pipelineConfig.Steps {
		child := allJobs[s.Name]
		for _, d := range s.DependsOn {
			parent := allJobs[d]
			if parent.GetName() == child.GetName() {
				return p, errors.New("Self loops not allowed!")
			}
			child.AddParent(&parent)
			parent.AddChild(&child)
		}
	}

	return p, nil

}
