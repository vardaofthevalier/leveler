package grpc

import (
	"log"
	"bytes"
	"context"
	data "leveler/data"
	util "leveler/util"
	//callbacks "leveler/callbacks"  // TODO: create callbacks to run in the resource CRUD functions below (type dependent)
	jsonpb "github.com/golang/protobuf/jsonpb"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type EndpointServer struct {
	Database data.Database
}

// ACTION ENDPOINTS

func (s *EndpointServer) CreateResource(ctx context.Context, obj Resource) (ResourceId, error) {
	log.Printf("Creating %s: %v", obj.Type, obj)

	var result ResourceId

	m, err := util.ConvertProtoToMap(obj.Details)
	if err != nil {
		return result, err
	}

	result.Id, err = s.Database.Create(obj.Type, m)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) GetResource(ctx context.Context, obj ResourceId) (Resource, error) {
	log.Printf("Retrieving %s: %s", obj.Type, obj.Id)

	var jsonString []byte
	var result Resource

	r, err := s.Database.Get(obj.Type, obj.Id)
	if err != nil {
		return result, err
	}

	jsonString, err = util.ConvertToJson(r)
	if err != nil {
		log.Printf("Error converting map to JSON: %v", err)
		return result, err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), result.Details)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) ListResources(ctx context.Context, query Query) (ResourceList, error) {
	log.Printf("Retrieiving %s list", query.Type)

	var jsonString []byte
	var result ResourceList

	list, err := s.Database.List(query.Type, query.Query)
	if err != nil {
		return result, err
	}

	//var r Resource  // TODO: start here
	// for k, v := range list {
	// 	details, err = util.ConvertMapToJson(v)
	// 	if err != nil {
	// 		return result, err
	// 	}
	// 	r = &Resource{
	// 		Id: 
	// 		Type:
	// 		Details: 
	// 	}
	// 	result.Results = append(result.Results, &Resource)
	// }

	jsonString, err = util.ConvertToJson(list)
	if err != nil {
		return result, err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) UpdateResource(ctx context.Context, obj Resource) (*empty.Empty, error) {
	log.Printf("Updating %s: %s", obj.Type, obj.Id)

	var result *empty.Empty

	m, err := util.ConvertProtoToMap(obj)
	if err != nil {
		return result, err
	}

	err = s.Database.Update(obj.Type, obj.Id, m)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) DeleteResource(ctx context.Context, obj ResourceId) (*empty.Empty, error) {
	log.Printf("Deleting %s: %s", obj.Type, obj.Id)

	var result *empty.Empty

	err := s.Database.Delete(obj.Type, obj.Id)
	if err != nil {
		return result, err
	}

	return result, nil
}