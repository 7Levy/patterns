package structural

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
	X      int
	Y      int
	Radius int
	Color  string
}

func NewCircle(color string) *Circle {
	return &Circle{Color: color}
}

//SetX 设置圆形类的x轴
func (c *Circle) SetX(x int) {
	c.X = x
}

//SetY 设置圆形类的y轴
func (c *Circle) SetY(y int) {
	c.Y = y
}

//SetRadius 设置圆形类的半径
func (c *Circle) SetRadius(radius int) {
	c.Radius = radius
}

//Draw 圆形类的Draw方法，实现Shape接口
func (c *Circle) Draw() {
	fmt.Printf("Circle Draw() [Color: %s, x: %d, y: %d, radius: %d] \n",
		c.Color,
		c.X,
		c.Y,
		c.Radius)
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]Shape),
	}
}

//GetCircle 获取一个圆形实例
//color: 需要获取实例的颜色
func (sf *ShapeFactory) GetCircle(color string) Shape {
	circle := sf.circleMap[color]
	if circle == nil {
		circle = NewCircle(color)
		sf.circleMap[color] = circle
		fmt.Println("Creating circleof color: ", color)
	}
	return circle
}
