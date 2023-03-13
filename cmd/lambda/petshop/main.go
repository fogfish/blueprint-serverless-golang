package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fogfish/blueprint-serverless-golang/cmd"
	httpd "github.com/fogfish/gouldian/v2/server/aws/apigateway"
)

func main() {
	api := cmd.NewPetShopAPI()

	lambda.Start(
		httpd.Serve(
			api.List(),
			api.Create(),
			api.Lookup(),
		),
	)
}
