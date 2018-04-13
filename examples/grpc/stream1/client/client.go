package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/giancarlopetrini/golang-learning/examples/grpc/stream1/protobuf"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't dial gRPC Server: %s", err)
	}
	defer conn.Close()

	c := pb.NewMessageClient(conn)

	stream, err := c.Chat(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive: %v \n", err)
			}
			fmt.Printf("message recieved: %s \n", in.Res)
		}
	}()

	messageCount := 10
	for i := 0; i < messageCount; i++ {
		if err := stream.Send(&pb.Request{Req: "SAMPLE REQUEST"}); err != nil {
			log.Println("Couldn't send message....", err)
		}
	}

	<-waitc
}
