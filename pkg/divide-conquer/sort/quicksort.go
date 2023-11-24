package main

import (
	"fmt"
)

func QuickSort[T comparable](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr) - 1

	// Pick a pivot
	pivotIndex := len(arr) / 2

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		// Need to define a way to compare elements of type T
		if less(arr[i], arr[right]) {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	QuickSort(arr[:left])
	QuickSort(arr[left + 1:])

	return arr
}

// less needs to be defined for comparing elements of type T
func less[T comparable](a, b T) bool {
	// Must be implemented based on the type T's characteristics
	// For example, if T is an int, then we can do:
	// return a < b

	return true
}

func main() {
	// Test with a specific type, like []int
	sample := []int{33, 21, 45, 64, 55, 34, 11, 8, 3, 2, 1}
	fmt.Println("Before sort: ", sample)

	sample = QuickSort(sample)
	fmt.Println("After sort: ", sample)
}


