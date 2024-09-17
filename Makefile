build: ## Build Go file
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap ./src/main.go
archive:
	zip myFunction.zip bootstrap
deploy:
	aws lambda update-function-code --function-name ${FUNC_NAME} --zip-file fileb://myFunction.zip