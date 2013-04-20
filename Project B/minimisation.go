package main

import (
	// "fmt"
	"github.com/swook/go.iccp/matrix"
	"math"
	"sort"
)

// xysort satisfies sort.Interface
type xysort struct {
	X, Y []float64
}

func (xy xysort) Len() int {
	return len(xy.X)
}

func (xy xysort) Less(i, j int) bool {
	return xy.X[i] < xy.X[j]
}

func (xy xysort) Swap(i, j int) {
	_x, _y := xy.X[i], xy.Y[i]
	xy.X[i] = xy.X[j]
	xy.Y[i] = xy.Y[j]
	xy.X[j] = _x
	xy.Y[j] = _y
}

// parabolicMin uses parabolic minimisation to find the minimum of a function
// given an initial x
func parabolicMin(x0 float64, N int, f func(float64, int) float64) float64 {
	x := []float64{x0, x0 + .1, x0 + .2}
	y := make([]float64, 3)
	var A, B float64
	var maxx int

	// Populate y value list
	for i, _ := range y {
		y[i] = f(x[i], N)
	}

	// Minimise until condition
	for {
		A = (y[1] - y[0]) / (x[1] - x[0])
		B = ((y[2]-y[0])/(x[2]-x[0]) - (y[1]-y[0])/(x[1]-x[0])) / (x[2] - x[1])
		x = append(x, .5*((x[0]+x[1])*B-A)/B)
		y = append(y, f(x[3], N))

		// Find highest y(x)
		maxx = 0
		for i := 1; i < 3; i++ {
			if y[i] > y[maxx] {
				maxx = i
			}
		}
		// Remove highest element
		x = append(x[:maxx], x[maxx+1:]...)
		y = append(y[:maxx], y[maxx+1:]...)

		// Sort slice in descending order
		sort.Sort(xysort{x, y})

		// If difference between lowest 2 x-values small
		if x[1]-x[0] < 1e-7 {
			break
		}
	}
	return x[0]
}

// QuasiNewtonMin uses the Quasi-Newton method to carry out multi-dimensional minimisation
func QuasiNewtonMin(x0 matrix.Matrix, N int, f func(matrix.Matrix, int) float64) (x1 matrix.Matrix) {
	// Create G matrix, initially identity matrix
	G, _ := matrix.Identity(x0.Size().Y)

	// Central difference scheme to calculate gradient
	CDS := func(in matrix.Matrix) matrix.Matrix {
		res, _ := matrix.New(in.Size())
		var p, m float64
		res.Foreach(func(r, c int) {
			in[r][c] += 1e-8
			p = f(in, N)
			in[r][c] -= 2e-8
			m = f(in, N)
			in[r][c] += 1e-8
			res[r][c] = (p - m) / 2e-8
		})
		return res
	}

	a := 1.0
	r0 := CDS(x0)
	x1, _ = matrix.New(x0)
	var cond int
	var dx, ddx, tmp, tmp1, tmp2, r1, dr matrix.Matrix

	// Main minimisation loop
	for i := 0; i < 1000; i++ {

		// Determine a using Wolfe Conditions
		a, cond = 1.0, 0
		for {
			// tmp = p_k
			tmp, _ = G.Multiply(r0)
			tmp, _ = tmp.Multiply(-1)

			tmp1, _ = tmp.Multiply(a)
			tmp1, _ = x0.Add(tmp1)

			// Check last term only if condition not yet met
			if cond == 0 {
				tmp2, _ = tmp.Transpose().Multiply(r0)

				// Sufficient Decrease Condition
				if f(tmp1, N) <= f(x0, N)+a*1e-4*tmp2[0][0] {
					cond++
				}
			}

			// If first condition met, try second condition
			if cond > 0 {
				tmp1, _ = tmp.Transpose().Multiply(CDS(tmp1))
				tmp2, _ = tmp.Transpose().Multiply(r0)

				// Curvature Condition
				// fmt.Println(tmp1[0][0], 0.9*tmp2[0][0])
				if math.Abs(tmp1[0][0]) >= 0.9*math.Abs(tmp2[0][0]) {
					break
				}
			}

			// Backtracking line search
			a *= 0.5
			cond = 0 // Reset checked cond count
		}

		// Find next x
		tmp, _ = G.Multiply(a)
		tmp, _ = tmp.Multiply(r0)
		x1, _ = x0.Subtract(tmp)
		// fmt.Println(i, "x1:", x1)

		// Calculate change in x and grad
		r1 = CDS(x1)
		// fmt.Println(i, "r1:", r1)

		// Convergence condition:
		// when gradient no longer changes
		if r0[0][0] == r1[0][0] && r0[1][0] == r1[1][0] {
			break
		}

		dx, _ = x1.Subtract(x0)
		dr, _ = r1.Subtract(r0)

		// First term in increment of G
		tmp, _ = dr.Transpose().Multiply(dx)
		ddx, _ = dx.OuterProduct(dx)
		tmp1, _ = ddx.Multiply(1 / tmp[0][0])
		// fmt.Println(i, "G:", G)

		// Second term in increment of G
		tmp, _ = dr.Transpose().Multiply(G)
		tmp, _ = tmp.Multiply(dr)
		tmp2, _ = G.Multiply(dr)
		tmp2, _ = tmp2.OuterProduct(dr)
		tmp2, _ = tmp2.OuterProduct(G)
		tmp2, _ = tmp2.Multiply(1 / tmp[0][0])

		// Update G for next x
		G, _ = G.Add(tmp1)
		// fmt.Println(i, "Add -> G:", G)
		G, _ = G.Subtract(tmp2)
		// fmt.Println(i, "Subtract -> G:", G)

		// Update old x and grad for later dx and dgrad
		x0 = x1
		r0 = r1
	}
	return
}
