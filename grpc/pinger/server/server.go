package main

import (
	"log"
	"net"

	"github.com/giancarlopetrini/golang-learning/examples/grpc/pinger/api"
	"google.golang.org/grpc"
)

func main() {
	// create a listener on tcp port 7777
	lis, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// creates a server instance
	s := api.Server{}
	// create grpc server
	grpcServer := grpc.NewServer()

	// attach pinger service (from api/api.proto) to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
