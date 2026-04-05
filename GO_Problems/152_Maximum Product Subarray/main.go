/*
Maximum Product Subarray
Given an integer array nums, find a subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer.
Note that the product of an array with a single element is the value of that element.
Example 1:
	Input: nums = [2,3,-2,4]
	Output: 6
	Explanation: [2,3] has the largest product 6.
Example 2:
	Input: nums = [-2,0,-1]
	Output: 0
	Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
Constraints:
	1 <= nums.length <= 2 * 104
	-10 <= nums[i] <= 10
	The product of any subarray of nums is guaranteed to fit in a 32-bit integer.

Maximum Product Subarray
Given an array arr[] consisting of positive, negative, and zero values, find the maximum product that can be obtained from
any contiguous subarray of arr[].
Example 1:
	Input: arr[] = [-2, 6, -3, -10, 0, 2]
	Output: 180
	Explanation: The subarray with maximum product is [6, -3, -10] with product = 6 * (-3) * (-10) = 180.
Example 2:
	Input: arr[] = [-1, -3, -10, 0, 6]
	Output: 30
	Explanation: The subarray with maximum product is [-3, -10] with product = (-3) * (-10) = 30.
Example 3:
	Input: arr[] = [2, 3, 4] 
	Output: 24 
	Explanation: For an array with all positive elements, the result is product of all elements. 

*/	


package main

import "fmt"

/*Brute
| Complexity | Value |
| ---------- | ----- |
| Time       | O(n²)  |
| Space      | O(1)  |
*/
func maxProductBrute(nums []int) int {
	n := len(nums)
	maxProd  := nums[0]

	for i := 0; i < n; i++ {
		currProd := 1

		for j := i; j < n; j++ {
			currProd = currProd * nums[j]		

			if currProd > maxProd  {
				maxProd  = currProd
			}
		}
	}

	return maxProd 
}

// -------- Main Function --------
func main() {
	// Test cases
	nums1 := []int{2, 3, -2, 4}
	nums2 := []int{-2, 0, -1}
	nums3 := []int{-2, 3, -4}
	nums4 := []int{0, 2}
	nums5 := []int{-2}
	nums6 := []int{-2, -3, -4}

	fmt.Println("Brute Force Results:")
	fmt.Println(maxProductBrute(nums1)) // 6
	fmt.Println(maxProductBrute(nums2)) // 0
	fmt.Println(maxProductBrute(nums3)) // 24
	fmt.Println(maxProductBrute(nums4)) // 2
	fmt.Println(maxProductBrute(nums5)) // -2
	fmt.Println(maxProductBrute(nums6)) // 12
}