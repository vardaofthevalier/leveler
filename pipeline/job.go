package pipeline 

import (
	"fmt"
	"os/exec"
)

type PipelineJob struct {
	Container string
	RepoPath string
	Command string
	Children []*PipelineJob
}

var pipelineJobScriptTemplate = `
#!/bin/bash

# use the vault token to get bitbucket secrets
# clone the repository
# cd to Workdir
# run the command 
# if the command succeeded, report success 
# otherwise report failure
`

func (j *PipelineJob) GenerateScript() error {

}

func (j *PipelineJob) Run() error {
	// write the pipeline job script (parameterized template) to the current working directory
	// run the script
}