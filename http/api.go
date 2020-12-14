package http

import (
	scud "github.com/fogfish/blueprint-serverless-golang"
	µ "github.com/fogfish/gouldian"
	"github.com/fogfish/gouldian/path"
)

// ServiceScud is and example RESTfull Service
type ServiceScud struct{}

// Lookup is an example implementation of endpoint
func (api ServiceScud) Lookup() µ.Endpoint {
	return µ.GET(
		µ.Path(path.Is("scud")),
		µ.FMap(func() error {
			return µ.Ok().JSON(scud.NewStubs())
		}),
	)
}
