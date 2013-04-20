package main

import "fmt"

func Euler(mkgraph bool) float64 {
	updateab()

	var u2, v2, E float64
	t, u, v := t0, u0, v0
	N := int((tf-t0)/h) + 1

	var tdata, udata, vdata, Kdata, Vdata, Edata []float64
	if mkgraph {
		tdata = make([]float64, N)
		udata = make([]float64, N)
		vdata = make([]float64, N)
		Kdata = make([]float64, N)
		Vdata = make([]float64, N)
		Edata = make([]float64, N)
	}

	for i := 0; i < N; i++ {
		u2 = u + fu(u, v)*h
		v2 = v + fv(u, v)*h

		if i == N-1 {
			E = KE(v) + PE(u)
		}

		if mkgraph {
			tdata[i] = t
			udata[i] = u
			vdata[i] = v
			Kdata[i] = KE(v)
			Vdata[i] = PE(u)
			Edata[i] = Kdata[i] + Vdata[i]
		}

		u, v = u2, v2
		t += h
	}

	if mkgraph {
		title := fmt.Sprintf("Euler with gamma = %v", gamma)
		Plot("euler θ", title, []string{"Time / s", "theta / rad", "theta"}, tdata, udata)
		Plot("euler ω", title, []string{"Time / s", "omega / rad/s", "omega"}, tdata, vdata)
		Plot("euler VKE", title, []string{"Time / s", "Energy / J", "Potential Energy", "Kinetic Energy", "Total Energy"}, tdata, Vdata, Kdata, Edata)
	}

	return E
}
