package main

import (
	"aws-infra/online_shop"
	pwa_movies "aws-infra/pwa-movies"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	env := &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
	stackProps := &awscdk.StackProps{
		Env: env,
	}
	online_shop.NewOnlineShopStack(app, "OnlineShop", stackProps)
	pwa_movies.NewPwaMoviesStack(app, "PwaMovies", stackProps)
	app.Synth(nil)
}
