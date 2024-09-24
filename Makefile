build: ## Build Go file
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap ./src/main.go

archive: ## Create archive for AWS Lambda
	zip myFunction.zip bootstrap

deploy: ## Deploy archive to AWS Lambda
	aws lambda update-function-code --function-name ${FUNC_NAME} --zip-file fileb://myFunction.zip

update: ## Run update process for AWS Lambda by one command
	$(MAKE) build
	$(MAKE) archive
	$(MAKE) deploy FUNC_NAME=myTestFunction
