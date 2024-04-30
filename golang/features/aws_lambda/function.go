package main

import (
	"context"
	"fmt"

	track "github.com/middleware-labs/golang-apm/tracker"

	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

var ctx = context.Background()

func HandleLambdaEvent(event *MyEvent) (*MyResponse, error) {
	if event == nil {
		return nil, fmt.Errorf("received nil event")
	}
	track.SetAttribute(ctx, "user.name", event.Name)
	track.SetAttribute(ctx, "user.age", fmt.Sprint(event.Age))
	return &MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	config, _ := track.Track(
		track.WithConfigTag("target", "<target url of middleware account>"),
		track.WithConfigTag("service", "<service name>"),
		track.WithConfigTag("projectName", "<project name>"),
		track.WithConfigTag("accessToken", "<access token here>"),
	)

	lambda.Start(otellambda.InstrumentHandler(HandleLambdaEvent, otellambda.WithFlusher(config.Tp)))
}
