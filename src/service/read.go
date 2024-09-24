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

func ReadWorkout(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// func ReadWorkout(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// GET (Read) の処理
	method := request.HTTPMethod
	pathparam := request.PathParameters["userid"]

	// DB接続
	sess, err := session.NewSession()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	db := dynamodb.New(sess)

	// 検索条件を用意
	getParam := &dynamodb.GetItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				N: aws.String(pathparam),
			},
		},
	}

	// 検索
	result, err := db.GetItem(getParam)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 404,
		}, err
	}

	// 結果を構造体にパース
	item := entity.Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// httpレスポンスを作成
	res := entity.Response{
		RequestMethod: method,
		Result:        item,
	}
	jsonBytes, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{Body: string(jsonBytes), StatusCode: 200}, nil
}
