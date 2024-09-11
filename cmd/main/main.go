package main

import (
	// "fmt"
	"fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/nonlinear"
)

func main() {
	tree := new(nonlinear.BinarySearchTree)
	tree.InsertElement(8,8)
	tree.InsertElement(3,3)
	tree.InsertElement(10,10)
	tree.InsertElement(1,1)
	tree.InsertElement(6,6)
	tree.String()
	fmt.Println(tree.RemoveNode(3))
	tree.String()
}
