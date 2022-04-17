package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/ponyo877/grpc_tutorial/client_service"
	pb "github.com/ponyo877/grpc_tutorial/tutorialpb"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051"
	flag.Parse()
	log.Printf("dial server: %s", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("cannot dial server: %v ", err)
	}
	tutorialClient := pb.NewTurtorialServiceClient(conn)
	cs := &client_service.ClientService{
		TutorialClient: tutorialClient,
	}

	e := echo.New()
	e.POST("/save_image/:placeId", cs.SaveImage)
	e.GET("/get_image_url/:placeId", cs.GetImageUrl)
	e.POST("/save_color_code/:placeId", cs.SaveColorCode)
	e.GET("/get_color_code/:placeId", cs.GetColorCode)

	e.Start(":8080")
}
