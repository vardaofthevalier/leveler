package main

import (
	"fmt"
	"log"
	"net"
	data "leveler/data"
	server "leveler/grpc"
	config "leveler/config"
	grpc "google.golang.org/grpc"
)

func main() {
	// TODO: parse command line options and server configuration
	var opts []grpc.ServerOption

	// read the configuration (default file if not overridden on the command line)
	var c = &config.Config{}
	var err error

	err = config.Read("", "server", c)
	if err != nil {
		log.Fatalf("Couldn't read config file: %v", err)
	}
	
	// listen on the specified interface and port
	protocol := "tcp"

	lis, err := net.Listen(protocol, fmt.Sprintf("%s:%d", c.Host, c.Port))  
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var db data.RedisDatabase 

	if c.Database.Type == "redis" {
		// create a Redis db object
		db = data.NewRedisDatabase(protocol, c.Database.Host, c.Database.Port, c.Database.GetOptions().GetPoolSize()) 
	} else {
		log.Fatalf("Unknown database type '%s'", c.Database.Type)
	}
	
	// register endpoints
	grpcServer := grpc.NewServer(opts...)
	s := &server.EndpointServer{db}
	server.RegisterResourceEndpointServer(grpcServer, s)

	// start the server
	grpcServer.Serve(lis)
}