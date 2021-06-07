package main

import (
	"fmt"
	"math"
	// "E:/Go/Go-lang/src/github.com/official.ShubhamRane/go_crash_course/03_packages/stringUtils"
)

func main() {
	fmt.Println("Into Packges")
	fmt.Println(math.Floor(2.7))

	var s string = "Shubham"
	fmt.Println(len(s))
	fmt.Println(reverseString("olleh"))
}

func reverseString(s string) string {
	fmt.Println("Attempting reversal of Stirng")

	strArr := []rune(s)
	fmt.Println("Length of an array: ", len(strArr))

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)

}
