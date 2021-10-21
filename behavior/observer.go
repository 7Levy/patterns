package behavior

import (
	"container/list"
	"fmt"
)

//被观察对象
type Subject struct {
	state    int
	observer *list.List
}

func NewSubject() *Subject {
	return &Subject{
		state:    0,
		observer: list.New(),
	}
}

func (s *Subject) GetState() int {
	return s.state
}

func (s *Subject) SetState(st int) {
	s.state = st
	s.NotifyAllObserver()
}

//Attach 将观察者绑定到自己
func (s *Subject) Attach(ob Observer) {
	s.observer.PushBack(ob)
}

//NotifyAllObserver 通知所有观察值
func (s *Subject) NotifyAllObserver() {
	for i := s.observer.Front(); i != nil; i = i.Next() {
		i.Value.(Observer).Update()
	}
}

//观察者而
type Observer interface {
	Update()
}

type BinaryObserver struct {
	subject *Subject
}

//NewBinaryObserver 实例化二进制观察者类
func NewBinaryObserver() *BinaryObserver {
	return &BinaryObserver{}
}

//Attach 将二进制观察者绑定到观察目标
func (bo *BinaryObserver) Attach(sub *Subject) {
	bo.subject = sub
	bo.subject.Attach(bo)
}

//Update 二进制观察者接受被观察目标的更新信息
func (bo *BinaryObserver) Update() {
	fmt.Printf("Binary String: %b \n", bo.subject.GetState())
}
