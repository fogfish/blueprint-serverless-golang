package suites

import (
	"os"

	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
)

var host = os.Getenv("INPUT_TARGET")

// TestNews endpoint
func TestNews() http.Arrow {
	return http.GET(
		ø.URI("%s/scud", host),
		ƒ.Status.OK,
	)
}
