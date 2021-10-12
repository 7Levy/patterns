package creational

import "sync"

//单例模式
//单例类型
type singleton map[string]string

var (
	once     sync.Once
	instance singleton
)

func NewSingleton() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
