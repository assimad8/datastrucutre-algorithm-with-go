package binarytree

import "fmt"

// Tree struct 
type Tree struct {
	LeftNode *Tree
	Value int
	RightNode *Tree
}

// Tree class method 
func (tree *Tree) insert(m int) {
	if tree !=nil{
		if tree.LeftNode==nil{
			tree.LeftNode = &Tree{nil,m,nil}
		}else{
			if tree.RightNode ==nil{
				tree.LeftNode = &Tree{nil,m,nil}
			}else{
				if tree.LeftNode!=nil{
					tree.LeftNode.insert(m)
				}else {
					tree.RightNode.insert(m)
				}
			} 
		}
	}else{
		tree = &Tree{nil,m,nil}
	}
}
// print method for printing a Tree
func (tree *Tree) print() {
	if tree !=nil {
		fmt.Println("Node: ")
		fmt.Println("Value ",tree.Value)
		fmt.Print("Tree Node Left ")
		tree.LeftNode.print()
		fmt.Print("Tree Node Right ")
		tree.RightNode.print()
	}else {
		fmt.Println("Nil")
	}
}

func BinaryTreeFunc(){
	tree := &Tree{nil,1,nil}
	fmt.Println("==============")
	tree.print()
	tree.insert(3)
	fmt.Println("==============")
	tree.print()
	tree.insert(5)
	fmt.Println("==============")
	tree.print()
	tree.LeftNode.insert(7)
	fmt.Println("==============")
	tree.print()
}