package main

import "fmt"

func main() {
	k := 1.0
	m := 1.0
	b := 1.0

	x := -0.5
	v := 0.0
	dt := 0.25

	fmt.Printf("%9s %9s %9s %9s %9s\n", "t", "x", "v", "x'", "v'")

	for i := 0; i < 19; i++ {
		// je kan hier denk ik closures van maken die je aan een
		// iterator meegeeft
		dxdt := v
		dvdt := -k*x/m - b*v/m
		fmt.Printf("%9.2f %9.2f %9.2f %9.2f %9.2f\n", float64(i)*dt, x, v, dxdt, dvdt)
		x = x + dt*dxdt
		v = v + dt*dvdt
	}
}
