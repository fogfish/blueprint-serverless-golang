package main

import (
	server "net/http"

	core "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/blueprint-serverless-golang/cache"
	"github.com/fogfish/blueprint-serverless-golang/http"
	"github.com/fogfish/blueprint-serverless-golang/mock"
	"github.com/fogfish/blueprint-serverless-golang/services/pets"
	"github.com/fogfish/gouldian/v2/server/httpd"
)

func main() {
	storePets := cache.New[core.Identity, core.Pet]()
	storePrice := cache.New[core.Category, core.Price]()

	fetcher := pets.NewFetcher(storePets, storePrice)
	creator := pets.NewCreator(storePets, storePrice)

	mock.SetupCreatorWithPets(creator)

	api := http.NewPetShopAPI(fetcher, creator)

	server.ListenAndServe(":8080",
		httpd.Serve(
			api.List(),
			api.Create(),
			api.Lookup(),
		),
	)
}
