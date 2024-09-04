package linear

import (
	"fmt"
)

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int,int) {
	return a*a, a*a*a
  }

func MyTuple() {
	var square int
	var cube int
	square,cube = powerSeries(3)
	fmt.Printf("Square %d\nCube %d\n",square,cube)
}