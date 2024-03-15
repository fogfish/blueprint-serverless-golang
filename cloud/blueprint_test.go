package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestBlueprint(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := awscdk.NewStack(app, jsii.String("Test"), nil)

	NewBlueprint(stack)

	require := map[*string]*float64{
		jsii.String("AWS::ApiGatewayV2::Api"):   jsii.Number(1),
		jsii.String("AWS::ApiGatewayV2::Stage"): jsii.Number(2),
		jsii.String("AWS::ApiGatewayV2::Route"): jsii.Number(1),
		jsii.String("AWS::IAM::Role"):           jsii.Number(2),
		jsii.String("AWS::Lambda::Function"):    jsii.Number(2),
		jsii.String("Custom::LogRetention"):     jsii.Number(1),
	}

	template := assertions.Template_FromStack(stack, nil)
	for key, val := range require {
		template.ResourceCountIs(key, val)
	}
}
