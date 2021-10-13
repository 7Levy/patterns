package creational

import (
	"errors"
	"fmt"
)

//操作符产品
type operation struct {
}

type cal interface {
	cal(int, int) (float64, error)
}

//加法产品
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

//一个操作符工厂
type operationFactory struct{}

//根据特性生产不同的运算产品
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
