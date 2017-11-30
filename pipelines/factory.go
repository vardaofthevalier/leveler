package pipelines

import (
	"fmt"
	"errors"
	"path/filepath"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
)

// PIPELINE FACTORY

type PipelineInputMapping struct {
	Name string
	SrcPath string
	DestPath string
	Link bool
	Integration string
	SrcJob string
}

type PipelineOutputMapping struct {
	Name string
	SrcPath string
	DestPath string
	Integration string
}

func generateInputMappings(serverConfig *config.ServerConfig, pipelineId string, jobName string, jobSpec *PipelineStep, inputs map[string]*PipelineInput, outputs map[string]*PipelineOutput) (map[string]*PipelineInputMapping, error) {
	var mappings = make(map[string]*PipelineInputMapping)

	for _, name := range jobSpec.Inputs {
		mapping := &PipelineInputMapping{
			Name: name,
			Link: false,
		}

		if _, ok := mappings[name]; ok {
			return mappings, errors.New(fmt.Sprintf("Input '%s' specified multiple times in configuration!", name))
		}

		if inputSpec, ok := inputs[name]; ok {
			mapping.Integration = inputSpec.Integration
			mapping.DestPath = filepath.Join(serverConfig.Datadir, "pipelines", pipelineId, jobName, name)

			if len(inputSpec.Integration) > 0 {
				mapping.SrcPath = inputSpec.From

			} else {
				if srcJob, ok := outputs[inputSpec.From]; ok {
					mapping.SrcJob = srcJob.From
					mapping.SrcPath = filepath.Join(serverConfig.Datadir, "pipelines", pipelineId, srcJob.From, inputSpec.From)
				} else {
					return mappings, errors.New(fmt.Sprintf("Input specified for dependent job '%s' doesn't exist in outputs map!", name))
				}
				
			}

			if inputSpec.Link {
				mapping.Link = true
			}
		} else {
			return mappings, errors.New(fmt.Sprintf("Input specified for job '%s' doesn't exist in inputs map!", name))
		}

		mappings[name] = mapping
	}

	return mappings, nil
}

func generateOutputMappings(serverConfig *config.ServerConfig, pipelineId string, jobName string, jobSpec *PipelineStep, outputs map[string]*PipelineOutput) (map[string]*PipelineOutputMapping, error) {
	var mappings = make(map[string]*PipelineOutputMapping)

	for _, name := range jobSpec.Outputs {
		mapping := &PipelineOutputMapping{
			Name: name,
		}

		if _, ok := mappings[name]; ok {
			return mappings, errors.New(fmt.Sprintf("Output '%s' specified multiple times in configuration!", name))
		}

		if outputSpec, ok := outputs[name]; ok {
			mapping.Integration = outputSpec.Integration 
			mapping.DestPath = outputSpec.To 
			mapping.SrcPath = filepath.Join(serverConfig.Datadir, "pipelines", pipelineId, jobName, name)

		} else {
			return mappings, errors.New(fmt.Sprintf("Output specified for job '%s' doesn't exist in outputs map!", name))
		}

		mappings[name] = mapping
	}

	return mappings, nil
}

func createJobsMap(serverConfig *config.ServerConfig, pipelineId string, pipelineConfig *BasicPipeline) (map[string]PipelineJob, *Pipeline, error) {
	var p = &Pipeline{}
	var allJobs = make(map[string]PipelineJob)

	// process jobs into a map for O(1) lookup later on, and also to verify that no duplicate names are found
	for name, s := range pipelineConfig.Steps {
		inputs, err := generateInputMappings(serverConfig, pipelineId, name, s, pipelineConfig.Inputs, pipelineConfig.Outputs)
		if err != nil {
			return allJobs, p, err
		}
	
		outputs, err := generateOutputMappings(serverConfig, pipelineId, name, s, pipelineConfig.Outputs)
		if err != nil {
			return allJobs, p, err
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
				j, err := NewLocalPipelineJob(serverConfig, pipelineId, name, s, inputs, outputs)
				if err != nil {
					return allJobs, p, err
				}

				allJobs[name] = &j

			default:
				return allJobs, p, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
			}
		}
	}

	return allJobs, p, nil
}

func NewBasicPipeline(serverConfig *config.ServerConfig, pipelineConfig *BasicPipeline) (*Pipeline, error) {
	// TODO: validate that integration configurations can be found for the user who submitted this pipeline
	// if not, return an error to the caller
	pipelineId := uuid.NewV4().String()
	
	allJobs, p, err := createJobsMap(serverConfig, pipelineId, pipelineConfig)
	if err != nil {
		return p, err
	}

	// TODO: recalculate dependencies using info from the functions above
 
	for name, s := range pipelineConfig.Steps {
		child := allJobs[name]
		

		var dependencies []string
		for _, i := range s.Inputs {
			if len(child.GetInputs()[i].Integration) == 0 {
				dependencies = append(dependencies, child.GetInputs()[i].SrcJob)
			}
		}

		if len(dependencies) == 0 {
			p.RootJobs = append(p.RootJobs, &child)
		} else {
			for _, d := range dependencies {
				parent := allJobs[d]
				if parent.GetName() == child.GetName() {
					return p, errors.New(fmt.Sprintf("Job '%s' contains a self loop!", child.GetName()))
				}
				
				child.AddParent(&parent)
				parent.AddChild(&child)
			}
		}
	}

	if p.hasCycle() {
		return p, errors.New("Cycle detected in graph!")
	}

	return p, nil
}



