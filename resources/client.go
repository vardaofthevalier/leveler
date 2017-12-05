package resources 

import (
	"fmt"
	"bytes"
	"context"
	//"google.golang.org/grpc"
	"github.com/golang/protobuf/jsonpb"
)

var jsonMarshaler = &jsonpb.Marshaler{
	Indent: "  ",
}

// INTEGRATION 

func Add(pb *proto.Message, client *ResourcesClient) error {
	_, err := (*client).AddIntegration(context.Background(), i)
	if err != nil {
		return err 
	}
	return nil
}	

func (i *Integration) Get(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).GetIntegration(context.Background(), i)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return fmt.Sprintf("%s", jsonString), err 
	}

	return jsonString.String(), nil
}

func (i *Integration) List(query string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).ListIntegrations(context.Background(), &Query{Query: query,})
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (i *Integration) Remove(client *ResourcesClient) error {
	_, err := (*client).RemoveIntegration(context.Background(), i)
	if err != nil {
		return err
	}
	return nil
}

// REPOSITORY 

func (r *Repository) Add(client *ResourcesClient) error {
	_, err := (*client).AddRepository(context.Background(), r)
	if err != nil {
		return err 
	}
	return nil
}

func (r *Repository) Get(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).GetRepository(context.Background(), r)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return fmt.Sprintf("%s", jsonString), nil
}

func (r *Repository) List(query string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).ListRepositories(context.Background(), &Query{Query: query,})
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (r *Repository) Remove(client *ResourcesClient) error {
	_, err := (*client).RemoveRepository(context.Background(), r)
	if err != nil {
		return err 
	}
	return nil
}

// PIPELINE


func (p *Pipeline) Run(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).RunPipeline(context.Background(), p)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (p *Pipeline) Get(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).GetPipeline(context.Background(), p)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (p *Pipeline) List(query string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).ListPipelines(context.Background(), &Query{Query: query,})
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (p *Pipeline) Cancel(client *ResourcesClient) error {
	_, err := (*client).CancelPipeline(context.Background(), p)
	if err != nil {
		return err 
	}
	return nil
}

// JOB

func (j *Job) Get(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).GetJob(context.Background(), j)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

func (j *Job) List(query string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	g, err := (*client).ListJobs(context.Background(), &Query{Query: query,})
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, g)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}

// LOG

func (l *Loggable) Get(client *ResourcesClient) (string, error) {
	logs, err := (*client).GetLogs(context.Background(), l)
	if err != nil {
		return "", err
	}

	for {
		log, err := logs.Recv()
		if err != nil {
			return "", err
		}

		fmt.Printf("%s\n", log.Message)
	}

	return "", nil
}

// USER

func (u *User) Add(client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	user, err := (*client).AddUser(context.Background(), u)
	if err != nil {
		return jsonString.String(), err
	}

	err = jsonMarshaler.Marshal(jsonString, user)
	if err != nil {
		return jsonString.String(), err 
	}

	return jsonString.String(), nil
}