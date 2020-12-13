package scud_test

import (
	"testing"

	"github.com/fogfish/it"
	"github.com/fogfish/scud-golang"
)

func TestStub(t *testing.T) {
	it.Ok(t).If(scud.NewStubs()).Should().Equal(
		[]scud.Stub{
			{ID: "a"},
			{ID: "b"},
			{ID: "c"},
		},
	)
}
