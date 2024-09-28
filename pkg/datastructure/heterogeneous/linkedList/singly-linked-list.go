package linkedlist

import (
	"fmt"
)

// Node struct
type Node struct {
	nextNode *Node
	property rune
}

// Create list method
func CreateLinkedList() *Node {
	headNode := &Node{nil,'a'}
	currNode := headNode
	
	for i:='b'; i<= 'z'; i++ {
		node := &Node{nil,i}
		currNode.nextNode = node
		currNode = node
	}
	return headNode
}
// ReverseLinkedList method
func ReverseLinkedList(nodeList *Node) *Node {
	currNode := nodeList
	topNode := new(Node)
	for {
		if currNode == nil {
			break
		}
		tempNode := currNode.nextNode
		currNode.nextNode = topNode
		topNode = currNode
		currNode = tempNode
	}
	return topNode
}
func StringifyList(list *Node) {
	for node := list;node !=nil;node=node.nextNode{
		fmt.Print(node.property," ")
	}
	fmt.Print("\n")
}


func SinglFunc() {
	linked := CreateLinkedList()
	StringifyList(linked)
	StringifyList(ReverseLinkedList(linked))
}