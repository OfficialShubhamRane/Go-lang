package main

import (
	"fmt"
)

func main() {
	fmt.Println("Into maps")

	/** Using built in method make()*/
	// marks := make(map[string]string)
	// marks["Shubham"] = "80"
	// marks["Vengu"] = "80"
	// marks["Rahul"] = "80"
	// fmt.Println(marks)

	/** Without using make()*/
	marks := map[string]string{"Bhramesh": "80"}
	marks["Shubham"] = "79"
	marks["Vengu"] = "80"
	marks["Rahul"] = "80"
	fmt.Println(marks)

	fmt.Println(len(marks))       //length
	fmt.Println(marks["Shubham"]) //retriving specific value

}
