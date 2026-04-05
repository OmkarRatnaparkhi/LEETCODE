package main

/*
49. Group Anagrams
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	Explanation:
	There is no string in strs that can be rearranged to form "bat".
	The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
	The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
Example 2:
	Input: strs = [""]
	Output: [[""]]
Example 3:
	Input: strs = ["a"]
	Output: [["a"]]


Constraints:
	1 <= strs.length <= 104
	0 <= strs[i].length <= 100
	strs[i] consists of lowercase English letters.
*/

import (
	"fmt"
	"slices"
)

/*
| Complexity | Value          |
| ---------- | -------------- |
| Time       | O(n * k log k) |
| Space      | O(n * k)       |
*/
func anagram(words []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		runes := []rune(word)

		// sort.Slice(runes, func(i, j int) bool {
		// 	return runes[i] < runes[j]
		// })
		slices.Sort(runes)

		key := string(runes)

		anagramMap[key] = append(anagramMap[key], word)
	}

	var result [][]string
	for _, val := range anagramMap {
		result = append(result, val)
	}

	return result
}

func main() {
	words := []string{"eat", "ate", "bat", "tan", "tea", "nat"}
	fmt.Println(anagram(words))
	fmt.Println()

	word := "eat"
	runes := []rune(word)
	bubbleSort(runes)
	fmt.Println("bubble sort: ", string(runes)) // aet
}

func bubbleSort(arr []rune) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				//swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
