package main

import (
	// "fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/linear"
)

func main() {
	linkedList := new(linear.LinkedList)
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(4)
	linkedList.AddToHead(5)
	linkedList.IterateList()
	linkedList.AddAfter(3,2)
	linkedList.IterateList()

}
