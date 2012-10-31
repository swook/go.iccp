package matrix

import (
	"fmt"
)

type operator func(int, int, float64) float64

func (m Matrix) Foreach(op operator) {
	// Iterate over rows
	for y, vy := range m {
		// Iterate over columns
		for x, vx := range vy {
			// Apply operation to matrix element
			op(y, x, vx)
		}
	}
	return
}

func (A Matrix) Add(B Matrix) (C Matrix, err error) {
	sA := A.Size()
	sB := B.Size()
	if (sA.Y != sB.Y) || (sA.X != sB.X) {
		err = fmt.Errorf("go.iccp/matrix: Matrix dimensions mis-match.")
		return
	}

	C, _ = New(sA)
	for y := 0; y < sA.Y; y++ {
		for x := 0; x < sA.X; x++ {
			C[y][x] = A[y][x] + B[y][x]
		}
	}
	return
}
