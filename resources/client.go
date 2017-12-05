package resources 

import (
	"fmt"
	"bytes"
	"errors"
	"context"
	"reflect"
	"github.com/golang/protobuf/jsonpb"
)

var jsonMarshaler = &jsonpb.Marshaler{
	Indent: "  ",
}

func Create(pb reflect.Type, pbType string, client *ResourcesClient) (string, error) {
	return "", errors.New("Create method not implemented for any types!")
}

func Update(pb reflect.Type, pbType string, client *ResourcesClient) error {
	return errors.New("Update method not implemented for any types!")
}

func Patch(pb reflect.Type, pbType string, client *ResourcesClient) error {
	return errors.New("Patch method not implemented for any types!")
}

func Apply(pb reflect.Type, pbType string, client *ResourcesClient) error {
	return errors.New("Apply method not implemented for any types!")
}

func Delete(pb reflect.Type, pbType string, client *ResourcesClient) error {
	return errors.New("Delete method not implemented for any types!")
}

func Add(pb reflect.Type, pbType string, client *ResourcesClient) error {
	var err error
	switch pbType {
	case "Integration":
		concretePb := reflect.ValueOf(pb).Interface().(Integration)
		_, err = (*client).AddIntegration(context.Background(), &concretePb)
	case "Repository":
		concretePb := reflect.ValueOf(pb).Interface().(Repository)
		_, err = (*client).AddRepository(context.Background(), &concretePb)
	case "User":
		concretePb := reflect.ValueOf(pb).Interface().(User)
		_, err = (*client).AddUser(context.Background(), &concretePb)
	default:
		// TODO: error
	}
	
	if err != nil {
		return err 
	}
	return nil
}	

func Get(pb reflect.Type, pbType string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	switch pbType {
	case "Integration":
		concretePb := reflect.ValueOf(pb).Interface().(Integration)
		g, err := (*client).GetIntegration(context.Background(), &concretePb)

		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return fmt.Sprintf("%s", jsonString), err 
		}
	case "Repository":
		concretePb := reflect.ValueOf(pb).Interface().(Repository)
		g, err := (*client).GetRepository(context.Background(), &concretePb)

		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return fmt.Sprintf("%s", jsonString), err 
		}
	case "Pipeline": 
		concretePb := reflect.ValueOf(pb).Interface().(Pipeline)
		g, err := (*client).GetPipeline(context.Background(), &concretePb)
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return fmt.Sprintf("%s", jsonString), err 
		}
	case "Job":
		concretePb := reflect.ValueOf(pb).Interface().(Job)
		g, err := (*client).GetJob(context.Background(), &concretePb)
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return fmt.Sprintf("%s", jsonString), err 
		}
	case "Loggable": 
		concretePb := reflect.ValueOf(pb).Interface().(Loggable)
		logs, err := (*client).GetLogs(context.Background(), &concretePb)
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

	default:
		// TODO: error
	}

	return jsonString.String(), nil
}

func List(query string, pbType string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	switch pbType {
	case "Integration":
		g, err := (*client).ListIntegrations(context.Background(), &Query{Query: query,})
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return jsonString.String(), err 
		}
	case "Repository":
		g, err := (*client).ListRepositories(context.Background(), &Query{Query: query,})
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return jsonString.String(), err 
		}
	case "Pipeline": 
		g, err := (*client).ListPipelines(context.Background(), &Query{Query: query,})
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return jsonString.String(), err 
		}
	case "Job":
		g, err := (*client).ListJobs(context.Background(), &Query{Query: query,})
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return jsonString.String(), err 
		}
	default:
		// TODO: error
	}

	return jsonString.String(), nil
}

func Remove(pb reflect.Type, pbType string, client *ResourcesClient) error {
	var err error
	switch pbType {
	case "Integration":
		concretePb := reflect.ValueOf(pb).Interface().(Integration)
		_, err = (*client).RemoveIntegration(context.Background(), &concretePb)
	case "Repository":
		concretePb := reflect.ValueOf(pb).Interface().(Repository)
		_, err = (*client).RemoveRepository(context.Background(), &concretePb)
	default:
		// TODO: error
	}
	
	if err != nil {
		return err 
	}
	return nil
}

func Run(pb reflect.Type, pbType string, client *ResourcesClient) (string, error) {
	var jsonString *bytes.Buffer

	switch pbType {
	case "Pipeline": 
		concretePb := reflect.ValueOf(pb).Interface().(Pipeline)
		g, err := (*client).RunPipeline(context.Background(), &concretePb)
		if err != nil {
			return jsonString.String(), err
		}

		err = jsonMarshaler.Marshal(jsonString, g)
		if err != nil {
			return jsonString.String(), err 
		}
	default:
		// TODO: error
	}

	return jsonString.String(), nil
}

func Cancel(pb reflect.Type, pbType string, client *ResourcesClient) error {
	var err error
	switch pbType {
	case "Pipeline":
		concretePb := reflect.ValueOf(pb).Interface().(Pipeline)
		_, err = (*client).CancelPipeline(context.Background(), &concretePb)
	default:
		// TODO: error
	}
	
	if err != nil {
		return err 
	}
	return nil
}


