package pipelines

import (
	"fmt"
	"log"
	"sync"
	"os/exec"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
)

type LocalPipelineJob struct {
	Id string
	Name string
	Workdir string
	Script string
	Env map[string]string
	Inputs []*PipelineInput  // for local jobs, read from a socket
	Outputs []*PipelineOutput  // for local jobs, write to a socket
	Parents []*PipelineJob
	Children []*PipelineJob
	Color string  // for detecting cycles in the job graph -- not intended to be used within a job
}

func NewLocalPipelineJob(serverConfig *config.ServerConfig, jobConfig *PipelineStep) (LocalPipelineJob, error) {
	// LEVELER_DATA default location:  /var/lib/leveler
	// create workdir under <LEVELER_DATA>/pipelines/<pipeline-id>/<job-name>
	// resolve inputs (i.e., create links) <LEVELER_DATA>/pipelines/<pipeline-id>/<dependency-name>/outputs/<output-name> -> /var/lib/leveler/pipelines/<pipeline-id>/<job-name>/inputs/<input-name>
	// create directory for outputs <LEVELER_DATA>/pipelines/<pipeline-id>/outputs/<output-name>
	// generate script <LEVELER_DATA>/pipelines/<pipeline-id>/<job-name>/run.sh

	k := LocalPipelineJob{
		Id: uuid.NewV4().String(),
		Name: jobConfig.Name,
		Script: script,
		Color: "white",   // for cycle detection -- not meant to be used within a job
	}

	return k, nil
}

func (j *LocalPipelineJob) SetColor(color string) {
	j.Color = color
}

func (j *LocalPipelineJob) GetColor() string {
	return j.Color
}
 
func (j *LocalPipelineJob) GetId() string {
	return j.Id
}

func (j *LocalPipelineJob) GetName() string {
	return j.Name
}

func (j *LocalPipelineJob) GetChildren() []*PipelineJob {
	return j.Children
}

func (j *LocalPipelineJob) GetParents() []*PipelineJob {
	return j.Parents
}

func (j *LocalPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *LocalPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *LocalPipelineJob) Watch(statuses chan *PipelineJobStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	// watch for the job to complete
	// put status on the channel
	_ = j.Cleanup()
}

func (j *LocalPipelineJob) Run() error { 
	// create job script
	// resolve inputs 
	// run script, generate outputs
	return nil
}

func (j *LocalPipelineJob) Cleanup() error {
	// delete working directory
	return nil
}