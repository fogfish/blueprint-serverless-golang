package api

import (
	"fmt"
	"path/filepath"

	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/schemaorg"
)

type Pet struct {
	ID       schemaorg.Identifier `json:"id"`
	Category schemaorg.Category   `json:"category"`
	Price    schemaorg.Price      `json:"price"`
	Url      schemaorg.Url        `json:"url"`
}

func (pet Pet) ToCore() core.Pet {
	return core.Pet{
		ID:       string(pet.ID),
		Category: string(pet.Category),
		Price:    float64(pet.Price),
	}
}

func NewPet(pet core.Pet) Pet {
	return Pet{
		ID:       schemaorg.Identifier(pet.ID),
		Category: schemaorg.Category(pet.Category),
		Price:    schemaorg.Price(pet.Price),
		Url:      schemaorg.Url(filepath.Join("/", core.PETSHOP, core.PETS, pet.ID)),
	}
}

type Pets struct {
	Pets []Pet          `json:"pets,omitempty"`
	Next *schemaorg.Url `json:"next,omitempty"`
}

func NewPets(size int, pets []core.Pet) Pets {
	cursor := ""
	if len(pets) > size {
		path := filepath.Join("/", core.PETSHOP, core.PETS)
		pets, cursor = pets[:size], fmt.Sprintf("%s?cursor=%s", path, pets[size].ID)
	}

	seq := make([]Pet, len(pets))
	for i, pet := range pets {
		seq[i] = NewPet(pet)
	}

	return Pets{
		Pets: seq,
		Next: (*schemaorg.Url)(&cursor),
	}
}
