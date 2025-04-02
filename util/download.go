package util

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Download(ctx context.Context, bucketName string, objectkey string, filename string, sdkConfig aws.Config) error {
	s3Client := s3.NewFromConfig(sdkConfig)
	simpleBucket := BucketBasics{
		S3Client: s3Client,
	}
	err := simpleBucket.DownloadFile(ctx, bucketName, objectkey, filename)
	if err != nil {
		return fmt.Errorf("Error when downloading file %s from Bucket %s at location %s", objectkey, bucketName, filename)
	}
	return nil
}
