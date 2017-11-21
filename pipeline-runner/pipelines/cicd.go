package pipelines

import (
	"fmt"
)

func (p *Pipeline) BuildBasicPipelineGraph() error {
	// validate that integration configurations can be found for the user who submitted this pipeline
	// if not, return an error to the caller
	// otherwise, build the pipeline graph
}

func (p *Pipeline) BuildCicdPipelineGraph() error {
	// build a pipeline graph based on the contents of the Config and any admin constraints applied to the repository 
	// the graph is rooted at Root
	// if there are multiple "roots" in the Config (i.e., the first stage is parallelized), then create a "dummy" root to simplify the tree
	// all job ids, parent/child relationships should be set here
}