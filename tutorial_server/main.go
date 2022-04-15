package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/ponyo877/grpc_tutorial/tutorialpb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPrinterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) PrintStr(ctx context.Context, in *pb.PrintStrRequest) (*pb.PrintStrResponce, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PrintStrResponce{Message: "Hello " + in.GetMessage()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrinterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
