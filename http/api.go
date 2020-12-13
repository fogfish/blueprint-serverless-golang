package http

import (
	µ "github.com/fogfish/gouldian"
	"github.com/fogfish/gouldian/path"
	"github.com/fogfish/scud-golang"
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
