package main

import (
	"fmt"
	"math"
)

type testCase struct {
	List       []float64
	Func       func(b bool) float64
	E0         float64
	Stopping_h float64
	Stopping_i int
	Stopped    bool
	LineX      []float64
	LineY      []float64
}

func Stability() {
	// Cache h
	ho := h

	// Set dE limit
	dElim := 0.05

	fmt.Println("> Starting Stability Test")
	fmt.Println("> Algorithms run until values deviate more than", dElim, "away from first value")

	// Step through varying values of h to plot graph E vs h in log10 scale
	h0 := 0.00001
	hf := 1.0
	Nh := 2000
	dh := (math.Log10(hf) - math.Log10(h0)) / float64(Nh)
	Nh++

	h_ := make([]float64, Nh)
	E_E := &testCase{
		List: make([]float64, Nh),
		Func: Euler,
	}
	E_L := &testCase{
		List: make([]float64, Nh),
		Func: Leapfrog,
	}
	E_R := &testCase{
		List: make([]float64, Nh),
		Func: RK4,
	}

	// test tests if the next E value has deviated enough from the initial E value
	var test = func(i int, tcase *testCase) {
		newE := tcase.Func(false)
		tcase.List[i] = math.Log10(newE)

		// If not yet marked as stopped, and dE above limit
		if !tcase.Stopped {
			if math.Abs(newE-tcase.E0) > dElim {
				tcase.Stopped = true
				tcase.Stopping_h = h
				tcase.Stopping_i = i
				tcase.LineX = make([]float64, 100)
				tcase.LineY = make([]float64, 100)

				// Plot vertical line at this limit
				logh := math.Log10(h)
				logE := math.Log10(newE)
				for x := 0; x < 100; x++ {
					tcase.LineX[x] = logh
					tcase.LineY[x] = logE - (float64(x)-50.0)*0.3
				}
			}
		}
	}

	for i := 0; i < Nh; i++ {
		// Step through in log10
		h = h0 * math.Pow(10, float64(i)*dh)
		h_[i] = math.Log10(h)

		if i > 0 {
			// Clear current line
			fmt.Print("\033[2K\r")
		}
		// Announce which step of the test the loop is at
		fmt.Print("Step ", i, " of ", (Nh - 1), ": h = ", h, ", steps = ", int(tf/h), " - ")

		if i > 0 {
			fmt.Print("Euler... ")
			test(i, E_E)

			fmt.Print("Leapfrog... ")
			test(i, E_L)

			fmt.Print("Runge-Kutta-4...")
			test(i, E_R)
		} else {
			fmt.Print("Euler... ")
			E_E.E0 = Euler(false)
			E_E.List[0] = math.Log10(E_E.E0)

			fmt.Print("Leapfrog... ")
			E_L.E0 = Leapfrog(false)
			E_L.List[0] = math.Log10(E_L.E0)

			fmt.Print("Runge-Kutta-4...")
			E_R.E0 = RK4(false)
			E_R.List[0] = math.Log10(E_R.E0)
		}
	}
	fmt.Print("\n")

	Plot(
		"Evh",
		fmt.Sprintf("Energy vs Δt with gamma = %v (log10-scale)", gamma),
		[]string{
			"Δt",
			"Energy / J",
			fmt.Sprintf("Euler stable until h=%.4f", E_E.Stopping_h),
			fmt.Sprintf("Leapfrog stable until h=%.4f", E_L.Stopping_h),
			fmt.Sprintf("RK4 stable until h=%.3f", E_R.Stopping_h),
		},
		h_,
		E_E.List,
		E_L.List,
		E_R.List,
		E_E.LineX,
		E_E.LineY,
		E_L.LineX,
		E_L.LineY,
		E_R.LineX,
		E_R.LineY,
	)

	// Restore original h
	h = ho
}
