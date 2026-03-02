package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list. You may assume the two
numbers do not contain any leading zero, except the number 0 itself.
Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode defines the structure for a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}

// Helper function to print the linked list
func printList(n *ListNode) {
	for n != nil {
		fmt.Printf("%d", n.Val)
		if n.Next != nil {
			fmt.Print(" -> ")
		}
		n = n.Next
	}
	fmt.Println()
}

func main() {
	// Representing number 342: (2 -> 4 -> 3)
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}

	// Representing number 465: (5 -> 6 -> 4)
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	fmt.Print("Input List 1: ")
	printList(l1)
	fmt.Print("Input List 2: ")
	printList(l2)

	result := addTwoNumbers(l1, l2)

	fmt.Print("Result List:  ")
	printList(result)
	// Output should be 7 -> 0 -> 8 (which represents 807)
}

/*----------------------------------------------------*/

package main

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

func lengthOfLongestSubstring(s string) int {
	// Map to store the last index of each character
	charIndex := make(map[byte]int)

	maxLen := 0
	start := 0 // left boundary of the sliding window

	for i := 0; i < len(s); i++ {
		if idx, found := charIndex[s[i]]; found && idx >= start {
			// Move the start to one position after the repeated character
			start = idx + 1
		}
		// Update the last index of the current character
		charIndex[s[i]] = i
		// Calculate window length
		if i-start+1 > maxLen {
			maxLen = i - start + 1
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

