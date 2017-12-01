package pipelines

import (
	"io"
	"os"
	"fmt"
	"log"
	"sync"
	"context"
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
	Parents []*PipelineJob 						`json:"-" yaml:"-"`
	ParentsList []string 						`json:"dependencies" yaml:"dependencies"`
	Children []*PipelineJob 					`json:"-" yaml:"-"`
	ChildrenList []string						`json:"dependents" yaml:"dependents"`
	Inputs []*PipelineInputMapping				`json:"inputs" yaml:"inputs"`
	Outputs []*PipelineOutputMapping			`json:"outputs" yaml:"outputs"`
	DockerContext context.Context 				`json:"-" yaml:"-"`
	DockerClient *docker.Client  				`json:"-" yaml:"-"`
	ContainerId string 							`json:"containerId" yaml:"containerId"`
	JobConfig *PipelineStep 					`json:"-" yaml:"-"`
	ServerConfig *config.ServerConfig 			`json:"-" yaml:"-"`
	Notifications chan *PipelineJobStatus 		`json:"-" yaml:"-"`
	Logger *log.Logger 							`json:"-" yaml:"-"`
	LogLock *sync.Mutex 						`json:"-" yaml:"-"`
	Status *PipelineJobStatus 					`json:"status" yaml:"status"`
	Color string  								`json:"-" yaml:"-"`
}

func NewDockerPipelineJob(serverConfig *config.ServerConfig, pipelineId string, jobName string, jobConfig *PipelineStep, pipelineInputs []*PipelineInput, pipelineOutputs []*PipelineOutput, dockerClient *docker.Client) (DockerPipelineJob, error) {
	jobDataDir := filepath.Join("/data", jobName)

	inputs, err := GenerateInputMapping(jobDataDir, jobSpec, pipelineInputs, pipelineOutputs)
	if err != nil {
		return err, nil
	}

	outputs, err := GenerateOutputMapping(jobDataDir, jobSpec, pipelineOutputs)
	if err != nil {
		return err, nil
	}

	k := DockerPipelineJob{
		Id: uuid.NewV4().String(),
		PipelineId: pipelineId,
		Name: jobName,
		ServerConfig: serverConfig,
		JobConfig: jobConfig,
		Inputs: inputs,
		Outputs: outputs,
		Children: []*PipelineJob{},
		Parents: []*PipelineJob{},
		DockerClient: dockerClient,
		DockerContext: context.Background(),
		Volumes: []*types.Volume{},
		Notifications: make(chan *PipelineJobStatus),
		Status: &PipelineJobStatus{},
		LogLock: &sync.Mutex{},
		Color: "white",   // for cycle detection -- not meant to be used within a job
	}

	return k, nil
}

func (j *DockerPipelineJob) Quit(status int64, message string) {
	// TODO
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

func (j *DockerPipelineJob) GetInputs() map[string]*PipelineInputMapping {
	return j.Inputs
}

func (j *DockerPipelineJob) GetOutputs() map[string]*PipelineOutputMapping {
	return j.Outputs
}

func (j *LocalPipelineJob) GetJson() (string, error) {
	js, err := json.MarshalIndent(j, "", "	")
	if err != nil {
		return fmt.Sprintf("%s", js), err
	}

	return fmt.Sprintf("%s", js), nil
}

func (j *LocalPipelineJob) GetStatus() *PipelineJobStatus {
	return j.Status
}

func (j *DockerPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *DockerPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *DockerPipelineJob) Init() error {
	serverData := filepath.Join(j.ServerConfig.Datadir, "pipelines", "docker", j.PipelineId, j.Name)
	// generate input sync script
	// - read template from file
	// - evaluate template
	// - write to server data directory
	// generate output sync script
	// - read template from file
	// - evaluate template
	// - write to server data directory
	// generate job script
	// - read template from file
	// - evaluate template
	// - write to server data directory
	// create a volume containing scripts and add to volumes list
	// add volumes from upstream jobs if necessary and to volumes list
	// create an empty volume for outputs if necessary and add to volumes list
	// render env map into a slice of "<KEY>=<VALUE> strings"

	var volumes string // TODO
	var script string // TODO

	// generate env strings
	var env = []string{} // TODO
	
	// create configuration
	containerConfig := &container.Config{
		Env: env,
		Cmd: (*j.Config).Command,
		Image: (*j.Config).Image,
		Volumes: "", // TODO
		WorkingDir: "", // TODO
	}

	hostConfig := &container.HostConfig{
		// TODO -- relevant values are likely going to be set in the server config
	}

	networkConfig := &network.NetworkingConfig{
		// TODO -- relevant values are likely going to be set in the server config
	}

	c, err := j.DockerClient.ContainerCreate(j.DockerContext, containerConfig, hostConfig, networkConfig, fmt.Sprintf("%s-%s", j.Name, j.Id))
	if err != nil {
		return err
	}

	j.ContainerId = c.ID
	return nil
}

func (j *DockerPipelineJob) Run(quit chan int8) { 
	err := j.Init()
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error initializing job: %v", err))
		return
	}

	err = j.DockerClient.ContainerStart(j.DockerContext, j.ContainerId, &types.ContainerStartOptions{})

	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error starting container: %v", err))
		return
	}

	// TODO: apply timeout?
	resultC, errC := j.DockerClient.ContainerWait(j.DockerContext, j.ContainerId, container.WaitConditionNextExit)
	j.Status.Status = RUNNING

	for {
		select {
        case err := <-errC:
        	j.Quit(FAILED, fmt.Sprintf("[runner] Job failed: %v", err))
			return
        case result := <-resultC:
        	if result.Error != nil {
        		j.Quit(FAILED, "[runner] Job failed: %v", result.Error.Message)
        	} else {
        		j.Quit(SUCCEEDED, "[runner] OK")
        	}
        	return
        case <-quit:
			j.Status.Status = CANCELLED
        	m := "[runner] Job cancelled!"

        	err := j.DockerClient.ContainerKill(j.DockerContext, j.ContainerId, "KILL")
        	if err != nil {
        		m += fmt.Sprintf(" Also couldn't kill the job's container: %v", err)
        	}

        	j.Quit(CANCELLED, m)
        	return
		}
	}
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

// func (j *LocalPipelineJob) Logs(stream *io.Pipe) error {
	
// 		// TODO:
// 		// - check if logfile still exists; 
// 		//   - if not, stream from Fluentd (or error if fluentd isn't integrated yet)
// 		//   - otherwise grab the mutex lock; tail log file; lock file again
// 		//   - need to have a reasonable timeout for this
	
// 	return nil
// }

func (j *DockerPipelineJob) Cleanup() error {
	j.LogfileLock.Lock()
	removeOpts := &types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks: false,
		Force: false,
	}

   	err := j.DockerClient.ContainerRemove(j.DockerContext, j.ContainerId, removeOpts)
   	j.LogfileLock.Unlock()

   	if err != nil {
   		return err
   	}

	return nil 
}