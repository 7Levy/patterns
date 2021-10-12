package creational

//饿汉单例模式
// 构建一个结构体，用来实例化单例
type HungerSingleton struct {
	name string
}

// 声明一个私有变量，作为单例
var instance2 *HungerSingleton

// init函数将在包初始化时执行，实例化单例
func init() {
	instance2 = new(HungerSingleton)
	instance2.name = "初始化单例模式"
}

func NewHungerSingleton() *HungerSingleton {
	return instance2
}
