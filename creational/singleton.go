package creational

import "sync"

//单例模式

//1.懒汉式
type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

func NewLazySingleton() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	return instance
}

//2.饿汉式
