package grpc

import (
	"fmt"
	"testing"
	"context"
	"reflect"
	resources "leveler/resources"
	uuid "github.com/satori/go.uuid"
)

type MockDatabase struct {}

func (m *MockDatabase) Create(kind string, keys map[string]interface{}, data string) (string, error) {
	return uuid.NewV4().String(), nil
}

func (m *MockDatabase) Get(kind string, id string) (string, error) {
	return fmt.Sprintf("Kind: %s, Id: %s", kind, id), nil
}

func (m *MockDatabase) List(kind string, query string) ([]string, error){
	return []string{fmt.Sprintf("Kind: %s, Query: %s", kind, query)}, nil
}

func (m *MockDatabase) Update(kind string, id string, data string) error {
	return nil
}

func (m *MockDatabase) Delete(kind string, id string) error {
	return nil
}

func (m *MockDatabase) Flush(db string) error {
	return nil
}

func TestCreateResource(t *testing.T) {
	var result *resources.Resource
	var err error 

	endpointServer := &EndpointServer{
		Database: &MockDatabase{},
	}

	r := &resources.Resource{
		Type: "resource",
	}

	result, err = endpointServer.CreateResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error running test: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.Resource{}) {
		t.Errorf("Incorrect type returned!")
	}
}

func TestGetResource(t *testing.T) {

}

func TestListResource(t *testing.T) {

}

func TestUpdateResource(t *testing.T) {

}

func TestDeleteResource(t *testing.T) {

}