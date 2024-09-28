package linkedlist

import (
	"fmt"
)

// Circular queue
type CircularQueue struct {
	size int
	nodes []interface{}
	head int
	last int
}

func NewCircularQueue(num int) *CircularQueue {
	circularQueue := CircularQueue{size:num+1,head:0,last:0}
	circularQueue.nodes = make([]interface{},num+1)
	return &circularQueue
}
// ISUnUsed method
func (circularQueue CircularQueue) ISUnUsed() bool {
	return circularQueue.head == circularQueue.last
}
// IsCompleted method
func (circularQueue CircularQueue) IsCompleted() bool {
	return circularQueue.head == (circularQueue.last+1)%circularQueue.size
}
// Add method
func (circularQueue *CircularQueue) Add(element interface{}) {
	if circularQueue.IsCompleted() {
		panic("Queue is Completely Utilized")
	}
	circularQueue.nodes[circularQueue.last] = element
	circularQueue.last = (circularQueue.last+1)%circularQueue.size
}
// MoveOneStep method
func (circularQueue *CircularQueue) MoveOneStep()(element interface{}) {
	if circularQueue.ISUnUsed(){
		return nil
	}
	element = circularQueue.nodes[circularQueue.head]
	circularQueue.head = (circularQueue.head+1)%circularQueue.size
	return 
}

func CircularFunc() {
	circularQueue := NewCircularQueue(5)
	circularQueue.Add(1)
	circularQueue.Add(2)
	circularQueue.Add(3)
	circularQueue.Add(4)
	circularQueue.Add(5)
	fmt.Println(circularQueue.nodes)
}