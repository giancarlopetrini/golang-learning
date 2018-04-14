package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/satori/go.uuid"

	pb "github.com/giancarlopetrini/golang-learning/examples/grpc/stream1/protobuf"
	"google.golang.org/grpc"
)

func main() {
	//TODO add UUID to context or implement auth
	clientUUID := uuid.Must(uuid.NewV4()).String()
	fmt.Printf("Client UUID: \t %v \n", clientUUID)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't dial gRPC Server: %s", err)
	}
	defer conn.Close()

	c := pb.NewMessageClient(conn)

	stream, err := c.Chat(context.Background())

	// TODO add send, recv channels and range over them
	go func() {
		for {
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Println("Enter chat message:::")
			var txt string
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) == 0 {
					break
				}
				txt = line
				fmt.Printf("Scanned in message: %s \n", txt)
			}

			if err := stream.Send(&pb.Request{Req: txt, Clientuuid: clientUUID}); err != nil {
				log.Println("Couldn't send message....", err)
			}
		}
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
	// blocks code here, since do data being written to waitc
	<-waitc
}
