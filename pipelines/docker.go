package pipelines

import (
	"io"
	// "os"
	"fmt"
	"log"
	"sync"
	"context"
	"path/filepath"
	"encoding/json"
	"leveler/config"
	"leveler/resources"
	uuid "github.com/satori/go.uuid"
	docker "github.com/docker/docker/client"
	types "github.com/docker/docker/api/types"
	//volume "github.com/docker/docker/api/types/volume"
	network "github.com/docker/docker/api/types/network"
	container "github.com/docker/docker/api/types/container"
)

type DockerPipelineJob struct {
	PipelineId string 							`json:"pipeline_id" yaml:"pipeline_id"`
	Id string 									`json:"id" yaml:"id"`
	Name string 								`json:"name" yaml:"name"`
	Datadir string 								`json:"datadir" yaml:"datadir"`
	Parents []*PipelineJob 						`json:"-" yaml:"-"`
	ParentsList []string 						`json:"dependencies" yaml:"dependencies"`
	Children []*PipelineJob 					`json:"-" yaml:"-"`
	ChildrenList []string						`json:"dependents" yaml:"dependents"`
	Inputs map[string]*PipelineInputMapping		`json:"inputs" yaml:"inputs"`
	Outputs map[string]*PipelineOutputMapping	`json:"outputs" yaml:"outputs"`
	DockerContext context.Context 				`json:"-" yaml:"-"`
	DockerClient *docker.Client  				`json:"-" yaml:"-"`
	ContainerId string 							`json:"container_id" yaml:"container_id"`
	JobConfig *resources.Job 		 			`json:"-" yaml:"-"`
	ServerConfig *config.ServerConfig 			`json:"-" yaml:"-"`
	Notifications chan *PipelineJobStatus 		`json:"-" yaml:"-"`
	Logger *log.Logger 							`json:"-" yaml:"-"`
	LogLock *sync.Mutex 						`json:"-" yaml:"-"`
	Status *PipelineJobStatus 					`json:"status" yaml:"status"`
	Color string  								`json:"-" yaml:"-"`
}

func NewDockerPipelineJob(serverConfig *config.ServerConfig, pipelineId string, jobName string, jobConfig *resources.Job, pipelineInputs map[string]*resources.PipelineInput, pipelineOutputs map[string]*resources.PipelineOutput) (DockerPipelineJob, error) {
	jobDataDir := filepath.Join("/data", jobName)

	inputs, err := GenerateInputMappings(jobDataDir, jobConfig, pipelineId, pipelineInputs, pipelineOutputs)
	if err != nil {
		return DockerPipelineJob{}, err
	}

	outputs, err := GenerateOutputMappings(jobDataDir, jobConfig, pipelineId, pipelineOutputs)
	if err != nil {
		return DockerPipelineJob{}, nil
	}

	client, err := docker.NewEnvClient()
	if err != nil {
		return DockerPipelineJob{}, err
	}

	k := DockerPipelineJob{
		Id: uuid.NewV4().String(),
		PipelineId: pipelineId,
		Name: jobName,
		Datadir: jobDataDir,
		ServerConfig: serverConfig,
		JobConfig: jobConfig,
		Inputs: inputs,
		Outputs: outputs,
		Children: []*PipelineJob{},
		Parents: []*PipelineJob{},
		DockerClient: client,
		DockerContext: context.Background(),
		//Volumes: []*types.Volume{},
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

func (j *DockerPipelineJob) GetJson() (string, error) {
	js, err := json.MarshalIndent(j, "", "	")
	if err != nil {
		return fmt.Sprintf("%s", js), err
	}

	return fmt.Sprintf("%s", js), nil
}

func (j *DockerPipelineJob) GetStatus() *PipelineJobStatus {
	return j.Status
}

func (j *DockerPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *DockerPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *DockerPipelineJob) Init() error {
	//serverData := filepath.Join(j.ServerConfig.Datadir, "pipelines", "docker", j.PipelineId, j.Name)
	//for _, i := range j.Inputs {
		// if input comes from a non-local source, create sync script 
		// otherwise, if it comes from a local source, copy the file to the server data dir and bind mount it at runtime
	//}

	//for _, i := range j.Outputs {
		// if output is going to a (non-local) integration destination, create sync script and empty volume
		// otherwise create an empty file and bind mount at runtime
	//}
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

	//var volumes string // TODO
	//var script string // TODO

	// generate env strings
	var env = []string{} // TODO
	
	// create configuration
	containerConfig := &container.Config{
		Env: env,
		Cmd: []string{
			"bash",
			"-c",
			(*j.JobConfig).Command,
		},
		Image: (*j.JobConfig).Image,
		//Volumes: , // TODO
		WorkingDir: j.Datadir,
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

	err = j.DockerClient.ContainerStart(j.DockerContext, j.ContainerId, types.ContainerStartOptions{})

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
        		j.Quit(FAILED, fmt.Sprintf("[runner] Job failed: %v", result.Error.Message))
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

func (j *DockerPipelineJob) Watch(wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(j.Notifications)

	// watch for the job to complete
	for {
        select {
        case <-j.Notifications:
        	return
        }
    }
}

func (j *DockerPipelineJob) Logs(follow, stdout, stderr bool) (io.ReadCloser, error) {
	
		// TODO:
		// - check if logfile still exists; 
		//   - if not, return an error -- the caller will be responsible for any next steps, ie, looking up logs in the log collector
		//   - timeout?
	logs, err := j.DockerClient.ContainerLogs(j.DockerContext, j.ContainerId, types.ContainerLogsOptions{Follow: follow, ShowStdout: stdout, ShowStderr: stderr,})
	if err != nil {
		return logs, err
	}
	return logs, nil
}

func (j *DockerPipelineJob) Cleanup() error {
	j.LogLock.Lock()
	removeOpts := types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks: false,
		Force: false,
	}

   	err := j.DockerClient.ContainerRemove(j.DockerContext, j.ContainerId, removeOpts)
   	j.LogLock.Unlock()

   	if err != nil {
   		return err
   	}

	return nil 
}