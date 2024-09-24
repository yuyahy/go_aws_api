package entity

// Response Lambdaが返答するデータ
type Response struct {
	RequestMethod string `json:RequestMethod`
	Result        Item   `json:Result`
}
