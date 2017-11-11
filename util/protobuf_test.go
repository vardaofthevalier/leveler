package util

import (
	"testing"
	"reflect"
	"leveler/config"
	proto "github.com/golang/protobuf/proto"
)

func TestConvertProtoToMap(t *testing.T) {
	expected := map[string]interface{} {
		"Host": "127.0.0.1",
		"Port": 80,
		"Database": {
			"Host": "127.0.0.1",
			"Port": 81,
			"Type": "redis"
		}
	}

	db := &config.Database{
		Host: "127.0.0.1",
		Port: 81,
		Type: "redis",
	}
	message := &config.Config{
		Host: "127.0.0.1",
		Port: 80,
		Database: db,
	}

	actual, err := ConvertProtoToMap(message)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	eq := reflect.DeepEqual(actual, expected)

	if !eq {
		t.Errorf("ERROR: got %v, want %v", actual, expected)
	}

}

func TestConvertJsonToProto(t *testing.T) {
	db := &config.Database{
		Host: "127.0.0.1",
		Port: 81,
		Type: "redis",
	}
	expected := &config.Config{
		Host: "127.0.0.1",
		Port: 80,
		Database: db,
	}

	var actual config.Config
	jsonString := []byte(`{"Host":"127.0.0.1","Port":80,"Database":{"Host": "127.0.0.1","Port":81,"Type":"Redis"}}`)

	err := ConvertJsonToProto(jsonString, &actual)

	actual, err := ConvertProtoToMap(message)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	eq := reflect.DeepEqual(actual, expected)

	if !eq {
		t.Errorf("ERROR: got %v, want %v", actual, expected)
	}
}

func TestConvertToJson(t *testing.T) {
	// vars: interface{}, []byte
}

func TestConvertFromJson(t *testing.T) {
	// vars: []byte, interface{}
}

func TestConvertFromYaml(t *testing.T) {
	// vars: []byte, interface{}
}

func TestGenerateProtoAny(t *testing.T) {
	//
}

func TestFormatProto(t *testing.T) {
	// vars: proto.Message, string
}