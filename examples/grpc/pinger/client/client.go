package main

import (
	"log"

	"github.com/giancarlopetrini/golang-learning/examples/grpc/pinger/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	// create connection of type *grpc.ClientConn
	var conn *grpc.ClientConn

	// dial grpc server, using WithInsecure for simple use case (DON'T do this in production!)
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not dial server: %s", err)
	}
	// close the connection to server once client has completed requests
	defer conn.Close()

	// create grpc client using proto file for registered ping service
	// using grpc connection created above
	c := api.NewPingClient(conn)

	// send GreetMessage, using CheckIn Method in api.proto
	res, err := c.CheckIn(context.Background(), &api.GreetMessage{Greeting: "Hi"})
	if err != nil {
		log.Fatalf("error when calling CheckIn: %s", err)
	}
	// log response from server (configured in api/handler.go)
	log.Printf("Response from CheckIn call to server: %s", res.Greeting)

	// call GetDate method
	res2, err := c.GetDate(context.Background(), &api.RequestDate{Req: "Requesting Date...."})
	if err != nil {
		log.Fatalf("failed to get date from server: %s", err)
	}
	log.Printf("Response from GetDate request: %v-%v-%v", res2.Day, res2.Month, res2.Year)
}
