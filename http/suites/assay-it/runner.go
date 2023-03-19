package main

import (
	"os"
	"github.com/fogfish/gurl/v2/http"
	suites "github.com/fogfish/blueprint-serverless-golang/http/suites"
)

func main() {
	http.WriteOnce(os.Stdout, http.New(http.WithMemento(), http.WithDefaultHost(os.Args[1])), suites.TestPetShopList, suites.TestPetShopListWithCursor, suites.TestPetShopLookup, suites.TestPetShopCreate, suites.TestPetShopCreateUnauthorized)
}
