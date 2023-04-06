package engine

import (
	"context"
	"github.com/google/wire"
)

type Engine interface {
	Container

	Mount(process ...Process)

	Run() error

	Shutdown() error
}

type engine struct {
	ctx         context.Context
	cancel      context.CancelFunc
	container   Container
	processes   []Process
	rootProcess Process
}

var WireEngineSet = wire.NewSet(NewEngine, WireContainerSet, WireProcessSet)

//var WireEngineSet = wire.NewSet(wire.InterfaceValue(new(Engine), wire.NewSet(NewEngine, WireContainerSet, WireProcessSet)))

func NewEngine(
	ctx context.Context,
	container2 Container,
	process Process,
) Engine {
	ctx, cancel := context.WithCancel(ctx)
	return &engine{
		ctx:         ctx,
		cancel:      cancel,
		container:   container2,
		rootProcess: process,
	}
}

func (e engine) Get(key string) (any, bool) {
	//TODO implement me
	panic("implement me")
}

func (e engine) Put(key string, val any) {
	//TODO implement me
	panic("implement me")
}

func (e engine) Add(key string, val any) bool {
	//TODO implement me
	panic("implement me")
}

func (e engine) Delete(key string) {
	//TODO implement me
	panic("implement me")
}

func (e engine) Exists(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (e engine) Clean() {
	//TODO implement me
	panic("implement me")
}

func (e engine) Mount(process ...Process) {
	//TODO implement me
	panic("implement me")
}

func (e engine) Run() error {
	//TODO implement me
	panic("implement me")
}

func (e engine) Shutdown() error {
	//TODO implement me
	panic("implement me")
}
