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
// IterateList method
func (linkedList *LinkedList) IterateList(){
	for node := linkedList.HeadNode;node !=nil;node=node.NextNode{
		fmt.Print(node.Property)
	}
	fmt.Print("\n")
}
// LastNode method
func (linkedList *LinkedList) LastNode() *Node {
	lastNode := new(Node)
	for node := linkedList.HeadNode;node!=nil;node=node.NextNode {
		lastNode = node
	}
	return lastNode
}
// AddToEnd method
func (linkedList *LinkedList) AddToEnd(property int) {
	lastNode := new(Node)
	lastNode.Property = property
	oldLast := linkedList.LastNode()
	oldLast.NextNode = lastNode
}
// NodeWithValue method
func (linkedList *LinkedList) NodeWithValue(value int) *Node {
	for node := linkedList.HeadNode;node!=nil;node=node.NextNode{
		if node.Property == value{
			return node
		}
	}
	return nil
}
// AddAfter mothod
func (linkedList *LinkedList) AddAfter(nodeProperty,property int) {
	node := new(Node)
	node.Property = property
	node.NextNode = nil
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil{
		node.NextNode = nodeWith.NextNode
		nodeWith.NextNode = node
	}
}