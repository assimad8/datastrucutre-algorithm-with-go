package homogeneous



// Basic 2D matrix operations
// Add method
func Add(matrix1,matrix2 [2][2]int) [2][2]int {
	var sum [2][2]int
	for l:=0;l<2;l++{
		for m:=0;m<2;m++{
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}
// Subtract method
func Subtract(matrix1,matrix2 [2][2]int) [2][2]int {
	var difference [2][2]int
	for l:=0;l<2;l++{
		for m:=0;m<2;m++{
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}
// Multiply method
func Multiply(matrix1,matrix2 [2][2]int) [2][2]int {
	var product  [2][2]int
	for l:=0;l<2;l++{
		for m:=0;m<2;m++{
			product [l][m] = matrix1[l][m] * matrix2[l][m]
		}
	}
	return product 
}
// transpose method
func Transpose(matrix1 [2][2]int) [2][2]int {
    var transMatrix [2][2]int
    for l := 0; l < 2; l++ {
        for m:=0; m <2; m++ {
            transMatrix[l][m] = matrix1[m][l]
        }
    }
    return transMatrix
}
// determinant method
func Determinant(matrix1 [2][2]int) int {
    det := matrix1[0][0]*matrix1[1][1]-matrix1[0][1]*matrix1[1][0]
    return det
}
//inverse method
func Inverse(matrix [2][2]int) [2][2]int {
	var invmatrix [2][2]int
	det := Determinant(matrix)
	invmatrix[0][0] = matrix[1][1]/det
	invmatrix[0][1] = -1*matrix[0][1]/det
	invmatrix[1][0] = -1*matrix[1][0]/det
	invmatrix[1][1] = matrix[0][0]/det
	return invmatrix
}