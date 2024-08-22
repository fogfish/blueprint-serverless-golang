package pets

import (
	"context"

	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/faults"
)

const (
	errCreator = faults.Safe1[core.Identity]("failed to create pet (%s)")
)

type WriterPet interface {
	core.Setter[core.Identity, core.Pet]
}

type WriterPrice interface {
	core.Setter[core.Category, core.Price]
}

type Creator struct {
	pets  WriterPet
	price WriterPrice
}

func NewCreator(pets WriterPet, price WriterPrice) Creator {
	return Creator{pets: pets, price: price}
}

func (lib Creator) CreatePet(ctx context.Context, pet core.Pet) error {
	err := lib.price.Set(ctx, pet.Category, pet.Price)
	if err != nil {
		return errCreator.New(err, pet.ID)
	}

	err = lib.pets.Set(ctx, pet.ID, pet)
	if err != nil {
		return errCreator.New(err, pet.ID)
	}

	return nil
}
