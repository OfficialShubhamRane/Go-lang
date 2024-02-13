package main

import (
	"fmt"
)

func main()  {
	var x int;
	var y int;
	var oper string;
	fmt.Println("enter 1 st number")
	fmt.Scan(&x)
	fmt.Println("enter 2 nd number")
	fmt.Scan(&y)
	fmt.Println("choose operation fomr + - / * %")
	fmt.Scan(&oper)

	if (oper == "+") {
		fmt.Println(x+y)
	}else if (oper == "-") {
		fmt.Println(x-y)
	}else if (oper == "*") {
		fmt.Println(x*y)
	}else if (oper == "/") {
		fmt.Println(x/y)
	}else if (oper == "%") {
		fmt.Println(x%y)
	}


}