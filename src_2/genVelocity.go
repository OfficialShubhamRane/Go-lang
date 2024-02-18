package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + vo*t + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Println("Enter acceleration:")
	fmt.Scanln(&a)

	fmt.Println("Enter initial velocity:")
	fmt.Scanln(&vo)

	fmt.Println("Enter initial displacement:")
	fmt.Scanln(&so)

	fn := GenDisplaceFn(a, vo, so)

	fmt.Println("Enter time:")
	fmt.Scanln(&t)

	fmt.Println("Displacement after", t, "seconds:", fn(t))
}
