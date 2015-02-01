package main

import "fmt"

func main() {
	k := 1.0
	m := 1.0
	b := 1.0

	x := -0.5
	v := 0.0
	h := 0.25

	T := 0.0

	fmt.Printf("%9s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	for T <= 3 {

		// je kan hier denk ik closures van maken die je aan een
		// iterator meegeeft
		dxdt := v
		dvdt := -k*x/m - b*v/m
		fmt.Printf("%9.2f %9.2f %9.2f %9.2f %9.2f\n", T, x, v, dxdt, dvdt)
		x = x + h*dxdt
		v = v + h*dvdt
		T += h
	}

	x = -0.5
	v = 0.0
	T = 0.0

	f_x := func(x, v, t float64) float64 { return v }
	f_v := func(x, v, t float64) float64 { return -k*x/m - b*v/m }

	for T <= 3 {
		k1, k2 := euler(x, v, T, h, f_x, f_v)
		x += k1
		v += k2
		T += h
	}

	x = -0.5
	v = 0.0
	T = 0.0

	f_x2 := func(x []float64, t float64) float64 { return x[1] }
	f_v2 := func(x []float64, t float64) float64 { return -k*x[0]/m - b*x[1]/m }

	x_n := []float64{x, v}
	fx_n := []func(x_n []float64, t float64) float64{f_x2, f_v2}

	for T <= 3 {

		kk := euler2(x_n, T, h, fx_n)
		for i, k := range kk {
			x_n[i] += k
		}
		T += h
	}
}

func euler(x1_n, x2_n, t_n, h float64, dx1dt, dx2dt func(x1, x2, t float64) float64) (k1, k2 float64) {
	delta_x1 := dx1dt(x1_n, x2_n, t_n)
	delta_x2 := dx2dt(x1_n, x2_n, t_n)
	fmt.Printf("=== %9.2f %9.2f %9.2f %9.2f %9.2f\n", t_n, x1_n, x2_n, delta_x1, delta_x2)
	return h * delta_x1, h * delta_x2
}

func euler2(x_n []float64, t_n, h float64, dxdt []func(x_n []float64, t float64) float64) (k []float64) {
	delta := make([]float64, len(x_n))
	for i, fun := range dxdt {
		delta[i] = fun(x_n, t_n)
	}

	x_vals := ""
	for _, x := range x_n {
		x_vals += fmt.Sprintf(" %9.2f", x)
	}

	delta_vals := ""
	for _, d := range delta {
		delta_vals += fmt.Sprintf(" %9.2f", d)
	}

	fmt.Printf("--- %9.2f%s%s\n", t_n, x_vals, delta_vals)

	increments := make([]float64, len(x_n))
	for i, d := range delta {
		increments[i] = h * d
	}

	return increments
}
