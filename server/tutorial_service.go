package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/oklog/ulid/v2"
	"github.com/ponyo877/grpc_tutorial/model"
	pb "github.com/ponyo877/grpc_tutorial/tutorialpb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var maxImageSize = 1048576

type TutorialServer struct {
	Db            *gorm.DB
	StorageClient *storage.Client
	pb.UnimplementedTurtorialServiceServer
}

func (ts *TutorialServer) PrintStr(ctx context.Context, in *pb.PrintStrRequest) (*pb.PrintStrResponce, error) {
	log.Printf("Received: %v", in.GetMessage())
	return &pb.PrintStrResponce{Message: "Hello " + in.GetMessage()}, nil
}

func GetULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func (ts *TutorialServer) SaveImageUnary(ctx context.Context, in *pb.SaveImageUnaryRequest) (*pb.SaveImageUnaryResponce, error) {

	placeId := in.GetPlaceId()
	userId := in.GetUserId()
	logUrlId := GetULID()

	imageBinary := in.GetImageBinary()

	tutorial_recode := &model.Tutorial{
		PlaceID:   placeId,
		LogoUrlID: logUrlId,
		UpdatedBy: userId,
		CreatedBy: userId,
	}

	ts.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "place_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"logo_url_id", "updated_by"}),
	}).Create(tutorial_recode)

	bucket := "tutorial_backet"
	object := logUrlId + ".jpg"
	writer := ts.StorageClient.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err := writer.Write(imageBinary); err != nil {
		return &pb.SaveImageUnaryResponce{}, err
	}
	log.Print("finish saving file to GCS")
	if err := writer.Close(); err != nil {
		panic(err)
	}

	return &pb.SaveImageUnaryResponce{}, nil
}

func (ts *TutorialServer) SaveImage(stream pb.TurtorialService_SaveImageServer) error {
	ctx := context.Background()

	imageData := bytes.Buffer{}
	imageSize := 0
	req, err := stream.Recv()
	placeId := req.GetInfo().GetPlaceId()
	userId := req.GetInfo().GetUserId()
	logUrlId := GetULID()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("cannot read chunk to buffer: %v", err)
			return err
		}
		chunkData := resp.GetImageBinary()
		chunkSize := len(chunkData)
		log.Printf("received a chunk with size: %d", chunkSize)

		imageSize += chunkSize
		if imageSize > maxImageSize {
			log.Printf("file size limit exceeded: %v", err)
			return errors.New("file size limit exceeded")
		}
		if _, err = imageData.Write(chunkData); err != nil {
			log.Printf("cannot write chunk data: %v", err)
			return err
		}
	}
	log.Print("finish writing file to bytes.Buffer")

	tutorial_recode := &model.Tutorial{
		PlaceID:   placeId,
		LogoUrlID: logUrlId,
		UpdatedBy: userId,
		CreatedBy: userId,
	}

	ts.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "place_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"logo_url_id", "updated_by"}),
	}).Create(tutorial_recode)

	bucket := "tutorial_backet"
	object := logUrlId + ".jpg"
	writer := ts.StorageClient.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err := writer.Write(imageData.Bytes()); err != nil {
		return err
	}
	log.Print("finish saving file to GCS")
	if err := writer.Close(); err != nil {
		panic(err)
	}

	if err = stream.SendAndClose(&pb.SaveImageResponce{}); err != nil {
		log.Printf("cannot write chunk data: %v", err)
		return err
	}
	return nil
}

func (ts *TutorialServer) GetImageUrl(ctx context.Context, in *pb.GetImageUrlRequest) (*pb.GetImageUrlResponce, error) {
	tutorial_recode := &model.Tutorial{}
	ts.Db.First(tutorial_recode, "place_id = ?", in.PlaceId)
	if tutorial_recode.LogoUrlID == "" {
		log.Printf("log image recode is nothing")
		return &pb.GetImageUrlResponce{ImageUrl: ""}, nil
	}
	bucket := "tutorial_backet"
	object := tutorial_recode.LogoUrlID + ".jpg"
	var gcsCredential model.GCSCredential

	credentialFilePath := "assets/key/gcp_key.json"
	jsonFile, err := os.Open(credentialFilePath)
	if err != nil {
		log.Printf("cannot open gcs credential file: %v", err)
		return &pb.GetImageUrlResponce{}, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &gcsCredential); err != nil {
		log.Printf("cannot parse gcs credential file: %v", err)
		return &pb.GetImageUrlResponce{}, err
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		GoogleAccessID: gcsCredential.ClientEmail,
		PrivateKey:     []byte(gcsCredential.PrivateKeyStr),
		Method:         "GET",
		Expires:        time.Now().Add(15 * time.Minute),
	}

	signedUrl, err := storage.SignedURL(bucket, object, opts)
	if err != nil {
		return &pb.GetImageUrlResponce{}, fmt.Errorf("Bucket(%q).SignedURL: %v", bucket, err)
	}
	return &pb.GetImageUrlResponce{ImageUrl: signedUrl}, nil
}

func (ts *TutorialServer) SaveColorCode(ctx context.Context, in *pb.SaveColorCodeRequest) (*pb.SaveColorCodeResponce, error) {
	placeId := in.GetPlaceId()
	userId := in.GetUserId()
	colorCode := in.GetColorCode()
	log.Printf("[Debug] SaveColorCode colorCode: %v", colorCode)
	tutorial_recode := &model.Tutorial{
		PlaceID:   placeId,
		ColorCode: colorCode,
		UpdatedBy: userId,
		CreatedBy: userId,
	}

	ts.Db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "place_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"color_code", "updated_by"}),
	}).Create(tutorial_recode)
	return &pb.SaveColorCodeResponce{}, nil
}

func (ts *TutorialServer) GetColorCode(ctx context.Context, in *pb.GetColorCodeRequest) (*pb.GetColorCodeResponce, error) {
	tutorial_recode := &model.Tutorial{}
	ts.Db.First(tutorial_recode, "place_id = ?", in.PlaceId)
	log.Printf("[Debug] GetColorCode ColorCode: %v", tutorial_recode.ColorCode)
	return &pb.GetColorCodeResponce{ColorCode: tutorial_recode.ColorCode}, nil
}
