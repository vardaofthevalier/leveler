package mocks 

import (
	resources "leveler/resources"
	proto "github.com/golang/protobuf/proto"
)

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