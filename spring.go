package main

import "fmt"

// see http://jordanorelli.com/post/42369331748/function-types-in-go-golang

// num as short for float64
type num float64

// the spring is described by two first order ode
type ode func([]num, num) num

// euler and midpoint are integrators
type integrator func([]num, num, num, []ode) []num

func main() {
	var k, m, b num

	k = 1
	m = 1
	b = 1

	dxdt := func(xx []num, t num) num { v := xx[1]; return v }
	dvdt := func(xx []num, t num) num { x, v := xx[0], xx[1]; return -k*x/m - b*v/m }

	odes := []ode{dxdt, dvdt}

	fixed_step(euler, odes, -0.5, 0, 3, 0.25)
	fixed_step(midpoint, odes, -0.5, 0, 3, 0.25)
}

func fixed_step(method integrator, dxdt []ode, x0, v0, tmax, h num) {
	var x, v, T num

	x = x0
	v = v0

	T = 0.0

	x_n := []num{x, v}

	fmt.Printf("%9s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	for T <= tmax {
		kk := method(x_n, T, h, dxdt)

		fmt.Printf("%9.3f %9.2f %9.2f %9.2f %9.2f\n",
			T, x_n[0], x_n[1], kk[0]/h, kk[1]/h)

		for i, k := range kk {
			x_n[i] += k
		}

		T += h
	}

}

func euler(x_n []num, t_n, h num, dxdt []ode) (k []num) {
	d_n := make([]num, len(x_n))
	for i, f := range dxdt {
		d_n[i] = f(x_n, t_n)
	}

	k = make([]num, len(x_n))
	for i, d := range d_n {
		k[i] = h * d
	}

	return k
}

func midpoint(x_n []num, t_n, h num, dxdt []ode) (k []num) {
	d_n := make([]num, len(x_n))

	for i, f := range dxdt {
		d_n[i] = f(x_n, t_n)
	}

	x_2 := make([]num, len(x_n))

	for i, x := range x_n {
		x_2[i] = x + d_n[i]*h/2
	}

	for i, f := range dxdt {
		d_n[i] = f(x_2, t_n+h/2)
	}

	k = make([]num, len(x_n))
	for i, d := range d_n {
		k[i] = h * d
	}

	return k
}
