package pipelines

import (
	"fmt"
)

type ExternalData interface {
	Sync() error
}

func (data *S3Data) Sync() error {

}

func (data *NexusData) Sync() error {

}

func (data *SCMData) Sync() error {

}

func (data *LocalData) Sync() error {

}

func (data *StreamData) Sync() error {
	
}