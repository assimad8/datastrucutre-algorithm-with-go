package linear

import (
	"strconv"
)
// Element class
type Element struct {
	value int
}

// String method on Element class
func (element *Element) String() string{
	return strconv.Itoa(element.value)
}
// Set method on Element class
func (element *Element) Set(v int){
	element.value = v
}
// Stack class
type Stack struct{
	elements []*Element
	count int
} 
// New method on Stack class 
func (stack *Stack) New(){
	stack.elements = make([]*Element,0)
}
// Push method on Stack class
func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements[:stack.count],element)
	stack.count += 1
}
// Pop method on Stack class
func(stack *Stack) Pop() *Element {
	if stack.count == 0 {
		return nil
	}
	lenght := len(stack.elements)
	element := stack.elements[lenght-1]
	stack.count -= 1
	if lenght>1 {
		stack.elements = stack.elements[:lenght-1]
	}else{
		stack.elements = stack.elements[0:]
	}
	return element
}