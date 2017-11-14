package grpc

import (
	"fmt"
	"testing"
	"context"
	"reflect"
	resources "leveler/resources"
	proto "github.com/golang/protobuf/proto"
)

type MockDatabase struct {}

var endpointServer = &EndpointServer{
	Database: &MockDatabase{},
}

func (m *MockDatabase) Create(kind string, keys map[string]interface{}, data string) (string, error) {
	return "MockID", nil
}

func (m *MockDatabase) Get(kind string, id string) (string, error) {
	r := &resources.Resource{
		Type: "resource",
	}


	return proto.MarshalTextString(r), nil
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
	var result *resources.Resource 
	var err error 

	// WITH ID
	r := &resources.Resource{
		Type: "resource",
		Id: "MockId",
	}

	result, err = endpointServer.GetResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error running test: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.Resource{}) {
		t.Errorf("Incorrect type returned!")
	}

	// WITHOUT ID
	r = &resources.Resource{
		Type: "resource",
	}

	result, err = endpointServer.GetResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass! Shouldn't pass without ID")
	}
}

func TestListResource(t *testing.T) {

}

func TestUpdateResource(t *testing.T) {

}

func TestDeleteResource(t *testing.T) {

}