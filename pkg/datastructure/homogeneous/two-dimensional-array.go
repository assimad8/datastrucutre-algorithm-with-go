package homogeneous

import "fmt"

// // Two dimension array
// var arr = [4][5]int{
// 	{1,2,3,4,5},
// 	{1,2,3,4,5},
// 	{1,2,3,4,5},
// 	{1,2,3,4,5},
// }
// // Roww matrix
// var rmatrix = [1][3] int {
// 	{1,2,3},
// }
// // Column matrix
// var cmatrix = [4][1]int {
// 	{1},
// 	{2},
// 	{3},
// 	{4},
// }
// // Lower triangular matrix
// var lmatrix = [3][3]int{
// 	{1,0,0},
// 	{1,1,0},
// 	{2,1,1},
// }
// // Uper triangular matrix
// var umatrix = [3][3]int{
// 	{1,2,3},
// 	{0,1,4},
// 	{0,0,1},
// }
// // Null matrix
// var nmatrix = [3][3]int{
// 	{0,0,0},
// 	{0,0,0},
// 	{0,0,0},
// }
// identity method
func Identity(order int) [][]float64{
	matrix := make([][]float64,order)
	for i:=0;i<order;i++{
		temp := make([]float64,order)
		temp[i] = 1
		matrix[i] = temp
	}
	return matrix
}
// Prints the matrix in zig-zag fashion
func PrintZigZag(n int) []int {
    var zigzag []int
    zigzag = make([]int, n*n)
    var i int
    i = 0
    var m int
    m = n * 2
    var  p int
    for p = 1; p <= m; p++ {
        var x int
        x = p - n
        if x < 0 {
           x = 0
        }
        var y int
        y = p - 1
        if y > n-1 {
            y = n - 1
       }
       var j int
       j = m - p
        if j > p {
            j = p
        }
        var k int
        for k = 0; k < j; k++ {
           if p&1 == 0 {
               zigzag[(x+k)*n+y-k] = i
           } else {
               zigzag[(y-k)*n+x+k] = i
            }
           i++
        }
   }
   return zigzag
}
// PrintSpiral method
func PrintSpiral(n int) []int {
    var left int
    var top int
    var right int
    var bottom int
    left =0
    top =0
    right = n-1
    bottom = n-1
    var size int
size = n * n
    var s []int
    s = make([]int, size)
    var i int
    i = 0
    for left < right {
        var c int
        for c = left; c <= right; c++ {
            s[top*n+c] = i
            i++
        }
        top++
        var r int
        for r = top; r <= bottom; r++ {
            s[r*n+right] = i
            i++
        }
        right--
        if top == bottom {
            break
        }
        for c = right; c >= left; c-- {
            s[bottom*n+c] = i
            i++
        }
        bottom--
        for r = bottom; r >= top; r-- {
            s[r*n+left] = i
            i++
        }
        left++
    }
    s[top*n+left] = i
    return s
}
// Boolean matrix
// changeMatrix method
func ChangeMatrix(matrix [3][3]int) [3][3]int {
	var i,j int
	var Rows,Columns [3]int
	var matrixChanged[3][3]int
	for i=0;i<3;i++{
		if matrix[i][j]==1{
			Rows[i]=1
			Columns[j]=1
		}
	}
	for i=0;i<3;i++{
		for j=0;j<3;j++{
			if Rows[i]==1 || Columns[j]==1{
				matrixChanged[i][j]=1
			}
		}
	}
	return matrixChanged
}
// PrintMatrix method
func PrintMatrix(matrix [3][3]int){
	for i:=0;i<3;i++{
		for j:=0; j<3; j++ {
			fmt.Printf("%d",matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

