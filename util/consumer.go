package util

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func Poll(ctx context.Context, messageQueueLink string, sdkConfig aws.Config) (string, error) {
	sqsClient := sqs.NewFromConfig(sdkConfig)
	sqsActions := SqsActions{
		SqsClient: sqsClient,
	}

	messages, err := sqsActions.GetMessages(ctx, messageQueueLink, 1, 20)
	if err != nil {
		return "", fmt.Errorf("Error when getting message from sqs: %s", err)
	}

	return *messages[0].Body, nil
}
