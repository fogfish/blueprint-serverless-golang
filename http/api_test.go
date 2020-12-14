package http_test

import (
	"testing"

	"github.com/fogfish/blueprint-serverless-golang/http"
	µ "github.com/fogfish/gouldian"
	"github.com/fogfish/gouldian/mock"
	"github.com/fogfish/it"
)

func TestLookup(t *testing.T) {
	api := http.ServiceScud{}

	req := mock.Input(
		mock.Method("GET"),
		mock.URL("/scud"),
	)

	expect := µ.Ok().
		With("Content-Type", "application/json").
		Bytes([]byte(`[{"id":"a"},{"id":"b"},{"id":"c"}]`))

	it.Ok(t).
		If(api.Lookup()(req)).Should().Equal(expect)
}
