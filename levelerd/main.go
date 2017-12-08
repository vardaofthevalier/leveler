package main

import (
	"fmt"
	"log"
	"net"
	//"leveler/data"
	"leveler/config"
	"leveler/server"
	"google.golang.org/grpc"
)

func main() {
	// TODO: parse command line options and server configuration
	var opts []grpc.ServerOption

	// read the configuration (default file if not overridden on the command line)
	var c = &server.ServerConfig{}
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

	//var db data.RedisDatabase 

	if c.Database.Type == "redis" {
		// create a Redis db object
		// db = data.NewRedisDatabase(c.Database.Protocol, c.Database.Host, c.Database.Port, c.Database.GetOptions().GetPoolSize()) 
	} else if c.Database.Type == "sql" {
		// create a MySql db object
		// db = data.NewSqlDatabase(protocol, )
		// defer db.Close()
	} else {
		log.Fatalf("Unknown database type '%s'", c.Database.Type)
	}
	
	// register endpoints
	grpcServer := grpc.NewServer(opts...)
	//s := &server.EndpointServer{db}
	s := &server.EndpointServer{}
	server.RegisterResourcesServer(grpcServer, s)

	// start the server
	grpcServer.Serve(lis)

}