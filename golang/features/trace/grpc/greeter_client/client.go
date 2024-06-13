package main

import (
	"context"
	"flag"
	"log"
	"time"

	mw_grpc "github.com/middleware-labs/golang-apm-grpc/grpc"
	track "github.com/middleware-labs/golang-apm/tracker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	track.Track(
		track.WithConfigTag(track.Service, "your service name"),
		track.WithConfigTag(track.Project, "your project name"),
		track.WithConfigTag(track.Token, "your access token"),
	)
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw_grpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(mw_grpc.StreamClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
