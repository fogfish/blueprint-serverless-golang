package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"github.com/fogfish/blueprint-serverless-golang/internal/awspetshop"
)

func vsn(app awscdk.App) string {
	switch val := app.Node().TryGetContext(jsii.String("vsn")).(type) {
	case string:
		return val
	default:
		return "latest"
	}
}

func main() {
	//
	// Global config
	//
	app := awscdk.NewApp(nil)
	config := &awscdk.StackProps{
		Env: &awscdk.Environment{
			Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
			Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
		},
	}

	//
	// Stack
	//
	vsn := vsn(app)
	awspetshop.NewPetShop(app, jsii.String(fmt.Sprintf("blueprint-golang-%s", vsn)),
		&awspetshop.PetShopProps{
			StackProps: config,
		},
	)

	app.Synth(nil)
}
