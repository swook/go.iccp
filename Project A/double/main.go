package main

import "math"

// u: pendulum 1's anglular position
// v: pendulum 1's angular velocity
// o: pendulum 2's anglular position
// p: pendulum 2's angular velocity

// alpha: time scale elmination factor

var (
	t0 = 0.0  // Initial t
	tf = 40.0 // Final t

	u0 = 0.1 // Initial θ
	v0 = 0.0 // Initial φ
	o0 = 0.0 // Initial ω
	p0 = 0.0 // Initial ν

	g = 9.81 // Gravitational acceleration

	m = 1.0 // Mass of pendulum 1
	M = 1.0 // Mass of pendulum 2

	l     = 10.0 // Length of pendulum
	gamma = 1.0  // Damping factor

	h = 0.001 // Step size

	R     = M / m
	G     = gamma / (m * math.Sqrt(g*l))
	alpha = math.Sqrt(l / g) // Time scale

	dElim    = 0.0
	barscale = 0.0
)

// updateRGa updates the R and G variables used in calculations
func updateRGa() {
	R = M / m
	G = gamma / (m * math.Sqrt(g*l))
	alpha = math.Sqrt(l / g) // Time scale
}

func main() {
	// Run with
	//   m = 1.0kg
	//   M = 1.0kg
	//   γ = 0.0
	m, M, gamma, dElim, barscale = 1.0, 1.0, 0.0, 0.008, 0.000005
	RK4(true)
	Stability()
	//   γ = 1.0
	gamma, barscale = 1.0, 0.0000005
	RK4(true)
	Stability()

	// Run with
	//   m = 1.0kg
	//   M = 100.0kg
	//   γ = 0.0
	m, M, gamma, dElim, barscale = 1.0, 100.0, 0.0, 100.0, 0.2
	RK4(true)
	Stability()
	//   γ = 1.0
	gamma = 1.0
	RK4(true)
	Stability()

	// Run with
	//   m = 100.0kg
	//   M = 1.0kg
	//   γ = 0.0
	m, M, gamma, dElim, barscale = 100.0, 1.0, 0.0, 0.4, 0.000001
	RK4(true)
	Stability()
	//   γ = 1.0
	gamma, barscale = 1.0, 0.0000007
	RK4(true)
	Stability()
}

func fu(u, v, o, p float64) float64 {
	return v
}

func fv(u, v, o, p float64) float64 {
	return -(R+1)*u/alpha/alpha + R*o/alpha/alpha - G*v/alpha
}

func fo(u, v, o, p float64) float64 {
	return p
}

func fp(u, v, o, p float64) float64 {
	return (R+1)*u/alpha/alpha - (R+1)*o/alpha/alpha + G*(1-1/R)*v/alpha - G/R*p/alpha
}

// Calculation of Energies

func KE(u, v, o, p float64) float64 {
	return .5 * l * l * ((m+M)*v*v + M*(p*p+2*v*p*(1-(u-o)*(u-o))))
}

func PE(u, o float64) float64 {
	return g * l * ((m+M)*(1+.5*u*u) - M*(1-.5*o*o))
}
