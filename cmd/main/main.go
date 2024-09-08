package main

import (
	// "fmt"

	"fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/linear"
)

func main() {
	set := new(linear.Set)
	set2 := new(linear.Set)
	set.New()
	set2.New()
	set.AddElement(1)
	set.AddElement(2)
	set.AddElement(3)
	set.AddElement(4)
	set2.AddElement(2)
	set2.AddElement(3)
	set2.AddElement(5)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
	fmt.Println(set.ContainsElement(4))
	fmt.Println(set.Intersect(set2))
	fmt.Println(set.Union(set2))

}
