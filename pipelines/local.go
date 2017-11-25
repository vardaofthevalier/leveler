package pipelines

import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
	"context"
	"os/exec"
	"filepath"
	"leveler/config"
	"github.com/fsnotify/fsnotify"
	uuid "github.com/satori/go.uuid"
)

type LocalPipelineJob struct {
	Id string
	Name string
	Datadir string
	Workdir string
	Command string
	Env map[string]string
	Inputs []*PipelineInput 
	Outputs []*PipelineOutput
	Parents []*PipelineJob
	Children []*PipelineJob
	Watcher *fsnotify.Watcher
	Process *os.Process
	Logger *log.Logger
	Color string  // for detecting cycles in the job graph -- not intended to be used within a job
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
		Watcher: *fsnotify.NewWatcher(),
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

func (j *LocalPipelineJob) SyncInputs() error {
	j.Logger.Println("[init] Syncing inputs...")

	for _, i := range jobConfig.Inputs {
		// if input is from an upstream pipeline job:
		// - verify that a parent actually produces the output
		// - create symlink to output in jobdir
		// otherwise if input is from an external source:
		// - init external datasource
		// - sync external datasource
	}
}

func (j *LocalPipelineJob) SyncOutputs() error {
	j.Logger.Println("[cleanup] Syncing outputs...")

	for _, o := range jobConfig.Outputs {
		// if output is going to an external destination:
		// - init external datasource
		// - sync external datasource
	}
}

func (j *LocalPipelineJob) Run() { 
	var quit = func(err error) {
		_, err2 := os.Create(filepath.Join(j.Datadir, "failure"))
		if err2 != nil {
			j.Logger.Printf("[runner] Error creating status file: %v", err)
		}
		return 
	}

	// TODO: figure out how to cancel this goroutine when the cancel function is run
	err = j.SyncInputs(ctx)
	if err != nil {
		j.Logger.Printf("[runner] Error syncing inputs: %v", err)
		quit(err)
	}

	proc = exec.Command("bash", "-c", j.Command)
	err = proc.Run()
	if err != nil {
		j.Logger.Printf("[runner] Error executing command: %v", err)
		quit(err)
	}

	j.Process = proc.Process

	err = j.SyncOutputs(ctx)
	if err != nil {
		log.Printf("[runner] Error syncing outputs: %v", err)
		quit(err)
	}

	// create success file to give the Watch function something to watch for
	_, err = os.Create(filepath.Join(j.Datadir, "success"))
	if err != nil {
		log.Printf("[runner] Error creating status file: %v", err)
		quit(err)
	}
}

func (j *LocalPipelineJob) Watch(ctx context.Context, statuses chan *PipelineJobStatus, wg *sync.WaitGroup) {
	defer wg.Done()
	defer j.Watcher.Close()

	// watch for the job to complete
	var status *PipelineJobStatus
	for {
        select {
        case event := <- j.Watcher.Events:
            if event.Op == fsnotify.Create {
            	if event.Name == filepath.Join(j.Datadir, "success") {
            		log.Println("[watcher] Job succeeded!")
            		status.Status = SUCCEEDED
            	} elif event.Name == filepath.Join(j.Datadir, "failure") {
            		log.Println("[watcher] Job failed!")
            		status.Status = FAILED
            	}

            	statuses <- status
                return	
            }
        case <-ctx.Done():
        	log.Println("[watcher] Job cancelled!")
        	j.Process.Kill()
        	status.Status = CANCELLED
        	return
        case err := <-watcher.Errors:
            log.Printf("[watcher] Error watching files: %s\n", err)
            j.Process.Kill()
            status.Status = KILLED
            return
        }

        time.Sleep(time.Duration(1 * time.Second))
    }
}

func (j *LocalPipelineJob) Cleanup() error {
	// delete working directory
	return nil
}