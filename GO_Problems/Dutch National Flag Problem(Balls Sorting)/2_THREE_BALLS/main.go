package main

/*
//https://leetcode.com/discuss/post/4150257/sorting-an-array-containing-only-0s-1s-a-oswb/
Problem Statement : 

Input: ["G", "B", "R", "R", "B", "G", "B"]
Goal: Sort them as [Red, Green, Blue] in $O(n)$ time and $O(1)$ space.
------------------------------------------------------------------------------------------------------

Given an array containing only 3 values (0, 1, 2)
Sort the array in one pass (O(n)) and in-place (no extra space)
------------------------------------------------------------------------------------------------------

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, 
with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function. 

Example 1:

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
Example 2:

Input: nums = [2,0,1]
Output: [0,1,2]
 

Constraints:

n == nums.length
1 <= n <= 300
nums[i] is either 0, 1, or 2.
 
*/

/*
nums = [2, 0, 2, 1, 1, 0]

Initial Pointers
low = 0
mid = 0
high = 5

Iterations 1
mid = 0, value = 2
	nums[mid] = 2
	swap(nums[mid], nums[high]) → swap(2, 0)
	nums = [0, 0, 2, 1, 1, 2]

	high-- → 4
	mid stays same

Iterations 2
mid = 0, value = 0
	nums[mid] = 0
	swap(nums[low], nums[mid]) → swap(0, 0)
	nums = [0, 0, 2, 1, 1, 2]

	low++ → 1
	mid++ → 1

Iterations 3
mid = 1, value = 0
	nums[mid] = 0
	swap(nums[low], nums[mid]) → swap(0, 0)
	nums = [0, 0, 2, 1, 1, 2]

	low++ → 2
	mid++ → 2

Iterations 4
mid = 2, value = 2
	nums[mid] = 2
	swap(nums[mid], nums[high]) → swap(2, 1)
	nums = [0, 0, 1, 1, 2, 2]

	high-- → 3
	mid stays same

Iterations 5
mid = 2, value = 1
	nums[mid] = 1
	mid++ → 3


Iterations 6
mid = 3, value = 1
	nums[mid] = 1
	mid++ → 4
	
Loop Ends
mid = 4 > high = 3 → stop

Final Output : [0, 0, 1, 1, 2, 2]

Key Pattern
0 → swap with low, move low & mid
1 → just move mid
2 → swap with high, move high only

Push 0s left, 2s right, and let 1s stay in middle
*/

/*
while mid <= high
0 → swap with low, low++, mid++
1 → mid++
2 → swap with high, high--
*/


package main

import "fmt"

//nums = [2, 0, 2, 1, 1, 0]
func sortColors(nums []int) {
	low, mid := 0, 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {

		case 0:
			// place 0 at beginning
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++

		case 1:
			// correct position
			mid++

		case 2:
			// place 2 at end
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
			// NOTE: do NOT increment mid here
		}
	}
}

func main() {
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println(nums1) // [0 0 1 1 2 2]

	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	fmt.Println(nums2) // [0 1 2]
}

