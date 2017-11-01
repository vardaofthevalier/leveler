package main

import (
	"context"
	endpoints "leveler/endpoints"
	redis_pool "github.com/mediocregopher/radix.v2/pool"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

// ACTION ENDPOINTS

type endpointServer struct {
	DatabaseConnectionPool redis_pool.Pool
}

func (s *endpointServer) CreateAction(ctx context.Context, action *endpoints.Action) (*endpoints.Action, error) {
	p := s.DatabaseConnectionPool.Get()

}

func (s *endpointServer) GetAction(ctx context.Context, actionId *endpoints.ActionId) (*endpoints.Action, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) ListActions(ctx context.Context, query *endpoints.Query) (*endpoints.Action, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) UpdateAction(ctx context.Context, action *endpoints.Action) (*endpoints.Action, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) DeleteAction(ctx context.Context, actionId *endpoints.ActionId) (*google_protobuf.Empty, error) {
	p := s.DatabaseConnectionPool.Get()
}

// REQUIREMENT ENDPOINTS

func (s *endpointServer) CreateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.Requirement, error) {
	p := s.DatabaseConnectionPool.Get()

}

func (s *endpointServer) GetRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*endpoints.Requirement, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) ListRequirements(ctx context.Context, query *endpoints.Query) (*endpoints.Requirement, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) UpdateRequirement(ctx context.Context, requirement *endpoints.Requirement) (*endpoints.Requirement, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) DeleteRequirement(ctx context.Context, requirementId *endpoints.RequirementId) (*google_protobuf.Empty, error) {
	p := s.DatabaseConnectionPool.Get()
}

// ROLE ENDPOINTS 

func (s *endpointServer) CreateRole(ctx context.Context, role *endpoints.Action) (*endpoints.Action, error) {
	p := s.DatabaseConnectionPool.Get()

}

func (s *endpointServer) GetRole(ctx context.Context, roleId *endpoints.RoleId) (*endpoints.Role, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) ListRoles(ctx context.Context, query *endpoints.Query) (*endpoints.Role, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) UpdateRole(ctx context.Context, role *endpoints.Role) (*endpoints.Role, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) DeleteRole(ctx context.Context, roleId *endpoints.RoleId) (*google_protobuf.Empty, error) {
	p := s.DatabaseConnectionPool.Get()
}

// HOST ENDPOINTS

func (s *endpointServer) CreateHost(ctx context.Context, host *endpoints.Host) (*endpoints.Host, error) {
	p := s.DatabaseConnectionPool.Get()

}

func (s *endpointServer) GetHost(ctx context.Context, hostId *endpoints.HostId) (*endpoints.Host, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) ListHosts(ctx context.Context, query *endpoints.Query) (*endpoints.Host, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) UpdateHost(ctx context.Context, host *endpoints.Host) (*endpoints.Host, error) {
	p := s.DatabaseConnectionPool.Get()
}

func (s *endpointServer) DeleteHost(ctx context.Context, hostId *endpoints.HostId) (*google_protobuf.Empty, error) {
	p := s.DatabaseConnectionPool.Get()
}


func main() {
	// TODO: parse command line options and server configuration
	var opts []grpc.ServerOption
	
	// listen on the specified port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))  // TODO: get port number from configuration
	if err != nil {
	        log.Fatalf("failed to listen: %v", err)
	}

	// create a Redis connection pool
	pool := redis_pool.New("tcp", "localhost:6379", 10)  // TODO: move connection info to configuration

	// register endpoints
	grpcServer := grpc.NewServer(opts...)
	endpoints.RegisterActionEndpointServer(grpcServer, &endpointServer{pool})
	endpoints.RegisterRequirementEndpointServer(grpcServer, &endpointServer{pool})
	endpoints.RegisterRoleEndpointServer(grpcServer, &endpointServer{pool})
	endpoints.RegisterHostEndpointServer(grpcServer, &endpointServer{pool})

	// start the server
	grpcServer.Serve(lis)
}