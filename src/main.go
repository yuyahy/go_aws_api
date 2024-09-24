package main

import (
	"context"
	"go-api-lambda/src/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "POST":
		return service.CreateWorkout(ctx, request)
	case "GET":
		return service.ReadWorkout(ctx, request)
	case "PUT":
		return service.UpdateWorkout(ctx, request)
	case "DELETE":
		return service.DeleteWorkout(ctx, request)
	default:
		// requestBody, _ := json.Marshal(request)
		return events.APIGatewayProxyResponse{StatusCode: 405, Body: "Method Not Allowed"}, nil
	}
}

func main() {
	lambda.Start(handler)
}
