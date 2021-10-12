package creational

import (
	"errors"
	"fmt"
)

type operation struct {
}

type cal interface {
	cal(int, int) (float64, error)
}

type operationAdd struct {
	operation
}

func (o *operationAdd) cal(num1, num2 int) (float64, error) {
	return float64(num1 + num2), nil
}

type operationMinus struct {
	operation
}

func (o *operationMinus) cal(num1 int, num2 int) (float64, error) {
	return float64(num1 - num2), nil
}

type operationMulti struct {
	operation
}

func (o *operationMulti) cal(num1 int, num2 int) (float64, error) {
	return float64(num1 * num2), nil
}

type operationDiv struct {
	operation
}

func (o *operationDiv) cal(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("除法运算中除数不能为0")
	}
	return float64(num1 / num2), nil
}

type operationFactory struct{}

func (o *operationFactory) createOperation(op string) cal {
	switch op {
	case "+":
		return new(operationAdd)
	case "-":
		return new(operationMinus)
	case "*":
		return new(operationMulti)
	case "/":
		return new(operationDiv)
	}
	return nil
}

func NewFactory(data string) {
	o := &operationFactory{}
	cal := o.createOperation(data)
	fmt.Printf("%T\n", cal)
}
