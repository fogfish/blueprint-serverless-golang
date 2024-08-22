package suites

import (
	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/blueprint-serverless-golang/internal/mock"
	"github.com/fogfish/blueprint-serverless-golang/pkg/api"
	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
)

func TestPetShopList() http.Arrow {
	return http.GET(
		ø.URI("/%s/%s", core.PETSHOP, core.PETS),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPets(3, mock.Pets[0:4])),
	)
}

func TestPetShopListWithCursor() http.Arrow {
	return http.GET(
		ø.URI("/%s/%s", core.PETSHOP, core.PETS),
		ø.Param("cursor", mock.Pets[4].ID),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPets(3, mock.Pets[4:8])),
	)
}

func TestPetShopLookup() http.Arrow {
	return http.GET(
		ø.URI("/%s/%s/%s", core.PETSHOP, core.PETS, mock.Pets[16].ID),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPet(mock.Pets[16])),
	)
}

func TestPetShopCreate() http.Arrow {
	return http.POST(
		ø.URI("/%s/%s/%s", core.CONSOLE, core.PETSHOP, core.PETS),
		ø.ContentType.ApplicationJSON,
		ø.Send(mock.Pets[16]),
		ƒ.Status.Created,
	)
}

func TestPetShopCreateUnauthorized() http.Arrow {
	return http.POST(
		ø.URI("/%s/%s", core.PETSHOP, core.PETS),
		ø.ContentType.ApplicationJSON,
		ø.Send(mock.Pets[16]),
		ƒ.Status.Unauthorized,
	)
}
