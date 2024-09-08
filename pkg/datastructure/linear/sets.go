package linear

// import "fmt"

// Set class
type Set struct {
	integerMap map[int]bool
}
// create the map of integer
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}
//checks if element is in the set
func (set *Set) ContainsElement(element int) bool{
	var exists bool
	_, exists = set.integerMap[element]
	return exists
   }
// AddElement method
func (set *Set) AddElement(e int) {
	if !set.ContainsElement(e) {
		set.integerMap[e] = true
	}
}
//deletes the element from the set
func (set *Set) DeleteElement(element int) {
    delete(set.integerMap,element)
}
// Intersect method 
func (set *Set) Intersect(otherSet *Set) *Set {
	intersectSet := new(Set)
	intersectSet.New()
	for value,_:=range set.integerMap {
		if otherSet.ContainsElement(value) {
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}
// Union method 
func (set *Set) Union(otherSet *Set) *Set {
	unionSet := new(Set)
	unionSet.New()
	for value,_:=range set.integerMap {
		unionSet.AddElement(value)
	}
	for value,_:=range otherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}