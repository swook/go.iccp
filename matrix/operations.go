package matrix

import (
	"fmt"
	"math"
)

// Foreach performs an operation over all elements in a Matrix.
// It takes a parameter of type func(int, int) error
func (m Matrix) Foreach(op func(int, int)) {
	// Iterate over rows
	for y, vy := range m {
		// Iterate over columns
		for x, _ := range vy {
			// Apply operation to matrix element
			op(y, x)
		}
	}
	return
}

// Add adds an int, float64 or Matrix to an existing Matrix.
// For int and float64, the operation is performed on each element.
func (A Matrix) Add(in ...interface{}) (Matrix, error) {
	var mat Matrix
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if (sA.Y != sB.Y) || (sA.X != sB.X) {
			return mat, fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}

		mat, _ = New(sA)
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] + B[y][x]
		})
	case float64:
		a := in[0].(float64)
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] + a
		})
	case int:
		a := float64(in[0].(int))
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] + a
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Add.")
	}
	return mat, err
}

// Subtract subtracts an int, float64 or Matrix from an existing Matrix.
// For int and float64, the operation is performed on each element.
func (A Matrix) Subtract(in ...interface{}) (Matrix, error) {
	var mat Matrix
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if (sA.Y != sB.Y) || (sA.X != sB.X) {
			return mat, fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}

		mat, _ = New(sA)
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] - B[y][x]
		})
	case float64:
		a := in[0].(float64)
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] - a
		})
	case int:
		a := float64(in[0].(int))
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = A[y][x] - a
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Subtract.")
	}
	return mat, err
}

// Multiply multiplies an int, float64 or Matrix to an existing Matrix.
// For int and float64, the operation is performed on each element.
func (A Matrix) Multiply(in ...interface{}) (Matrix, error) {
	var mat Matrix
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if sA.X != sB.Y {
			return mat, fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}
		// New matrix of new dimensions
		mat, _ = New(sA.Y, sB.X)

		// Loop through rows of new matrix
		for y, yv := range mat {
			// columns
			for x, _ := range yv {
				// Multiplication operation
				for i := 0; i < sA.X; i++ {
					mat[y][x] += A[y][i] * B[i][x]
				}
			}
		}
	case float64:
		a := in[0].(float64)
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = a * A[y][x]
		})
	case int:
		a := float64(in[0].(int))
		mat, _ = New(A.Size())
		mat.Foreach(func(y, x int) {
			mat[y][x] = a * A[y][x]
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Subtract.")
	}
	return mat, err
}

// Transpose returns the transpose of the current matrix.
func (A Matrix) Transpose() Matrix {
	sA := A.Size()
	B, _ := New(sA.X, sA.Y)
	A.Foreach(func(r, c int) {
		B[c][r] = A[r][c]
	})
	return B
}

// OuterProduct calculates the outer product of two matrices.
func (A Matrix) OuterProduct(B Matrix) (Matrix, error) {
	return A.Multiply(B.Transpose())
}

// ScalarProduct calculates the scalar product of two matrices.
func (A Matrix) ScalarProduct(B Matrix) (float64, error) {
	var sum float64
	sA := A.Size()
	sB := B.Size()
	if (sA.Y != sB.Y) || (sA.X != sB.X) {
		return sum, fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
	}

	A.Foreach(func(r, c int) {
		sum += A[r][c] * B[r][c]
	})
	return sum, nil
}

// Inverse uses Gauss-Jordan Elimination to calculate the inverse of a given matrix.
func (A Matrix) Inverse() (Matrix, error) {
	var err error
	sA := A.Size()
	if sA.X != sA.Y {
		return nil, fmt.Errorf("go.iccp/matrix: Matrix dimensions invalid")
	}
	A = A.Duplicate()

	var max int
	var rat float64
	B, _ := Identity(sA.X)
	for r, _ := range A {
		max = r
		for r_ := r + 1; r_ < sA.Rows; r_++ {
			if math.Abs(A[r_][r]) > math.Abs(A[max][r]) {
				max = r_
			}
		}
		A[max], A[r] = A[r], A[max]
		B[max], B[r] = B[r], B[max]
	}
	for r, _ := range A {
		for r_, _ := range A {
			rat = A[r_][r] / A[r][r]
			for c := 0; c < sA.Columns; c++ {
				if r == r_ {
					continue
				}
				A[r_][c] -= A[r][c] * rat
				B[r_][c] -= B[r][c] * rat
			}
		}
	}
	for r, _ := range A {
		rat = A[r][r]
		for c := 0; c < sA.Columns; c++ {
			A[r][c] /= rat
			B[r][c] /= rat
		}
	}
	return B, err
}
