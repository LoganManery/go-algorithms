package main

import (
	"fmt"
)

// BinarySearch function takes a sorted array and a target value
// It returns the index of the target value if found, or -1 if not found
func BinarySearch(array []int, target int) int {
	// Define starting and ending index
	low := 0
	high := len(array) - 1

	// While the range is valid
	for low <= high {
		// find the middle index
		mid := low + (high - low) / 2

		// check if the middle element is our target
		if array[mid] == target {
			return mid
		}

		// if the target is greater, ignore the left half
		if array[mid] < target {
			low = mid + 1
		} else {
			// if the target is smaller, ignore the right half
			high = mid - 1
		}
	}
	// target is not present in the array
	return -1
}

func main() {
	array := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	target := 9

	result := BinarySearch(array, target)

	if result != -1 {
		fmt.Println("Element found at index", result)
	} else {
		fmt.Println("Element not found in the array")
	}
}