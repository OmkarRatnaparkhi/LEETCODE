/*
Rearrange an array in maximum minimum form in O(1) extra space
Given a sorted array of positive integers, rearrange the array alternately i.e first element should be the maximum value, 
second minimum value, third-second max, fourth-second min and so on. 
Examples:
	Input: arr[] = {1, 2, 3, 4, 5, 6, 7} 
	Output: arr[] = {7, 1, 6, 2, 5, 3, 4}
	Explanation: First 7 is the max value, then 1 is the min value, then 6 is the second max value, then 2 is the second 
	min value, then 5 is the third max value, then 3 is the third min value and at last 4 is the fourth max value.
Examples:
	Input: arr[] = {1, 2, 3, 4, 5, 6} 
	Output: arr[] = {6, 1, 5, 2, 4, 3}
*/

/*
Given sorted array: [1, 2, 3, 4, 5, 6, 7]
You need:			[max, min, 2nd max, 2nd min, ...]
Result:				[7, 1, 6, 2, 5, 3, 4]

Use two pointers:
	left → smallest element
	right → largest element

Alternate picking:
	right → left → right → left ...
	
Time: O(n)	Single pass
Space: O(n)	Extra array used
*/

package main

import "fmt"

func rearrange(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	left := 0
	right := n - 1
	idx := 0

	for left <= right {
		if idx < n {
			result[idx] = arr[right]
			right--
			idx++
		}

		if idx < n {
			result[idx] = arr[left]
			left++
			idx++
		}
	}

	return result
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arr2 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(rearrange(arr1)) // [7 1 6 2 5 3 4]
	fmt.Println(rearrange(arr2)) // [6 1 5 2 4 3]
}