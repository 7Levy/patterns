package creational

import "fmt"

type Worker interface {
	Cal(int, int) (float64, error)
}

type WorkerCreator interface {
	Create() Worker
}

//数学运算产品族
type ProgrammerOperation struct {
}

//数学运算产品族-加法
func (p *ProgrammerOperation) Cal(num1 int, num2 int) (float64, error) {
	fmt.Println("result", num1+num2)
	return float64(num1 + num2), nil
}

//数学运算产品族-减法
func (p *ProgrammerOperation) Minus(num1 int, num2 int) (float64, error) {
	fmt.Println("result", num1+num2)
	return float64(num1 - num2), nil
}

//抽象工厂
type ProgrammerCreator struct {
}

//创建数学运算产品的工厂
func (pc *ProgrammerCreator) CreateOperation() Worker {
	s := new(ProgrammerOperation)
	return s
}
