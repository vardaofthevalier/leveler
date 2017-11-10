package test

import (
	"fmt"
	util "leveler/util"
	data "leveler/data"
	grpc "leveler/grpc"
	ptypes "github.com/golang/protobuf/ptypes"
	//any "github.com/golang/protobuf/ptypes/any"
)

var db = data.NewRedisDatabase("tcp", "127.0.0.1", 6379, 10) 

func TestCreate(t string, m map[string]interface{}) (string, error) {
	id, err := db.Create(t, m)
	if err != nil {
		return id, err
	}

	return id, nil
}

// func TestGet(kind string, id string) error {

// }

// func TestList() error {

// }

// func TestUpdate() error {

// }

// func TestDelete() error {

// }
// type ActionDetails struct {
// 	Name string
// 	Description string
// 	Command string
// 	Shell string
// }

func RunTests() error {
	// var testObjects []*Resource

	var id string
	//var err error
	//var r map[string]string

	d := map[string]interface{} {
		"Name": "foo",
		"Description": "a really sweet foo",
		"Command": "echo FOOOOOOO",
		"Shell": "/bin/bash",
	}

	// jsonByteArray, err := util.ConvertMapToJson(d)
	// if err != nil {
	// 	return err
	// }
	
	// message, err := util.ConvertJsonStringToProto(jsonByteArray)
	// if err != nil {
	// 	return err
	// }

	message, err := util.GenerateProto(d)

	a, err := ptypes.MarshalAny(message)
	if err != nil {
		return err
	}

	t := &grpc.Resource{
		Type: "action",
		Details: a,
	}

	m, err := util.ConvertProtoToMap(t.Details)
	if err != nil {
		return err
	}

	id, err = TestCreate(t.Type, m)
	if err != nil {
		return err
	}

	fmt.Printf("Created '%s'", id)

	// r, err = TestGet("action", id)
	// if err != nil {
	// 	return err
	// }

	// fmt.Printf("Got %v", r)

	return nil
}