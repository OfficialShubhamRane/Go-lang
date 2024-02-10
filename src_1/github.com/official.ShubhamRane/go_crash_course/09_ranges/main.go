package main

import (
	"fmt"
)

var fruits = []string{"Apple", "Mango", "Banana", "Strawberry", "Orange", "Jackfruit"}

var countryCapitals = map[string]string{"France": "Paris",
	"Germany": "Berlin",
	"Britain": "London",
	"India":   "New Delhi"}

func main() {
	fmt.Println("Into Ranges")

	/**Implementing for each loop using range*/
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	/** Range with Maps*/
	for k, v := range countryCapitals {
		fmt.Printf("%s's capital is %s. \n", k, v)
	}
}
