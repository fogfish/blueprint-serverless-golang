package restapi

import (
	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/blueprint-serverless-golang/internal/http"
	"github.com/fogfish/blueprint-serverless-golang/internal/mock"
	"github.com/fogfish/blueprint-serverless-golang/internal/services/pets"
	cache "github.com/fogfish/blueprint-serverless-golang/internal/storage"
)

func NewPetShopAPI() *http.PetShopAPI {
	storePets := cache.New[core.Identity, core.Pet]()
	storePrice := cache.New[core.Category, core.Price]()

	fetcher := pets.NewFetcher(storePets, storePrice)
	creator := pets.NewCreator(storePets, storePrice)

	mock.SetupCreatorWithPets(creator)

	return http.NewPetShopAPI(fetcher, creator)
}
