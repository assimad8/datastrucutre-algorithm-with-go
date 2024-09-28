package linkedlist

import (
	"container/list"
	"fmt"
)

func DoublyFunc() {
	linkedList := new(list.List)
	element := linkedList.PushBack(14)
	frontElement := linkedList.PushFront(1)
	linkedList.InsertBefore(6,element)
	linkedList.InsertAfter(5,frontElement)
	linkedList.InsertAfter(5,element)

	for currElement := linkedList.Front();currElement !=nil;currElement=currElement.Next() {
		fmt.Println(currElement.Value)
	}
}