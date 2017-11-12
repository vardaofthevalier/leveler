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

func (s *EndpointServer) CreateResource(ctx context.Context, obj *Resource) (*ResourceMetadata, error) {
	log.Printf("%v", obj)
	log.Printf("Creating %s: %v", obj.Type, obj)

	var result = &Resource{}

	m, err := util.ConvertProtoToMap(obj)
	if err != nil {
		return result, err
	}

	result.Id, err = s.Database.Create(obj.Type, m)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) GetResource(ctx context.Context, obj *Resource) (*Resource, error) {
	log.Printf("Retrieving %s: %s", obj.Type, obj.Id)

	var jsonString []byte
	var result = &Resource{}

	r, err := s.Database.Get(obj.Type, obj.Id)
	if err != nil {
		return result, err
	}

	jsonString, err = util.ConvertToJsonString(r)
	if err != nil {
		log.Printf("Error converting map to JSON: %v", err)
		return result, err
	}

	err = jsonpb.Unmarshal(bytes.NewReader(jsonString), result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *EndpointServer) ListResources(ctx context.Context, query *Query) (*ResourceList, error) {
	log.Printf("Retrieiving %s list", query.Type)

	var result = &ResourceList{}

	list, err := s.Database.List(query.Type, query.Query)
	if err != nil {
		return result, err
	}

	for _, v := range list {
		s, err := util.ConvertToJsonString(v)
		if err != nil {
			return result, err
		}

		var r *Resource
		err = util.ConvertJsonToProto(s, &r)
		result.Results = append(result.Results, r)
	}

	return result, nil
}

func (s *EndpointServer) UpdateResource(ctx context.Context, obj *Resource) (*empty.Empty, error) {
	log.Printf("Updating %s: %s", obj.Metadata.Type, obj.Metadata.Id)

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

func (s *EndpointServer) DeleteResource(ctx context.Context, obj *Resource) (*empty.Empty, error) {
	log.Printf("Deleting %s: %s", obj.Type, obj.Id)

	var result *empty.Empty

	err := s.Database.Delete(obj.Type, obj.Id)
	if err != nil {
		return result, err
	}

	return result, nil
}