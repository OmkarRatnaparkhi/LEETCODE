package main

import "fmt"

func maxSumSubarray(arr []int, k int) int {
	windowSum, maxSum := 0, 0

	for i := 0; i < len(arr); i++ {
		windowSum += arr[i]

		if i >= k-1 {
			if windowSum > maxSum {
				maxSum = windowSum
			}
			windowSum -= arr[i-(k-1)]
		}
	}
	return maxSum
}

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Println("Max sum:", maxSumSubarray(nums, k)) // Output: 9
}
