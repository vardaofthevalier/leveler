package main 

import (
	"os"
	"fmt"
	"io/ioutil"
	"leveler/config"
	"leveler/pipelines"
	"github.com/ghodss/yaml"
)

func main() {
	serverConfig := &config.ServerConfig{
		Datadir: "/home/abby/.leveler",
		Platform: &config.ContainerPlatform{
			Name: "local",
			// Host: "localhost",
			// Port: 8001,
			// Opts: &config.ContainerPlatform_KubernetesOptions{
			// 	KubernetesOptions: &config.KubernetesOptions{
			// 		Namespace: "jobs",
			// 	},
			// },
		},
	}

	contents, err := ioutil.ReadFile("advanced-2.yml")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	pipelineConfig := &pipelines.BasicPipeline{}
	err = yaml.Unmarshal(contents, pipelineConfig)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} 

	p, err := pipelines.NewBasicPipeline(serverConfig, pipelineConfig)

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	
	p.Run(make(chan int8))

	p.PrettyPrint()
}