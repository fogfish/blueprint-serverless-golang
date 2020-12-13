package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	µ "github.com/fogfish/gouldian"
	"github.com/fogfish/scud-golang/http"
)

func main() {
	api := http.ServiceScud{}

	lambda.Start(
		µ.Serve(
			api.Lookup(),
		),
	)
}
