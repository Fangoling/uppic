package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func create_s3uploader() (*s3manager.Uploader, error) {
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	})
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(s3Session)
	return uploader, nil
}

func read_file(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Upload(filename, bucket, key string) error {
	file, err := read_file(filename)
	if err != nil {
		return fmt.Errorf("Failed to read file: %v", err)
	}

	uploader, err := create_s3uploader()
	if err != nil {
		return fmt.Errorf("Failed to create uploader: %v", err)
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	file.Close()

	if err != nil {
		return fmt.Errorf("Failed to upload image: %v", err)
	}

	fmt.Printf("File uploaded to %s\n", result.Location)
	return nil
}
