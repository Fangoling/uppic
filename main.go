package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/Fangoling/uppic/util"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Operation string

const (
	Resize Operation = "resize"
	Crop   Operation = "crop"
	Rotate Operation = "rotate"
)

func (o *Operation) String() string {
	return string(*o)
}

func (o *Operation) Set(value string) error {
	// List of valid operation
	validOperations := []Operation{Resize, Crop, Rotate}

	for _, op := range validOperations {
		if value == string(op) {
			*o = op
			return nil
		}
	}
	return errors.New("invalid opeation: must be one of the following: Resize, Crop, Rotate")
}

func main() {
	var operation Operation
	var filename string

	flag.StringVar(&filename, "img", "", "Image to be uploaded")
	flag.Var(&operation, "op", "Image operation (resize, crop, rotate)")

	flag.Parse()

	if operation == "" {
		fmt.Println("Error: no operation provided")
		return
	}

	fmt.Printf("Selected image: %s\n", filename)
	fmt.Printf("Selected operation: %s\n", operation)

	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Error when loading default config: ", err)
		return
	}
	filenameInput := filename

	_, err = util.Upload(ctx, filename, "uppic-image-input", filenameInput, sdkConfig)
	if err != nil {
		fmt.Println("Error when uploading file: ", err)
		return
	}

	_, err = util.Poll(ctx, "https://sqs.eu-central-1.amazonaws.com/393937408116/uppic-queue", sdkConfig)
	if err != nil {
		fmt.Println("Error when polling for finished processing message: ", err)
		return
	}

	err = util.Download(ctx, "uppic-image-output", filenameInput, "output/"+filenameInput, sdkConfig)
	if err != nil {
		fmt.Println("Error when downloading file from output bucket: ", err)
		return
	}

	fmt.Println("Succesfully processed file")
}
