package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fogfish/blueprint-serverless-golang/http/curl"
	"github.com/fogfish/gurl/awsapi"
	"github.com/fogfish/gurl/v2/http"
)

// go run main.go https://XXXXXXXXXX.execute-api.eu-west-1.amazonaws.com
func main() {
	fmt.Println(os.Args)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	curl := curl.NewPetShop(
		http.New(awsapi.WithSignatureV4(cfg)),
		os.Args[1],
	)

	pets, err := curl.List(context.Background())
	if err != nil {
		panic(err)
	}
	output(pets)

	if pets.Next != nil {
		pets, err = curl.Continue(context.Background(), *pets.Next)
		if err != nil {
			panic(err)
		}
		output(pets)
	}

	if len(pets.Pets) > 0 {
		pet, err := curl.Pet(context.Background(), pets.Pets[0].Url)
		if err != nil {
			panic(err)
		}
		output(pet)
	}
}

func output(x any) {
	b, _ := json.MarshalIndent(x, "|", "  ")
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n\n"))
}
