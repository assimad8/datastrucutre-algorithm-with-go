package main

import (
	// "fmt"

	"fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/linear"
)

func main() {
	stack := new(linear.Stack)
	stack.New()
	e1 := new(linear.Element)
	e1.Set(3)
	e2 := new(linear.Element)
	e2.Set(5)
	e3 := new(linear.Element)
	e3.Set(7)
	e4 := new(linear.Element)
	e4.Set(9)
	stack.Push(e1)
	stack.Push(e2)
	stack.Push(e3)
	stack.Push(e4)
	fmt.Println(stack.Pop(),stack.Pop(),stack.Pop(),stack.Pop())
}
