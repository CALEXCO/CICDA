package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/CALEXCO/CICDA/examples/learning_grpc/greeter_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
	helloAgain  = "Hello Again jijijijij"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	rAgain, errAgain := c.SayHelloAgain(ctx, &pb.HelloRequest{Name: helloAgain})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if errAgain != nil {
		log.Fatalf("could not greet Again, %v", errAgain)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	log.Printf("Greeting Again: %s", rAgain.GetMessage())
}
