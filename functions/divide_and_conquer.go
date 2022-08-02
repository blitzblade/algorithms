package functions

func MatrixMultiplyRecursive(A[][]int, B[][]int, C[][]int, n int) [][]int{
	//Base case
	if n == 1 {
		C[0][0] = C[0][0] + A[0][0] * B[0][0]
		return C
	}

	//Divide
	a11, a12, a21, a22 := split(A)
	b11, b12, b21, b22 := split(B)
	c11, c12, c21, c22 := split(C)

	//conquer
	//[A11 A12] x [B11 B12]  = [A11*B11+A12*B21  A11*B12+A12*B22]
	//[A21 A22]   [B21 B22]    [A21*B11+A22*B22 A21*B12+A22*B22]
	MatrixMultiplyRecursive(a11,b11,c11,n/2)
	MatrixMultiplyRecursive(a11, b12, c12, n/2)
	MatrixMultiplyRecursive(a21,b11,c21,n/2)
	MatrixMultiplyRecursive(a21,b12,c22, n/2)
	MatrixMultiplyRecursive(a12,b21,c12,n/2)
	MatrixMultiplyRecursive(a12,b22,c12,n/2)
	MatrixMultiplyRecursive(a22,b21,c22,n/2)
	MatrixMultiplyRecursive(a22,b22,c22,n/2)
return C
}
//[             
// [1,2,3],
// [4,5,6],
// [7,8,9]
//]




func split(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(matrix)
	n1 := n/2
	a11 := make([][]int, n1)
	a12 := make([][]int, n1)
	a21 := make([][]int, n1)
	a22 := make([][]int, n1)

	for i := 0; i < n1; i++ {
		a11[i] = matrix[i][0:n1];
		a12[i] = matrix[i][n1:n]
		a21[i] = matrix[n1+i][0:n1]
		a22[i] = matrix[n1+i][n1:n]
	}

	return a11, a12, a21, a22
}