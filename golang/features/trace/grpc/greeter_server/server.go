package main

import (
	"context"
	"flag"
	"fmt"
	mw_grpc "github.com/middleware-labs/golang-apm-grpc/grpc"
	track "github.com/middleware-labs/golang-apm/tracker"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	track.Track(
		track.WithConfigTag("service", "Your service name"),
		track.WithConfigTag("projectName", "Your project name"),
		track.WithConfigTag("accessToken", "your access token"),
	)
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(mw_grpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(mw_grpc.StreamServerInterceptor()))
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
