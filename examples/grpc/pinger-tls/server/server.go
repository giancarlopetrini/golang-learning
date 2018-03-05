package main

import (
	"log"
	"net"

	"github.com/giancarlopetrini/golang-learning/examples/grpc/pinger-tls/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/* TODO, add to make file?:

$ openssl genrsa -out cert/server.key 2048
$ openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
$ openssl req -new -sha256 -key cert/server.key -out cert/server.csr
$ openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
*/

func main() {
	// create a listener on tcp port 7777
	lis, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// creates a server instance
	s := api.Server{}

	// creates TLS creds on the server
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalf("Unable to load server certs: %s", err)
	}

	// create gRPC options for variadic function call
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a gRPC server object using variadic function
	grpcServer := grpc.NewServer(opts...)

	// attach pinger service (from api/api.proto) to the server
	api.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
