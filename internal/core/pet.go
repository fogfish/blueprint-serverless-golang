package core

type Identity = string

type Category = string

type Price = float64

// Pet is an example domain type
type Pet struct {
	ID       Identity
	Category Category
	Price    Price
}
