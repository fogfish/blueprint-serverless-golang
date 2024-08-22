package core

import "context"

//
// Example definition of data access objects.
// Interfaces abstracts capabilities of the storage layer(s).
//

type Getter[K, V any] interface {
	Get(context.Context, K) (V, error)
}

type GetterSeq[K, V any] interface {
	Seq(context.Context, K, int) ([]V, error)
}

type Setter[K, V any] interface {
	Set(context.Context, K, V) error
}
