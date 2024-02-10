package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	var stringValue string

	fmt.Println("enter a string");

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    stringValue = scanner.Text()

	fmt.Println(stringValue);

	stringValue = strings.ToLower(stringValue)

	if strings.HasPrefix(stringValue, "i") && strings.HasSuffix(stringValue, "n") && strings.Contains(stringValue, "a") {
		fmt.Println("Found");
	}else{
		fmt.Println("Not Found");
	}
}