package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fogfish/blueprint-serverless-golang/internal/services/restapi"
	httpd "github.com/fogfish/gouldian/v2/server/aws/apigateway"
)

func main() {
	api := restapi.NewPetShopAPI()

	lambda.Start(
		httpd.Serve(
			api.List(),
			api.Lookup(),
		),
	)
}
