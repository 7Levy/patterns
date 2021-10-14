package creational

import "fmt"

type animal struct {
}

//动物行为接口
type Behavior interface {
	say()
}

type dog struct {
	animal
}

type cat struct {
	animal
}

type pig struct {
	animal
}

func (d *dog) say() {
	fmt.Println("I am a dog.")
}

func (c *cat) say() {
	fmt.Println("I am a cat.")
}

func (p *pig) say() {
	fmt.Println("I am a pig.")
}

//工厂接口
type Factory interface {
	Create() Behavior
}

//生产狗的工厂
type dogFactory struct {
}

func (df *dogFactory) Create() Behavior {
	return new(dog)
}

//生产猫的工厂
type catFactory struct {
}

func (cf *catFactory) Create() Behavior {
	return new(cat)
}

//生产猪的工厂
type pigFactory struct {
}

func (pf *pigFactory) Create() Behavior {
	return new(pig)
}

func New() {
	df := &dogFactory{}
	cf := &catFactory{}
	pg := &pigFactory{}
	dog := df.Create()
	cat := cf.Create()
	pig := pg.Create()
	dog.say()
	cat.say()
	pig.say()
}
