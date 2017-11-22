package pipelines

import (
	"fmt"
	"errors"
	"leveler/config"
)

func BuildBasicPipelineGraph(serverConfig *config.ServerConfig, pipelineConfig *BasicPipeline) (*Pipeline, error) {
	// validate that integration configurations can be found for the user who submitted this pipeline
	// if not, return an error to the caller
	// var p = &Pipeline{}
	// var allJobs = make(map[string]PipelineJob)
	// var visited = make(map[string]int8)

	// // process jobs into a map for O(1) lookup later on, and also to verify that no duplicate names are found
	// for _, s := range pipelineConfig.Steps {
	// 	if _, ok := allJobs[s.Name]; ok {
	// 		return p, errors.New("Duplicate job names found in pipeline!")
	// 	} else {
	// 		switch serverConfig.Platform.Name {
	// 		case "kubernetes":
	// 			j, err := NewKubernetesPipelineJob(serverConfig, s)
	// 			if err != nil {
	// 				return p, err
	// 			}
	// 			allJobs[s.Name] = &j

	// 		default:
	// 			return p, errors.New(fmt.Sprintf("Unknown platform '%s'", serverConfig.Platform.Name))
	// 		}
			
	// 		if len(s.DependsOn) == 0 {
	// 			ptr := allJobs[s.Name]
	// 			p.RootJobs = append(p.RootJobs, &ptr)
	// 			visited[s.Name] = 0
	// 		}
	// 	}
	// }


	// for len(visited) < len(allJobs) {
	// 	for _, s := range pipelineConfig.Steps {
	// 		name := s.Name
	// 		if _, ok := visited[name]; !ok {
	// 			parentsRemaining := []string{}
	// 			for _, d := range s.DependsOn {
	// 				if name == d {
	// 					return p, errors.New("No self loops allowed!")
	// 				} else if _, ok := allJobs[name]; !ok {
	// 					return p, errors.New(fmt.Sprintf("Job '%s' specifies unknown dependency '%s'!", name, d))
	// 				} else if _, ok := visited[d]; ok {
	// 					parent := allJobs[d]
	// 					child := allJobs[name]
	// 					parent.AddChild(&child)
	// 					child.AddParent(&parent)
	// 					parentsVisited += 1
	// 				} else {

	// 				}
	// 			}

	// 			visited[name] = 0

	// 		} else {
				
	// 		}
	// 	}
	// }
	
	// return p, nil

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
			child.AddParent(&parent)
			parent.AddChild(&child)
		}
	}

	return p, nil

}

// func BuildCicdPipelineGraph(config *CicdPipeline) (Pipeline, error) {
// 	// validate that integration configurations can be found for the user who submitted this pipeline
// 	// if not, return an error to the caller
// 	// otherwise, build a pipeline graph based on the contents of the Config and any admin constraints applied to the repository 
// 	// the graph is rooted at Root
// 	// if there are multiple "roots" in the Config (i.e., the first stage is parallelized), then create a "dummy" root to simplify the tree
// 	// all job ids, parent/child relationships should be set here

// 	// for now, don't worry about integrations since they aren't completely fleshed out yet


// }