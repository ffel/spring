package main

import "fmt"

type num float64
type ode func([]num, num) num

func main() {
	// k, m, b are system constants
	// x, v will be calculated
	// h is interval
	// T is time
	var k, m, b, x, v, h, T num

	k = 1.0
	m = 1.0
	b = 1.0

	x = -0.5
	v = 0.0

	h = 0.25
	T = 0.0

	f_x := func(x []num, t num) num { return x[1] }
	f_v := func(x []num, t num) num { return -k*x[0]/m - b*x[1]/m }

	x_n := []num{x, v}
	f := []ode{f_x, f_v}

	fmt.Printf("%9s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	for T <= 3 {
		kk := euler(x_n, T, h, f)
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

	print(t_n, x_n, d_n)

	k = make([]num, len(x_n))
	for i, d := range d_n {
		k[i] = h * d
	}

	return k
}

func print(t_n num, x_n []num, d_n []num) {
	x_vals := ""
	for _, x := range x_n {
		x_vals += fmt.Sprintf(" %9.2f", x)
	}

	d_vals := ""
	for _, d := range d_n {
		d_vals += fmt.Sprintf(" %9.2f", d)
	}

	fmt.Printf("%9.2f%s%s\n", t_n, x_vals, d_vals)
}
