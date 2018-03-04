package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server struct represent gRPC server created in server/server.go
type Server struct{}

// CheckIn - generates resposne to CheckIn rpc call
func (s *Server) CheckIn(ctx context.Context, in *GreetMessage) (*GreetMessage, error) {
	log.Printf("Received message: %s", in.Greeting)
	return &GreetMessage{Greeting: "Hello, there."}, nil
}
