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