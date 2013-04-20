package main

import "fmt"

func Leapfrog(mkgraph bool) float64 {
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

	// First step: Euler
	u_, v_ := u, v
	u = u_ + fu(u_, v_)*h
	v = v_ + fv(u_, v_)*h
	if mkgraph {
		tdata[0] = t
		udata[0] = u_
		vdata[0] = v_
		Kdata[0] = KE(v_)
		Vdata[0] = PE(u_)
		Edata[0] = Kdata[0] + Vdata[0]
	}
	t += h

	// Subsequent steps use leapfrog
	for i := 1; i < N; i++ {
		u2 = u_ + 2.0*fu(u, v)*h
		v2 = v_ + 2.0*fv(u, v)*h

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

		// NOTE: Order important.
		u_, v_ = u, v
		u, v = u2, v2
		t += h
	}

	if mkgraph {
		title := fmt.Sprintf("Leapfrog with gamma = %v", gamma)
		Plot("leapfrog θ", title, []string{"Time / s", "theta / rad", "theta"}, tdata, udata)
		Plot("leapfrog ω", title, []string{"Time / s", "omega / rad/s", "omega"}, tdata, vdata)
		Plot("leapfrog VKE", title, []string{"Time / s", "Energy / J", "Potential Energy", "Kinetic Energy", "Total Energy"}, tdata, Vdata, Kdata, Edata)
	}

	return E
}
