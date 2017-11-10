package leveler

import (
	grpc "google.golang.org/grpc"
)

func RegisterEndpoints(grpcServer *grpc.Server, s *EndpointServer) {
	RegisterResourceEndpointServer(grpcServer, s)
}