package main

import (
	"github.com/swook/go.iccp/matrix"
	"github.com/swook/gogsl/specfunc"
	"math"
)

// NLLFunc is the fit function used in this project
func fitFunc(t, tau, sigma float64) float64 {
	sig_tau := sigma / tau
	erfc, _ := specfunc.Erfc(1 / math.Sqrt(2) * (sig_tau - t/sigma))
	return 0.5 / tau * math.Exp(0.5*sig_tau*sig_tau-t/tau) * erfc
}

// calcFitFunc returns a list of t and fitFunc results which extend until
// y < 1e-5, ie. close to x-axis
func calcFitFunc(t, dt, tau, sigma float64) (x, y []float64) {
	x = make([]float64, 0)
	y = make([]float64, 0)
	for i := 0; ; i++ {
		if len(y) > 0 && y[i-1] < 1e-5 {
			break
		}
		x = append(x, t)
		y = append(y, fitFunc(t, tau, sigma))
		t += dt
	}
	//y = normalise(dt, y)
	return
}

// NLLFunc is the Negative-Log-Likelihood function used in this project
func NLLFunc(tau float64, N int) (NLL float64) {
	for i := 0; i < N; i++ {
		NLL += math.Log(fitFunc(data[0][i], tau, data[1][i]))
	}
	NLL = -NLL
	return
}

// tauFromNLL uses the Newton-Raphson method to find solutions of NLLFunc-NLLFunc(minimum)-0.5
func tauFromNLL(NLL float64, N int) (tau_plus, min, tau_minus float64) {
	min = parabolicMin(0.5, N, NLLFunc)

	// Offset NLL function by NLL + y at minimum
	NLL += NLLFunc(min, N)

	// grad finds the gradient of NLLFunc at point x using the
	// central difference scheme method
	grad := func(x float64) float64 {
		return (NLLFunc(x+0.001, N) - NLLFunc(x-0.001, N)) / 0.002
	}

	// NR is a Newton-Raphson method implementation to search for roots of (NLLFunc()-NLL)
	NR := func(x0 float64) float64 {
		x1 := x0 + 0.001
		var x_ float64
		for i := 0; i < 20; i++ {
			x_ = x1
			x1 = x0 - (NLLFunc(x0, N)-NLL)/grad(x0)
			x0 = x_
			// Stop if change in x smaller than 1e-12
			if x1-x0 < 1e-12 {
				break
			}
		}
		return x1
	}
	tau_minus = NR(min - 0.2)
	tau_plus = NR(min + 0.2)
	return
}

// tauFromNLLBG has the same function as tauFromNLL, but uses NLLFuncWithBG
func tauFromNLLBG(min matrix.Matrix, NLL float64, N int) (tau_plus, tau_minus float64) {
	// Offset NLL function by NLL + y at minimum
	NLL += NLLFuncWithBG(min, N)

	// grad finds the gradient of NLLFunc at point x using the
	// central difference scheme method
	grad := func(in matrix.Matrix) float64 {
		var p, m float64
		in[0][0] += 1e-8
		p = NLLFuncWithBG(in, N)
		in[0][0] -= 2e-8
		m = NLLFuncWithBG(in, N)
		in[0][0] += 1e-8
		return (p - m) / 2e-8
	}

	// NR is a Newton-Raphson method implementation to search for roots of (NLLFunc()-NLL)
	NR := func(x0 matrix.Matrix) float64 {
		x0 = x0.Duplicate()
		x1 := x0.Duplicate()
		x1[0][0] += 0.001
		var x_ float64
		for i := 0; i < 20; i++ {
			x_ = x1[0][0]
			x1[0][0] = x0[0][0] - (NLLFuncWithBG(x0, N)-NLL)/grad(x0)
			x0[0][0] = x_
			// Stop if change in x smaller than 1e-12
			if x1[0][0]-x0[0][0] < 1e-12 {
				break
			}
		}
		return x1[0][0]
	}
	min[0][0] -= 0.2
	tau_minus = NR(min)
	min[0][0] += 0.4
	tau_plus = NR(min)
	return
}

// normalise takes a given list of data, and attempts to normalise it by first
// calculating the area under the curve, and dividing all values by the area
func normalise(dx float64, ydata []float64) []float64 {
	A := 0.0            // Normalisation constant
	f := len(ydata) - 2 // Last index for Simpsons'

	// Use Simpsons' Rule to integrate
	for i := 0; i < f; i += 2 {
		A += ydata[i] + 4.0*ydata[i+1] + ydata[i+2]
	}

	if A == 0 {
		return ydata
	}
	A *= dx / 3.0
	// println(A)

	// Normalise
	for i, _ := range ydata {
		ydata[i] /= A
	}

	return ydata
}

// BGFunc is the background Gaussian
func BGFunc(t, sigma float64) float64 {
	return math.Exp(-.5*t*t/sigma/sigma) / (sigma * math.Sqrt(2*math.Pi))
}

// NLLFuncWithBG calculates the new fit function given the new parameters which includes a,
// the proportion of non-background contribution to measurements
func NLLFuncWithBG(in matrix.Matrix, N int) (NLL float64) {
	tau := in[0][0]
	a := in[1][0]
	a_ := 1 - a

	var t, sigma, next float64
	for i := 0; i < N; i++ {
		t = data[0][i]
		sigma = data[1][i]
		next = a*fitFunc(t, tau, sigma) + a_*BGFunc(t, sigma)
		if next > 0 {
			NLL += math.Log(next)
		}
	}
	NLL = -NLL
	return
}
