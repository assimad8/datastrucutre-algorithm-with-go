package homogeneous

import (
	"fmt"
	"math/rand"
)

func TensorFunc() {
	var array [3][3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				array[i][j][k] = rand.Intn(3)
			}
		}
	}
	fmt.Println("The array: ",array)
	fmt.Println("Zero mode unfold : ")
	for j:=0;j<3;j++{
		for k:=0;k<3;k++{
			fmt.Printf("%d ",array[0][j][k])
		}
		fmt.Println("")
	}
	fmt.Println("1-mode unfold : ")
	for j:=0;j<3;j++{
		for k:=0;k<3;k++{
			fmt.Printf("%d ",array[1][j][k])
		}
		fmt.Println("")
	}
	fmt.Println("2-mode unfold : ")
	for j:=0;j<3;j++{
		for k:=0;k<3;k++{
			fmt.Printf("%d ",array[2][j][k])
		}
		fmt.Println("")
	}
}