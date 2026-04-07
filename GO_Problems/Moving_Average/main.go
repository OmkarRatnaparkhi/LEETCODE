/*
Problem Statement : We need to design a system that returns the moving average of the last k elements from a stream of integers.
*/

/*
I use a sliding window approach where I maintain the last k elements and a running sum to compute the average efficiently
Data Structure Design
type MovingAverage struct {
	window []int
	size   int
	sum    int
}

window → stores last k elements
size → maximum window size (k)
sum → running sum of elements in window

Using a running sum avoids recalculating sum every time, reducing time complexity
*/

/*
Step-by-Step Logic of Next(val int)
Step 1: Add new value
	m.window = append(m.window, val)
	m.sum += val
	Whenever a new value comes, I add it to the window and update the running sum

Step 2: Maintain window size
	if len(m.window) > m.size {
		m.sum = m.sum - m.window[0]		//m.sum = m.sum - oldestValue
		m.window = m.window[1:]
	}
	If the window exceeds size k, I remove the oldest element to maintain only the last k elements.
	m.window[0] = oldest element
	Subtract it from sum before removing

Step 3: Calculate average
return float64(m.sum) / float64(len(m.window))
Finally, I compute the average using the current window size. Initially, it may be less than k
*/

/*
Dry Run
Step 1: Next(1)
	window = [1]
	sum = 1
	avg = 1 / 1 = 1.0
Step 2: Next(3)
	window = [1, 3]
	sum = 4
	avg = 4 / 2 = 2.0
Step 3: Next(5)
	window = [1, 3, 5]
	sum = 9
	avg = 9 / 3 = 3.0
Step 4: Next(7)
	window = [1, 3, 5, 7] → exceeds size
	remove 1
	Final:
	window = [3, 5, 7]
	sum = 15
	avg = 15 / 3 = 5.0
Step 5: Next(9)
	window = [3, 5, 7, 9] → exceeds size
	remove 3
	Final:
	window = [5, 7, 9]
	sum = 21
	avg = 21 / 3 = 7.0
*/

/*
Time Complexity: O(1) per operation
👉 because we only add/remove one element
Space Complexity: O(k)
👉 window stores at most k elements
*/

package main

import "fmt"

type MovingAverage struct {
	window []int
	size   int
	sum    int
}

// Method to process next value
func (m *MovingAverage) Next(val int) float64 {
	// Add new value
	m.window = append(m.window, val)
	m.sum += val

	// Remove oldest if window exceeds size
	if len(m.window) > m.size {
		m.sum -= m.window[0]
		m.window = m.window[1:]
	}

	// Return average
	return float64(m.sum) / float64(len(m.window))
}

func main() {
	// Direct initialization (no constructor)
	m := MovingAverage{
		window: []int{},
		size:   3,
		sum:    0,
	}

	fmt.Println(m.Next(1)) // 1.0
	fmt.Println(m.Next(3)) // 2.0
	fmt.Println(m.Next(5)) // 3.0
	fmt.Println(m.Next(7)) // 5.0
	fmt.Println(m.Next(9)) // 7.0
}