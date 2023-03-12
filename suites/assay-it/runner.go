package main

import (
	"os"

	"github.com/fogfish/blueprint-serverless-golang/suites"
	"github.com/fogfish/gurl/v2/http"
)

func main() {
	http.WriteOnce(os.Stdout,
		http.New(
			http.WithMemento(),
			http.WithDefaultHost("http://127.1:8080"),
		),
		suites.TestPetShopList,
		suites.TestPetShopListCursor,
	)
}
