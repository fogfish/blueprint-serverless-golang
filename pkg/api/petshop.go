package api

import (
	"context"

	"github.com/fogfish/blueprint-serverless-golang/internal/core"
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

func (c *PetShop) List(ctx context.Context) (*Pets, error) {
	return http.IO[Pets](c.WithContext(ctx),
		http.GET(
			ø.URI("%s/%s/%s", c.host, core.PETSHOP, core.PETS),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}

func (c *PetShop) Continue(ctx context.Context, cursor schemaorg.Url) (*Pets, error) {
	return http.IO[Pets](c.WithContext(ctx),
		http.GET(
			ø.URI("%s%s", c.host, ø.Path(cursor)),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}

func (c *PetShop) Pet(ctx context.Context, url schemaorg.Url) (*Pet, error) {
	return http.IO[Pet](c.WithContext(ctx),
		http.GET(
			ø.URI("%s%s", c.host, ø.Path(url)),
			ø.Accept.ApplicationJSON,
			ƒ.Status.OK,
		),
	)
}

func (c *PetShop) Create(ctx context.Context, pet *Pet) error {
	return c.IO(ctx,
		http.POST(
			ø.URI("%s/%s/%s/%s", c.host, core.CONSOLE, core.PETSHOP, core.PETS),
			ø.ContentType.ApplicationJSON,
			ø.Send(pet),
			ƒ.Status.Created,
		),
	)
}
