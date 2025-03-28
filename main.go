package main

import (
	"errors"
	"flag"
	"fmt"
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
}
