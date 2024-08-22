package awspetshop

import (
	"path/filepath"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
	"github.com/fogfish/blueprint-serverless-golang/internal/core"
	"github.com/fogfish/scud"
)

type PetShopProps struct {
	*awscdk.StackProps
}

type PetShop struct {
	awscdk.Stack
	api     *scud.Gateway
	console *scud.AuthorizerIAM
}

func NewPetShop(app awscdk.App, id *string, props *PetShopProps) *PetShop {
	stack := awscdk.NewStack(app, id, props.StackProps)

	ps := &PetShop{Stack: stack}
	ps.createGateway()
	ps.createReader()
	ps.createWriter()
	return ps
}

func (ps *PetShop) createGateway() {
	ps.api = scud.NewGateway(ps.Stack, jsii.String("Gateway"),
		&scud.GatewayProps{},
	)
	ps.console = ps.api.NewAuthorizerIAM()

	awscdk.NewCfnOutput(ps.Stack, jsii.String("Endpoint"),
		&awscdk.CfnOutputProps{
			Value: ps.api.RestAPI.ApiEndpoint(),
		},
	)
}

func (ps *PetShop) createReader() {
	handler := scud.NewFunctionGo(ps.Stack, jsii.String("Reader"),
		&scud.FunctionGoProps{
			SourceCodeModule: "github.com/fogfish/blueprint-serverless-golang",
			SourceCodeLambda: "cmd/lambda/petshop",
		},
	)
	ps.api.AddResource(filepath.Join("/", core.PETSHOP), handler)
}

func (ps *PetShop) createWriter() {
	// handler := scud.NewContainerGo(ps.Stack, jsii.String("Writer"),
	// 	&scud.ContainerGoProps{
	// 		SourceCodeModule: "github.com/fogfish/blueprint-serverless-golang",
	// 		SourceCodeLambda: "cmd/lambda/console",
	// 	},
	// )
	handler := scud.NewFunctionGo(ps.Stack, jsii.String("Writer"),
		&scud.FunctionGoProps{
			SourceCodeModule: "github.com/fogfish/blueprint-serverless-golang",
			SourceCodeLambda: "cmd/lambda/console",
		},
	)

	principal := awsiam.NewRole(ps.Stack, jsii.String("Principal"),
		&awsiam.RoleProps{
			AssumedBy: awsiam.NewAccountPrincipal(awscdk.Aws_ACCOUNT_ID()),
		},
	)

	ps.console.AddResource(filepath.Join("/", core.CONSOLE, core.PETSHOP), handler, principal)
}
