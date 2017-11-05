package leveler

import (
	grpc "google.golang.org/grpc"
)

func RegisterEndpoints(grpcServer *grpc.Server, s *EndpointServer) {
	RegisterActionEndpointServer(grpcServer, s)
	RegisterRequirementEndpointServer(grpcServer, s)
	RegisterRoleEndpointServer(grpcServer, s)
	RegisterHostEndpointServer(grpcServer, s)
}