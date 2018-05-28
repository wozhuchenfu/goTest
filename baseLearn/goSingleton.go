package baseLearn

import (
	"sync"
	"fmt"
)

type singleton map[string]string

var (
	once sync.Once
	instance singleton
)

func new() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}
//单例模式主要考察知识点是sync.once的幂等。
func SingletonTest()  {
	var m = new()
	m["name"] = "张三"
	fmt.Println(m)
	var m2 = new()
	m2["name2"] = "李四"
	fmt.Println(m,m2) //m和m2都是同一个map
}