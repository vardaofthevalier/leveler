package pipelines

import (
	"fmt"
	"log"
	"sync"
	"bytes"
	"context"
	"os/exec"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
	docker "github.com/docker/docker/client"
	types "github.com/docker/docker/api/types"
	volume "github.com/docker/docker/api/types/volume"
	network "github.com/docker/docker/api/types/network"
	container "github.com/docker/docker/api/types/container"
)

type DockerPipelineJob struct {
	Id string
	Name string
	Workdir string
	Script string
	Env map[string]string
	Volumes []volume.volumetypes.VolumesCreateBody
	Parents []*PipelineJob
	Children []*PipelineJob
	DockerContext context.Context
	DockerClient *docker.Client
	ContainerId string
	Config *PipelineStep
	Color string  // for detecting cycles in the job graph -- not intended to be used within a job
}

func NewDockerPipelineJob(serverConfig *config.ServerConfig, jobConfig *PipelineStep) (DockerPipelineJob, error) {
	// LEVELER_DATA default location:  /var/lib/leveler
	// create directory under <LEVELER_DATA>/pipelines/docker/<pipeline-id>/<job-name>
	// resolve inputs (i.e., check for existence of named volumes) 
	// create new named volumes for outputs
	// generate script <LEVELER_DATA>/pipelines/docker/<pipeline-id>/<job-name>/run.sh
	
	var k DockerPipelineJob
	client, err := docker.NewEnvClient()
	if err != nil {
		return k, err
	}

	// generate script
	var script string // TODO

	// generate env map
	var env map[string]string  // TODO

	// create volume requests
	var volumes []volume.volumetypes.VolumesCreateBody  // TODO: should include a volume for the script to run
	
	k := DockerPipelineJob{
		Id: uuid.NewV4().String(),
		Name: jobConfig.Name,
		Script: script,
		Env: env,
		Volumes: volumes,
		DockerClient: client,
		DockerContext: context.Background(),
		Config: jobConfig,
		Color: "white",   // for cycle detection -- not meant to be used within a job
	}

	return k, nil
}

func (j *DockerPipelineJob) SetColor(color string) {
	j.Color = color
}

func (j *DockerPipelineJob) GetColor() string {
	return j.Color
}
 
func (j *DockerPipelineJob) GetId() string {
	return j.Id
}

func (j *DockerPipelineJob) GetName() string {
	return j.Name
}

func (j *DockerPipelineJob) GetChildren() []*PipelineJob {
	return j.Children
}

func (j *DockerPipelineJob) GetParents() []*PipelineJob {
	return j.Parents
}

func (j *DockerPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *DockerPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *DockerPipelineJob) Watch(statuses chan *PipelineJobStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	var jobStatus = &PipelineJobStatus{
		Config: j.Config,
	}
	// wait for the container -- TODO: apply a timeout to the context
	status, err := j.DockerClient.ContainerWait(j.DockerContext, j.ContainerId, container.WaitConditionNextExit)

	if err != nil {
		jobStatus.Status = status.StatusCode
		jobStatus.Message = status.Error
		statuses <- jobStatus
		wg.Done()
	}

	jobStatus.Status = status.StatusCode

	// get logs for the container
	logsOptions := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: true,
		Details: true,
	}

	logs, error := j.DockerClient.ContainerLogs(j.DockerContext, j.ContainerId, logsOptions)
	if err != nil {
		jobStatus.Message = fmt.Sprintf("WARNING: couldn't get logs for job: %v", err)
	} else {
		buf := new(bytes.Buffer)
	    buf.ReadFrom(logs)
		jobStatus.Message = buf.String()
	}

	// put status on the channel
	statuses <- jobStatus
}

func (j *DockerPipelineJob) Run(quit chan *PipelineJobStatus) { 
	var errorResponse *PipelineJobStatus
	var quitPipeline = func(err error) {
		errorResponse.Status = err.StatusCode
		errorResponse.Message = err.Error
		quit <- errorResponse
		return
	}

	// create new volumes for outputs
	var volumes []types.Volume
	for _, v := range j.Volumes {
		newVolume, err := j.DockerClient.VolumeCreate(j.DockerContext, v)
		if err != nil {
			quitPipeline(err)
		}

		volumes = append(volumes, newVolumes)
	}

	// TODO: sync data to volumes using info from PipelineData and Integrations
	// The integration name should map to a type of PipelineData and we should provide standard images for using those integrations 
	// The name of the pipeline data should map to input and output names
	// if errors occur, quitPipeline


	// start docker container using info in j and volume info
	containerConfig := &container.Config{
		Env:
		Cmd:
		Image:
		Volumes:
		WorkingDir:
	}

	hostConfig := &container.HostConfig{

	}

	networkConfig := &network.NetworkConfig{

	}

	c, err := j.DockerClient.ContainerCreate(j.DockerContext, containerConfig, hostConfig, networkConfig, fmt.Sprintf("%s-%s", j.Name, j.Id))
	if err != nil {
		quitPipeline(err)
	}

	j.ContainerId = c.ID

	err = j.DockerClient.ContainerStart(j.DockerContext, j.ContainerId, &types.ContainerStartOptions{})

	if err != nil {
		quitPipeline(err)
	}
}