package main 

import (
	"fmt"
	"leveler/config"
	"leveler/pipelines"
)

func main() {
	serverConfig := &config.ServerConfig{
		Platform: &config.ContainerPlatform{
			Name: "kubernetes",
			Host: "localhost",
			Port: 8001,
			Opts: &config.ContainerPlatform_KubernetesOptions{
				KubernetesOptions: &config.KubernetesOptions{
					Namespace: "jobs",
				},
			},
		},
	}

	p1_1 := &pipelines.PipelineStep{
		Name: "p1_1",
		Workdir: "foo/bar",
		Command: "ls -al",
		Image: "ubuntu",
	}

	p1_2 := &pipelines.PipelineStep{
		Name: "p1_2",
		Workdir: "foo/bar",
		Command: "ls -al",
		Image: "ubuntu",
		DependsOn: []string{"p1_2"},
	}

	p1_3 := &pipelines.PipelineStep{
		Name: "p1_3",
		Workdir: "foo/bar",
		Command: "ls -al",
		Image: "ubuntu",
		DependsOn: []string{"p1_1", "p1_2"},
	}

	pipelineConfigNoCycle := &pipelines.BasicPipeline{
		Steps: []*pipelines.PipelineStep{p1_1, p1_2, p1_3},
	}

	g1, err := pipelines.BuildBasicPipelineGraph(serverConfig, pipelineConfigNoCycle)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	if g1.HasCycle() {
		fmt.Printf("Yikes!")
	}

	// for _, r := range g1.RootJobs {
	// 	fmt.Printf("%v\n", r)
	// 	children := (*r).GetChildren()
	// 	fmt.Printf("%v\n", children)
	// 	for _, c := range children {
	// 		fmt.Printf("%v\n", (*c).GetChildren())
	// 	}
	// }

}