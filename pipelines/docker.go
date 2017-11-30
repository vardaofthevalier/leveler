package pipelines

import (
	"os"
	"fmt"
	"log"
	"sync"
	"context"
	//"os/exec"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
	docker "github.com/docker/docker/client"
	types "github.com/docker/docker/api/types"
	volume "github.com/docker/docker/api/types/volume"
	network "github.com/docker/docker/api/types/network"
	container "github.com/docker/docker/api/types/container"
)

type DockerPipelineJob struct {
	Id string 									`json:"id" yaml:"id"`
	Name string 								`json:"name" yaml:"name"`
	Parents []*PipelineJob 						`json:"dependencies" yaml:"dependencies"`
	Children []*PipelineJob 					`json:"dependents" yaml:"dependents"`
	Inputs []*PipelineInput						`json:"inputs" yaml:"inputs"`
	Outputs []*PipelineOutput 					`json:"outputs" yaml:"outputs"`
	DockerContext context.Context 				`json:"-" yaml:"-"`
	DockerClient *docker.Client  				`json:"-" yaml:"-"`
	ContainerId string 							`json:"containerId" yaml:"containerId"`
	Config *PipelineStep 						`json:"-" yaml:"-"`
	Volumes []*types.Volume 						`json:"volumes" yaml:"volumes"`
	Notifications chan *PipelineJobStatus 		`json:"-" yaml:"-"`
	Process *os.Process 						`json:"-" yaml:"-"`
	Logger *log.Logger 							`json:"-" yaml:"-"`
	Status *PipelineJobStatus 					`json:"status" yaml:"status"`
	Color string  								`json:"-" yaml:"-"`
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

	k = DockerPipelineJob{
		Id: uuid.NewV4().String(),
		Name: jobConfig.Name,
		DockerClient: client,
		DockerContext: context.Background(),
		Volumes: []*types.Volume{},
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

func (j *DockerPipelineJob) Init() error {
	// generate script
	var script string // TODO

	// generate env map
	var env map[string]string  // TODO
	
	// create configuration
	containerConfig := &container.Config{
		// TODO:
		// Env:
		// Cmd:
		// Image:
		// Volumes:
		// WorkingDir:
	}

	hostConfig := &container.HostConfig{
		// TODO
	}

	networkConfig := &network.NetworkingConfig{
		// TODO
	}

	c, err := j.DockerClient.ContainerCreate(j.DockerContext, containerConfig, hostConfig, networkConfig, fmt.Sprintf("%s-%s", j.Name, j.Id))
	if err != nil {
		return err
	}

	j.ContainerId = c.ID
	return nil
}

func (j *DockerPipelineJob) SyncInputs() error {
	// TODO: create new volumes for new inputs and attach integration containers to volumes to perform download
	var volumeCreate volume.VolumesCreateBody
	for _, i := range j.Inputs {
		// TODO: volume create body
		newVolume, err := j.DockerClient.VolumeCreate(j.DockerContext, volumeCreate)
		if err != nil {
			return err
		}

		// TODO: run container to populate the volume using the correct integration
		j.Volumes = append(volumes, newVolume)
	}
}

func (j *DockerPipelineJob) SyncOutputs() error {
	// TODO: attach integration containers to output volumes to perform upload

	return nil
}

func (j *DockerPipelineJob) Run(ctx context.Context) { 
	var quit= func() {
		j.Status.Status = FAILED
		j.Logger.Println(j.Status.Message)
		j.Notifications <- j.Status
		return
	}

	err = j.SyncInputs(ctx)
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error syncing inputs: %v", err)
		quit()
	}

	err := j.Init()
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error initializing job: %v", err)
		quit()
	}

	err = j.DockerClient.ContainerStart(j.DockerContext, j.ContainerId, &types.ContainerStartOptions{})

	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error starting container: %v", err)
		quit()
	}

	// TODO: apply timeout?
	resultC, errC := j.DockerClient.ContainerWait(j.DockerContext, j.ContainerId, container.WaitConditionNextExit)
	j.Status.Status = RUNNING

	for {
		select {
        case errC <- err:
        	result := <-resultC
			j.Status.Status = FAILED
			j.Status.Message = fmt.Sprintf("[runner] Job failed: %v", result.Error.Message)
			break
        case resultC <- result:
        	j.Status.Status = SUCCEEDED
        	j.Status.Message = "[runner] OK"
        	break
        case <-ctx.Done():
			j.Status.Status = CANCELLED
        	j.Status.Message = "[watcher] Job cancelled!"

        	err := j.DockerClient.ContainerKill(j.DockerContext, j.ContainerId, "KILL")
        	if err != nil {
        		j.Status.Message += fmt.Sprintf(" Also couldn't kill the job's container: %v", err)
        	}
        	break
		}
	}

	j.Logger.Println(j.Status.Message)
	j.Notifications <- j.Status
}

func (j *DockerPipelineJob) Watch(report chan *PipelineJobStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(j.Notifications)

	// watch for the job to complete
	var status *PipelineJobStatus
	for {
        select {
        case j.Notifications <- n:
        	report <- n
        	return
        }
    }
}

func (j *DockerPipelineJob) Cleanup() error {
	// TODO: decide when this will run and how volumes and links should be treated
	removeOpts := &types.ContainerRemoveOptions{
        		RemoveVolumes: false,
        		RemoveLinks: false,
        		Force: false,
        	}

   	err := j.DockerClient.ContainerRemove(j.DockerContext, j.ContainerId, removeOpts)
   	if err != nil {
   		return err
   	}

	return nil 
}