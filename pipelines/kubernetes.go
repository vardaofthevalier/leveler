package pipelines

// import (
// 	//"fmt"
// 	"sync"
// 	//"os/exec"
// 	"leveler/config"
// 	uuid "github.com/satori/go.uuid"
// )

// // type NewJobFn func(*config.ServerConfig, *PipelineStep) (PipelineJob, error) 

// // type KubernetesPipelineJobInit struct {
// // 	Credentials []*ExternalDataProvider
// // 	Data []*ExternalDataProvider
// // }

// type KubernetesPipelineJob struct {
// 	Id string
// 	Name string
// 	Namespace string
// 	Container string
// 	Workdir string
// 	Script string
// 	Env *map[string]string
// 	VolumeSizeGi int
// 	Parents []*PipelineJob
// 	Children []*PipelineJob
// 	Color string
// }

// var pipelineJobPvcTemplate = `
// {
// 	"apiVersion": "v1",
// 	"kind": "PersistentVolumeClaim",
// 	"metadata": {
// 		"name": "{{.Name}}-{{.Id}}",
// 		"namespace": "{{.Namespace}}"
// 	},
// 	"spec": {
// 		"resources": {
// 			"requests": {
// 				"storage": "{{.VolumeSizeGi}}Gi"
// 			}
// 		},
// 		"accessModes": [
// 			"ReadWriteOnce"
// 		]
// 	}
// }
// `

// var pipelineJobPodTemplate = `
// {
// 	"apiVersion": "v1",
// 	"kind": "Pod",
// 	"metadata": {
// 		"name": "{{.Name}}-{{.Id}}",
// 		"namespace": "{{.Namespace}}"
// 	},
// 	"spec": {
// 		{{ if .Init != nil -}}
// 		"initContainers": [
// 		{{ if .Init.Credentials != nil -}}
// 		{{ for $i, $e := range .Init.Credentials -}}
// 			{{ if $i < len(.Init.Credentials)-1 -}}
// 			{
// 				"image": "{{ $e.GetContainer() }}",
// 				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
// 				"command": "/configmap/auth{{ $i }}.sh"
// 			},
// 			{{ else if .Init.Data != nil -}}
// 			{
// 				"image": "{{ $e.GetContainer() }}",
// 				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
// 				"command": "/configmap/auth{{ $i }}.sh"
// 			},
// 			{{ else -}}
// 			{
// 				"image": "{{ $e.GetContainer() }}",
// 				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
// 				"command": "/configmap/auth{{ $i }}.sh"
// 			}
// 			{{- end }}
// 		{{- end }}
// 		{{- end }}
// 		{{ if .Init.Data != nil -}}
// 		{{ for $i, $e := range .Init.Data -}}
// 			{{ if $i < len(.Init.Data)-1 -}}
// 			{
// 				"image": "{{ $e.GetContainer() }}",
// 				"name": "data{{ $i }}-{{.Name}}-{{.Id}}",
// 				"command": "/configmap/data{{ $i }}.sh"
// 			},
// 			{{ else -}}
// 			{
// 				"image": "{{ $e.GetContainer() }}",
// 				"name": "data{{ $i }}-{{.Name}}-{{.Id}}",
// 				"command": "/configmap/data{{ $i }}.sh"
// 			}
// 			{{- end }}
// 		{{- end }}
// 		{{- end }}
// 		],
// 		{{- end }}
// 		"containers": [
// 			"image": "{{.Container}}",
// 			"name": "{{.Name}}-{{.Id}}",
// 			"command": "/configmap/job.sh",
// 			"volumeMounts": [
// 				{
// 					"name": "data",
// 					"mountPath": "/data"
// 				},
// 				{
// 					"name": "configmap",
// 					"mountPath": "/configmap"
// 				}
// 			],
// 			"workingDir": "/data/{{.Workdir}}",
// 			"lifecyle"
// 		],
// 		"volumes": [
// 			{
// 				"name": "data",
// 				"persistentVolumeClaim": {
// 					"claimName": "{{.Name}}-{{.Id}}"
// 				},
// 				"name": "configmap",
// 				"configMap": {
// 					"name": {{.Name}}-{{.Id}}
// 				}
// 			}
// 		]
// 	}
// }
// `

// var pipelineJobScriptConfigMapTemplate = `
// {
// 	"apiVersion": "v1",
// 	"kind": "ConfigMap",
// 	"metadata": {
// 		"name": "{{.Name}}-{{.Id}}",
// 		"namespace": "{{.Namespace}}"
// 	},
// 	"data": {
// 		{{ if .Init == nil -}}
// 		"job.sh": "{{.Script}}"
// 		{{ else -}}
// 		"job.sh": "{{.Script}}",
// 		{{ end -}}
// 		{{ if .Init.Credentials != nil -}}
// 		{{ for i, e := range .Init.Credentials -}}
// 		{{ if i < len(.Init.Credentials)-1 -}}
// 		"auth{{ $i }}.sh": "{{ $e.GetScript()}}",
// 		{{ else if .Init.Data != nil -}}
// 		"auth{{ $i }}.sh": "{{ $e.GetScript()}}",
// 		{{ else -}}
// 		"auth{{ $i }}.sh": "{{ $e.GetScript()}}"
// 		{{- end }}
// 		{{- end }}
// 		{{- end }}
// 		{{ if .Init.Data != nil -}}
// 		{{ for i, e := range .Init.Data -}}
// 		{{ if i < len(.Init.Data)-1 -}}
// 		"data{{ $i }}.sh": "{{ $e.GetScript() }}",
// 		{{ else -}}
// 		"data{{ $i }}.sh": "{{ $e.GetScript() }}"
// 		{{- end }}
// 		{{- end }}
// 		{{- end }}
// 	}
// }
// `

// var pipelineDataSyncScriptTemplate = `
// #!/bin/bash

// `

// var pipelineJobScriptTemplate = `
// #!/bin/bash

// cd {{.Workdir}}
// {{.Command}}
// exit $?
// `

// func NewKubernetesPipelineJob(serverConfig *config.ServerConfig, jobConfig *PipelineStep) (KubernetesPipelineJob, error) {
// 	k := KubernetesPipelineJob{
// 		Id: uuid.NewV4().String(),
// 		Name: jobConfig.Name,
// 		Color: "white",
// 	}

// 	// TODO: fully flesh out the job object before returning
// 	return k, nil
// }

// func (j *KubernetesPipelineJob) SetColor(color string) {
// 	j.Color = color
// }

// func (j *KubernetesPipelineJob) GetColor() string {
// 	return j.Color
// }
 
// func (j *KubernetesPipelineJob) GetId() string {
// 	return j.Id
// }

// func (j *KubernetesPipelineJob) GetName() string {
// 	return j.Name
// }

// func (j *KubernetesPipelineJob) GetChildren() []*PipelineJob {
// 	return j.Children
// }

// func (j *KubernetesPipelineJob) GetParents() []*PipelineJob {
// 	return j.Parents
// }

// func (j *KubernetesPipelineJob) AddChild(child *PipelineJob) {
// 	j.Children = append(j.Children, child)
// }

// func (j *KubernetesPipelineJob) AddParent(parent *PipelineJob) {
// 	j.Parents = append(j.Parents, parent)
// }

// func (j *KubernetesPipelineJob) Watch(statuses chan *PipelineJobStatus, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// watch for the job to complete
// 	// put status on the channel
// 	_ = j.Cleanup()
// }

// func (j *KubernetesPipelineJob) Init() error {_
// 	return nil
// }

// func (j *KubernetesPipelineJob) SyncInputs(ctx context.Context) error {
// 	return nil 
// }

// func (j *KubernetesPipelineJob) SyncOutputs(ctx context.Context) error {
// 	return nil
// }

// func (j *KubernetesPipelineJob) Run(ctx context.Context) error { 
// 	// parameterize job pod, config and storage templates
// 	// create job storage
// 	// create job configuration
// 	// start job pod
// 	return nil
// }

// func (j *KubernetesPipelineJob) Cleanup() error {
// 	// delete the pod 
// 	return nil
// }