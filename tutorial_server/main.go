package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/storage"
	"github.com/ponyo877/grpc_tutorial/server"
	pb "github.com/ponyo877/grpc_tutorial/tutorialpb"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	dsn := "root@tcp(127.0.0.1:3306)/app_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	credentialFilePath := "assets/key/gcp_key.json"
	ctx := context.Background()
	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatalf("cannot open grpc client: %v", err)
	}

	pb.RegisterTurtorialServiceServer(s, &server.TutorialServer{
		Db:            db,
		StorageClient: storageClient,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
