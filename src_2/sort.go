package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	numbers := make([]int, 0, 10)

	fmt.Println("Enter up to 10 integers (press enter after each integer):")
	for i := 0; i < 10; i++ {
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer.")
			continue
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Original numbers:", numbers)

	BubbleSort(numbers)

	fmt.Println("Sorted numbers:", numbers)
}

func Swap(i int, j int, numbers []int) {
	numbers[j], numbers[i] = numbers[i], numbers[j]
}

func BubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if numbers[i] > numbers[j] {
				Swap(i, j, numbers)
			}
		}
	}
}

func bubbleSort() {
	sli := []int{23, 65, 87, 34, 12}

	for i := 0; i < len(sli); i++ {
		for j := i; j < len(sli); j++ {
			if sli[i] > sli[j] {
				fmt.Printf("swapping [%d] and [%d]\n", sli[i], sli[j])
				temp := sli[i]
				sli[i] = sli[j]
				sli[j] = temp
			}
		}
	}

	for i := 0; i < len(sli); i++ {
		fmt.Println(sli[i])
	}
}
