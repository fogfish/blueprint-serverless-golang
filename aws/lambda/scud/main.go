package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fogfish/blueprint-serverless-golang/http"
	µ "github.com/fogfish/gouldian"
)

func main() {
	api := http.ServiceScud{}

	lambda.Start(
		µ.Serve(
			api.Lookup(),
		),
	)
}
