package matrix

import (
	"fmt"
)

// Foreach performs an operation over all elements in a Matrix.
// It takes a parameter of type func(int, int) error
func (m Matrix) Foreach(op func(int, int) error) {
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
func (A Matrix) Add(in ...interface{}) error {
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if (sA.Y != sB.Y) || (sA.X != sB.X) {
			return fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}

		A.Foreach(func(y, x int) error {
			A[y][x] += B[y][x]
			return nil
		})
	case float64:
		a := in[0].(float64)
		A.Foreach(func(y, x int) error {
			A[y][x] += a
			return nil
		})
	case int:
		a := float64(in[0].(int))
		A.Foreach(func(y, x int) error {
			A[y][x] += a
			return nil
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Add.")
	}
	return err
}

// Subtract subtracts an int, float64 or Matrix from an existing Matrix.
// For int and float64, the operation is performed on each element.
func (A Matrix) Subtract(in ...interface{}) error {
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if (sA.Y != sB.Y) || (sA.X != sB.X) {
			return fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}

		A.Foreach(func(y, x int) error {
			A[y][x] -= B[y][x]
			return nil
		})
	case float64:
		a := in[0].(float64)
		A.Foreach(func(y, x int) error {
			A[y][x] -= a
			return nil
		})
	case int:
		a := float64(in[0].(int))
		A.Foreach(func(y, x int) error {
			A[y][x] -= a
			return nil
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Subtract.")
	}
	return err
}

// Multiply multiplies an int, float64 or Matrix to an existing Matrix.
// For int and float64, the operation is performed on each element.
func (A Matrix) Multiply(in ...interface{}) error {
	var err error
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if sA.X != sB.Y {
			return fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		}
		// New matrix of new dimensions
		C, _ := New(sA.Y, sB.X)

		// Loop through rows of new matrix
		for cy, cyv := range C {
			// columns
			for cx, _ := range cyv {
				// Multiplication operation
				for i := 0; i < sA.X; i++ {
					C[cy][cx] += A[cy][i] * B[i][cx]
				}
			}
		}

		// Update A to have values of new multiplied Matrix
		for cy, _ := range C {
			A[cy] = C[cy]
		}
	case float64:
		a := in[0].(float64)
		A.Foreach(func(y, x int) error {
			A[y][x] *= a
			return nil
		})
	case int:
		a := float64(in[0].(int))
		A.Foreach(func(y, x int) error {
			A[y][x] *= a
			return nil
		})
	default:
		err = fmt.Errorf("go.iccp/matrix: Invalid input for method Subtract.")
	}
	return err
}

func (A Matrix) Transpose() error {
	var err error
	return err
}

func (A Matrix) OuterProduct(B Matrix) error {
	var err error
	return err
}
