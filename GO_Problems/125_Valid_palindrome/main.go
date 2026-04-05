/*
125. Valid Palindrome
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric 
characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
Example 1:
	Input: s = "A man, a plan, a canal: Panama"
	Output: true
	Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:
	Input: s = "race a car"
	Output: false
	Explanation: "raceacar" is not a palindrome.
Example 3:
	Input: s = " "
	Output: true
	Explanation: s is an empty string "" after removing non-alphanumeric characters.
	Since an empty string reads the same forward and backward, it is a palindrome.
*/
package main 

import "fmt"

func isPalindrome(str string) bool {
	// Step 1: clean string
	nstr := []byte{}
	for i := 0; i<len(str); i++ {
		c:=str[i]
		
		if (c >='a' && c <= 'z') || (c >='A' && c <= 'Z') || (c >='0' && c <= '9') {
			if c >='A' && c <= 'Z' {
				c = c + ('a'-'A')
			}
			nstr = append(nstr, c)
		}
			
	}
	
	// Step 2: check palindrome
	left := 0
	right := len(nstr) -1
	for left < right {
		if nstr[left] != nstr[right] {
			return false
		} 
		left++
		right--
		
	}
	return true
}

func main() {
	// Test cases
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "

	fmt.Println(isPalindrome(s1)) // true
	fmt.Println(isPalindrome(s2)) // false
	fmt.Println(isPalindrome(s3)) // true
}