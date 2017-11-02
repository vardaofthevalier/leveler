package main

import (
	"fmt"
	"log"
	"net"
	"context"
	json "encoding/json"
	data "leveler/data"
	grpc "google.golang.org/grpc"
	endpoints "leveler/endpoints"
	proto "github.com/golang/protobuf/proto"
	jsonpb "github.com/golang/protobuf/jsonpb"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

type endpointServer struct {
	Database data.Database
}

func toJson(m proto.Message) (map[string]interface{}, error) {
	var jsonMap interface{}
	jsonString := jsonpb.MarshalToString(m)

	err := json.Unmarshal(jsonString, &jsonMap)
	if err != nil {
		// TODO
		return jsonMap, err
	} 

	return jsonMap
}

// ACTION ENDPOINTS

func (s endpointServer) CreateAction(ctx context.Context, action *endpoints.Action) (*endpoints.Action, error) {
	log.Print(action)
	// TODO: logging

	err := s.Database.Create("action", toJson(action))
	if err != nil {
		// TODO
	}

	return 201, nil
}

func (s *endpointServer) GetAction(ctx context.Context, actionId *endpoints.ActionId) (*endpoints.Action, error) {
	var action endpoints.Action
	return &action, nil
}

func (s *endpointServer) ListActions(ctx context.Context, query *endpoints.Query) (*endpoints.ActionList, error) {
	var actionList endpoints.ActionList
	return &actionList, nil
}

func (s *endpointServer) UpdateAction(ctx context.Context, action *endpoints.Action) (*endpoints.Action, error) {
	return action, nil
}

func (s *endpointServer) DeleteAction(ctx context.Context, actionId *endpoints.ActionId) (*google_protobuf.Empty, error) {
	return nil, nil
}

// REQUIREMENT ENDPOINTS

func (s *endpointServer) CreateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.Requirement, error) {
	return requirement, nil
}

func (s *endpointServer) GetRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*endpoints.Requirement, error) {
	var requirement endpoints.Requirement
	return &requirement, nil
}

func (s *endpointServer) ListRequirements(ctx context.Context, query *endpoints.Query) (*endpoints.RequirementList, error) {
	var requirementList endpoints.RequirementList
	return &requirementList, nil
}

func (s *endpointServer) UpdateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.Requirement, error) {
	return requirement, nil
}

func (s *endpointServer) DeleteRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*google_protobuf.Empty, error) {
	return nil, nil
}

// ROLE ENDPOINTS 

func (s *endpointServer) CreateRole(ctx context.Context, role *endpoints.Role) (*endpoints.Role, error) {
	return role, nil
}

func (s *endpointServer) GetRole(ctx context.Context, roleId *endpoints.RoleId) (*endpoints.Role, error) {
	var role endpoints.Role
	return &role, nil
}

func (s *endpointServer) ListRoles(ctx context.Context, query *endpoints.Query) (*endpoints.RoleList, error) {
	var roleList endpoints.RoleList
	return &roleList, nil
}

func (s *endpointServer) UpdateRole(ctx context.Context, role *endpoints.Role) (*endpoints.Role, error) {
	return role, nil
}

func (s *endpointServer) DeleteRole(ctx context.Context, roleId *endpoints.RoleId) (*google_protobuf.Empty, error) {
	return nil, nil
}

// HOST ENDPOINTS

func (s *endpointServer) CreateHost(ctx context.Context, host *endpoints.Host) (*endpoints.Host, error) {
	return host, nil
}

func (s *endpointServer) GetHost(ctx context.Context, hostId *endpoints.HostId) (*endpoints.Host, error) {
	var host endpoints.Host
	return &host, nil
}

func (s *endpointServer) ListHosts(ctx context.Context, query *endpoints.Query) (*endpoints.HostList, error) {
	var hostList endpoints.HostList
	return &hostList, nil
}

func (s *endpointServer) UpdateHost(ctx context.Context, host *endpoints.Host) (*endpoints.Host, error) {
	return host, nil
}

func (s *endpointServer) DeleteHost(ctx context.Context, hostId *endpoints.HostId) (*google_protobuf.Empty, error) {
	return nil, nil
}


func main() {
	// TODO: parse command line options and server configuration
	var opts []grpc.ServerOption
	
	// listen on the specified port
	protocol := "tcp"
	host := "127.0.0.1"
	port := 8080

	lis, err := net.Listen(protocol, fmt.Sprintf("%s:%d", host, port))  // TODO: get port number from configuration
	if err != nil {
	        log.Fatalf("failed to listen: %v", err)
	}

	// create a Redis db object
	port = 6379
	pool_size := 10
	db := data.NewRedisDatabase(protocol, host, port, pool_size)  // TODO: move connection info to configuration

	// register endpoints
	grpcServer := grpc.NewServer(opts...)
	s := &endpointServer{db}
	
	endpoints.RegisterActionEndpointServer(grpcServer, s)
	endpoints.RegisterRequirementEndpointServer(grpcServer, s)
	endpoints.RegisterRoleEndpointServer(grpcServer, s)
	endpoints.RegisterHostEndpointServer(grpcServer, s)

	// start the server
	grpcServer.Serve(lis)
}