package leveler

import (
	json "encoding/json"
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
