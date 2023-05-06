package pwa_movies

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsamplify"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewPwaMoviesStack(scope constructs.Construct, id string, props *awscdk.StackProps) {
	stack := awscdk.NewStack(scope, &id, props)
	app := awsamplify.NewCfnApp(stack, jsii.String("Web"), &awsamplify.CfnAppProps{
		Name:        jsii.String("pwa-movies"),
		Repository:  jsii.String("https://github.com/akin-tekeoglu/pwa-movies"),
		Platform:    jsii.String("WEB_COMPUTE"),
		AccessToken: jsii.String(os.Getenv("GITHUB_ACCESS_TOKEN")),
	})
	awsamplify.NewCfnBranch(stack, jsii.String("WebProd"), &awsamplify.CfnBranchProps{
		AppId:           app.AttrAppId(),
		EnableAutoBuild: true,
		BranchName:      jsii.String("master"),
		Stage:           jsii.String("PRODUCTION"),
	})
}
