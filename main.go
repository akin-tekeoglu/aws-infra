package main

import (
	"aws-infra/online_shop"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	env := &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
	online_shop.NewOnlineShopLambda(app, "OnlineShopLambda",
		&awscdk.StackProps{
			Env: env,
		},
	)

	app.Synth(nil)
}
