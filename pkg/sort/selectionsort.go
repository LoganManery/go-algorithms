package main

import (
	"fmt"
)

func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// find the minimum element in the unsorted part of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the first elemebt of the unsorted part
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	sample := []int{64, 25, 12, 22, 11}
	fmt.Println("Original array:", sample)

	SelectionSort(sample)
	fmt.Println("Sorted array:  ", sample)
}