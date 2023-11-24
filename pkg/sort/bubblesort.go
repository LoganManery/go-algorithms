package main

import (
	"fmt"
)

// Bubble sorts an array using the bubble sort algorithm
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	sample := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Before sort: ", sample)

	BubbleSort(sample)
	fmt.Println("After sort: ", sample)
}