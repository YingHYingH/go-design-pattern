package _1_singleton

import "sync"

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{a: "test"}
		})
	}
	return lazySingleton
}
