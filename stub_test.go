package scud_test

import (
	"testing"

	scud "github.com/fogfish/blueprint-serverless-golang"
	"github.com/fogfish/it"
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
