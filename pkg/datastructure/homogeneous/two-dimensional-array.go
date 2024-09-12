package homogeneous

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





