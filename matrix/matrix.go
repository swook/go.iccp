package matrix

import (
	"errors"
	"fmt"
	//"math"
	"strconv"
	"strings"
)

type Matrix [][]float64

type matrixSize struct {
	Y, X, Rows, Columns int
}

// _make creates a Matrix of y*x dimensions
// or y*y dimensions depending on x parameter
func _make(y int, _x ...int) Matrix {
	x := y
	if len(_x) == 1 {
		x = _x[0]
	}

	mat := make([][]float64, y)
	for i, _ := range mat {
		mat[i] = make([]float64, x)
	}
	return mat
}

// New returns a new Matrix created based on the input parameters
// (int) produces a Matrix of int x int
// (r int, c int) produces a matrix of r x c
// (string) produces a Matrix using parseMatlab
func New(input ...interface{}) (Matrix, error) {
	var mat Matrix
	var err error

	l := len(input)
	types := make([]string, l)
	for i := range input {
		types[i] = fmt.Sprintf("%T", input[i])
	}
	e := errors.New("go.iccp/matrix: Unknown input type (" + strings.Join(types, ", ") + ")")

	switch l {
	case 1:
		switch input[0].(type) {
		default:
			err = e
		case string:
			mat, err = parseMatlab(input[0].(string))
		case int:
			mat = _make(input[0].(int))
		case Matrix:
			mat = input[0].(Matrix).Duplicate()
		case matrixSize:
			s := input[0].(matrixSize)
			mat = _make(s.Y, s.X)
		}
	case 2:
		switch input[0].(type) {
		case int:
			switch input[1].(type) {
			case int:
				mat = _make(input[0].(int), input[1].(int))
			default:
				err = e
			}
		default:
			err = e
		}
	default:
		err = e
	}
	return mat, err
}

// Identity returns an identity matrix of NxN
func Identity(N int) (Matrix, error) {
	if N < 1 {
		return nil, fmt.Errorf("go.iccp/matrix: Provided size invalid (%f)", N)
	}
	mat, _ := New(N)
	// iterate over row
	for r, _ := range mat {
		// Set as identity matrix
		mat[r][r] = 1.0
	}
	return mat, nil
}

// parseMatlab parses a matrix in string form.
// Valid row delimiters are ';' and '\n'.
// Valid column delims are ' ', ',', '\n'.
// All row delimiters are counted, but column delimiters are only counted if a value is present.
func parseMatlab(str string) (Matrix, error) {
	var mat Matrix
	var err error
	var snum []rune
	var num float64
	var rows, cols int

	// Pass 1 to determine dimensions
	var r, c int = 1, 0
	for _, v := range str {
		// When number
		if v == '.' || (v > 47 && v < 58) {
			snum = append(snum, v)
			continue
		}
		if len(snum) == 0 {
			continue
		}

		switch v {
		case ' ', ',', '\t':
			c += 1
		case ';', '\n':
			r += 1
			if c > cols {
				cols = c
			}
			c = 0
		}
		snum = snum[0:0]
	}
	rows = r
	if c > cols {
		cols = c
	}
	cols += 1
	mat = _make(rows, cols)

	// Pass 2 to populate matrix
	r, c = 0, 0
	snum = snum[0:0]
	for _, v := range str {
		// When number
		if v == '.' || (v > 47 && v < 58) {
			snum = append(snum, v)
			continue
		}
		if len(snum) == 0 {
			continue
		}
		num, _ = strconv.ParseFloat(string(snum), 64)

		switch v {
		case ' ', ',', '\t':
			mat[r][c] = num
			c += 1
		case ';', '\n':
			mat[r][c] = num
			r += 1
			c = 0
		}
		snum = snum[0:0]
	}
	if len(snum) > 0 {
		mat[r][c], _ = strconv.ParseFloat(string(snum), 64)
	}
	return mat, err
}

// Duplicate returns a copy of the current Matrix
func (m Matrix) Duplicate() Matrix {
	s := m.Size()
	m2, _ := New(s.Y, s.X)
	for y := 0; y < s.Y; y++ {
		for x := 0; x < s.X; x++ {
			m2[y][x] = m[y][x]
		}
	}
	return m2
}

func (m Matrix) Size() matrixSize {
	return matrixSize{
		len(m),
		len(m[0]),
		len(m),
		len(m[0]),
	}
}

// String defines the format in which a Matrix is stringified.
func (m Matrix) String() string {
	out := "[\n  "
	size := m.Size()
	stopY := size.Y - 1
	stopX := size.X - 1
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			out += fmt.Sprintf("%f", m[y][x])
			if x < stopX {
				out += " "
			}
		}
		if y < stopY {
			out += "\n  "
		}
	}
	out += "\n]"
	return out
}
