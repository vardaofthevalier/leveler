package leveler

import (
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

func ConvertJsonToProto(b []byte) (proto.Message, error) {
	var message proto.Message

	r := bytes.NewReader(b)
	err := jsonUnmarshaler.Unmarshal(r, message)
	if err != nil {
		return message, err
	}

	return message, nil
}

func ConvertToJson(i interface{}) ([]byte, error) {
	jsonString, err := json.Marshal(i)
	if err != nil {
		return jsonString, err
	}

	return jsonString, nil
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

	b, err := ConvertToJson(m)
	if err != nil {
		return a, err
	}

	message, err = ConvertJsonToProto(b)
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
