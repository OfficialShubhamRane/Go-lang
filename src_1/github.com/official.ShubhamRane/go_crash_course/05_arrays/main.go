package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inside arrays")

	/** Simple declarations and initialization */
	var fruitArr [2]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Mango"

	/** Shorthand method */
	// fruitArr := [2]string{"Apple", "Mango"}
	fmt.Println(fruitArr)

	fmt.Println("\nInside Slices")

	/** Simple slice declaration and initialization*/
	// var fruitSlice []string
	// fruitSlice[0] = "Banana"
	// fruitSlice[1] = "Watermelon"
	// fruitSlice[2] = "Pear"

	/** Shorthand method*/
	fruitSlice := []string{"Banana", "Watermelon", "Pear", "Grapes"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:3])

	/** Using built-in make(T, len, cap(opt)) mothod*/
	// var fruitSlice []string
	// fruitSlice = make([]string, 0, 10)
	// fruitSlice[0] = "Orange"
	// fruitSlice[1] = "Guava"
	// fruitSlice[2] = "Jackfruit"

	// fmt.Println(fruitSlice)

}
