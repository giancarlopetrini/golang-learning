package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

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

	sendc := make(chan struct{})
	go func() {
		for i := 0; i < 2; i++ {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("Enter sting:::")
			var txt string
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) == 0 {
					break
				}
				txt = line
				fmt.Printf("Scanned in message: %s \n", txt)
			}

			if err := stream.Send(&pb.Request{Req: txt}); err != nil {
				log.Println("Couldn't send message....", err)
			}
		}
		close(sendc)
	}()

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
	<-waitc

	stream.CloseSend()
}
