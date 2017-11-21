package data

import (
	"fmt"
	"templates"
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

func (d *CredentialsProvider) GetContainer() {
	return d.ContainerImage
}

func (d *CredentialsProvider) GetScript() string, error {
	script := `
#!/bin/bash

export VAULT_TOKEN={{.VaultToken}}
mkdir -p /data/scripts
value=$(vault read -field=value {{.VaultPath}})
echo $value > /data/secret/{{.OutputFilename}}
`

	t := template.Must(template.New("script").Parse(script))
	err := t.Execute(d)
	if err != nil {
		return "", err 
	}

	return script, nil
}

func (d *DataProvider) GetContainer() {
	return d.ContainerImage
}

func (d *DataProvider) GetScript() {
	script :=  `
#!/bin/bash

cd /data
{{.DownloadCmd}} {{.DownloadUrl}}
`
	t := template.Must(template.New("script").Parse(script))
	err := t.Execute(d)
	if err != nil {
		return "", err 
	}

	return script, nil
}

