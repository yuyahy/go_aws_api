package service

import (
	"context"
	"encoding/json"
	"go-api-lambda/src/entity"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func CreateWorkout(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Create (POST) の処理
	method := request.HTTPMethod

	// DBと接続
	session, err := session.NewSession()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	db := dynamodb.New(session)

	// リクエストボディのjsonから、Item構造体(DB用データの構造体)を作成
	reqBody := request.Body
	reqBodyJSONBytes := ([]byte)(reqBody)
	item := entity.Item{}
	if err := json.Unmarshal(reqBodyJSONBytes, &item); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// Item構造体から、inputするデータを用意
	inputAV, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String("user"),
		Item:      inputAV,
	}

	// insert実行
	_, err = db.PutItem(input)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// httpレスポンス作成
	response := entity.Response{
		RequestMethod: method,
	}
	jsonBytes, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}
