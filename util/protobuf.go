package util

import (
	"fmt"
	"bytes"
	json "encoding/json"
	yaml "gopkg.in/yaml.v2"
	proto "github.com/golang/protobuf/proto"
	jsonpb "github.com/golang/protobuf/jsonpb"
	ptypes "github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
)

var jsonMarshaler = jsonpb.Marshaler{
	EnumsAsInts: false,
	EmitDefaults: true,
	Indent: "  ",
	OrigName: true,
}

var jsonUnmarshaler = jsonpb.Unmarshaler{
	AllowUnknownFields: false,
}

func ConvertProtoToMap(m proto.Message) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	jsonString, err := jsonMarshaler.MarshalToString(m)
	if err != nil {
		return jsonMap, err
	}

	err = json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		return jsonMap, err
	} 

	return jsonMap, nil
}

func ConvertJsonToProto(b []byte, m interface{}) error {
	r := bytes.NewReader(b)
	err := jsonUnmarshaler.Unmarshal(r, m.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func ConvertFromJsonString(b []byte, i interface{}) error {
	err := json.Unmarshal(b, i)
	if err != nil {
		return err
	}

	return nil
}

func ConvertToJsonString(i interface{}) ([]byte, error) {
	jsonString, err := json.Marshal(i)
	if err != nil {
		return jsonString, err
	}

	return jsonString, nil
}

func ConvertJsonStringToMap(jsonString []byte) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal(jsonString, &jsonMap)
	if err != nil {
		return jsonMap, err
	} 

	return jsonMap, nil
}

func ConvertFromYaml(yml []byte, i interface{}) error {
	err := yaml.Unmarshal(yml, i)
	if err != nil {
		return err
	}

	return nil
}

func GenerateProtoAny(m map[string]interface{}) (*any.Any, error) {
	var message proto.Message
	var a *any.Any

	b, err := ConvertToJsonString(m)
	if err != nil {
		return a, err
	}

	fmt.Printf("%v", b)
	err = ConvertJsonToProto(b, message)
	if err != nil {
		return a, err
	}

	a, err = ptypes.MarshalAny(message)
	if err != nil {
		return a, err
	}

	return a, nil
}

func FormatProto(p *proto.Message) {
	// TODO: transform messages into human readable output and print them
}
