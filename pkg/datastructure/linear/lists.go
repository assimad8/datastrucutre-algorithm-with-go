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

// Doubly Linked List

// Node Class
type DNode struct {
	property int
	nextNode *DNode
	previousNode *DNode
}
type DLinkedList struct {
	HeadNode *DNode
}
// NodeBetweenValues mothods of LinedList
func (linkedList *DLinkedList) NodeBetweenValues(firstProperty,secondProperty int) *DNode {
	for node := linkedList.HeadNode; node != nil; node = node.nextNode{
		if node.previousNode!=nil && node.nextNode != nil{
			if node.previousNode.property==firstProperty&&node.nextNode.property==secondProperty {
				return node
			}
		}
	}
	return nil
}
// NodeWithValue method 
func (linkedList *DLinkedList) NodeWithValue(value int) *DNode {
	for node := linkedList.HeadNode;node!=nil;node=node.nextNode{
		if node.property==value{
			return node
		}
	}
	return nil
}
// AddToHead method
func (linkedList *DLinkedList) AddToHead(property int) {
	node := &DNode{property,nil,nil}
	if linkedList.HeadNode!=nil {
		node.nextNode = linkedList.HeadNode
		linkedList.HeadNode.previousNode=node
	}
	linkedList.HeadNode = node
}
// AddAfter method
func (linkedList *DLinkedList) AddAfter(nodeProperty,property int) {
	node := &DNode{property,nil,nil}
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}

}
// LastNode method
func (linkedList *DLinkedList) LastNode() *DNode {

	for node := linkedList.HeadNode;node!=nil;node=node.nextNode {
		return node
	}
	return nil
}
// AddToEnd method
func (linkedList *DLinkedList) AddToEnd(property int) {
	node := new(DNode)
	node.property = property
	lastNode := linkedList.LastNode()
	if lastNode!=nil{
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
} 
// IterateList method
func (linkedList *DLinkedList) IterateList(){
	for node := linkedList.HeadNode;node !=nil;node=node.nextNode{
		fmt.Print(node.property)
	}
	fmt.Print("\n")
}