package main

import (
	server "net/http"

	"github.com/fogfish/blueprint-serverless-golang/internal/services/restapi"
	"github.com/fogfish/gouldian/v2/server/httpd"
)

func main() {
	api := restapi.NewPetShopAPI()

	server.ListenAndServe(":8080",
		httpd.Serve(
			api.List(),
			api.Create(),
			api.Lookup(),
		),
	)
}
