package main

// u: anglular position
// v: angular velocity

// alpha: time scale elmination factor
// a = -g / (l * alpha^2)
// b = -gamma / (m * l * alpha)

var (
	t0 = 0.0   // Initial t
	tf = 50.0 // Final t

	u0 = 0.1 // Initial θ
	v0 = 0.0 // Initial φ

	g     = 9.81 // Gravitational acceleration
	m     = 1.0  // Mass of pendulum
	l     = 1.0  // Length of pendulum
	alpha = 1.0  // Time scale
	gamma = 0.0  // Damping factor

	h = 0.001 // Step size

	a = -g / (alpha * alpha * l)
	b = -gamma / (m * alpha * l)
)

func updateab() {
	a = -g / (alpha * alpha * l)
	b = -gamma / (m * alpha * l)
}

func main() {
	// Run with
	//   γ = 0.0
	gamma = 0.0
	Euler(true)
	Leapfrog(true)
	RK4(true)
	Stability()

	// Run with
	//   γ = 0.2
	gamma = 0.2
	Euler(true)
	Leapfrog(true)
	RK4(true)
	Stability()
}

func fu(u, v float64) float64 {
	return v
}

func fv(u, v float64) float64 {
	return a*u + b*v
}

// Kinetic Energy
func KE(v float64) float64 {
	return .5 * m * l * l * v * v
}

// Potential Energy
func PE(u float64) float64 {
	return .5 * m * g * l * u * u
}
