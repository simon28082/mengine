package engine

import (
	"github.com/google/wire"
	"sync"
)

type Container interface {
	Get(key string) any

	Put(key string, val any)

	Add(key string, val any) bool

	Delete(key string)

	Exists(key string) bool

	Clean()
}

type container struct {
	box sync.Map
}

var WireContainerSet = wire.NewSet(NewContainer)

//var WireContainerSet = wire.NewSet(NewContainer, wire.Bind(new(Container), new(*container)))

func NewContainer() *container {
	return &container{}
}

func (c *container) Get(key string) any {
	v, ok := c.box.Load(key)
	if ok {
		return v
	}
	return nil
}

func (c *container) Put(key string, val any) {
	c.box.Store(key, val)
}

func (c *container) Add(key string, val any) bool {
	_, ok := c.box.LoadOrStore(key, val)
	return ok
}

func (c *container) Delete(key string) {
	c.box.Delete(key)
}

func (c *container) Exists(key string) bool {
	_, ok := c.box.Load(key)
	return ok
}

func (c *container) Clean() {
	c.box = sync.Map{}
}
