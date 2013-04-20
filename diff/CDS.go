package diff

import (
//	"github.com/swook/go.iccp"
//	"math"
)

type y func(float64)float64

// Error of order h
func FDS(f y, x0 float64, h float64) float64 {
	return (f(x0 + h) - f(x0)) / h
}

// Error of order h
func BDS(f y, x0 float64, h float64) float64 {
	return (f(x0) - f(x0 - h)) / h
}

// Error of order h^2
func CDS(f y, x0 float64, h float64) float64 {
	return (f(x0 + h) - f(x0 - h)) / 2 / h
}