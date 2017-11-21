package pipelines

import (
	"fmt"
)

func (p *Pipeline) BuildBatchJobGraph() error {

}

func (p *Pipeline) BuildCicdPipelineGraph() error {
	// build a pipeline graph based on the contents of the Config and any admin constraints applied to the repository 
	// the graph is rooted at Root
	// if there are multiple "roots" in the Config (i.e., the first stage is parallelized), then create a "dummy" root to simplify the tree
	// all job ids, parent/child relationships should be set here
}