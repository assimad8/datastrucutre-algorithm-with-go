package patterns


import "fmt"

// IComposite interface
type IComposite interface{
	perform()
}

// LeafLet struct
type LeafLet struct{
	name string
}
// LeafLet class method perform
func (leaf *LeafLet) perform(){
	fmt.Println("LeafLet "+leaf.name)
}

// Branch struct
type Branch struct{
	leafs []LeafLet
	name string
	branches []Branch
}
func (branch *Branch) perform() {
	fmt.Println("Branch: "+branch.name)
	for _,leaf := range branch.leafs{
		leaf.perform()
	}
	for _,branch := range branch.branches{
		branch.perform()
	} 
}

// Branch class method add leaflet
func (branch *Branch) add(leaf LeafLet) {
	branch.leafs = append(branch.leafs, leaf)
}
// Branch class method add branch
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// Branch class method getLeaflts
func (branch *Branch) getLeaflts() []LeafLet {
	return branch.leafs
}

func Composite() {
	var branch = &Branch{name:"branch 1"}
	var leaf1 = LeafLet{name:"leaf 1"}
	var leaf2 = LeafLet{name:"leaf 2"}
	var branch2 = Branch{name:"branch 2"}
	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
	branch.getLeaflts()
}

