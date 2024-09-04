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
