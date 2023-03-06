package http_test

import (
	"testing"

	"github.com/fogfish/blueprint-serverless-golang/http"
	"github.com/fogfish/gouldian/v2/mock"
	ø "github.com/fogfish/gouldian/v2/output"
	"github.com/fogfish/it"
)

func TestLookup(t *testing.T) {
	api := http.ServiceScud{}

	req := mock.Input(
		mock.Method("GET"),
		mock.URL("/scud"),
	)

	endpoint := mock.Endpoint(api.Lookup())
	expect := ø.Status.OK(
		ø.ContentType.ApplicationJSON,
		ø.Send([]struct {
			ID string `json:"id"`
		}{{ID: "a"}, {ID: "b"}, {ID: "c"}}),
	)

	it.Ok(t).
		If(endpoint(req)).Should().Equal(expect)
}
