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
