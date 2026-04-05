/*
BOTH PROBLEMS ARE SAME
Leetcode Problem Statement]
645. Set Mismatch
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of
the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another
number.
You are given an integer array nums representing the data status of this set after the error.
Find the number that occurs twice and the number that is missing and return them in the form of an array.
Example 1:
	Input: nums = [1,2,2,4]
	Output: [2,3]
Example 2:
	Input: nums = [1,1]
	Output: [1,2]
Constraints:
	2 <= nums.length <= 104
	1 <= nums[i] <= 104

GeeksforGeeks Problem Statement]
Given an unsorted array arr[] of size n, containing elements from the range 1 to n, it is known that one number in
this range is missing, and another number occurs twice in the array, find both the duplicate number and the missing number.
Example1:
	Input: arr[] = [3, 1, 3]
	Output: [3, 2]
	Explanation: 3 is occurs twice and 2 is missing.
Example2:
	Input: arr[] = [4, 3, 6, 2, 1, 1]
	Output: [1, 5]
	Explanation: 1 is occurs twice and 5 is missing.
*/

/*
Brute Force Approach
For every number from 1 → n:
Count how many times it appears in the array

Logic
	If count == 2 → duplicate
	If count == 0 → missing

Steps
	Loop from i = 1 to n
	For each i, scan whole array
	Count occurrences
	Identify:
	duplicate
	missing

| Type  | Value       |
| ----- | ----------- |
| Time  | O(n²) ❌	  |
| Space | O(1)        |

*/

package main

import (
	"fmt"
)

func findErrorNumsBrute(nums []int) []int {
	n := len(nums)

	duplicate, missing := -1, -1

	// Check each number from 1 to n
	for i := 1; i <= n; i++ {
		count := 0

		// Count occurrences of i
		for j := 0; j < n; j++ {
			if nums[j] == i {
				count++
			}
		}

		if count == 2 {
			duplicate = i
		} else if count == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

/*
Better Approach using Hashing
Count how many times each number appears, then identify:

	the one that appears twice → duplicate
	the one that appears zero times → missing

Why this works
Given: Numbers are in range 1 to n

	Array size = n

So ideally: Each number should appear exactly once
But due to error:

	One number appears twice
	One number appears zero times

Data Structure Used : freq := make([]int, n+1)

	Because numbers go from 1 to n

So we use:

	index → number
	freq[i] → how many times i appears

Example:
nums = [1,2,2,4]
freq = [0,1,2,0,1]

	↑ ↑ ↑ ↑
	1 2 3 4

| Type  | Value |
| ----- | ----- |
| Time  | O(n)  |
| Space | O(n)  |
*/
func findErrorNums(nums []int) []int {
	n := len(nums)
	freq := make([]int, n+1)

	// count frequency
	for _, num := range nums {
		freq[num]++
	}

	duplicate, missing := -1, -1

	for i := 1; i <= n; i++ {
		if freq[i] == 2 {
			duplicate = i
		}
		if freq[i] == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

func main() {

	nums := []int{1, 2, 2, 4}
	fmt.Println(nums)
	fmt.Println(findErrorNumsBrute(nums))
	fmt.Println()
	nums1 := []int{1, 2, 2, 4}
	fmt.Println(nums1)
	fmt.Println(findErrorNums(nums1))

	fmt.Println()
	fmt.Println()

	nums = []int{4, 3, 6, 2, 1, 1}
	fmt.Println(nums)
	fmt.Println(findErrorNumsBrute(nums))
	fmt.Println()
	nums1 = []int{4, 3, 6, 2, 1, 1}
	fmt.Println(nums1)
	fmt.Println(findErrorNums(nums1))
}
