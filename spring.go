package main

import (
	"fmt"

	"github.com/ffel/ode"
)

func main() {
	var k, m, b ode.Num

	k = 1
	m = 1
	b = 0.4

	dxdt := func(xx []ode.Num, t ode.Num) ode.Num { v := xx[1]; return v }
	dvdt := func(xx []ode.Num, t ode.Num) ode.Num { x, v := xx[0], xx[1]; return -k*x/m - b*v/m }

	odes := []ode.Ode{dxdt, dvdt}

	fmt.Printf("#%8s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	print(ode.FixedStep(ode.Euler, odes, []ode.Num{-0.5, 0}, 0, 15, 0.25), odes)
	fmt.Println("")
	print(ode.AdaptiveStep(ode.Euler, odes, []ode.Num{-0.5, 0}, 0, 15, 0.01, 0.5), odes)
	fmt.Println("")
	print(ode.FixedStep(ode.MidPoint, odes, []ode.Num{-0.5, 0}, 0, 15, 0.25), odes)
	fmt.Println("")
	print(ode.AdaptiveStep(ode.MidPoint, odes, []ode.Num{-0.5, 0}, 0, 15, 0.01, 0.5), odes)
	fmt.Println("")
	print(ode.FixedStep(ode.Rk4, odes, []ode.Num{-0.5, 0}, 0, 15, 0.25), odes)
	fmt.Println("")
	print(ode.AdaptiveStep(ode.Rk4, odes, []ode.Num{-0.5, 0}, 0, 15, 0.01, 0.5), odes)
}

func print(data []ode.Result, dxdt []ode.Ode) {
	// assume system with two ode
	for _, d := range data {
		fmt.Printf("%9.3f %9.3f %9.3f %9.3f %9.3f\n",
			d.T, d.XX[0], d.XX[1], dxdt[0](d.XX, d.T), dxdt[1](d.XX, d.T))
	}
}
