package linear

import (
	"fmt"
	"container/list"
)

func MyList(){
	var inList list.List
	inList.PushBack(11)
	inList.PushBack(21)
	inList.PushBack(31)
	for element := inList.Front();element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

// LinkedList

// the Node class

type Node struct {
	Property int
	NextNode *Node
}
// LinkedList class
type LinkedList struct {
	HeadNode *Node
}
// AddToHead method
func (linkedList *LinkedList) AddToHead(property int){
	node := new(Node)
	node.Property = property
	node.NextNode = linkedList.HeadNode
	linkedList.HeadNode = node
}
