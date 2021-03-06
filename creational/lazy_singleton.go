package creational

import "sync"

//懒汉单例模式

type LazySingleton map[string]string

var (
	once     sync.Once
	instance LazySingleton
	mu       sync.Mutex
)

//1.懒汉实现-双重检查且线程安全
func SafeLazySingleton1() LazySingleton {
	//sync.Once控制只执行一次创建实例
	once.Do(func() {
		instance = make(LazySingleton)
	})
	return instance
}

//2.懒汉实现-线程安全
func SafeLazySingleton2() LazySingleton {
	//加锁
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = make(LazySingleton)
	}
	return instance
}

//3.懒汉实现-非线程安全
func UnsafeLazySingleton() LazySingleton {
	if instance == nil {
		instance = make(LazySingleton)
	}
	return instance
}
