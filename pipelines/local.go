package pipelines

import (
	"io"
	"os"
	"fmt"
	"log"
	"sync"
	"errors"
	"strings"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"leveler/config"
	uuid "github.com/satori/go.uuid"
)

type LocalPipelineJob struct {
	Id string 										`json:"id" yaml:"id"`
	Name string 									`json:"name" yaml:"name"`
	Datadir string 									`json:"data_directory" yaml:"data_directory"`
	Workdir string 									`json:"working_directory" yaml:"working_directory"`
	Command string 									`json:"command" yaml:"command"`
	Env map[string]string 							`json:"env" yaml:"env"`
	Inputs map[string]*PipelineInputMapping			`json:"inputs" yaml:"inputs"`
	Outputs map[string]*PipelineOutputMapping 		`json:"outputs" yaml:"outputs"`
	ParentsList []string 							`json:"dependencies" yaml:"dependencies"`
	Parents []*PipelineJob 							`json:"-" yaml:"-"`
	ChildrenList []string 							`json:"dependents" yaml:"dependents"`
	Children []*PipelineJob 						`json:"-" yaml:"-"`
	Notifications chan *PipelineJobStatus 			`json:"-" yaml:"-"`
	Process *os.Process 							`json:"-" yaml:"-"`
	Logger *log.Logger 								`json:"-" yaml:"-"`
	LogLock *sync.Mutex 							`json:"-" yaml:"-"`
	Status *PipelineJobStatus 						`json:"status" yaml:"status"`
	Color string  									`json:"-" yaml:"-"`
	JobConfig *PipelineStep							`json:"-" yaml:"-"`
	ServerConfig *config.ServerConfig 				`json:"-" yaml:"-"`
} 

type SyncStatus struct {
	Status int64
	Message string
	Name string
}

type ProcessStatus struct {
	State *os.ProcessState
	Error error
	Stdout string
	Stderr string
}


func NewLocalPipelineJob(serverConfig *config.ServerConfig, pipelineId string, jobName string, jobConfig *PipelineStep, pipelineInputs []*PipelineInput, pipelineOutputs []*PipelineOutput) (LocalPipelineJob, error) {
	jobDataDir := filepath.Join(serverConfig.Datadir, "pipelines", "local", pipelineId, jobName)

	inputs, err := GenerateInputMapping(jobDataDir, jobSpec, pipelineInputs, pipelineOutputs)
	if err != nil {
		return err, nil
	}

	outputs, err := GenerateOutputMapping(jobDataDir, jobSpec, pipelineOutputs)
	if err != nil {
		return err, nil
	}

	k := LocalPipelineJob{
		Id: uuid.NewV4().String(),
		PipelineId: pipelineId,
		Name: jobName,
		JobConfig: jobConfig,
		ServerConfig: serverConfig,
		Inputs: inputs,
		Outputs: outputs,
		Children: []*PipelineJob{},
		Parents: []*PipelineJob{},
		Command: jobConfig.Command,
		Datadir: filepath.Join(serverConfig.Datadir, "pipelines", pipelineId, jobName),
		Notifications: make(chan *PipelineJobStatus),
		Status: &PipelineJobStatus{},
		LogLock: &sync.Mutex{},
		Color: "white",   // for cycle detection -- not meant to be used within a job
	}

	return k, nil
}

func (j *LocalPipelineJob) Quit(status int64, message string) {
	(*j.Status).Status = status
	(*j.Status).Message = message
	if j.Logger != nil {
		(*j.Logger).Println((*j.Status).Message)
	}

	j.Notifications <- j.Status
	close(j.Notifications)
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

func (j *LocalPipelineJob) GetInputs() map[string]*PipelineInputMapping {
	return j.Inputs
}

func (j *LocalPipelineJob) GetOutputs() map[string]*PipelineOutputMapping {
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

func (j *LocalPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
	j.ChildrenList = append(j.ChildrenList, (*child).GetName())
}

func (j *LocalPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
	j.ParentsList = append(j.ParentsList, (*parent).GetName())
}

func (j *LocalPipelineJob) Init() error {
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

func (j *LocalPipelineJob) SyncInputs(quit chan int8) error {
	(*j.Logger).Println("[data] Syncing inputs...")
	done := make(chan *SyncStatus)

	go func() {
		for name, config := range j.Inputs {
			status := &SyncStatus{
				Name: name,
			}

			if len(config.Integration) == 0 {
				if config.Link {
					err := os.Symlink(config.SrcPath, config.DestPath)
					if err != nil {
						status.Status = FAILED 
						status.Message = fmt.Sprintf("Couldn't create symlink: %v", err)
						done <- status
					} else {
						status.Status = SUCCEEDED 
						done <- status
					}
				} else {
					cmd := exec.Command("cp", "-R", config.SrcPath, config.DestPath)
					out, err := cmd.CombinedOutput()
					if err != nil {
						status.Status = FAILED
						status.Message = fmt.Sprintf("Couldn't copy input from upstream job: %s", out)
						done <- status
					} else {
						status.Status = SUCCEEDED 
						done <- status
					}
				}
				
			} else if config.Integration == "local" {
				cmd := exec.Command("cp", "-R", config.SrcPath, config.DestPath)
				out, err := cmd.CombinedOutput()
				if err != nil {
					status.Status = FAILED 
					status.Message = fmt.Sprintf("Couldn't copy local input: %s", out)
					done <- status
				} else {
					_, err := os.Stat(config.DestPath)
					if err != nil {
						fmt.Println("stat error! %s", err)
					}
					status.Status = SUCCEEDED 
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

				status.Status = FAILED 
				status.Message = "Integration logic not yet implemented!"
				done <- status
			}
		}
		close(done)
	}()

	var messages = []string{}
	for s := range done {
		if s.Status == FAILED {
			messages = append(messages, s.Message)
		}
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n"))	
	} else {
		return nil
	}
}

func (j *LocalPipelineJob) SyncOutputs(quit chan int8) error {
	(*j.Logger).Println("[data] Syncing outputs...")
	done := make(chan *SyncStatus)

	// broadcaster := make(map[string]*chan int8)

	go func() {
		for name, config := range j.Outputs {
			status := &SyncStatus{
				Name: name,
			}

			if config.Integration == "local" {
				cmd := exec.Command("cp", "-R", config.SrcPath, config.DestPath)
				out, err := cmd.CombinedOutput()
				if err != nil {
					status.Status = FAILED
					status.Message = fmt.Sprintf("Couldn't copy output to local destination: %s", out)
					done <- status
					break
				}
			} else if len(config.Integration) != 0 {
				// add channel to broadcaster for propagating quit messages to downloaders
				// upload the output using the integration data
				// err := i.GetExternal().Sync(ctx)
				// if err != nil {
				// 	status.Status = FAILED 
				// 	status.Message = fmt.Sprintf("Couldn't sync external data: %v", err)
				// 	done <- status
				// }
				status.Status = FAILED 
				status.Message = "Integration logic not yet implemented!"
				done <- status
				break
			} else {
				status.Status = SUCCEEDED
				status.Message = fmt.Sprintf("No integration specified... nothing to do!")
				done <- status
			}
		}
		close(done)
	}()

	var messages = []string{}
	for s := range done {
		if s.Status == FAILED {
			messages = append(messages, s.Message)
		}
	}

	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n"))	
	} else {
		return nil
	}
}

func (j *LocalPipelineJob) Run(cancel chan int8) { 
	(*j.Status).Status = INITIALIZING

	err := j.Init()
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error initializing job: %v", err))
		return
	}

	err = j.SyncInputs(cancel)
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error syncing inputs: %v", err))
		return
	}

	proc := exec.Command("/bin/bash", "-c", j.Command)
	stderr, err := proc.StderrPipe()
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error attaching to stderr pipe: %v", err))
		return
	}

	stdout, err := proc.StdoutPipe()
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error attaching to stdout pipe: %v", err))
		return
	}

	proc.Dir = filepath.Join(j.Datadir, j.Workdir)

	err = proc.Start()
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error executing command: %v", err))
		return
	}

	(*j.Logger).Println("[runner] Job started!")
	(*j.Status).Status = RUNNING

	waiter := make(chan *ProcessStatus)
	go func(stdout, stderr io.ReadCloser) {
		state, err := proc.Process.Wait()

		stdoutBytes, _ := ioutil.ReadAll(stdout)
		stderrBytes, _ := ioutil.ReadAll(stderr)

		status := &ProcessStatus{
			State: state,
			Error: err,
			Stdout: fmt.Sprintf("%s", stdoutBytes),
			Stderr: fmt.Sprintf("%s", stderrBytes),
		}

		waiter <- status
		close(waiter)
	}(stdout, stderr)

	var breakFor bool
	for {
		if breakFor {
			break
		} else {
			select {
			case <-cancel:
	        	m := "[runner] Job cancelled!"
	        	err := proc.Process.Kill()
	        	if err != nil {
	        		m += fmt.Sprintf(" Also couldn't kill the underlying process: %v", err)
	        	}
	        	j.Quit(CANCELLED, m)
	        	return
	        case status, ok := <-waiter:
	        	if ok {
	        		if status.State.Exited() {
		        		if status.State.Success() {
		        			j.Logger.Println("[runner] Job succeeded!")
		        			breakFor = true
		        			break
		        		} else {
		        			j.Quit(FAILED, fmt.Sprintf("[runner] Job failed: status=%+v, error=%v, stdout=%s, stderr=%s", status.State, status.Error, status.Stdout, status.Stderr))
		        			return
		        		}
		        	} 
	        	}
			}
		}
	}

	err = j.SyncOutputs(cancel)
	if err != nil {
		j.Quit(FAILED, fmt.Sprintf("[runner] Error syncing outputs: %v", err))
		return
	}

	j.Quit(SUCCEEDED, "[runner] OK")
}

func (j *LocalPipelineJob) Watch(wg *sync.WaitGroup) {
	defer wg.Done()

	// watch for the job to complete
	for {
        select {
        case <-j.Notifications:
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

func (j *LocalPipelineJob) Cleanup() error {
	
	j.LogLock.Lock()
	err := os.RemoveAll(j.Datadir)
	j.LogLock.Unlock()

	if err != nil {
		return err
	}

	return nil
}