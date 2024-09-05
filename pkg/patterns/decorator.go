package patterns

import "fmt"

// IProcess interface
type Iprocess interface{
	process()
}

type ProcessClass struct{}
func (p *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}
type ProcessDecorator struct{
	processInstance *ProcessClass
}
func (decorator *ProcessDecorator) process(){
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	}else{
		fmt.Print("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func DecoratorFunc() {
	process := new(ProcessClass)
	decorator := new(ProcessDecorator)
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}