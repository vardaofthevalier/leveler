package server

import (
	"fmt"
	"errors"
	"path/filepath"
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

func GenerateInputMappings(datadir string, jobSpec *JobConfig, pipelineId string, inputs map[string]*PipelineInputConfig, outputs map[string]*PipelineOutputConfig) (map[string]*PipelineInputMapping, error) {
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
			mapping.DestPath = filepath.Join(datadir, name)

			if len(inputSpec.Integration) > 0 {
				mapping.SrcPath = inputSpec.From

			} else {
				if srcJob, ok := outputs[inputSpec.From]; ok {
					mapping.SrcJob = srcJob.From
					mapping.SrcPath = filepath.Join(datadir, pipelineId, srcJob.From, inputSpec.From)
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

func GenerateInputSyncScript(datadir string, integration *PipelineIntegrationConfig, input *PipelineInputMapping) (string, error) {
	return "", nil
}

func GenerateOutputSyncScript(datadir string, integration *PipelineIntegrationConfig, output *PipelineInputMapping) (string, error) {
	return "", nil
}

func GenerateOutputMappings(datadir string, jobSpec *JobConfig, pipelineId string, outputs map[string]*PipelineOutputConfig) (map[string]*PipelineOutputMapping, error) {
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
			mapping.SrcPath = filepath.Join(datadir, name)

		} else {
			return mappings, errors.New(fmt.Sprintf("Output specified for job '%s' doesn't exist in outputs map!", name))
		}

		mappings[name] = mapping
	}

	return mappings, nil
}

func createJobsMap(serverConfig *ServerConfig, pipelineId string, pipelineConfig *PipelineConfig) (map[string]PipelineJob, error) {
	var allJobs = make(map[string]PipelineJob)

	// process jobs into a map for O(1) lookup later on, and also to verify that no duplicate names are found
	for name, s := range pipelineConfig.Steps {
		if _, ok := allJobs[name]; ok {
			return allJobs, errors.New("Duplicate job names found in pipeline!")
		} else {
			switch serverConfig.Platform.Name {
			// case "kubernetes":
			//  TODO: create instance of k8s client
			// 	j, err := NewKubernetesPipelineJob(serverConfig, name, s, inputs, outputs)
			// 	if err != nil {
			// 		return allJobs, p, err
			// 	}
			// 	allJobs[name] = &j

			case "docker":
				j, err := NewDockerPipelineJob(serverConfig, pipelineId, name, s, pipelineConfig.Inputs, pipelineConfig.Outputs)
				if err != nil {
					return allJobs, err
				}
				allJobs[name] = &j

			case "local":
				j, err := NewLocalPipelineJob(serverConfig, pipelineId, name, s, pipelineConfig.Inputs, pipelineConfig.Outputs)
				if err != nil {
					return allJobs, err
				}

				allJobs[name] = &j

			default:
				return allJobs, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
			}
		}
	}

	return allJobs, nil
}

func NewPipeline(serverConfig *ServerConfig, pipelineConfig *PipelineConfig) (*Pipeline, error) {
	pipelineId := uuid.NewV4().String()
	p := &Pipeline{}
	
	allJobs, err := createJobsMap(serverConfig, pipelineId, pipelineConfig)
	if err != nil {
		return p, err
	}

	p.Id = pipelineId
	p.JobsMap = allJobs
 
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



