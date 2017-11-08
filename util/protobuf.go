package leveler

import (
	json "encoding/json"
	yaml "gopkg.in/yaml.v2"
	proto "github.com/golang/protobuf/proto"
	jsonpb "github.com/golang/protobuf/jsonpb"
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

func ConvertProtoToJsonMap(m proto.Message) (map[string]interface{}, error) {
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

func ConvertMapToJson(m map[string]string) ([]byte, error) {
	jsonString, err := json.Marshal(m)
	if err != nil {
		return jsonString, err
	}

	return jsonString, nil
}

func ConvertYamlToProto(yml []byte, p *proto.Message) error {
	err := yaml.Unmarshal(yml, p)
	if err != nil {
		return err
	}

	return nil
}

func ConvertYamlToMap(yml []byte, m map[string]interface) error {
	err := yaml.Unmarshal(yml, m)
	if err != nil {
		return err
	}

	return nil
}

func FormatProto(p *proto.Message) {
	// TODO: transform messages into human readable output and print them
}
