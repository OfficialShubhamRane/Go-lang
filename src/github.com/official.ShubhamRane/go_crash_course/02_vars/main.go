package main

import "fmt"

var topic string = "My Info"

func main() {

	fmt.Println(topic)
	const name string = "Shubham Rane"
	var age int = 23
	// var age = 0;
	// age := 0;

	email, gender, age := "shubham@gmail.com", "Male", 24

	fmt.Println(name, age, email, gender)
	fmt.Printf("%T %T", name, age) // %T return type of variable

}
