/*
Leetcode]
26. Remove Duplicates from Sorted Array
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element 
appears only once. The relative order of the elements should be kept the same. Consider the number of unique elements in 
nums to be k​​​​​​​​​​​​​​. After removing duplicates, return the number of unique elements k. The first k elements of nums should contain 
the unique numbers in sorted order. The remaining elements beyond index k - 1 can be ignored.

Custom Judge: The judge will test your solution with the following code:
	int[] nums = [...]; // Input array
	int[] expectedNums = [...]; // The expected answer with correct length

	int k = removeDuplicates(nums); // Calls your implementation

	assert k == expectedNums.length;
	for (int i = 0; i < k; i++) {
		assert nums[i] == expectedNums[i];
	}
If all assertions pass, then your solution will be accepted.
Example 1:
	Input: nums = [1,1,2]
	Output: 2, nums = [1,2,_]
	Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).
Example 2:
	Input: nums = [0,0,1,1,1,2,2,3,3,4]
	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
	Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).
Constraints:
	1 <= nums.length <= 3 * 104
	-100 <= nums[i] <= 100
	nums is sorted in non-decreasing order.
	
geeksforgeeks]
Given a sorted array arr[] of size n, the goal is to rearrange the array so that all distinct elements appear at the beginning 
in sorted order. Additionally, return the length of this distinct sorted subarray.
Note: The elements after the distinct ones can be in any order and hold any value, as they don't affect the result.
Example 1: 
	Input: arr[] = [2, 2, 2, 2, 2]
	Output: [2]
	Explanation: All the elements are 2, So only keep one instance of 2.
Example 2: 
	Input: arr[] = [1, 2, 2, 3, 4, 4, 4, 5, 5]
	Output: [1, 2, 3, 4, 5]
Example 3: 
	Input: arr[] = [1, 2, 3]
	Output: [1, 2, 3]
	Explanation : No change as all elements are distinct.
*/

package main

import "fmt"

/*
| i | nums[i] | nums[idx] | Action        | idx | Array                 |
| - | ------- | --------- | ------------- | --- | --------------------- |
| 1 | 0       | 0         | skip          | 0   | [0,0,1,1,1,2,2,3,3,4] |
| 2 | 1       | 0         | idx++ → place | 1   | [0,1,1,1,1,2,2,3,3,4] |
| 3 | 1       | 1         | skip          | 1   | same                  |
| 4 | 1       | 1         | skip          | 1   | same                  |
| 5 | 2       | 1         | idx++ → place | 2   | [0,1,2,1,1,2,2,3,3,4] |
| 6 | 2       | 2         | skip          | 2   | same                  |
| 7 | 3       | 2         | idx++ → place | 3   | [0,1,2,3,1,2,2,3,3,4] |
| 8 | 3       | 3         | skip          | 3   | same                  |
| 9 | 4       | 3         | idx++ → place | 4   | [0,1,2,3,4,2,2,3,3,4] |


| Feature  | Value       |
| -------- | ----------- |
| Pattern  | Two Pointer |
| Time     | O(n)        |
| Space    | O(1)        |
| In-place | ✅ Yes       |

*/

// Pattern Identification : Two Pointer Pattern
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0 // slow pointer

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}

	return idx + 1
}

func main() {

	// ✅ Test Case 1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Println("Test 1:")
	fmt.Println("k =", k1)
	fmt.Println("Array =", nums1[:k1])
	fmt.Println("Array =", nums1)
	fmt.Println()

	// ✅ Test Case 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Println("Test 2:")
	fmt.Println("k =", k2)
	fmt.Println("Array =", nums2[:k2])
	fmt.Println("Array =", nums2)
	fmt.Println()

	// ✅ Test Case 3 (All same)
	nums3 := []int{2, 2, 2, 2, 2}
	k3 := removeDuplicates(nums3)
	fmt.Println("Test 3:")
	fmt.Println("k =", k3)
	fmt.Println("Array =", nums3[:k3])
	fmt.Println("Array =", nums3)
	fmt.Println()

	// ✅ Test Case 4 (Already unique)
	nums4 := []int{1, 2, 3, 4, 5}
	k4 := removeDuplicates(nums4)
	fmt.Println("Test 4:")
	fmt.Println("k =", k4)
	fmt.Println("Array =", nums4[:k4])
	fmt.Println("Array =", nums4)
	fmt.Println()

	// ✅ Test Case 5 (Single element)
	nums5 := []int{10}
	k5 := removeDuplicates(nums5)
	fmt.Println("Test 5:")
	fmt.Println("k =", k5)
	fmt.Println("Array =", nums5[:k5])
	fmt.Println("Array =", nums5)
	fmt.Println()

	// ✅ Test Case 6 (Empty array)
	nums6 := []int{}
	k6 := removeDuplicates(nums6)
	fmt.Println("Test 6:")
	fmt.Println("k =", k6)
	fmt.Println("Array =", nums6)
}