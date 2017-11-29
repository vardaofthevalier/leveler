package pipelines

import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
	"errors"
	"context"
	"os/exec"
	"path/filepath"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
)

type LocalPipelineJob struct {
	Id string 								`json:"id" yaml:"id"`
	Name string 							`json:"name" yaml:"name"`
	Datadir string 							`json:"data_directory" yaml:"data_directory"`
	Workdir string 							`json:"working_directory" yaml:"working_directory"`
	Command string 							`json:"command" yaml:"command"`
	Env map[string]string 					`json:"env" yaml:"env"`
	Inputs map[string]*PipelineInput		`json:"inputs" yaml:"inputs"`
	Outputs map[string]*PipelineOutput 		`json:"outputs" yaml:"outputs"`
	Parents []*PipelineJob 					`json:"dependencies" yaml:"dependencies"`
	Children []*PipelineJob 				`json:"dependents" yaml:"dependents"`
	Notifications chan *PipelineJobStatus 	`json:"-" yaml:"-"`
	Process *os.Process 					`json:"-" yaml:"-"`
	Logger *log.Logger 						`json:"-" yaml:"-"`
	Status *PipelineJobStatus 				`json:"status" yaml:"status"`
	Color string  							`json:"-" yaml:"-"`
	Config *PipelineStep					`json:"-" yaml:"-"`
} 

type SyncStatus struct {
	Status int64
	Message string
	Name string
}

type ProcessStatus struct {
	State *os.ProcessState
	Error error
}

func NewLocalPipelineJob(serverConfig *config.ServerConfig, jobName string, jobConfig *PipelineStep, inputs map[string]*PipelineInput, outputs map[string]*PipelineOutput) (LocalPipelineJob, error) {
	// LEVELER_DATA default location:  /var/lib/leveler
	// create workdir under <LEVELER_DATA>/pipelines/jobs/<job-id>/
	// resolve inputs (i.e., create links) <LEVELER_DATA>/pipelines/jobs/<dependency-id>/outputs/<output-name> -> /var/lib/leveler/pipelines/jobs/<job-id>/inputs/<input-name>
	// create directory for outputs <LEVELER_DATA>/pipelines/jobs/<job-id>/outputs/<output-name>
	// generate script <LEVELER_DATA>/pipelines/jobs/<job-id>/run.sh
	jobId := uuid.NewV4().String()

	k := LocalPipelineJob{
		Id: jobId,
		Name: jobName,
		Config: jobConfig,
		Inputs: inputs,
		Outputs: outputs,
		Command: jobConfig.Command,
		Datadir: filepath.Join(serverConfig.Datadir, "pipelines/jobs", jobId),
		Notifications: make(chan *PipelineJobStatus),
		Status: &PipelineJobStatus{},
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

func (j *LocalPipelineJob) GetConfig() *PipelineStep {
	return j.Config
}

func (j *LocalPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *LocalPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *LocalPipelineJob) Init() error {
	fmt.Printf("Inititializing job: %+v\n", j.Name)
	
	err := os.MkdirAll(j.Datadir, 0700)
	if err != nil {
		return err
	}

	logFile, err := os.OpenFile(filepath.Join(j.Datadir, "log"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	j.Logger = log.New(logFile, fmt.Sprintf("[job %s]", j.Id), log.LUTC)

	return nil
}

func (j *LocalPipelineJob) SyncInputs(ctx context.Context) error {
	(*j.Logger).Println("[init] Syncing inputs...")
	done := make(chan *SyncStatus)

	for name, config := range j.Inputs {
		go func() {
			status := &SyncStatus{
				Name: name,
			}

			if len(config.Integration) == 0 {
				// the input is from another job -- need to test that the file exists and create a symlink
				err := os.Symlink(filepath.Join(filepath.Dir(j.Datadir), j.Outputs[config.From].From, config.From), filepath.Join(j.Datadir, name))
				if err != nil {
					status.Status = FAILED 
					status.Message = fmt.Sprintf("Couldn't create symlink: %v", err)
					done <- status
				}
			} else {
				// download the input using the integration data
				// err := i.GetExternal().Sync(ctx)
				// if err != nil {
				// 	status.Status = FAILED 
				// 	status.Message = fmt.Sprintf("Couldn't sync external data: %v", err)
				// 	done <- status
				// }
			}
		}()
	}

	for s := range done {
		if s.Status == FAILED {
			return errors.New(s.Message)
		}
	}

	return nil
}

func (j *LocalPipelineJob) SyncOutputs(ctx context.Context) error {
	(*j.Logger).Println("[data] Syncing outputs...")
	done := make(chan *SyncStatus)

	for name, config := range j.Outputs {
		go func() {
			status := &SyncStatus{
				Name: name,
			}

			if len(config.Integration) != 0 {
				// upload the output using the integration data
				// err := i.GetExternal().Sync(ctx)
				// if err != nil {
				// 	status.Status = FAILED 
				// 	status.Message = fmt.Sprintf("Couldn't sync external data: %v", err)
				// 	done <- status
				// }
			} else {
				status.Status = SUCCEEDED
				status.Message = fmt.Sprintf("No integration specified... nothing to do!")
				done <- status
			}
		}()
	}

	for s := range done {
		if s.Status == FAILED {
			return errors.New(s.Message)
		}
	}

	return nil
}

func (j *LocalPipelineJob) Run(ctx context.Context) { 
	var quit = func() {
		(*j.Status).Status = FAILED
		if j.Logger != nil {
			(*j.Logger).Println((*j.Status).Message)
		}
		j.Notifications <- j.Status
		return 
	}

	(*j.Status).Status = INITIALIZING

	err := j.Init()
	if err != nil {
		(*j.Status).Message = fmt.Sprintf("[runner] Error initializing job: %v", err)
		quit()
	}

	err = j.SyncInputs(ctx)
	if err != nil {
		(*j.Status).Message = fmt.Sprintf("[runner] Error syncing inputs: %v", err)
		quit()
	}


	proc := exec.Command("bash", "-c", j.Command)
	err = proc.Start()
	if err != nil {
		(*j.Status).Message = fmt.Sprintf("[runner] Error executing command: %v", err)
		quit()
	}

	(*j.Status).Status = RUNNING

	waiter := make(chan *ProcessStatus)
	go func() {
		state, err := proc.Process.Wait()
		status := &ProcessStatus{
			State: state,
			Error: err,
		}

		waiter <- status
	}()

	for {
		select {
		case <-ctx.Done():
			(*j.Status).Status = CANCELLED
        	(*j.Status).Message = "[runner] Job cancelled!"
        	err := proc.Process.Kill()
        	if err != nil {
        		(*j.Status).Message += fmt.Sprintf(" Also couldn't kill the underlying process: %v", err)
        	}

        	j.Logger.Println((*j.Status).Message)
        	j.Notifications <- j.Status
        	return
        case status := <-waiter:
        	if status.State.Exited() {
        		if status.State.Success() {
        			(*j.Status).Message = "[runner] Job succeeded!"
        			j.Logger.Println((*j.Status).Message)
        			j.Notifications <- j.Status
        			break
        		} else {
        			(*j.Status).Status = FAILED
        			(*j.Status).Message = fmt.Sprintf("[runner] Job failed: %v", status.Error)
        			j.Logger.Println((*j.Status).Message)
        			j.Notifications <- j.Status
        			return
        		}
        	}
		}
		time.Sleep(time.Duration(1)*time.Second)
	}

	err = j.SyncOutputs(ctx)
	if err != nil {
		(*j.Status).Message = fmt.Sprintf("[runner] Error syncing outputs: %v", err)
		quit()
	}

	(*j.Status).Status = SUCCEEDED
	(*j.Status).Message = "[runner] OK"
	j.Notifications <- j.Status
}

func (j *LocalPipelineJob) Watch(report chan *PipelineJobStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(j.Notifications)

	// watch for the job to complete
	for {
        select {
        case n:= <-j.Notifications:
        	report <- n
        	return
        }
    }
}

func (j *LocalPipelineJob) Cleanup() error {
	// delete working directory
	err := os.RemoveAll(j.Datadir)
	if err != nil {
		return err
	}

	return nil
}