/*
Problem Statement : Longest Substring Without Repeating Characters
Given a string s, find the length of the longest substring without duplicate characters.
Example 1:
	Input: s = "abcabcbb"
	Output: 3
	Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.
Example 2:
	Input: s = "bbbbb"
	Output: 1
	Explanation: The answer is "b", with the length of 1.
Example 3:
	Input: s = "pwwkew"
	Output: 3
	Explanation: The answer is "wke", with the length of 3.
	Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Longest Substring Without Repeating Characters
Given a string s having lowercase characters, find the length of the longest substring without repeating characters. 
Example 1:
	Input: s = "geeksforgeeks"
	Output: 7 
	Explanation: The longest substrings without repeating characters are "eksforg” and "ksforge", with lengths of 7.
Example 2:
	Input: s = "aaa"
	Output: 1
	Explanation: The longest substring without repeating characters is "a"
Example 3:
	Input: s = "abcdefabcbb"
	Output: 6
	Explanation: The longest substring without repeating characters is "abcdef".

What the problem is asking
You have a string (a bunch of letters, numbers, or symbols).
You need to find the longest piece of it (substring) where no character repeats.

Think of it like this:
	You’re walking through the string, letter by letter.
	You’re allowed to collect letters until you hit a duplicate.
	Once you hit a duplicate, you stop, and that’s one substring.
	You keep checking all possible substrings and find the longest one.
	
Key words
	Substring = continuous part of the string (no skipping).
	Example: In "abcde", "abc" and "cde" are substrings. "ace" is not (because it skips).
	Without repeating characters = every letter in that substring must be unique.

Examples in plain words
1) Input: "abcabcbb"
	Start: "abc" → all unique.
	Then "abca" → has two "a"s, not allowed.
	Longest unique substring is "abc" (length 3).
	Answer: 3
2) Input: "bbbbb"
	Only "b" everywhere.
	Longest substring without repeats is just "b" (length 1).
	Answer: 1
3) Input: "pwwkew"
	"pw" → unique.
	"pww" → duplicate "w".
	"wke" → unique, length 3.
	"kew" → also unique, length 3.
	Longest is 3.
	Answer: 3
Super simple analogy
	Imagine you’re collecting cards from a deck.
	You can keep picking cards until you get a duplicate.
	Once you get a duplicate, you stop and start again.
	The goal is to find the largest set of unique cards you ever collected in one go.
*/


import (
	"fmt"
)

//Brute Force Solution
//Time Complexity: O(n²)
//Space Complexity: O(1)
func lengthOfLongestSubstringBrute(s string) int {
	n := len(s)
	maxLen := 0

	for i := 0; i < n; i++ {
		seen := make(map[byte]bool)

		for j := i; j < n; j++ {
			if seen[s[j]] {
				break // duplicate found
			}

			seen[s[j]] = true

			if j-i+1 > maxLen {
				maxLen = j - i + 1
			}
		}
	}

	return maxLen
}

/*Sliding Window
Use a map to store last seen index of characters
Maintain a window [left, right]
If duplicate found → move left
Track max length

Time Complexity: O(n)			Space Complexity: O(1)

| Step | right | char | left (before) | Action                         | left (after) | Window | maxLen |
| ---- | ----- | ---- | ------------- | ------------------------------ | ------------ | ------ | ------ |
| 1    | 0     | p    | 0             | not seen → add                 | 0            | p      | 1      |
| 2    | 1     | w    | 0             | not seen → add                 | 0            | pw     | 2      |
| 3    | 2     | w    | 0             | duplicate → move left to `1+1` | 2            | w      | 2      |
| 4    | 3     | k    | 2             | not seen → add                 | 2            | wk     | 2      |
| 5    | 4     | e    | 2             | not seen → add                 | 2            | wke    | 3      |
| 6    | 5     | w    | 2             | duplicate → move left to `2+1` | 3            | kew    | 3      |

Step 1: Initialize left := 0 maxLen := 0 map := make(map[byte]int) 
Step 2: Expand Window (Move right)		for right := 0; right < len(s); right++ 	Always move right forward
Step 3: For every character from string handle duplicate
		ch := str[right]
		if idx, found := map[s[right]]; found && idx >= left {
			left = idx + 1
		}
Step 4: Update map 	charIndex[ch] = right	// Update last seen index 
Step 5: Update Answer 
		windowLen := right - left + 1
		maxLen = max(maxLen, windowLen)

*/
func lengthOfLongestSubstring(str string) int {
	charIndex := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(str); right++ {
		ch := str[right]

		// If character already seen and inside window
		if idx, found := charIndex[ch]; found && idx >= left {
			left = idx + 1
		}

		// Update last seen index
		charIndex[ch] = right

		// Update max length
		windowLen := right - left + 1
		if windowLen > maxLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
	fmt.Println(lengthOfLongestSubstring(" "))        // Output: 1
}

