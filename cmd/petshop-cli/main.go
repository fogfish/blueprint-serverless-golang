package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fogfish/blueprint-serverless-golang/pkg/api"
	"github.com/fogfish/gurl/awsapi"
	"github.com/fogfish/gurl/v2/http"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("petshop-cli https://XXXXXXXXXX.execute-api.eu-west-1.amazonaws.com [pet.json]\n")
	}

	host := os.Args[1]
	if len(os.Args) == 3 {
		create(host, os.Args[2])
	} else {
		list(host)
	}
}

func list(host string) {
	curl := api.NewPetShop(http.New(), host)

	pets, err := curl.List(context.Background())
	if err != nil {
		panic(err)
	}
	output(pets)

	for pets.Next != nil && len(*pets.Next) != 0 {
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

func create(host string, file string) {
	var pet api.Pet

	fd, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	if err := json.NewDecoder(fd).Decode(&pet); err != nil {
		panic(err)
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic(err)
	}

	curl := api.NewPetShop(
		http.New(awsapi.WithSignatureV4(cfg)),
		host,
	)

	if err := curl.Create(context.Background(), &pet); err != nil {
		panic(err)
	}

	fmt.Printf("==> pet created\n")
	output(pet)
}

func output(x any) {
	b, _ := json.MarshalIndent(x, "|", "  ")
	os.Stdout.Write(b)
	os.Stdout.Write([]byte("\n\n"))
}
