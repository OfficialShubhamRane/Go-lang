package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	var stringValue string

	fmt.Println("enter a string"); // ask for input

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // keep reading regardless of whitespaces
    stringValue = scanner.Text()

	// fmt.Println(stringValue);

	stringValue = strings.ToLower(stringValue) // convert everything to lower case for ease of comparison

	if strings.HasPrefix(stringValue, "i") && strings.HasSuffix(stringValue, "n") && strings.Contains(stringValue, "a") { // compare
		fmt.Println("Found");
	}else{
		fmt.Println("Not Found");
	}
}

// package main

// import "fmt"
// import s "strings"

// func main() {
// 	var s_str string
// 	var str string
// 	var has_it int

//     _, err := fmt.Scan(&s_str)
//     has_it = 0

//     if err != nil {
//         fmt.Println("error scanning value:", err)
//     } else {
// 	    str = s.ToUpper(s_str)

//       	if s.HasPrefix(str, "I") {
//         	if s.HasSuffix(str, "N") {
//         		if s.Contains(str, "A") {
//         			has_it = 1
//         		}
//         	}
//         }  
//     }

//     if has_it == 1 {
//     	fmt.Println("Found!")
//     } else {
//     	fmt.Println("Not Found!")
//     }
// }