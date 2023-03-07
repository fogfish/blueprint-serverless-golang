package suites

import (
	"github.com/fogfish/gurl/v2/http"
	ƒ "github.com/fogfish/gurl/v2/http/recv"
	ø "github.com/fogfish/gurl/v2/http/send"
)

// TestNews endpoint
func TestNews() http.Arrow {
	return http.GET(
		ø.URI("/scud"),
		ƒ.Status.OK,
	)
}
