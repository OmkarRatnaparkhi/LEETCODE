package main

import "fmt"

/*Problem :
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]*/

//Solution 1] Brute Force
func twoSum1(nums []int, target int) []int {
	results := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				results = append(results, i, j)
				return results
			}
		}
	}

	return results
}

//Solution 2]
/*
	nums = []int{10, 20, 30, 40, 50}
	target = 50

	numMap - map[KeyType]ValueType
	map[10]0
	map[20]1
	map[30]2
	map[40]3
	map[50]4

	for 0 index and value 10 numMap[50-10] = 40 and j = 3 available in map

*/
func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int) // Step 1: Create a map to store numbers and their indices.
	// Step 2: Add each number and its index to the map.
	for i, num := range nums {
		numMap[num] = i
	}
	// Step 3: Check for each number's complement in the map.
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok && i != j {
			return []int{i, j} // Step 4: Return the indices if the complement exists.
		}
	}
	// Step 5: If no complement is found, return nil.
	return nil
}

func main() {
	nums := []int{10, 20, 30, 40, 50}
	result := twoSum2(nums, 50)
	fmt.Println(result)
}
