package main

import (
	"fmt"
)

func main() {

	/** FizzBuzz problem with for loop */
	fizzC := 1
	buzzC := 1
	for i := 1; i <= 100; i++ {
		if fizzC == 3 && buzzC == 5 {
			fmt.Println("FizzBuzz")
			fizzC = 0
			buzzC = 0
		} else if fizzC == 3 {
			fmt.Println("Fizz")
			fizzC = 0
		} else if buzzC == 5 {
			fmt.Println("Buzz")
			buzzC = 0
		} else {
			fmt.Println(i)
		}
		fizzC++
		buzzC++
	}

}
