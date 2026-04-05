package main 

import "fmt"


//Basic Palindrome Check (Two Pointer Method) 
func isPalindrome(s string) {
	left := 0
	right := len(s) - 1
		
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	
	return true	
}

func reverseString(s string) string {
	runes := []rune(s) // convert to rune for Unicode safety

	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {
	fmt.Println("isPalindrome: ", isPalindrome("madam"))  // true
	fmt.Println("isPalindrome: ", isPalindrome("hello"))  // false
	
	fmt.Println()
	fmt.Println("reverseString:", reverseString("madam"))
}