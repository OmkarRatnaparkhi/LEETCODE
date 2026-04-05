/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
Example 1:
	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: The subarray [4,-1,2,1] has the largest sum 6.
Example 2:
	Input: nums = [1]
	Output: 1
	Explanation: The subarray [1] has the largest sum 1.
Example 3:
	Input: nums = [5,4,-1,7,8]
	Output: 23
	Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Maximum Subarray Sum - Kadane's Algorithm
Given an integer array arr[], find the subarray (containing at least one element) which has the maximum possible sum, and return that sum.
Note: A subarray is a continuous part of an array.

Examples:
	Input: arr[] = [2, 3, -8, 7, -1, 2, 3]
	Output: 11
	Explanation: The subarray [7, -1, 2, 3] has the largest sum 11.
Examples:
	Input: arr[] = [-2, -4]
	Output: -2
	Explanation: The subarray [-2] has the largest sum -2.
Examples:
	Input: arr[] = [5, 4, 1, 7, 8]
	Output: 25
	Explanation: The subarray [5, 4, 1, 7, 8] has the largest sum 25.

*/	

/*
subarray = n*(n+1)/2
//https://www.youtube.com/watch?v=9IZYqostl2M 
*/

package main

import "fmt"

/*Brute
| Complexity | Value |
| ---------- | ----- |
| Time       | O(n²)  |
| Space      | O(1)  |
*/
func maxSubArrayBrute(nums []int) int {
	n := len(nums)
	maxSum := nums[0]

	for i := 0; i < n; i++ {
		currSum := 0

		for j := i; j < n; j++ {
			currSum = currSum + nums[j]		

			if currSum > maxSum {
				maxSum = currSum
			}
		}
	}

	return maxSum
}

/*
Kadane's Algorithm
| Complexity | Value |
| ---------- | ----- |
| Time       | O(n)  |
| Space      | O(1)  |
*/
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currSum < 0 {
			currSum = nums[i]
		} else {
			currSum = currSum + nums[i]		//currSum += nums[i]
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}

func main() {
	// LeetCode test cases
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums2 := []int{1}
	nums3 := []int{5, 4, -1, 7, 8}

	// GFG test cases
	nums4 := []int{2, 3, -8, 7, -1, 2, 3}
	nums5 := []int{-2, -4}
	nums6 := []int{5, 4, 1, 7, 8}
	
	fmt.Println("LeetCode Brute Cases:")
	fmt.Println(maxSubArray(nums1)) // 6
	fmt.Println(maxSubArray(nums2)) // 1
	fmt.Println(maxSubArray(nums3)) // 23

	fmt.Println("\nGFG Cases:")
	fmt.Println(maxSubArray(nums4)) // 11
	fmt.Println(maxSubArray(nums5)) // -2
	fmt.Println(maxSubArray(nums6)) // 25

	fmt.Println("\n\nLeetCode Kadane Cases:")
	fmt.Println("LeetCode Cases:")
	fmt.Println(maxSubArray(nums1)) // 6
	fmt.Println(maxSubArray(nums2)) // 1
	fmt.Println(maxSubArray(nums3)) // 23

	fmt.Println("\nGFG Cases:")
	fmt.Println(maxSubArray(nums4)) // 11
	fmt.Println(maxSubArray(nums5)) // -2
	fmt.Println(maxSubArray(nums6)) // 25
}

