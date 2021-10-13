package creational

import "fmt"

type Worker interface {
	Work(task *string)
}

type WorkerCreator interface {
	Create() Worker
}

type Programmer struct {
}

func (p *Programmer) Work(task *string) {
	fmt.Println("Programmer process", *task)
}

type ProgrammerCreator struct {
}

func (pc *ProgrammerCreator) Create() Worker {
	s := new(Programmer)
	return s
}

type Farmer struct {
}

func (f *Farmer) Work(task *string) {
	fmt.Println("Farmer process", *task)
}

type FarmerCreator struct {
}

func (fc *FarmerCreator) Create() Worker {
	s := new(Farmer)
	return s
}
