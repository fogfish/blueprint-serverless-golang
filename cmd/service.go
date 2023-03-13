package cmd

import (
	core "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/blueprint-serverless-golang/http"
	"github.com/fogfish/blueprint-serverless-golang/internal/cache"
	"github.com/fogfish/blueprint-serverless-golang/internal/mock"
	"github.com/fogfish/blueprint-serverless-golang/internal/services/pets"
)

func NewPetShopAPI() *http.PetShopAPI {
	storePets := cache.New[core.Identity, core.Pet]()
	storePrice := cache.New[core.Category, core.Price]()

	fetcher := pets.NewFetcher(storePets, storePrice)
	creator := pets.NewCreator(storePets, storePrice)

	mock.SetupCreatorWithPets(creator)

	return http.NewPetShopAPI(fetcher, creator)
}
