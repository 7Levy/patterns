package creational

import "fmt"

//一个抽象工厂生产mac和windows工厂，定义好工厂接口，生产系列产品

//键盘接口
type Keyboard interface {
	input()
}

//鼠标接口
type Mouse interface {
	click()
}

//Mac产品族
type MacKeyboard struct {
}

type MacMouse struct {
}

//windows产品族
type WindowsKeyboard struct {
}

type WindowsMouse struct {
}

func (macKeyboard *MacKeyboard) input() {
	fmt.Println("Create Mac keyboard")
}

func (macMouse *MacMouse) click() {
	fmt.Println("Create Mac mouse")
}

func (windowsKeyboard *WindowsKeyboard) input() {
	fmt.Println("Create Windows keyboard")
}

func (windowsMouse *WindowsMouse) click() {
	fmt.Println("Create Windows mouse")
}

//定义工厂接口
type HardWare interface {
	CreateKeyboard() Keyboard
	CreateMouse() Mouse
}

type FactoryProducer struct {
}

//Mac工厂，生产Mac产品族
type MacFactory struct {
}

//生产产品等级键盘
func (mf *MacFactory) CreateKeyboard() Keyboard {
	return new(MacKeyboard)
}

//生产产品等级鼠标
func (mf *MacFactory) CreateMouse() Mouse {
	return new(MacMouse)
}

//Windows工厂，生产windows产品族
type WindowsFactory struct {
}

func (wf *WindowsFactory) CreateKeyboard() Keyboard {
	return new(WindowsKeyboard)
}

func (wf *WindowsFactory) CreateMouse() Mouse {
	return new(WindowsMouse)
}

func (fp *FactoryProducer) CreateFactory(factoryName string) HardWare {
	switch factoryName {
	case "mac":
		return &MacFactory{}
	case "windows":
		return &WindowsFactory{}
	default:
		return nil
	}
}
func NewAbstract() {
	fp := &FactoryProducer{}
	macFactory := fp.CreateFactory("mac")
	macKeyboard := macFactory.CreateKeyboard()
	macMouse := macFactory.CreateMouse()
	macKeyboard.input()
	macMouse.click()
}
