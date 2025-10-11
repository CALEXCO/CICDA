package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "examples.CICDA/examples/greeter_proto"
	"google.golang.org/grpc"
)

type greeter_server struct {
	pb.UnimplementedGreeterServer
}

var (
	port = flag.Int("port", 50051, "Server Port that is listening")
)

// Example usage to avoid unused type error
func NewGreeterServer() *greeter_server {
	return &greeter_server{}
}

func (s *greeter_server) SayHello(ctx context.Context, clientRequest *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v\nContext: %v", clientRequest.GetName(), ctx.Done())
	return &pb.HelloReply{Message: "Hello" + clientRequest.GetName()}, nil
}

func (s *greeter_server) SayHelloAgain(ctx context.Context, clientRequest *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("The Server received: %v\nContext: %v", clientRequest.GetName(), ctx.Done())
	return &pb.HelloReply{Message: "Hello again!!" + clientRequest.GetName()}, nil
}

func main() {
	flag.Parse()
	serverPortListen, errorPortListen := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if errorPortListen != nil {
		log.Fatalf("Error listening to port: %v", errorPortListen)
	}

	gRPCServer := grpc.NewServer()

	pb.RegisterGreeterServer(gRPCServer, NewGreeterServer())

	log.Printf("Server listening at: %v", serverPortListen.Addr())

	if errorgRPCServer := gRPCServer.Serve(serverPortListen); errorPortListen != nil {
		log.Fatalf("Failed to serve: %v", errorgRPCServer)
	}

}
