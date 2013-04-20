package main

import "fmt"

func RK4(mkgraph bool) float64 {
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

	var ku1, ku2, ku3, ku4, kv1, kv2, kv3, kv4 float64

	for i := 0; i < N; i++ {
		ku1 = h * fu(u, v)
		kv1 = h * fv(u, v)
		ku2 = h * fu(u+ku1/2.0, v+kv1/2.0)
		kv2 = h * fv(u+ku1/2.0, v+kv1/2.0)
		ku3 = h * fu(u+ku2/2.0, v+kv2/2.0)
		kv3 = h * fv(u+ku2/2.0, v+kv2/2.0)
		ku4 = h * fu(u+ku3, v+kv3)
		kv4 = h * fv(u+ku3, v+kv3)
		u2 = u + (ku1+2.0*ku2+2.0*ku3+ku4)/6.0
		v2 = v + (kv1+2.0*kv2+2.0*kv3+kv4)/6.0

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
		title := fmt.Sprintf("RK4 with gamma = %v", gamma)
		Plot("rk4 θ", title, []string{"Time / s", "theta / rad", "theta"}, tdata, udata)
		Plot("rk4 ω", title, []string{"Time / s", "omega / rad/s", "omega"}, tdata, vdata)
		Plot("rk4 VKE", title, []string{"Time / s", "Energy / J", "Potential Energy", "Kinetic Energy", "Total Energy"}, tdata, Vdata, Kdata, Edata)
	}

	return E
}
