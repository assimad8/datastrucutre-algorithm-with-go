package main

import (
	// "fmt"

	"github.com/assimad8/golang-algo/pkg/datastructure/homogeneous"
)

func main() {
	matrix := [3][3]int{{1,0,0},{0,0,0},{0,0,0},}
	homogeneous.PrintMatrix(matrix)
	matrix = homogeneous.ChangeMatrix(matrix)
	homogeneous.PrintMatrix(matrix)
}
