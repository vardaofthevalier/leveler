package grpc

import (
	"os"
	"fmt"
	"flag"
	"reflect"
	"testing"
	"context"
	data "leveler/data"
	resources "leveler/resources"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
)

var getEndpointServer = func () *EndpointServer {
	var mock = flag.Bool("mock", false, "Specify whether or not to use mock objects (for unit testing) or real objects (for integration testing)")
	var db = flag.String("db", "", "The database type to use for integration tests -- currently only 'redis' is supported")
	var dbHost = flag.String("dbhost", "127.0.0.1", "The database host to connect to for integration tests -- 127.0.0.1 is the default")
	var dbPort = flag.Int("dbport", -1, "The database port to connect to for integration tests -- no default")
	flag.Parse()

	var e *EndpointServer

	if *mock {
		e = &EndpointServer{
			Database: &MockDatabase{},
		}
	} else {
		if len(*db) == 0 {
			fmt.Println("Database type not specified -- please rerun with the '-db' option!")
			os.Exit(1)
		} else {
			if *dbPort == -1 {
				fmt.Println("Database port not specified -- please rerun with the '-dbport' option!")
				os.Exit(1)
			}

			e = &EndpointServer{
				Database: data.NewRedisDatabase("tcp", *dbHost, int32(*dbPort), 10),
			}
		}
	}

	return e
}

var endpointServer = getEndpointServer()

type MockDatabase struct {}

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
	r := &resources.Resource{
		Type: "resource",
	}

	return []string{proto.MarshalTextString(r)}, nil
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

func TestCreateResource_WithType(t *testing.T) {
	r := &resources.Resource{
		Type: "resource",
	}

	result, err := endpointServer.CreateResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.Resource{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestCreateResource_WithoutType(t *testing.T) {
	r := &resources.Resource{}

	_, err := endpointServer.CreateResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without type!")
	} 
}

func TestGetResource_WithId_WithType(t *testing.T) {
	r := &resources.Resource{
		Type: "resource",
		Id: "MockId",
	}

	result, err := endpointServer.GetResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.Resource{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestGetResource_WithoutType(t *testing.T) {
	r := &resources.Resource{
		Id: "MockId",
	}

	_, err := endpointServer.CreateResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without type!")
	} 
}

func TestGetResource_WithoutId(t *testing.T) {
	var err error 

	// WITHOUT ID
	r := &resources.Resource{
		Type: "resource",
	}

	_, err = endpointServer.GetResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass! Shouldn't pass without ID")
	}
}

func TestListResource_QueryWithType(t *testing.T) {
	query := &resources.Query{
		Type: "resource",
		Query: "name == quirky_turkey",
	}

	result, err := endpointServer.ListResources(context.Background(), query)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.ResourceList{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestListResource_QueryWithoutType(t *testing.T) {
	query := &resources.Query{
		Query: "name == quirky_turkey",
	}

	_, err := endpointServer.ListResources(context.Background(), query)
	if err == nil {
		t.Errorf("Unexpected pass: query should fail without type!")
	}
}

func TestListResource_QueryWithoutQueryString(t *testing.T) {
	query := &resources.Query{
		Type: "resource",
	}

	result, err := endpointServer.ListResources(context.Background(), query)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&resources.ResourceList{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestUpdateResource_WithId_WithType(t *testing.T) {
	r := &resources.Resource{
		Type: "resource",
		Id: "MockId",
	}

	result, err := endpointServer.UpdateResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&empty.Empty{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestUpdateResource_WithoutType(t *testing.T) {
	r := &resources.Resource{
		Id: "MockId",
	}

	_, err := endpointServer.UpdateResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without type!")
	}
}

func TestUpdateResource_WithoutId(t *testing.T) {
	r := &resources.Resource{}

	_, err := endpointServer.UpdateResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without ID!")
	}
}

func TestDeleteResource_WithId_WithType(t *testing.T) {
	r := &resources.Resource{
		Type: "resource",
		Id: "MockId",
	}

	result, err := endpointServer.DeleteResource(context.Background(), r)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if reflect.TypeOf(result) != reflect.TypeOf(&empty.Empty{}) {
		t.Errorf("Error: incorrect type returned!")
	}
}

func TestDeleteResource_WithoutType(t *testing.T) {
	r := &resources.Resource{
		Id: "MockId",
	}

	_, err := endpointServer.DeleteResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without type!")
	}
}

func TestDeleteResource_WithoutId(t *testing.T) {
	r := &resources.Resource{}

	_, err := endpointServer.DeleteResource(context.Background(), r)
	if err == nil {
		t.Errorf("Unexpected pass -- Shouldn't pass without ID!")
	}
}