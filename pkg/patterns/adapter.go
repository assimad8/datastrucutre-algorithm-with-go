package patterns

import (
	"fmt"
)

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process")
	adapter.adaptee.convert()
}
type Adaptee struct {
	adapterType int
}

func (adapter Adaptee) convert() {
	fmt.Println("Adaptee convert")
}

func AdapterFun() {
	var processor IProcess = Adapter{}
	processor.process()
}