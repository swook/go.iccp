package matrix

import (
	"fmt"
)

type operator func(int, int) error

func (m Matrix) Foreach(op operator) {
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

func (A Matrix) Add(in ...interface{}) (err error) {
	switch in[0].(type) {
	case Matrix:
		B := in[0].(Matrix)
		sA := A.Size()
		sB := B.Size()
		if (sA.Y != sB.Y) || (sA.X != sB.X) {
			err = fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
			return
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
	return
}
