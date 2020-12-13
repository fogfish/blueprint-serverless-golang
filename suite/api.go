package suite

import (
	"github.com/assay-it/sdk-go/assay"
	ç "github.com/assay-it/sdk-go/cats"
	"github.com/assay-it/sdk-go/http"
	ƒ "github.com/assay-it/sdk-go/http/recv"
	ø "github.com/assay-it/sdk-go/http/send"
	scud "github.com/fogfish/blueprint-serverless-golang"
)

var host = assay.Host("")

// TestLookup endpoint
func TestLookup() assay.Arrow {
	var seq []scud.Stub

	return http.Join(
		ø.GET("%s/scud", host),
		ƒ.Code(http.StatusCodeOK),
		ƒ.ServedJSON(),
		ƒ.Recv(&seq),
	).Then(
		ç.Value(&seq).Is(&[]scud.Stub{
			{ID: "a"}, {ID: "b"}, {ID: "c"},
		}),
	)
}
