package main

import (
	"fmt"
	"sort"
)

func anagram(words []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		runes := []rune(word)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		key := string(runes)

		anagramMap[key] = append(anagramMap[key], word)
	}

	var result [][]string
	for _, val := range anagramMap {
		result = append(result, val)
	}

	return result
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

func main() {
	words := []string{"eat", "ate", "bat", "tan", "tea", "nat"}
	fmt.Println(anagram(words))
	fmt.Println()
	word := "eat"
	runes := []rune(word)
	bubbleSort(runes)
	fmt.Println("bubble sort: ", string(runes)) // aet
}
