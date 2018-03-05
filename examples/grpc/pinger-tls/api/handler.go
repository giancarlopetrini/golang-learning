package api

import (
	"log"
	"time"

	"golang.org/x/net/context"
)

// Server struct represent gRPC server created in server/server.go
type Server struct{}

// CheckIn - generates resposne to CheckIn rpc call
func (s *Server) CheckIn(ctx context.Context, in *GreetMessage) (*GreetMessage, error) {
	log.Printf("Received CheckIn message: %s", in.Greeting)
	return &GreetMessage{Greeting: "Hello, there."}, nil
}

// GetDate - response to GetDate rpc call
func (s *Server) GetDate(ctx context.Context, in *RequestDate) (*DateMessage, error) {
	log.Printf("Received RequestDate message from client: %s", in.Req)
	return &DateMessage{Day: int32(time.Now().Day()), Month: int32(time.Now().Month()), Year: int32(time.Now().Year())}, nil
}
