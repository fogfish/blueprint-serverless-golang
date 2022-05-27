package http

import (
	scud "github.com/fogfish/blueprint-serverless-golang"
	µ "github.com/fogfish/gouldian"
)

// ServiceScud is and example RESTfull Service
type ServiceScud struct{}

// Lookup is an example implementation of endpoint
func (api ServiceScud) Lookup() µ.Routable {
	return µ.GET(
		µ.URI(µ.Path("scud")),
		func(*µ.Context) error {
			return µ.Status.OK(
				µ.WithJSON(scud.NewStubs()),
			)
		},
	)
}
