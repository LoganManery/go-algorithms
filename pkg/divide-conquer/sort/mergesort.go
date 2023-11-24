package main

import (
	"fmt"
)

func MergeSort(arr []int) [] int {
	if len(arr) <= 1 {
		return arr
	}

	// Divide the array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the two halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	sample := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Before sort: ", sample)

	sample = MergeSort(sample)
	fmt.Println("After sort: ", sample)
}