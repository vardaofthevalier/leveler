package leveler

import (
	"log"
	"bytes"
	"reflect"
	data "leveler/data"
	util "leveler/util"
	grpc "google.golang.org/grpc"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	
)

type EndpointServer struct {
	Database data.Database
}

func (s *EndpointServer) Register (grpcServer *grpc.Server) {
	RegisterEndpoints(grpcServer, s)
}

// ACTION ENDPOINTS

func (s *EndpointServer) genericCreate(t string, obj proto.Message, dest interface{}) error {
	log.Printf("Creating %s: %v", t, obj)

	var id string

	m, err := util.ConvertProtoToJsonMap(obj)
	if err != nil {
		return err
	}

	id, err = s.Database.Create(t, m)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(dest).Elem().FieldByName("Id")
	ptr := v.Addr().Interface().(*string)
	*ptr = id

	return nil
}

func (s *EndpointServer) genericGet(t string, id string, dest interface{}) error {
	log.Printf("Retrieving %s: %s", t, id)

	var jsonString []byte

	result, err := s.Database.Get(t, id)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		log.Printf("Error converting map to JSON: %v", err)
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericList(t string, query string, dest interface{}) error {
	log.Printf("Retrieiving %s list", t)

	var jsonString []byte

	result, err := s.Database.List(t, query)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericUpdate(t string, id string, obj proto.Message, dest interface{}) error {
	log.Printf("Updating %s: %v", obj)

	var jsonString []byte

	m, err := util.ConvertProtoToJsonMap(obj)
	if err != nil {
		return err
	}

	result, err := s.Database.Update(t, id, m)
	if err != nil {
		return err
	}

	jsonString, err = util.ConvertMapToJson(result)
	if err != nil {
		return err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), dest.(proto.Message))
	if err != nil {
		return err
	}

	return nil
}

func (s *EndpointServer) genericDelete(t string, id string) error {
	log.Printf("Deleting %s: %s", t, id)

	err := s.Database.Delete(t, id)
	if err != nil {
		return err
	}

	return nil
}