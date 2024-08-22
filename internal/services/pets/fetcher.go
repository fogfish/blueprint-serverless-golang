package pets

import (
	"context"

	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/faults"
)

const (
	errFetcher = faults.Safe1[core.Identity]("failed to fetch pet (%s)")
)

type ReaderPet interface {
	core.Getter[core.Identity, core.Pet]
	core.GetterSeq[core.Identity, core.Pet]
}

type ReaderPrice interface {
	core.Getter[core.Category, core.Price]
}

type Fetcher struct {
	pets  ReaderPet
	price ReaderPrice
}

func NewFetcher(pets ReaderPet, price ReaderPrice) Fetcher {
	return Fetcher{pets: pets, price: price}
}

func (lib Fetcher) LookupPet(ctx context.Context, key core.Identity) (core.Pet, error) {
	pet, err := lib.pets.Get(ctx, key)
	if err != nil {
		return core.Pet{}, errFetcher.New(err, key)
	}

	pet.Price, err = lib.price.Get(ctx, pet.Category)
	if err != nil {
		return core.Pet{}, errFetcher.New(err, key)
	}

	return pet, nil
}

func (lib Fetcher) LookupPetsAfterKey(ctx context.Context, afterKey core.Identity, n int) ([]core.Pet, error) {
	pets, err := lib.pets.Seq(ctx, afterKey, n)
	if err != nil {
		return nil, errFetcher.New(err, afterKey)
	}

	for i := 0; i < len(pets); i++ {
		pets[i].Price, err = lib.price.Get(ctx, pets[i].Category)
		if err != nil {
			return nil, errFetcher.New(err, afterKey)
		}
	}

	return pets, nil
}
