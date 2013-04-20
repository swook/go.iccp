package main

import (
	"fmt"
	"github.com/swook/go.iccp/matrix"
)

// Read data into memory
// data is a 2-d slice of floats
// data[0] contains tau data
// data[1] contains sigma data
var data = readFile("lifetime.txt")

func main() {
	/*
		Plot histograms of sigma with varying bin widths
	*/
	createHistogram(data[0], 50)
	createHistogram(data[0], 100)
	createHistogram(data[0], 500)
	createHistogram(data[0], 1000)
	createHistogram(data[0], 5000)
	fmt.Println("Plotted histogram of source data")

	/*
		Plot fit function graphs
	*/
	var x, y []float64
	N := len(data[0])
	tau := 0.1
	sigma := 0.1

	// vary sigma
	xs := make([][]float64, 5)
	ys := make([][]float64, 5)
	yls := make([]string, 5)
	for i := 0; i < 5; i++ {
		xs[i], ys[i] = calcFitFunc(0.0, 0.01, tau, sigma)
		yls[i] = fmt.Sprintf("sigma = %.1f", sigma)
		sigma += 0.2
	}
	createLine(xs, ys, yls, fmt.Sprintf("Fit Function (tau = %.1f)", tau))

	// vary tau
	tau, sigma = 0.1, 0.1
	for i := 0; i < 5; i++ {
		xs[i], ys[i] = calcFitFunc(0.0, 0.01, tau, sigma)
		yls[i] = fmt.Sprintf("tau = %.1f", tau)
		tau += 0.2
	}
	createLine(xs, ys, yls, fmt.Sprintf("Fit Function (sigma = %.1f)", sigma))
	fmt.Println("Plotted fit function with varying variables")

	/*
		Plot NLL vs tau for idea of what we're minimising
	*/
	tau = 0.1
	x = make([]float64, 0, 100)
	y = make([]float64, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, tau)
		y = append(y, NLLFunc(tau, N))
		tau += 0.01
	}
	createScatter(x, y, "NLL vs tau")
	fmt.Println("Plotted NLL against tau")

	/*
		Parabolic Minimisation
	*/
	tau = parabolicMin(0.5, N, NLLFunc)
	fmt.Println("Minimum of NLL (no Background) found using parabolic method: tau =", tau, "NLL =", NLLFunc(tau, N))

	/*
		Plot std dev of 1-d fit result
	*/
	var min, tau_p, tau_m float64
	x = make([]float64, 0, 100)
	y = make([]float64, 0, 100)
	mins := make([]float64, 0, 100)
	for i := 0; i < 100; i++ {
		N = len(data[0]) - i*50
		tau_p, min, tau_m = tauFromNLL(0.5, N)
		x = append(x, float64(N))
		mins = append(mins, min)
		y = append(y, tau_p-tau_m)
	}
	createScatter(x, y, "Standard Deviation vs Sample Size")
	fmt.Println("Plotted graph of standard dev vs sample size")
	saveFile("stddev v samp size.txt", [][]float64{x, y})

	createScatter(x, mins, "Determined Minimum vs Sample Size")
	fmt.Println("Plotted graph of minimum found through parabolic minimisation vs sample size")
	saveFile("parabolic min v samp size.txt", [][]float64{x, mins})
	N = len(data[0])

	/*
		Find minimum tau and a for NLLFuncWithBG
	*/
	x0, _ := matrix.New("0.4;0.1")
	x0 = QuasiNewtonMin(x0, N, NLLFuncWithBG)
	fmt.Println("Minimum of NLL (w Background) found using Quasi-Newton DFP method: tau =",
		x0[0][0], "a =", x0[1][0], "NLL =", NLLFuncWithBG(x0, N))

	/*
		Plot of NLLFuncWithBG against tau to confirm minimum found in previous step
	*/
	tau = 0.1
	var mat matrix.Matrix
	x = make([]float64, 0, 100)
	y = make([]float64, 0, 100)
	for i := 0; i < 100; i++ {
		x = append(x, tau)
		mat, _ = matrix.New(fmt.Sprintf("%f;0.983684", tau))
		y = append(y, NLLFuncWithBG(mat, N))
		tau += 0.01
	}
	createScatter(x, y, fmt.Sprintf("NLL with BG, a = %f", x0[1][0]))

	/*
		Find error of NLLBG fit
	*/
	tau_p, tau_m = tauFromNLLBG(x0, 0.5, N)
	fmt.Println("Error of calculated NLL (w BG) is: ", tau_p-tau_m)
}
