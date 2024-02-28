package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type AppEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event *AppEvent) (*string, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	message := fmt.Sprintf("Hi %s", event.Name)
	return &message, nil
}

func main() {
	lambda.Start(HandleRequest)
}
