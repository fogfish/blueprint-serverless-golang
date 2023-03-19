package http

import (
	"context"

	core "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/faults"
	µ "github.com/fogfish/gouldian/v2"
	ø "github.com/fogfish/gouldian/v2/output"
	"github.com/fogfish/schemaorg"
)

//go:generate mockgen -destination ../mock/petshop.go -package mock . PetFetcher,PetCreator

type PetFetcher interface {
	LookupPet(context.Context, core.Identity) (core.Pet, error)
	LookupPetsAfterKey(context.Context, core.Identity, int) ([]core.Pet, error)
}

type PetCreator interface {
	CreatePet(context.Context, core.Pet) error
}

type PetShopAPI struct {
	fetcher PetFetcher
	creator PetCreator
}

func NewPetShopAPI(fetcher PetFetcher, creator PetCreator) *PetShopAPI {
	return &PetShopAPI{
		fetcher: fetcher,
		creator: creator,
	}
}

type reqPetShop struct {
	ID  schemaorg.Identifier
	Pet api.Pet
}

var (
	reqID   = µ.Optics1[reqPetShop, schemaorg.Identifier]()
	reqPet  = µ.Optics1[reqPetShop, api.Pet]()
	petshop = µ.Path("petshop")
	pets    = µ.Path("pets")
	petID   = µ.Path(reqID)
	petSeqN = 3
)

func (shop PetShopAPI) List() µ.Routable {
	return µ.GET(
		µ.URI(petshop, pets),
		µ.ParamMaybe("cursor", reqID),
		µ.Accept.ApplicationJSON,
		µ.FMap(func(ctx *µ.Context, req *reqPetShop) error {
			seq, err := shop.fetcher.LookupPetsAfterKey(ctx, string(req.ID), petSeqN+1)
			switch {
			case err == nil:
				pets := api.NewPets(petSeqN, seq)
				return ø.Status.OK(
					ø.ContentType.ApplicationJSON,
					// ø.L
					ø.Send(pets),
				)
			default:
				return ø.Status.InternalServerError(ø.Error(err))
			}
		}),
	)
}

func (shop PetShopAPI) Create() µ.Routable {
	return µ.POST(
		µ.URI(petshop, pets),
		AllowSecretCode(),
		µ.ContentType.ApplicationJSON,
		µ.Body(reqPet),
		µ.FMap(func(ctx *µ.Context, req *reqPetShop) error {
			pet := req.Pet.ToCore()
			err := shop.creator.CreatePet(ctx, pet)
			switch {
			case err == nil:
				return ø.Status.Created()
			default:
				return ø.Status.InternalServerError(ø.Error(err))
			}
		}),
	)
}

func (shop PetShopAPI) Lookup() µ.Routable {
	return µ.GET(
		µ.URI(petshop, pets, petID),
		µ.Accept.ApplicationJSON,
		µ.FMap(func(ctx *µ.Context, req *reqPetShop) error {
			pet, err := shop.fetcher.LookupPet(ctx, string(req.ID))
			switch {
			case err == nil:
				return ø.Status.OK(
					ø.ContentType.ApplicationJSON,
					ø.Send(api.NewPet(pet)),
				)
			case faults.IsNotFound(err):
				return ø.Status.NotFound(ø.Error(err))
			default:
				return ø.Status.InternalServerError(ø.Error(err))
			}
		}),
	)
}
