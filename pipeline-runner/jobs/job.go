package pipeline 

import (
	"fmt"
	"sync"
	"os/exec"
)

type KubernetesPipelineJobInit struct {
	Credentials []*ExternalDataProvider
	Data []*ExternalDataProvider
}

type Kubernetes

type KubernetesPipelineJob struct {
	Id string
	Name string
	Namespace string
	Container string
	RepoPath string
	Script string
	Env *map[string]string
	Init *KubernetesPipelineJobInit
	VolumeSizeGi int
	Parents []*PipelineJob
	Children []*PipelineJob
}

var pipelineJobPvcTemplate = `
{
	"apiVersion": "v1",
	"kind": "PersistentVolumeClaim",
	"metadata": {
		"name": "{{.Name}}-{{.Id}}",
		"namespace": "{{.Namespace}}"
	},
	"spec": {
		"resources": {
			"requests": {
				"storage": "{{.VolumeSizeGi}}Gi"
			}
		},
		"accessModes": [
			"ReadWriteOnce"
		]
	}
}
`

var pipelineJobPodTemplate = `
{
	"apiVersion": "v1",
	"kind": "Pod",
	"metadata": {
		"name": "{{.Name}}-{{.Id}}",
		"namespace": "{{.Namespace}}"
	},
	"spec": {
		{{ if .Init != nil -}}
		"initContainers": [
		{{ if .Init.Credentials != nil -}}
		{{ for $i, $e := range .Init.Credentials -}}
			{{ if $i < len(.Init.Credentials)-1 -}}
			{
				"image": "{{ $e.GetContainer() }}",
				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
				"command": "/configmap/auth{{ $i }}.sh"
			},
			{{ else if .Init.Data != nil -}}
			{
				"image": "{{ $e.GetContainer() }}",
				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
				"command": "/configmap/auth{{ $i }}.sh"
			},
			{{ else -}}
			{
				"image": "{{ $e.GetContainer() }}",
				"name": "auth{{ $i }}-{{.Name}}-{{.Id}}",
				"command": "/configmap/auth{{ $i }}.sh"
			}
			{{- end }}
		{{- end }}
		{{- end }}
		{{ if .Init.Data != nil -}}
		{{ for $i, $e := range .Init.Data -}}
			{{ if $i < len(.Init.Data)-1 -}}
			{
				"image": "{{ $e.GetContainer() }}",
				"name": "data{{ $i }}-{{.Name}}-{{.Id}}",
				"command": "/configmap/data{{ $i }}.sh"
			},
			{{ else -}}
			{
				"image": "{{ $e.GetContainer() }}",
				"name": "data{{ $i }}-{{.Name}}-{{.Id}}",
				"command": "/configmap/data{{ $i }}.sh"
			}
			{{- end }}
		{{- end }}
		{{- end }}
		],
		{{- end }}
		"containers": [
			"image": "{{.Container}}",
			"name": "{{.Name}}-{{.Id}}",
			"command": "/configmap/job.sh",
			"volumeMounts": [
				{
					"name": "data",
					"mountPath": "/data"
				},
				{
					"name": "configmap",
					"mountPath": "/configmap"
				}
			],
			"workingDir": "/data/{{.RepoPath}}",
			"lifecyle"
		],
		"volumes": [
			{
				"name": "data",
				"persistentVolumeClaim": {
					"claimName": "{{.Name}}-{{.Id}}"
				},
				"name": "configmap",
				"configMap": {
					"name": {{.Name}}-{{.Id}}
				}
			}
		]
	}
}
`

var pipelineJobScriptConfigMapTemplate = `
{
	"apiVersion": "v1",
	"kind": "ConfigMap",
	"metadata": {
		"name": "{{.Name}}-{{.Id}}",
		"namespace": "{{.Namespace}}"
	},
	"data": {
		{{ if .Init == nil -}}
		"job.sh": "{{.Script}}"
		{{ else -}}
		"job.sh": "{{.Script}}",
		{{ end -}}
		{{ if .Init.Credentials != nil -}}
		{{ for i, e := range .Init.Credentials -}}
		{{ if i < len(.Init.Credentials)-1 -}}
		"auth{{ $i }}.sh": "{{ $e.GetScript()}}",
		{{ else if .Init.Data != nil -}}
		"auth{{ $i }}.sh": "{{ $e.GetScript()}}",
		{{ else -}}
		"auth{{ $i }}.sh": "{{ $e.GetScript()}}"
		{{- end }}
		{{- end }}
		{{- end }}
		{{ if .Init.Data != nil -}}
		{{ for i, e := range .Init.Data -}}
		{{ if i < len(.Init.Data)-1 -}}
		"data{{ $i }}.sh": "{{ $e.GetScript() }}",
		{{ else -}}
		"data{{ $i }}.sh": "{{ $e.GetScript() }}"
		{{- end }}
		{{- end }}
		{{- end }}
	}
}
`

var pipelineDataSyncScriptTemplate = `
#!/bin/bash

`

var pipelineJobScriptTemplate = `
#!/bin/bash

cd {{.RepoPath}}
{{.Command}}
exit $?
`

func (j *PipelineJob) generateScripts() error {

}

func (j *KubernetesPipelineJob) Watch(statuses *chan map[string]interface{}, wg *sync.Waitgroup) {
	defer wg.Close()
	// watch for the job to complete
	// put status on the channel
	_ := j.Cleanup()
}

func (j *KubernetesPipelineJob) Run() error { 
	// parameterize job pod, config and storage templates
	// create job storage
	// create job configuration
	// start job pod
}

func (j *KubernetesPipelineJob) Cleanup() error {
	// delete the pod 
}