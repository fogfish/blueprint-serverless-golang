package curl

import (
	"context"

	"github.com/fogfish/blueprint-serverless-golang/http/api"
	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
	"github.com/fogfish/schemaorg"
)

type PetShop struct {
	http.Stack
	host ø.Authority
}

func NewPetShop(stack http.Stack, host string) *PetShop {
	return &PetShop{
		Stack: stack,
		host:  ø.Authority(host),
	}
}

func (c *PetShop) List(ctx context.Context) (*api.Pets, error) {
	return http.IO[api.Pets](c.WithContext(ctx),
		http.GET(
			ø.URI("%s/petshop/pets", c.host),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}

func (c *PetShop) Continue(ctx context.Context, cursor schemaorg.Url) (*api.Pets, error) {
	return http.IO[api.Pets](c.WithContext(ctx),
		http.GET(
			ø.URI("%s%s", c.host, ø.Path(cursor)),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}

func (c *PetShop) Pet(ctx context.Context, url schemaorg.Url) (*api.Pet, error) {
	return http.IO[api.Pet](c.WithContext(ctx),
		http.GET(
			ø.URI("%s%s", c.host, ø.Path(url)),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}
