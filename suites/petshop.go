package suites

import (
	"fmt"
	"os"

	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/blueprint-serverless-golang/mock"
	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
	"github.com/google/go-cmp/cmp"
)

// TODO: support diff at gurl
// check assay-sdk

type Pets api.Pets

func (pets *Pets) Validate(expect api.Pets) http.Arrow {
	return func(ctx *http.Context) error {
		v := cmp.Diff(api.Pets(*pets), expect)

		os.WriteFile("/tmp/abc", []byte(v), 0666)

		if v != "" {
			return fmt.Errorf("unexpected %s\n", v)
		}

		return nil
	}
}

func TestPetShopList() http.Arrow {
	var pets Pets

	return http.GET(
		ø.URI("/petshop/pets"),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Recv(&pets),
		pets.Validate(api.NewPets(3, mock.Pets[0:4])),
	)
}

func TestPetShopListCursor() http.Arrow {
	var pets Pets

	return http.GET(
		ø.URI("/petshop/pets"),
		ø.Param("cursor", mock.Pets[4].ID),
		ø.Accept.ApplicationJSON,
		ƒ.Status.OK,
		ƒ.Recv(&pets),
		pets.Validate(api.NewPets(3, mock.Pets[4:8])),
	)
}
