package main

import (
	server "net/http"

	"github.com/fogfish/blueprint-serverless-golang/cmd"
	"github.com/fogfish/gouldian/v2/server/httpd"
)

func main() {
	api := cmd.NewPetShopAPI()

	server.ListenAndServe(":8080",
		httpd.Serve(
			api.List(),
			api.Create(),
			api.Lookup(),
		),
	)
}
