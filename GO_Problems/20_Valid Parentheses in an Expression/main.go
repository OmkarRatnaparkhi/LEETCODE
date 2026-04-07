/*
LEETCODE Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
An input string is valid if:
	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.
	Every close bracket has a corresponding open bracket of the same type.
Example 1:
	Input: s = "()"
	Output: true
Example 2:
	Input: s = "()[]{}"
	Output: true
Example 3:
	Input: s = "(]"
	Output: false
Example 4:
	Input: s = "([])"
	Output: true
Example 5:
	Input: s = "([)]"
	Output: false
Constraints:
	1 <= s.length <= 104
	s consists of parentheses only '()[]{}'.

Geeksforgeeks
Given a string s containing three types of brackets {}, () and []. Determine whether the Expression are balanced or not.
An expression is balanced if each opening bracket has a corresponding closing bracket of the same type, the pairs are properly ordered and no bracket closes before its matching opening bracket.
	Balanced: "[()()]{}" → every opening bracket is closed in the correct order.
	Not balanced: "([{]})" → the ']' closes before the matching '{' is closed, breaking the nesting rule.

Example 1: 
	Input: s = "[{()}]"
	Output: true
	Explanation:  All the brackets are well-formed.
Example 2:
	Input:  s = "([{]})"
	Output: false
	Explanation: The expression is not balanced because there is a closing ']' before the closing '}'
*/

/*
stack - Last In, First Out (LIFO)
Key Characteristics of LIFO in Stacks
Push: Adds an item to the top of the stack.
Pop: Removes the item from the top of the stack.
Peek: Views the top item without removing it.
*/
package main

import "fmt"

/*
| Operation | Code                           | Meaning    |
| --------- | ------------------------------ | ---------- |
| Push      | `stack = append(stack, ch)`    | Add to top |
| Peek      | `stack[len(stack)-1]`          | See top    |
| Pop       | `stack = stack[:len(stack)-1]` | Remove top |
*/

func isValid(s string) bool {
	stack := []rune{}

	for _, ch := range s {
		// push opening brackets
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch) //Push - Add to top
		} else {
			// stack empty → invalid
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1] //Peek - See top

			// match manually
			if (ch == ')' && top != '(') || (ch == '}' && top != '{') || (ch == ']' && top != '[') {
				return false
			}

			// pop
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
func main() {
	fmt.Println(isValid("()"))       // true
	fmt.Println(isValid("()[]{}"))   // true
	fmt.Println(isValid("(]"))       // false
	fmt.Println(isValid("([])"))     // true
	fmt.Println(isValid("([)]"))     // false
	fmt.Println(isValid("[{()}]"))   // true
	fmt.Println(isValid("([{]})"))   // false
}