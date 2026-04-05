/*
Problem Statement 
Set Matrix Zeroes : Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1
 

Follow up:
A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?

*/

/*
Brute First Approach
Whenever you find a 0, immediately make:
	its entire row = 0
	its entire column = 0
Problem with direct approach If you directly convert to 0, you will:
	accidentally affect future cells
	create chain reactions
	👉 So we need a temporary marker
Use a special value (like -1) to mark cells that should become 0 later
Condition: Matrix should NOT already contain -1 (otherwise conflict)

Steps
Traverse matrix
	If matrix[i][j] == 0
	mark entire row and column as -1 (only if not already 0)
After traversal:
	convert all -1 → 0
	
Dry Run (Quick)
1 1 1
1 0 1
1 1 1

1 -1 1
-1 0 -1
1 -1 1

1 0 1
0 0 0
1 0 1

| Type  | Value                           |
| ----- | -----------------------------   |
| Time  | O(m × n × (m + n)) ❌ slow  |
| Space | O(1)                            |

*/
package main

import "fmt"

func setZeroesBrute(matrix [][]int)  {

    m := len(matrix)
    n := len(matrix[0])

	// Step 1: Mark with -1
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++{
            if matrix[i][j] == 0 {
                for row := 0; row < m; row++ {			// mark column
                    if matrix[row][j] != 0{
                        matrix[row][j] = -1
                    }
                }

                for col := 0; col < n; col++ {			// mark row
                    if matrix[i][col] != 0 {
                        matrix[i][col] = -1
                    }
                }
            }
        } 
    }

	// Step 2: Convert -1 to 0
    for i := 0; i < m; i ++ {
        for j := 0; j < n; j++{
            if matrix[i][j] == -1 {
                 matrix[i][j] = 0
            }
        } 
    }

}


/*

| Type  | Value      |
| ----- | ---------- |
| Time  | O(m × n) ✅ |
| Space | O(m + n)   |

Instead of modifying matrix while traversing, we store which rows & columns need to be zero

Use Two Arrays	
	row[m] → marks which rows should be zero
	col[n] → marks which columns should be zero

Steps
Traverse matrix
If matrix[i][j] == 0:
	mark row[i] = 1
	mark col[j] = 1
	
Traverse again
If:
	row[i] == 1 OR col[j] == 1
	→ set matrix[i][j] = 0

I first track which rows and columns contain zero using two arrays. Then I update the matrix in a second pass based on 
those markers. This avoids modifying the matrix during traversal and keeps the logic clean with O(m+n) space.
I use two arrays to track which rows and columns need to be zero. In the first pass, I mark them. In the second pass, 
I update the matrix based on those markers. This avoids modifying the matrix during traversal
*/

func setZeroesBetter(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	row := make([]int, m)
	col := make([]int, n)

	// Step 1: Mark rows & columns
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	// Step 2: Update matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(matrix)
	setZeroesBrute(matrix)
	fmt.Println(matrix)
	
	fmt.Println()
	
	matrix1 := [][]int{
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}
	fmt.Println(matrix1)
	setZeroesBrute(matrix1)
	fmt.Println(matrix1)
	
	fmt.Println()
	/*-----------------------------------------*/
	fmt.Println()
	
	m1 := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(m1)
	setZeroesBetter(m1)
	fmt.Println(m1)
	
	fmt.Println()
	
	m2 := [][]int{
		{0,1,2,0},
		{3,4,5,2},
		{1,3,1,5},
	}
	fmt.Println(m2)
	setZeroesBrute(m2)
	fmt.Println(m2)
	
	setZeroesBetter(m1)
}



func setZeroesOptimal(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	firstRowZero := false
	firstColZero := false

	// Step 1: check first row
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	// Step 2: check first column
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	// Step 3: mark rows & columns
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Step 4: fill inner matrix
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// Step 5: first row
	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Step 6: first column
	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}


