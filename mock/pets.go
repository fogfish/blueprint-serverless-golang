package mock

import (
	"context"

	core "github.com/fogfish/blueprint-serverless-golang"
)

var Pets = []core.Pet{
	{ID: "A01", Category: "elk", Price: 12249.99},
	{ID: "A02", Category: "ant", Price: 0.99},
	{ID: "A03", Category: "cow", Price: 3449.99},
	{ID: "A04", Category: "pig", Price: 749.99},
	{ID: "A05", Category: "dog", Price: 249.99},
	{ID: "A06", Category: "cat", Price: 124.99},
	{ID: "A07", Category: "bee", Price: 2.99},
	{ID: "A08", Category: "eel", Price: 49.99},
	{ID: "A09", Category: "owl", Price: 449.99},
	{ID: "A10", Category: "fox", Price: 349.99},
	{ID: "A11", Category: "hen", Price: 19.99},
	{ID: "A12", Category: "bat", Price: 249.99},
	{ID: "A13", Category: "rat", Price: 224.99},
	{ID: "A14", Category: "emu", Price: 924.99},
	{ID: "A15", Category: "gnu", Price: 11449.99},
	{ID: "A16", Category: "ape", Price: 9449.99},
	{ID: "A17", Category: "koi", Price: 74.99},
}

type Creator interface {
	CreatePet(context.Context, core.Pet) error
}

func SetupCreatorWithPets(creator Creator) error {
	for _, pet := range Pets {
		if err := creator.CreatePet(context.Background(), pet); err != nil {
			return err
		}
	}
	return nil
}
