package main

/*
The Problem Statement: "Partitioning Two Colors"

Given an array containing only two types of elements (for example, "G" for Green and "R" for Red), write a function to rearrange the 
elements in-place so that all "G" elements appear first, followed by all "R" elements.

Constraints & Requirements:Time Complexity: O(n) (You must solve it in a single pass).
Space Complexity: O(1)$
(You must modify the array in-place without using extra storage).Stability: The relative order of the same-colored balls doesn't strictly matter here, but the swap-based approach is a standard "partitioning" logic used in algorithms like Quicksort.
*/


package main

import "fmt"


//[R, G, R, G, G, R]
func rearrange(arr []string) {
	left := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == "G" {
	
			// 1. Save the value of arr[left] in a temporary holder ---> temp := arr[left]
			// 2. Move the value of arr[i] into arr[left]---> arr[left] = arr[i]
			// 3. Move the saved temp value into arr[i] ---> arr[i] = temp
		
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
}
/*
| Step | i | arr[i] | left | Action            | Array State        |
| ---- | - | ------ | ---- | ----------------- | ------------------ |
| 1    | 0 | R      | 0    | Do nothing        | [R, G, R, G, G, R] |
| 2    | 1 | G      | 0    | Swap(1,0), left++ | [G, R, R, G, G, R] |
| 3    | 2 | R      | 1    | Do nothing        | [G, R, R, G, G, R] |
| 4    | 3 | G      | 1    | Swap(3,1), left++ | [G, G, R, R, G, R] |
| 5    | 4 | G      | 2    | Swap(4,2), left++ | [G, G, G, R, R, R] |
| 6    | 5 | R      | 3    | Do nothing        | [G, G, G, R, R, R] |

Initial:
i=0, left=0
	[R, G, R, G, G, R]

i=1 → found G → swap with left=0
	[G, R, R, G, G, R]
	left→1

i=3 → found G → swap with left=1
	[G, G, R, R, G, R]
	left→2

i=4 → found G → swap with left=2
	[G, G, G, R, R, R]
	left→3
	
Steps:

i=1 → "G" → swap with left(0) → [G, R, R, G, G, R]
i=3 → "G" → swap with left(1) → [G, G, R, R, G, R]
i=4 → "G" → swap with left(2) → [G, G, G, R, R, R]

*/



func rearrangeBrute(arr []string) {
	countG := 0

	// Step 1: Count G
	for i := 0; i < len(arr); i++ {
		if arr[i] == "G" {
			countG++
		}
	}

	// Step 2: Fill G
	for i := 0; i < countG; i++ {
		arr[i] = "G"
	}

	// Step 3: Fill R
	for i := countG; i < len(arr); i++ {
		arr[i] = "R"
	}
}
/*
Dry Run (Step-by-Step)
Step 1: Count "G"
| i | arr[i] | countG |
| - | ------ | ------ |
| 0 | R      | 0      |
| 1 | G      | 1      |
| 2 | R      | 1      |
| 3 | G      | 2      |
| 4 | G      | 3      |
| 5 | R      | 3      |
Final countG = 3


Step 2: Fill "G"
| i | Action     | Array              |
| - | ---------- | ------------------ |
| 0 | arr[0] = G | [G, G, R, G, G, R] |
| 1 | arr[1] = G | [G, G, R, G, G, R] |
| 2 | arr[2] = G | [G, G, G, G, G, R] |

*/


func main() {
	arr := []string{"R", "G", "R", "G", "G", "R"}
	fmt.Println(arr)
	rearrange(arr)
	fmt.Println(arr)
	
	fmt.Println()
	arr1 := []string{"R", "G", "R", "G", "G", "R"}
	fmt.Println(arr1)
	rearrangeBrute(arr1)
	fmt.Println(arr1)
}


/*
TIME and SPACE Complexity
Brute Force (Counting Approach)
Code Flow Recap
	First loop → count "G"
	Second loop → fill "G"
	Third loop → fill "R"
Time Complexity (Detailed)
Let n = len(arr)
Step-by-step:
	Loop 1 → runs n times → O(n)
	Loop 2 → runs countG times → worst case n → O(n)
	Loop 3 → runs n - countG times → worst case n → O(n)
O(n) + O(n) + O(n) = O(3n)
Even though there are 3 loops, they are sequential (not nested), so complexity remains linear.

Space Complexity
We are using:
	countG → integer variable → O(1)
	No extra array or memory used.
	Space = O(1) (constant space)
We are overwriting entire array
Even if already sorted, still doing full work

Optimal Approach (2-Pointer / Partition)
Code Flow Recap
	Single loop
	Swap whenever "G" is found
Loop:
	Runs from i = 0 → n-1
	Each iteration does:
	Comparison → O(1)
	Swap (if needed) → O(1)
n iterations × O(1) work = O(n)

Space Complexity
	We are using:
	left pointer → O(1)
	Temporary variable during swap → O(1)
	Space = O(1)
	
| Feature          | Brute Force                   | Optimal (2-pointer)   |
| ---------------- | ----------------------------- | --------------------- |
| Passes           | 3                             | 1                     |
| Time Complexity  | O(n)                          | O(n)                  |
| Space Complexity | O(1)                          | O(1)                  |
| Writes           | Always overwrite entire array | Only swap when needed |
| Efficiency       | Slightly less efficient       | More efficient        |
| Interview Value  | Basic                         | Preferred             |

Brute force takes O(n) time with 3 passes, while optimal uses a single pass with two pointers. Both use O(1) space, but the 
optimal approach is more efficient due to fewer operations and better in-place handling
*/
