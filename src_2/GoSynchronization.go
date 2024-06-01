package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var n int
	fmt.Print("Enter the number of integers: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter integer #%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Original array:", arr)

	// Divide the array into 4 parts
	partSize := len(arr) / 4
	part1 := arr[:partSize]
	part2 := arr[partSize : 2*partSize]
	part3 := arr[2*partSize : 3*partSize]
	part4 := arr[3*partSize:]

	var wg sync.WaitGroup
	wg.Add(4)

	go sortSubarray(part1, &wg)
	go sortSubarray(part2, &wg)
	go sortSubarray(part3, &wg)
	go sortSubarray(part4, &wg)

	wg.Wait()

	// Merge the sorted subarrays
	mergeSortedArrays(arr, part1, part2, part3, part4)

	fmt.Println("Sorted array:", arr)
}

func sortSubarray(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sorting subarray:", arr)
	sort.Ints(arr)
}

func mergeSortedArrays(arr []int, part1, part2, part3, part4 []int) {
	i, j, k, l := 0, 0, 0, 0

	for i < len(part1) && j < len(part2) && k < len(part3) && l < len(part4) {
		if part1[i] <= part2[j] && part1[i] <= part3[k] && part1[i] <= part4[l] {
			arr[i+j+k+l] = part1[i]
			i++
		} else if part2[j] <= part1[i] && part2[j] <= part3[k] && part2[j] <= part4[l] {
			arr[i+j+k+l] = part2[j]
			j++
		} else if part3[k] <= part1[i] && part3[k] <= part2[j] && part3[k] <= part4[l] {
			arr[i+j+k+l] = part3[k]
			k++
		} else {
			arr[i+j+k+l] = part4[l]
			l++
		}
	}

	for i < len(part1) {
		arr[i+j+k+l] = part1[i]
		i++
	}
	for j < len(part2) {
		arr[i+j+k+l] = part2[j]
		j++
	}
	for k < len(part3) {
		arr[i+j+k+l] = part3[k]
		k++
	}
	for l < len(part4) {
		arr[i+j+k+l] = part4[l]
		l++
	}
}
