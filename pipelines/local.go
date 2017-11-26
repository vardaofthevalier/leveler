package pipelines

import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
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
	Inputs []*PipelineInput
	Outputs []*PipelineOutput 				
	Parents []*PipelineJob 					`json:"dependencies" yaml:"dependencies"`
	Children []*PipelineJob 				`json:"dependents" yaml:"dependents"`
	Notifications chan *PipelineJobStatus 	`json:"-" yaml:"-"`
	Process *os.Process 					`json:"-" yaml:"-"`
	Logger *log.Logger 						`json:"-" yaml:"-"`
	Status *PipelineJobStatus 				`json:"status" yaml:"status"`
	Color string  							`json:"-" yaml:"-"`
} 

func NewLocalPipelineJob(serverConfig *config.ServerConfig, jobConfig *PipelineStep) (LocalPipelineJob, error) {
	// LEVELER_DATA default location:  /var/lib/leveler
	// create workdir under <LEVELER_DATA>/pipelines/jobs/<job-id>/
	// resolve inputs (i.e., create links) <LEVELER_DATA>/pipelines/jobs/<dependency-id>/outputs/<output-name> -> /var/lib/leveler/pipelines/jobs/<job-id>/inputs/<input-name>
	// create directory for outputs <LEVELER_DATA>/pipelines/jobs/<job-id>/outputs/<output-name>
	// generate script <LEVELER_DATA>/pipelines/jobs/<job-id>/run.sh
	// create *fsnotify.Watcher for watching for the creation of a "success" file <LEVELER_DATA>/pipelines/jobs/<job-id>/success
	jobId := uuid.NewV4().String()

	k := LocalPipelineJob{
		Id: jobId,
		Name: jobConfig.Name,
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

func (j *LocalPipelineJob) AddChild(child *PipelineJob) {
	j.Children = append(j.Children, child)
}

func (j *LocalPipelineJob) AddParent(parent *PipelineJob) {
	j.Parents = append(j.Parents, parent)
}

func (j *LocalPipelineJob) Init() error {
	err := os.Mkdir(j.Datadir)
	if err != nil {
		return err
	}

	err = os.Mkdir(filepath.Join(j.Datadir, "inputs"))
	if err != nil {
		return err
	}

	for _, i := range j.Inputs {
		err = os.Mkdir(filepath.Join(j.Datadir, "inputs", i.Name))
		if err != nil {
			return err
		}
	}

	err = os.Mkdir(filepath.Join(j.Datadir, "outputs"))
	if err != nil {
		return err
	}

	for _, o := range j.Outputs {
		err = os.Mkdir(filepath.Join(j.Datadir, "outputs", o.Name))
		if err != nil {
			return err
		}
	}

	logFile, err := os.OpenFile(filepath.Join(j.Datadir, "log"), os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	j.Logger = log.New(logFile, log.LUTC, 0)

	err = j.Watcher.Add(filepath.Join(j.Datadir, "success"))
	if err != nil {
		return err
	}

	err = j.Watcher.Add(filepath.Join(j.Datadir, "failure"))
	if err != nil {
		return err
	}

	return nil
}

func (j *LocalPipelineJob) SyncInputs(ctx context.Context) error {
	j.Logger.Println("[init] Syncing inputs...")
	done := make(chan bool)

	for _, i := range jobConfig.Inputs {
		go func() {
			switch i.Src.(type) {
			case *PipelineInput.PipelineJob:
				var parentName string
				for _, p := range j.Parents {
					for _, o := range p.Outputs {
						if o.Name == i.Name {
							parentName = p.Name
						}
					}
				}

				if len(parentName) == 0 {
					return errors.New(fmt.Sprintf("Job '%s' has no parents producing input '%s'", j.Name, i.Name))
				}
				
				err := os.Symlink(filepath.Join(filepath.Dirname(j.Datadir), parentName, "outputs", i.Name), filepath.Join(j.Datadir, "inputs", i.Name))
				if err != nil {
					return err
				}

			case *PipelineInput.External:
				i.External.Datamap.Sync(ctx)
			default:
				return errors.New(fmt.Sprintf("Unknown type '%v' for input source"), i.Src.(type))
			}
			
			err = os.Symlink(filepath.Join(j.Datadir, "inputs", i.Name, i.Datamap.Dest), filepath.Join(j.Datadir, i.Datamap.Dest))

			done <- true
		}()
	}

	counter := len(jobConfig.Inputs)

	for {
		select {
			case <-done:
				if counter == 0 {
					return nil 
				} else {
					counter -= 1
				}
			case <-ctx.Done():
				return errors.New("Context cancelled")
		}
	}
}

func (j *LocalPipelineJob) SyncOutputs(ctx context.Context) error {
	j.Logger.Println("[data] Syncing outputs...")
	done := make(chan bool)

	for _, o := range jobConfig.Outputs {
		go func() {
			err = os.Symlink(filepath.Join(j.Datadir, o.Datamap.Src), filepath.Join(j.Datadir, "outputs", o.Name, o.Datamap.Src))
			if err != nil {
				return err
			}

			switch i.Dest.(type) {
			case *PipelineInput.External:
				i.External.Datamap.Sync(ctx)

			default:
				return errors.New(fmt.Sprintf("Unknown type '%v' for output destination"), i.Dest.(type))
			}
			
			done <- true
		}()
	}

	counter := len(jobConfig.Outputs)

	for {
		select {
			case <-done:
				if counter == 0 {
					return nil 
				} else {
					counter -= 1
				}
			case <-ctx.Done():
				return errors.New("Context cancelled")
		}
	}
}

func (j *LocalPipelineJob) Run(ctx context.Context) { 
	var quit = func() {
		j.Status.Status = FAILED
		j.Logger.Println(status.Message)
		j.Notifications <- j.Status
		return 
	}

	j.Status.Status = INITIALIZING

	err := j.Init()
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error initializing job: %v", err)
		quit()
	}

	err = j.SyncInputs(ctx)
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error syncing inputs: %v", err)
		quit()
	}


	proc = exec.Command("bash", "-c", j.Command)
	err = proc.Start()
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error executing command: %v", err)
		quit()
	}

	j.Status.Status = RUNNING

	for {
		select {
		case <-ctx.Done():
			j.Status.Status = CANCELLED
        	j.Status.Message = "[runner] Job cancelled!"
        	err := proc.Process.Kill()
        	if err != nil {
        		j.Status.Message += fmt.Sprintf(" Also couldn't kill the underlying process: %v", err)
        	}

        	j.Logger.Println(j.Status.Message)
        	j.Notifications <- j.Status
        	return
        default:
        	if proc.Process.Exited() {
        		if proc.Process.Success() {
        			j.Status.Message = "[runner] Job succeeded!"
        		} else {
        			j.Status.Status = FAILED
        			j.Status.Message = fmt.Sprintf("[runner] Job failed: %v", )
        		}
        		j.Logger.Println(j.Status.Message)
        		j.Notifications <- j.Status
        		return
        	}
		}
		time.Sleep(time.Duration(1)*time.Second)
	}

	err = j.SyncOutputs(ctx)
	if err != nil {
		j.Status.Message = fmt.Sprintf("[runner] Error syncing outputs: %v", err)
		quit()
	}

	j.Status.Status = SUCCEEDED
	j.Status.Message = "[runner] OK"
	j.Notifications <- j.Status
}

func (j *LocalPipelineJob) Watch(report chan *PipelineJobStatus, wg *sync.WaitGroup) {
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

func (j *LocalPipelineJob) Cleanup() error {
	// delete working directory
	err := os.RemoveAll(j.Datadir)
	if err != nil {
		return err
	}

	return nil
}