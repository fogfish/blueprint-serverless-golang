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
		jsii.String("AWS::ApiGateway::RestApi"):    jsii.Number(1),
		jsii.String("AWS::ApiGateway::Deployment"): jsii.Number(1),
		jsii.String("AWS::ApiGateway::Stage"):      jsii.Number(1),
		jsii.String("AWS::ApiGateway::Method"):     jsii.Number(5),
		jsii.String("AWS::IAM::Role"):              jsii.Number(3),
		jsii.String("AWS::Lambda::Function"):       jsii.Number(2),
		jsii.String("Custom::LogRetention"):        jsii.Number(1),
	}

	template := assertions.Template_FromStack(stack, nil)
	for key, val := range require {
		template.ResourceCountIs(key, val)
	}
}
