package main

import (
	"fmt"
	"github.com/assimad8/golang-algo/pkg/datastructure/linear"
)

func main() {
	linkedList := new(linear.LinkedList)
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println(linkedList.HeadNode.Property)
}
