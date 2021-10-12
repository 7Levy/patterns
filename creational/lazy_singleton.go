package creational

import "sync"

//单例模式

//1.懒汉式
type LazySingleton map[string]string

var (
	once     sync.Once
	instance LazySingleton
)

func NewLazySingleton() LazySingleton {
	//sync.Once控制只执行一次创建实例
	once.Do(func() {
		instance = make(LazySingleton)
	})
	return instance
}
