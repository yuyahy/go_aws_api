package service

import (
	"context"
	"encoding/json"
	"go-api-lambda/src/entity"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func UpdateWorkout(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// PUT (Update) の処理
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

	// リクエストボディののJSONからItem構造体を作成
	reqBody := request.Body
	resBodyJSONBytes := ([]byte)(reqBody)
	item := entity.Item{}
	if err := json.Unmarshal(resBodyJSONBytes, &item); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	// updateするデータを作る
	// 更新方法を指定している(今回は既存の値を更新するが、既存リストに要素を追加する事もできる)
	update := expression.UpdateBuilder{}
	if address := item.Address; address != "" {
		update = update.Set(expression.Name("address"), expression.Value(address))
	}
	if email := item.Email; email != "" {
		update = update.Set(expression.Name("email"), expression.Value(email))
	}
	if gender := item.Gender; gender != "" {
		update = update.Set(expression.Name("gender"), expression.Value(gender))
	}
	if name := item.Name; name != "" {
		update = update.Set(expression.Name("name"), expression.Value(name))
	}
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 500,
		}, err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("user"),
		Key: map[string]*dynamodb.AttributeValue{
			"userid": {
				N: aws.String(pathparam),
			},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	// update実行
	_, err = db.UpdateItem(input)
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

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(jsonBytes)}, nil
}
