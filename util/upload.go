package util

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func read_file(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Upload(ctx context.Context, filename string, bucket string, key string, sdkConfig aws.Config) (string, error) {
	file, err := read_file(filename)
	if err != nil {
		return "", fmt.Errorf("Failed to read file: %v", err)
	}
	defer file.Close()

	s3Client := s3.NewFromConfig(sdkConfig)
	uploader := manager.NewUploader(s3Client)
	s3Actions := S3Actions{
		S3Client:  s3Client,
		S3Manager: uploader,
	}

	outkey, err := s3Actions.UploadObject(
		ctx,
		"uppic-image-input",
		filename,
		file,
	)

	if err != nil {
		return "", fmt.Errorf("Failed to upload image: %v", err)
	}

	fmt.Printf("File uploaded, %s", outkey)
	return outkey, nil
}
