package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/fogfish/scud"
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
	stack := awscdk.NewStack(app,
		jsii.String(fmt.Sprintf("blueprint-golang-%s", vsn)),
		config,
	)

	NewBlueprint(stack)

	app.Synth(nil)
}

// NewBlueprint create example REST api
func NewBlueprint(scope constructs.Construct) {
	gateway := scud.NewGateway(scope, jsii.String("Gateway"),
		&scud.GatewayProps{
			HttpApiProps: &awsapigatewayv2.HttpApiProps{},
		},
	)
	gateway.WithAuthorizerIAM()

	handler := scud.NewFunctionGo(scope, jsii.String("Handler"),
		&scud.FunctionGoProps{
			SourceCodePackage: "github.com/fogfish/blueprint-serverless-golang",
			SourceCodeLambda:  "cmd/lambda/petshop",
		},
	)
	gateway.AddResource("/petshop", handler)
}
