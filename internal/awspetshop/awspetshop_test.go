package awspetshop_test

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
	"github.com/fogfish/blueprint-serverless-golang/internal/awspetshop"
)

func TestPetShop(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := awspetshop.NewPetShop(app, jsii.String("Test"),
		&awspetshop.PetShopProps{},
	)

	require := map[*string]*float64{
		jsii.String("AWS::ApiGatewayV2::Api"):   jsii.Number(1),
		jsii.String("AWS::ApiGatewayV2::Stage"): jsii.Number(2),
		jsii.String("AWS::ApiGatewayV2::Route"): jsii.Number(2),
		jsii.String("AWS::IAM::Role"):           jsii.Number(4),
		jsii.String("AWS::Lambda::Function"):    jsii.Number(3),
		jsii.String("Custom::LogRetention"):     jsii.Number(2),
	}

	template := assertions.Template_FromStack(stack.Stack, nil)
	for key, val := range require {
		template.ResourceCountIs(key, val)
	}
}
