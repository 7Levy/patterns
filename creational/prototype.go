package creational

import "fmt"

//克隆方法接口
type ShapeClone interface {
	Clone() ShapeClone
}

//模型管理类
type ShapeManager struct {
	ShapeList map[string]ShapeClone
}

//新建模型管理类
func NewShapeManager() *ShapeManager {
	return &ShapeManager{
		ShapeList: make(map[string]ShapeClone),
	}
}

//添加原型
func (s *ShapeManager) Set(name string, proto ShapeClone) {
	s.ShapeList[name] = proto
}

//获取原型
func (s *ShapeManager) Get(name string) ShapeClone {
	return s.ShapeList[name]
}

//Circle 圆形类，实现shapeclone接口
type Circle struct {
	Name string
}

//NewCircle 新建一个CIrcle类
func NewCircle(name string) *Circle {
	return &Circle{Name: name}
}

//GetName 获取circle的名字
func (c *Circle) GetName() string {
	return c.Name
}

//Clone 返回circle类的复制
func (c *Circle) Clone() ShapeClone {
	circleClone := c
	return circleClone
}

//Draw circle的方法
func (c *Circle) Draw() {
	fmt.Println("Circle Draw().")
}
