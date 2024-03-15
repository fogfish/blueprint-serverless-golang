package http

import (
	"fmt"
	µ "github.com/fogfish/gouldian/v2"
	"net/http"
)

// Customer Endpoint
func AllowSecretCode() µ.Endpoint {
	return func(ctx *µ.Context) error {
		code := ctx.Request.Header.Get("X-Secret-Code")
		if code == "" {
			out := µ.NewOutput(http.StatusUnauthorized)
			out.SetIssue(fmt.Errorf("unauthorized %s", ctx.Request.URL.Path))
			return out
		}

		if code != "cGV0c3RvcmU6b3duZXIK" {
			return µ.ErrNoMatch
		}

		return nil
	}
}
