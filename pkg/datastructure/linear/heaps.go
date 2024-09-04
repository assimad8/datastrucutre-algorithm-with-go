package linear

import (
	"container/heap"
	"fmt"
)

// integerHeap a type
type IntegerHeap []int

func (he IntegerHeap) Len() int { return len(he)}

func (he IntegerHeap) Less(i, j int) bool {
	return he[i] < he[j]
}

func (he IntegerHeap) Swap(i, j int) {
	he[i], he[j] = he[j], he[i]
}

func (he *IntegerHeap) Push(heapintf interface{}) {
	*he = append(*he,heapintf.(int))
}

func (he *IntegerHeap) Pop() interface{} {
	var n,x1 int
	var previous IntegerHeap = *he
	n = previous.Len()
	x1 = previous[n-1]
	*he = previous[0:n-1]
	return x1
}

func MyHeap() {
	iheap := &IntegerHeap{1,2,3}
	heap.Init(iheap)
	heap.Push(iheap,4)
	iheap.Push(5)
	fmt.Printf("minimum: %d\n", (*iheap)[0])
 	for iheap.Len() > 0 {
 		fmt.Printf("%d \n", heap.Pop(iheap))
	}
}