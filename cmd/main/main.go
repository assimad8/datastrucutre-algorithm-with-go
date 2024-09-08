package main

import (
	// "fmt"

	"fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/linear"
)

func main() {
	queue := make(linear.Queue,0)
	order1 := new(linear.Order)
	priority1,quantity1,product1,customerName1 := 2,20,"Computer","Greg"
	order1.New(priority1,quantity1,product1,customerName1)
	order2 := new(linear.Order)
	order2.New(1,10,"Monitor","emad lakhbizi")
	queue.Add(order1)
	queue.Add(order2)
	for i:=0;i<len(queue);i++{
		fmt.Println(queue[i])
	}
}
