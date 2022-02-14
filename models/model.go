package models

type Input struct {
	Num1 int `json:"number1"`
	Num2 int `json:"number2"`
}

type Response struct {
	Result int `json:"result"`
}
