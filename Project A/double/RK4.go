package main

import "fmt"

func RK4(mkgraph bool) float64 {
	updateRGa()

	var u2, v2, o2, p2, E float64
	t, u, v, o, p := t0, u0, v0, o0, p0
	N := int((tf-t)/h) + 1

	var tdata, udata, vdata, odata, pdata, Kdata, Vdata, Edata []float64
	if mkgraph {
		// return 0.0
		tdata = make([]float64, N)
		udata = make([]float64, N)
		vdata = make([]float64, N)
		odata = make([]float64, N)
		pdata = make([]float64, N)
		Kdata = make([]float64, N)
		Vdata = make([]float64, N)
		Edata = make([]float64, N)
	}

	var ku1, ku2, ku3, ku4, kv1, kv2, kv3, kv4 float64
	var ko1, ko2, ko3, ko4, kp1, kp2, kp3, kp4 float64
	var k_u, k_v, k_o, k_p float64

	for i := 0; i < N; i++ {
		ku1 = h * fu(u, v, o, p)
		kv1 = h * fv(u, v, o, p)
		ko1 = h * fo(u, v, o, p)
		kp1 = h * fp(u, v, o, p)

		k_u, k_v = u+ku1/2.0, v+kv1/2.0
		k_o, k_p = o+ko1/2.0, p+kp1/2.0

		ku2, kv2 = h*fu(k_u, k_v, k_o, k_p), h*fv(k_u, k_v, k_o, k_p)
		ko2, kp2 = h*fo(k_u, k_v, k_o, k_p), h*fp(k_u, k_v, k_o, k_p)

		k_u, k_v = u+ku2/2.0, v+kv2/2.0
		k_o, k_p = o+ko2/2.0, p+kp2/2.0

		ku3, kv3 = h*fu(k_u, k_v, k_o, k_p), h*fv(k_u, k_v, k_o, k_p)
		ko3, kp3 = h*fo(k_u, k_v, k_o, k_p), h*fp(k_u, k_v, k_o, k_p)

		k_u, k_v = u+ku3, v+kv3
		k_o, k_p = o+ko3, p+kp3

		ku4, kv4 = h*fu(k_u, k_v, k_o, k_p), h*fv(k_u, k_v, k_o, k_p)
		ko4, kp4 = h*fo(k_u, k_v, k_o, k_p), h*fp(k_u, k_v, k_o, k_p)

		u2 = u + (ku1+2.0*ku2+2.0*ku3+ku4)/6.0
		v2 = v + (kv1+2.0*kv2+2.0*kv3+kv4)/6.0
		o2 = o + (ko1+2.0*ko2+2.0*ko3+ko4)/6.0
		p2 = p + (kp1+2.0*kp2+2.0*kp3+kp4)/6.0

		if mkgraph {
			tdata[i] = t
			udata[i] = u2
			vdata[i] = v2
			odata[i] = o2
			pdata[i] = p2
			Kdata[i] = KE(u, v, o, p)
			Vdata[i] = PE(u, o)
			Edata[i] = Kdata[i] + Vdata[i]
		}

		if i == N-1 {
			E = KE(u, v, o, p) + PE(u, o)
		}

		u, v, o, p = u2, v2, o2, p2
		t += h
	}

	if mkgraph {
		title := fmt.Sprintf("RK4 with R = %v, gamma = %v", R, gamma)
		Plot("rk4 θφ", title, []string{"Time / s", "Angle / rad", "theta", "phi"}, tdata, udata, odata)
		Plot("rk4 ων", title, []string{"Time / s", "Angular velocity / rad/s", "omega", "nu"}, tdata, vdata, pdata)
		Plot("rk4 E", title, []string{"Time / s", "Energy / J", "Total Energy"}, tdata, Edata)
		Plot("rk4 VKE", title, []string{"Time / s", "Energy / J", "Potential Energy", "Kinetic Energy", "Total Energy"}, tdata, Vdata, Kdata, Edata)
	}

	return E
}
