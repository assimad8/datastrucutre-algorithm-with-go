package nonlinear

import (
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct{
	key int
	value int
	leftNode *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct{
	rootNode *TreeNode
	lock sync.RWMutex
}
// InsertElement method
func (tree *BinarySearchTree) InsertElement(key,value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	treeNode := &TreeNode{key,value,nil,nil}
	if tree.rootNode==nil{
		tree.rootNode = treeNode
	}else{
		insertTreeNode(tree.rootNode,treeNode)
	}
}

// insertTreeNode function
func insertTreeNode(rootNode,newTreeNode *TreeNode){
	if rootNode.key<newTreeNode.key{
		if rootNode.leftNode==nil{
			rootNode.leftNode = newTreeNode
		}else{
			insertTreeNode(rootNode.leftNode,newTreeNode)
		}
	}else{
		if rootNode.rightNode==nil{
			rootNode.rightNode = newTreeNode
		}else{
			insertTreeNode(rootNode.rightNode,newTreeNode)
		}
	}
}
// InOrderTraverseTree method
func (tree *BinarySearchTree)InOrderTraverseTree(function func(int)){
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode,function)
}
// inOrderTraverseTree  function
func inOrderTraverseTree (treeNode *TreeNode,function func(int)){
	if treeNode!=nil{
		inOrderTraverseTree(treeNode.leftNode,function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode,function)
	}
}
// PreOrderTraverseTree method
func (tree *BinarySearchTree)PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode,function)
}
// preOrderTraverseTree function
func preOrderTraverseTree(treeNode *TreeNode,function func(int)){
	if treeNode!=nil{
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode,function)
		preOrderTraverseTree(treeNode.rightNode,function)

	}
}
// PostOrderTraverseTree method
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.rootNode, function)
   }
//  postOrderTraverseTree method
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
	postOrderTraverseTree(treeNode.leftNode, function)
	postOrderTraverseTree(treeNode.rightNode, function)
	function(treeNode.value)
	}
   }
// MinNode method
func (tree *BinarySearchTree) MinNode()*int{
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	treeNode := tree.rootNode
	if treeNode == nil{
		return (*int) (nil)
	}
	for {
		if treeNode.leftNode==nil{
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}
// MaxNode method
func (tree *BinarySearchTree) MaxNode()*int{
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	treeNode := tree.rootNode
	if treeNode == nil{
		return (*int) (nil)
	}
	for {
		if treeNode.rightNode==nil{
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}
// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool{
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode,key)
}
// searchNode method
func searchNode(treeNode *TreeNode,key int) bool{
	if treeNode==nil{
		return false
	}
	if key <treeNode.key{
		return searchNode(treeNode.leftNode,key)
	}
	if key>treeNode.key{
		return searchNode(treeNode.leftNode,key)
	}
	return true
}
// RemoveNode method
func (tree *BinarySearchTree) RemoveNode(key int) *TreeNode {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return removeNode(tree.rootNode,key)
}

// removeNode function
func removeNode(treeNode *TreeNode,key int) *TreeNode {
	if treeNode==nil{
		return nil
	}
	if key<treeNode.key{
		treeNode.leftNode = removeNode(treeNode.leftNode,key)
		return treeNode
	}
	if key > treeNode.key{
		treeNode.rightNode = removeNode(treeNode.rightNode,key)
		return treeNode
	}
	// key == treeNode.key
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	leftmostrightnode := treeNode.rightNode
	for {
		// find smaller value on the right side
		if leftmostrightnode != nil && leftmostrightnode.leftNode != nil {
			leftmostrightnode = leftmostrightnode.leftNode
		}else{
			break
		}
	}
	treeNode.key,treeNode.value = leftmostrightnode.key,leftmostrightnode.value
	treeNode.rightNode = removeNode(treeNode.rightNode,treeNode.key)
	return treeNode
}
// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("------------------------------")
	stringify(tree.rootNode,0)
	fmt.Println("------------------------------")
}
// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
	format := ""
	for i := 0; i < level; i++ {
	format += " "
	}
	format += "---[ "
	level++
	stringify(treeNode.leftNode, level)
	fmt.Printf(format+"%d\n", treeNode.key)
	stringify(treeNode.rightNode, level)
	}
}

// AVL Tree

// KeyValue type
type KeyValue interface {
	LessThan(KeyValue) bool
	EqualTo(KeyValue) bool
}
// TreeNode class
type VTreeNode struct {
	KeyValue KeyValue
	BalanceValue int
	LinkedNodes [2]*VTreeNode
}
// singleRotation method
func singleRotation(rootNode *VTreeNode,nodeValue int) *VTreeNode {
	saveNode := rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]
	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}
// opposite method
func opposite(nodeValue int) int {
	return 1 - nodeValue
}
// double rotation
func doubleRotation(rootNode *VTreeNode,nodeValue int) *VTreeNode {
	saveNode := rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue]
	rootNode.LinkedNodes[opposite(nodeValue)].LinkedNodes[nodeValue] = saveNode.LinkedNodes[opposite(nodeValue)]
	saveNode.LinkedNodes[opposite(nodeValue)] = rootNode.LinkedNodes[opposite(nodeValue)]
	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode
	saveNode = rootNode.LinkedNodes[opposite(nodeValue)]

	rootNode.LinkedNodes[opposite(nodeValue)] = saveNode.LinkedNodes[nodeValue]

	saveNode.LinkedNodes[nodeValue] = rootNode
	return saveNode
}

// adjustBalance metho
func adjustBalance(rootNode *VTreeNode,nodeValue,balanceValue int) {
	node := rootNode.LinkedNodes[nodeValue]
	oppNode := node.LinkedNodes[opposite(nodeValue)]
	switch oppNode.BalanceValue {
	case 0 :
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		rootNode.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		rootNode.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	oppNode.BalanceValue = 0
}

// BalanceTree method
func BalanceTree(rootNode *VTreeNode,nodeValue int) *VTreeNode {
	node := rootNode.LinkedNodes[nodeValue]
	balance := 2 * nodeValue - 1
	if node.BalanceValue == balance {
		rootNode.BalanceValue = 0
		node.BalanceValue = 0
		return singleRotation(rootNode,opposite(nodeValue))
	}
	adjustBalance(rootNode,nodeValue,balance)
	return doubleRotation(rootNode,opposite(nodeValue))
}

// insertRnode method
func insertRNode(rootNode *VTreeNode,key KeyValue) (*VTreeNode,bool) {
	if rootNode == nil {
		return &VTreeNode{KeyValue: key},false
	}
	dir := 0
	if rootNode.KeyValue.LessThan(key) {
		dir = 1
	}
	var done bool
	rootNode.LinkedNodes[dir],done = insertRNode(rootNode.LinkedNodes[dir],key)
	if done {
		return rootNode, true
	}
	rootNode.BalanceValue = rootNode.BalanceValue + (2*dir-1)
	switch rootNode.BalanceValue {
	case 0:
		return rootNode,true
	case 1,-1:
		return rootNode,false
	}
	return BalanceTree(rootNode,dir),true
}

// InsertNode method
func InsertNode(treeNode **VTreeNode,key KeyValue) {
	*treeNode,_ = insertRNode(*treeNode,key)
}
