package main

import (
	"fmt"
	"log"
	"net"
	data "leveler/data"
	server "leveler/server"
	grpc "google.golang.org/grpc"
)

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
	s := &server.EndpointServer{db}
	s.RegisterEndpoints(grpcServer)

	// start the server
	grpcServer.Serve(lis)
}