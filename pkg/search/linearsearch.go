package main

import (
	"fmt"
)

func LinearSearch(array []int, target int) int {
	for i, value := range array {
		// check if the current element is the target
		if value == target {
			return i
		}
	}
	return -1
}

func main() {
	array := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 10
	result := LinearSearch(array, target)

	if result != -1 {
		fmt.Println("Element found at index", result)
	} else {
		fmt.Println("Element not found in the array")
	}
}