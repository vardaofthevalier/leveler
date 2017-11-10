package main

import (
	"fmt"
	util "leveler/util"
	data "leveler/data"
	grpc "leveler/grpc"
	ptypes "github.com/golang/protobuf/ptypes"
)

func main() {
	err := RunTests()

	if err != nil {
		fmt.Println(err)
	}
}

var db = data.NewRedisDatabase("tcp", "127.0.0.1", 6379, 10) 

func TestCreate(t string, m map[string]interface{}) (string, error) {
	id, err := db.Create(t, m)
	if err != nil {
		return id, err
	}

	return id, nil
}

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

	a, err := util.GenerateProtoAny(d)
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