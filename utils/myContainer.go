package utils

import (
	"sync"
)

type Container struct {
	mux sync.Mutex
	m map[string]interface{}
}

func (c *Container) Add(name string,object interface{}) {
	if c.m == nil {
		c.m = make(map[string]interface{})
	}
	c.mux.Lock()
	c.m[name] = object
	c.mux.Unlock()
}
func (c *Container) Remove(name string) {
	c.mux.Lock()
	delete(c.m,name)
	c.mux.Unlock()
}
func (c *Container) Get(name string) (object interface{},bool bool){
	c.mux.Lock()
	object,ok := c.m[name]
	c.mux.Unlock()
	return object, ok
}

