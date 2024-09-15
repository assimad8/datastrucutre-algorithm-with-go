package homogeneous

import (
	"fmt"
	"math/rand"
)

func MultiArray(){
	var treedarray [2][2][2]int
	for i:=0;i<2;i++{
		for j:=0;j<2;j++{
			for b:=0;b<2;b++{
				treedarray[i][j][b] = rand.Intn(3)
			}
		}
	}
	fmt.Println(treedarray)
}


