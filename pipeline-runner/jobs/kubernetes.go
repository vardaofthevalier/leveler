package pipeline 

import (
	"fmt"
	"sync"
	"os/exec"
)

type KubernetesPipelineJob struct {
	Id string
	Name string
	Namespace string
	Container string
	RepoPath string
	Script string
	DataProvider ExternalDataProvider
	CredentialsProvider ExternalDataProvider
	Volume 
	Parents []*PipelineJob
	Children []*PipelineJob
}

type ExternalDataProvider interface {
	GetScript()
	GetContainer()
}

var pipelineJobVolumeTemplate = `
{
	"apiVersion": "v1",
	"kind": ""
}
`

var pipelineJobPodTemplate = `
{
	"apiVersion": "v1",
	"kind": "Pod",
	"metadata": {
		"name": {{.Name}},
		"namespace": {{.Namespace}}
	},
	"spec": {
		"initContainers": [
			""
		],
		"containers": [

		]
	}
}
`

var pipelineJobScriptConfigMapTemplate = `
{
	"apiVersion": "v1",
	"kind": "ConfigMap",
	"metadata": {
		"name": "job-{{.Id}}",
		"namespace": "{{.Namespace}}"
	},
	"data": {
		"job.sh": "{{.Script}}",
		"credentials.sh": "{{.CredentialsProvider.Script}}"
		"data.sh" "{{.DataProvider.Script}}"
	}
}
`

var pipelineDataSyncScriptTemplate = `
#!/bin/bash

`

var pipelineJobScriptTemplate = `
#!/bin/bash

cd {{.RepoPath}}
/job-{{.Id}}/script.sh
exit $?
`

func (j *PipelineJob) generateScript() error {

}

func (j *KubernetesPipelineJob) Watch(statuses *chan map[string]interface{}, wg *sync.Waitgroup) {
	defer wg.Close()
	// watch for the job to complete
	// put status on the channel
	// cleanup
}
func (j *KubernetesPipelineJob) Run() error { 
	// parameterize job script template
	// parameterize the job pod template
	// update script value
	// write the pipeline job script (parameterized template) to the current working directory
	// run the script
}