package grpc

// import (
// 	"log"
// 	"errors"
// 	"context"
// 	data "leveler/data"
// 	resources "leveler/resources"
// 	//callbacks "leveler/callbacks"  // TODO: create callbacks to run in the resource CRUD functions below (type dependent)
// 	proto "github.com/golang/protobuf/proto"
// 	ptypes "github.com/golang/protobuf/ptypes"
// 	empty "github.com/golang/protobuf/ptypes/empty"
// )

// type EndpointServer struct {
// 	Database data.Database
// }

// func (s *EndpointServer) CreateResource(ctx context.Context, obj *resources.Resource) (*resources.Resource, error) {
// 	log.Printf("Creating %s: %v", obj.Type, obj)

// 	var result = &resources.Resource{}

// 	var keys = make(map[string]interface{})
// 	var err error

// 	var stringDetail resources.StringDetail
// 	var stringDetailErr error 

// 	var boolDetail resources.BoolDetail 
// 	var boolDetailErr error 

// 	var int64Detail resources.Int64Detail 
// 	var int64DetailErr error 

// 	for _, d := range obj.Details {
// 		stringDetailErr = ptypes.UnmarshalAny(d, &stringDetail)
// 		boolDetailErr = ptypes.UnmarshalAny(d, &boolDetail)
// 		int64DetailErr = ptypes.UnmarshalAny(d, &int64Detail)

// 		if stringDetailErr == nil {
// 			if stringDetail.IsSecondaryKey {
// 				keys[stringDetail.Name] = stringDetail.Value
// 			}
// 		} else if boolDetailErr == nil {
// 			if boolDetail.IsSecondaryKey {
// 				keys[boolDetail.Name] = boolDetail.Value
// 			}
// 		} else if int64DetailErr == nil {
// 			if int64Detail.IsSecondaryKey {
// 				keys[int64Detail.Name] = int64Detail.Value
// 			}
// 		} else {
// 			return result, errors.New("Malformed detail")
// 		}
// 	}

// 	pb := proto.MarshalTextString(obj)

// 	result.Id, err = s.Database.Create(obj.Type, keys, pb)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func (s *EndpointServer) GetResource(ctx context.Context, obj *resources.Resource) (*resources.Resource, error) {
// 	log.Printf("Retrieving %s: %s", obj.Type, obj.Id)

// 	var result = &resources.Resource{}

// 	r, err := s.Database.Get(obj.Type, obj.Id)
// 	if err != nil {
// 		return result, err
// 	}

// 	err = proto.UnmarshalText(r, result)
// 	if err != nil {
// 		log.Printf("Error converting text data to Protobuf: %v", err)
// 		return result, err
// 	}

// 	return result, nil
// }

// func (s *EndpointServer) ListResources(ctx context.Context, query *resources.Query) (*resources.ResourceList, error) {
// 	log.Printf("Retrieiving %s list", query.Type)

// 	var result = &resources.ResourceList{}
// 	var list []string
// 	var err error 

// 	list, err = s.Database.List(query.Type, query.Query)
// 	if err != nil {
// 		return result, err
// 	}

// 	var r *resources.Resource
// 	for _, v := range list {
// 		r = new(resources.Resource)
// 		err := proto.UnmarshalText(v, r)
// 		if err != nil {
// 			log.Printf("Error converting text data to Protobuf: %v", err)
// 			return result, err
// 		}

// 		result.Results = append(result.Results, r)
// 	}

// 	return result, nil
// }

// func (s *EndpointServer) UpdateResource(ctx context.Context, obj *resources.Resource) (*empty.Empty, error) {
// 	log.Printf("Updating %s: %s", obj.Type, obj.Id)

// 	var result *empty.Empty
// 	var err error 

// 	pb := proto.MarshalTextString(obj)

// 	err = s.Database.Update(obj.Type, obj.Id, pb)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func (s *EndpointServer) DeleteResource(ctx context.Context, obj *resources.Resource) (*empty.Empty, error) {
// 	log.Printf("Deleting %s: %s", obj.Type, obj.Id)

// 	var result *empty.Empty

// 	err := s.Database.Delete(obj.Type, obj.Id)
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }