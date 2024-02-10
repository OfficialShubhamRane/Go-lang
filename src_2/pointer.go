// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int
	x := new(int)

	fmt.Println(x)  // pointers own address
	fmt.Println(&x) // address that pointer holds
	fmt.Println(*x) // data at the address that pointer holds

	// ptr = *x
	*x = 3

	fmt.Println(x)  // pointers own address
	fmt.Println(&x) // address that pointer holds
	fmt.Println(*x) // data at the address that pointer holds

	// fmt.Println(ptr)  // pointers own address
	// fmt.Println(&ptr) // address that pointer holds
	// fmt.Println(*ptr) // data at the address that pointer holds

}
