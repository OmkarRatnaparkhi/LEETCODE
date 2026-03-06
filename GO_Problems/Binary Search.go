package main

import (
	"fmt"
	"slices"
	"sort"
)

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
	items := []int{1, 11, 4, 5, 8, 9, 3}
	target := 11

	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})

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
