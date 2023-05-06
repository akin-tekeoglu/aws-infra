clean: 
	@rm -rf cdk.out; \
	rm -rf lambda_bootstrap_functions/go/out

build: 
	@GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o lambda_bootstrap_functions/go/out/main lambda_bootstrap_functions/go/main.go; \
	cdk synth

deploy: clean build
	@cdk deploy --all