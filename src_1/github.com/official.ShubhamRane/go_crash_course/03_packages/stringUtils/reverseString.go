package stringutils

import (
	"fmt"
)

func reverseString(s string) string {
	fmt.Println("Attempting reversal of Stirng")

	strArr := []rune(s)
	fmt.Println("Length of an array", len(strArr))

	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return string(strArr)

}
