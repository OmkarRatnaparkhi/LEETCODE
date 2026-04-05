/*
344. Reverse String
Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.
Example 1:
	Input: s = ["h","e","l","l","o"]
	Output: ["o","l","l","e","h"]
Example 2:
	Input: s = ["H","a","n","n","a","h"]
	Output: ["h","a","n","n","a","H"]
Constraints:
	1 <= s.length <= 105
	s[i] is a printable ascii character.
*/

/*
Time Complexity: O(n), where $n$ is the length of the string. We visit each character exactly once.
Space Complexity: O(1). We only use two integer variables (left and right) regardless of how large the input array is.

Pattern : Two-Pointer Pattern.
How the Two-Pointer Pattern works here : In the string reversal problem, you are using the Opposite Ends variation of this 
pattern:
	Placement: One pointer starts at the beginning (left := 0) and the other at the end (right := len(s) - 1).
	Movement: They move toward each other (left++ and right--).
	Action: At each step, you perform an operation (the swap) on the elements at those two positions.
	Termination: The loop stops when the pointers meet or cross (left < right becomes false).
*/

package main

import (
	"fmt"
)

func reverseString(str []byte) {
	
	left := 0
	right := len(str)-1
	
	for left < right {
		str[left], str[right] = str[right],str[left]
		left++
		right--
	}
	
}

func main() {
	// Example 1
	input1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Printf("Input 1:  %s\n", string(input1))
	reverseString(input1)
	fmt.Printf("Output 1: %s\n\n", string(input1))

	// Example 2
	input2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Printf("Input 2:  %s\n", string(input2))
	reverseString(input2)
	fmt.Printf("Output 2: %s\n", string(input2))
}