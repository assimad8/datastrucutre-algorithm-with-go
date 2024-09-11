package datastructure

import (
	"fmt"
	"container/ring"
)

func LinkedLIstFun() {
	integers := []int{1,2,3,4,5}
	circular_list := new(ring.Ring)
	circular_list = ring.New(len(integers))
	for i:=0;i<circular_list.Len();i++{
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}
	circular_list.Do(func(element interface{}) {
		fmt.Print(element,",")
	})
	fmt.Println()
	// reverse of the circular list
	for i := 0; i < circular_list.Len(); i++ {
		fmt.Print(circular_list.Value,",")
		circular_list = circular_list.Prev()
		}
	fmt.Println()
	// move two elements forward in the circular list
	circular_list = circular_list.Move(2)
	circular_list.Do(func(element interface{}) {
	fmt.Print(element,",")
	})
	fmt.Println()

}

