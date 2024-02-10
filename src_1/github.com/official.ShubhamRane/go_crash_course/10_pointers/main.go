package main

import (
	"fmt"
)

func main() {
	fmt.Println("Into Pointers")

	a := 10
	b := &a

	fmt.Println(a, b, *b, &b, *&b)
}
