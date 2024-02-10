// package main

// import "fmt"

// func main(){
// 	var floatNum float64

// 	fmt.Println("Enter floating point number");
// 	fmt.Scan(&floatNum);
	
// 	fmt.Print( int(floatNum))
	
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var num float64
// 	fmt.Print("Please enter a floating point number: ")
// 	fmt.Scan(&num)
// 	fmt.Printf("%.0f", math.Trunc(num))
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter a floating point number:")

	var num float64

	fmt.Scanln(&num)

	var y int = int(num)

	fmt.Println(y)

}
