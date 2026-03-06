package main

import (
	"fmt"
	"slices"
)

// BinarySearch finds the index of the target in a sorted integer slice.
func BinarySearchManual(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		// Calculate the middle index (to avoid potential overflow with (low+high)/2 for very large slices)
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] < target {
			low = mid + 1 // Target is in the right half
		} else {
			high = mid - 1 // Target is in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	items := []int{1, 3, 4, 5, 8, 9, 11}
	target := 11

	// BinarySearch returns the index where the target is found,
	// and a boolean indicating if it was found.
	index, found := slices.BinarySearch(items, target)

	if found {
		fmt.Printf("%d found at index %d\n", target, index)
	} else {
		fmt.Printf("%d not found, would be at index %d\n", target, index)
	}

	index1 := BinarySearchManual(items, target)
	if index1 != -1 {
		fmt.Printf("%d is present at index %d\n", target, index1)
	} else {
		fmt.Printf("%d is not present in the array.\n", target)
	}
}
