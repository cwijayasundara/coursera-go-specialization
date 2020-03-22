package main

import (
	"fmt"
)

func main() {
	// s =½ a t2 + vot + so
	accelaration := readInputValues("acceleration:")
	velocity := readInputValues("initial velocity:")
	initialDisplacement := readInputValues("initial displacement:")
	fn := GenDisplaceFn(accelaration, velocity, initialDisplacement)
	t := readInputValues("time:")
	fmt.Println(fn(t))
}

func readInputValues(msg string) float64 {
	fmt.Print("pls enter ", msg)
	var result float64
	fmt.Scanln(&result)
	return result
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}
