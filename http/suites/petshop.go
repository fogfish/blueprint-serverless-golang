package suites

import (
	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/blueprint-serverless-golang/internal/mock"
	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
)

func TestPetShopList() http.Arrow {
	return http.GET(
		ø.URI("/petshop/pets"),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPets(3, mock.Pets[0:4])),
	)
}

func TestPetShopListWithCursor() http.Arrow {
	return http.GET(
		ø.URI("/petshop/pets"),
		ø.Param("cursor", mock.Pets[4].ID),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPets(3, mock.Pets[4:8])),
	)
}

func TestPetShopLookup() http.Arrow {
	return http.GET(
		ø.URI("/petshop/pets/"+mock.Pets[16].ID),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Expect(api.NewPet(mock.Pets[16])),
	)
}

func TestPetShopCreate() http.Arrow {
	return http.POST(
		ø.URI("/petshop/pets"),
		ø.Authorization.Set("Basic cGV0c3RvcmU6b3duZXIK"),
		ø.ContentType.ApplicationJSON,
		ø.Send(mock.Pets[16]),
		ƒ.Status.Created,
	)
}

func TestPetShopCreateUnauthorized() http.Arrow {
	return http.POST(
		ø.URI("/petshop/pets"),
		ø.ContentType.ApplicationJSON,
		ø.Send(mock.Pets[16]),
		ƒ.Status.Unauthorized,
	)
}
