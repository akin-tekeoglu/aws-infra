package online_shop

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewOnlineShopLambda(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)
	awslambda.NewFunction(stack, jsii.String("Product"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda-bootstrap-functions/go/out"), &awss3assets.AssetOptions{}),
		Handler: jsii.String("main"),
	})
	return stack
}
