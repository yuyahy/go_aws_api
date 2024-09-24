package service

import (
	"context"
	"encoding/json"
	"go-api-lambda/src/entity"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DeleteWorkout(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// DELETE (Delete) の処理
	method := request.HTTPMethod
	pathparam := request.PathParameters["userid"]

	// まずDBと接続するセッションを作る
	sess, err := session.NewSession()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	db := dynamodb.New(sess)

	// deleteするデータを指定
	deleteParam := &dynamodb.DeleteItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				N: aws.String(pathparam),
			},
		},
	}

	// delete実行
	_, err = db.DeleteItem(deleteParam)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// httpレスポンス作成
	res := entity.Response{
		RequestMethod: method,
	}
	jsonBytes, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{Body: string(jsonBytes), StatusCode: 200}, nil
}
