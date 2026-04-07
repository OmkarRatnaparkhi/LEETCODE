/*
Geeksforgeeks
Given an array arr[] of daily stock prices, the stock span for the i-th day is the count of consecutive days up to and 
including day i, such that each of those days had a stock price less than or equal to the price on day i.
Example 1:
	Input: arr[] = [100, 80, 60, 120]
	Output: [1, 1, 1, 4]
	Explanation: For 100, there are no previous higher prices, so span = 1. For 80 and 60, each is smaller than the previous, 
	so their spans remain 1. For 120, it is greater than all earlier prices (100, 80, 60), so the span extends back across all
	four days, giving span = 4.
Example 2:
	Input: arr[] = [10, 4, 5, 90, 120, 80]
	Output: [1, 1, 2, 4, 5, 1]
	Explanation: For 10 and 4, no earlier prices are smaller, so span = 1 each. For 5, it is greater than 4, so span = 2. 
	For 90, it is greater than 10, 5, and 4, so span = 4. For 120, it is greater than all previous prices, giving span = 5. 
	Finally, 80 is smaller than 120, so span = 1.
*/

package main

/*
Brute Force Approach
	For each day i:
	Go backwards (j = i-1)
	Count until you find a price greater than current
Complexity
Time: ❌ O(n²)
Space: ✅ O(1)
*/
func stockSpanBrute(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		span := 1

		for j := i - 1; j >= 0; j-- {
			if arr[j] <= arr[i] {
				span++
			} else {
				break
			}
		}

		result[i] = span
	}

	return result
}

/*
Optimal Stack Approach
Idea (VERY IMPORTANT) Use monotonic decreasing stack (store indices)
	Pop all smaller/equal elements
	Find nearest greater element on left
	Calculate span

⚡ Complexity
Time: ✅ O(n) (each element pushed/popped once)
Space: O(n)	
*/
/*
| Operation | Code                           | Meaning    |
| --------- | ------------------------------ | ---------- |
| Push      | `stack = append(stack, ch)`    | Add to top |
| Peek      | `stack[len(stack)-1]`          | See top    |
| Pop       | `stack = stack[:len(stack)-1]` | Remove top |
*/
func stockSpan(arr []int) []int {
	n := len(arr)
	result := make([]int, n)

	stack := []int{} // stores indices

	for i := 0; i < n; i++ {
		// pop all smaller or equal elements
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		// calculate span
		if len(stack) == 0 {
			result[i] = i + 1
		} else {
			result[i] = i - stack[len(stack)-1] //Peek
		}

		// push current index
		stack = append(stack, i) //Push
	}

	return result
}
//At first glance, it looks like O(n²) because of nested loops, but it’s actually O(n). Let’s break it down simply.
//The inner for loop does NOT run fully every time Because: Each element is pushed once and popped once only
//So total operations: O(n) + O(n) = O(n)
//Although there are nested loops, each element is pushed and popped at most once, so total operations are linear, 
//making time complexity O(n).
func main() {
	
}