package pipelines

import (
	//"fmt"
	"bytes"
	"text/template"
)

type ExternalDataProvider interface {
	GetScript()
	GetContainer()
}

type CredentialsProvider struct {
	ContainerImage string
	VaultToken string
	VaultPath string
	OutputFileName string
}

type DataProvider struct {
	ContainerImage string
	DownloadUrl string
	DownloadCmd string
}

func (d *CredentialsProvider) GetContainer() string {
	return d.ContainerImage
}

func (d *CredentialsProvider) GetScript() (string, error) {
	script := `
#!/bin/bash

export VAULT_TOKEN={{.VaultToken}}
mkdir -p /data/scripts
value=$(vault read -field=value {{.VaultPath}})
echo $value > /data/secret/{{.OutputFilename}}
`
	buf := new(bytes.Buffer)
	t := template.Must(template.New("script").Parse(script))
	err := t.Execute(buf, d)
	if err != nil {
		return "", err 
	}

	return buf.String(), nil
}

func (d *DataProvider) GetContainer() string {
	return d.ContainerImage
}

func (d *DataProvider) GetScript() (string, error) {
	script :=  `
#!/bin/bash

cd /data
{{.DownloadCmd}} {{.DownloadUrl}}
`	
	buf := new(bytes.Buffer)
	t := template.Must(template.New("script").Parse(script))
	err := t.Execute(buf, d)
	if err != nil {
		return "", err 
	}

	return buf.String(), nil
}

