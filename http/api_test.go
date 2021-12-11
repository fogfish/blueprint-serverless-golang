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

	endpoint := mock.Endpoint(api.Lookup())
	expect := µ.Status.OK(
		µ.WithHeader("Content-Type", "application/json"),
		µ.WithBytes([]byte(`[{"id":"a"},{"id":"b"},{"id":"c"}]`)),
	)

	it.Ok(t).
		If(endpoint(req)).Should().Equal(expect)
}
