package http

import (
	µ "github.com/fogfish/gouldian/v2"
)

func AllowSecretCode() µ.Endpoint {
	return µ.Authorization(
		func(kind, digest string) error {
			if kind != "Basic" {
				return µ.ErrNoMatch
			}

			if digest != "cGV0c3RvcmU6b3duZXIK" {
				return µ.ErrNoMatch
			}

			return nil
		},
	)
}
