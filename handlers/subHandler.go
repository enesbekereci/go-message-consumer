package handlers

import "fmt"

func SubHandler(input SubHandlerType) {
	fmt.Println(input.number1 - input.number2)
}

type SubHandlerType struct {
	number1 int `json:"number1"`
	number2 int `json:"number2"`
}
