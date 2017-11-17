package pipeline

import (
	"fmt"
	"os/exec"
)

type Pipeline struct {
	Config *PipelineJobConfig
	Root *PipelineJob
}

func (p *Pipeline) BuildPipelineGraph() error {
	// create root job object
	// recurse through the graph 
}

func (p *Pipeline) StartPipeline() error {

}