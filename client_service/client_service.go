package client_service

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	pb "github.com/ponyo877/grpc_tutorial/tutorialpb"
)

type ClientService struct {
	TutorialClient pb.TurtorialServiceClient
}

func (cs *ClientService) SaveImage(c echo.Context) error {
	ctx := context.Background()

	placeId := c.Param("placeId")
	userId := c.FormValue("userId")
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	f, err := file.Open()
	defer f.Close()
	if err != nil {
		log.Printf("cannot open file: %v ", err)
		return err
	}

	stream, err := cs.TutorialClient.SaveImage(ctx)
	if err != nil {
		log.Printf("fail tutorialClient.SaveImage: %v", err)
		return c.JSON(http.StatusOK, stream)
	}
	req := &pb.SaveImageRequest{
		Data: &pb.SaveImageRequest_Info{
			Info: &pb.ImageInfo{
				PlaceId: placeId,
				UserId:  userId,
			},
		},
	}
	err = stream.Send(req)
	buffer := make([]byte, 1024)
	for {
		n, err := f.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot send chunk to server: %v", err)
		}
		req := &pb.SaveImageRequest{
			Data: &pb.SaveImageRequest_ImageBinary{
				ImageBinary: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Printf("cannot read chunk to buffer: %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("cannot receive response: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (cs *ClientService) GetImageUrl(c echo.Context) error {
	ctx := context.Background()

	placeId := c.Param("placeId")

	resp, err := cs.TutorialClient.GetImageUrl(ctx, &pb.GetImageUrlRequest{PlaceId: placeId})
	if err != nil {
		log.Printf("fail tutorialClient.GetImageUrl: %v", err)
		return c.JSON(http.StatusOK, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

func (cs *ClientService) SaveColorCode(c echo.Context) error {
	ctx := context.Background()

	placeId := c.Param("placeId")
	userId := c.FormValue("userId")
	colorCode := c.FormValue("colorCode")

	resp, err := cs.TutorialClient.SaveColorCode(ctx, &pb.SaveColorCodeRequest{
		PlaceId:   placeId,
		UserId:    userId,
		ColorCode: colorCode,
	})
	if err != nil {
		log.Printf("fail tutorialClient.SaveColorCode: %v", err)
		return c.JSON(http.StatusNotFound, resp)
	}
	return c.JSON(http.StatusOK, resp)
}

func (cs *ClientService) GetColorCode(c echo.Context) error {
	ctx := context.Background()

	placeId := c.Param("placeId")

	resp, err := cs.TutorialClient.GetColorCode(ctx, &pb.GetColorCodeRequest{PlaceId: placeId})
	if err != nil {
		log.Printf("fail tutorialClient.GetColorCode: %v", err)
		return c.JSON(http.StatusOK, resp)
	}
	return c.JSON(http.StatusOK, resp)
}
