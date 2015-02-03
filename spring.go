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

	fmt.Printf("#%8s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	// fixed_step(euler, odes, -0.5, 0, 3, 0.25)
	adaptive_step(euler, odes, -0.5, 0, 15, 0.01, 0.5)

	fmt.Println("")

	// fixed_step(midpoint, odes, -0.5, 0, 3, 0.25)
	adaptive_step(midpoint, odes, -0.5, 0, 15, 0.01, 0.5)
}

// too much tied to particular x, v problem
func fixed_step(method integrator, dxdt []ode, x0, v0, tmax, h num) {
	var x, v, T num

	x = x0
	v = v0

	T = 0.0

	x_n := []num{x, v}

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

func adaptive_step(method integrator, dxdt []ode, x0, v0, tmax, hmin, h0 num) {
	var x, v, T num

	x = x0
	v = v0

	h := h0

	T = 0.0

	x_full := []num{x, v}

	var kk_full []num

	var H num

	for T <= tmax {

		// max 5 attempts
		for a := 0; a < 5; a++ {
			x_full_tmp := make([]num, len(x_full))
			x_half_tmp := make([]num, len(x_full))

			for i, x := range x_full {
				x_half_tmp[i] = x
			}

			// hier een inner loop waarin je de berekening obv h moet worden vergeleken met de h/2
			kk_full = method(x_full, T, h, dxdt)

			// fmt.Printf("%9.3f %9.2f %9.2f %9.2f %9.2f\n",
			// 	T, x_full[0], x_full[1], kk_full[0]/h, kk_full[1]/h)

			for i, k := range kk_full {
				x_full_tmp[i] = x_full[i] + k
			}

			var kk_half []num

			for halfs := 0; halfs <= 1; halfs++ {
				kk_half = method(x_half_tmp, T, h/2, dxdt)

				for i, k := range kk_half {
					x_half_tmp[i] += k
				}
			}

			q := quality(x_full_tmp, x_half_tmp, h)

			// store h as the used value
			H = h

			// fmt.Printf("%9d %9.3f %9.3f\n", a, h, q)

			if h < hmin {
				break
			} else if q > 0.01 {
				h /= 2
			} else if q < 0.005 {
				h *= 2
				break
			} else {
				break
			}
			// fmt.Printf("%9.3f %9.3f %9.3f %9.2f %9.2f (full)\n",
			// 	T, x_full[0], x_full[1], kk_full[0]/h, kk_full[1]/h)
			// fmt.Printf("%9.3f %9.3f %9.3f %9.2f %9.2f (halfs)\n",
			// 	T, x_half[0], x_half[1], kk_half[0]/h, kk_half[1]/h)

			// I guess (x_half - x_full)/h is a reasonable guess for how accurate the bigger one is.

			// fmt.Printf("%9s %9.3f %9.3f\n", "", Delta[0]/h, Delta[1]/h)

		}

		fmt.Printf("%9.3f %9.3f %9.3f %9.2f %9.2f\n",
			T, x_full[0], x_full[1], kk_full[0]/h, kk_full[1]/h)

		//increment x_full

		T += H

		for i, k := range kk_full {
			x_full[i] += k
		}

	}
}

func quality(x_full []num, x_half []num, h num) num {
	var q num = 0

	for i, full := range x_full {
		var c num

		if diff := full - x_half[i]; diff >= 0 {
			c = diff / h
		} else {
			c = -diff / h
		}

		if c > q {
			q = c
		}
	}

	return q
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
