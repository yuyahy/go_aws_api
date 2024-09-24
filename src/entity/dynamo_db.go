package entity

// Item DBに入れるデータ
type Item struct {
	UserID  int    `dynamodbav:"userid" json:"userid"`
	Address string `dynamodbav:"address" json:"address"`
	Email   string `dynamodbav:"email" json:"email"`
	Gender  string `dynamodbav:"gender" json:"gender"`
	Name    string `dynamodbav:"name" json:"name"`
}
