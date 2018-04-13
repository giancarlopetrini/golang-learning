package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/giancarlopetrini/golang-learning/examples/grpc/stream1/protobuf"
	"google.golang.org/grpc"
)

type messageServer struct{}

func (*messageServer) Chat(stream pb.Message_ChatServer) error {
	fmt.Println("Running Chat RPC")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("Received: %v \n", req.Req)
		res := &pb.Response{
			Res: "DEMO RESPONSE",
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	conn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln("Couldn't setup listener", err)
	}

	s := messageServer{}
	grpcServer := grpc.NewServer()

	pb.RegisterMessageServer(grpcServer, &s)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("gRPC Server failed to server: %s", err)
	}
}
