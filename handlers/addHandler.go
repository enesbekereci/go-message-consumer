package handlers

import (
	"encoding/json"
	"fmt"
)

type AddHandler[T any] struct {
	Handler func(*T)
}

func (m *AddHandler[T]) Set(f func(*T)) {
	m.Handler = f
}

func (m *AddHandler[T]) Parse(message []byte, input *T) error {
	err := json.Unmarshal(message, input)
	return err
}

func (m *AddHandler[T]) Handle(input *T) {
	m.Handler(input)
}

func AddOperation(input AddHandlerType) {
	fmt.Println(input.number1 + input.number2)
}

type AddHandlerType struct {
	number1 int `json:"number1"`
	number2 int `json:"number2"`
}
