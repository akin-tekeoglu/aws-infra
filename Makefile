clean: 
	@rm -rf cdk.out; \
	rm -rf lambda-bootstrap-functions/go/out

build: 
	@GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o lambda-bootstrap-functions/go/out/main lambda-bootstrap-functions/go/main.go; \
	cdk synth

deploy: clean build
	@cdk deploy