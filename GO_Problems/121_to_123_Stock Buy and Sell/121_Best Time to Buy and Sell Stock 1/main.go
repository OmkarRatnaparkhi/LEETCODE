/*
You are given an array prices where prices[i] is the price of a given stock on the i th day. You want to maximize your profit 
by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
Example 1:
	Input: prices = [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
	Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:
	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transactions are done and the max profit = 0.
Constraints:
1 <= prices.length <= 105
0 <= prices[i] <= 104

Geeks for Geeks ] Stock Buy and Sell - Max one Transaction Allowed
Given an array prices[] of non-negative integers, representing the prices of the stocks on different days, find the maximum 
profit possible by buying and selling the stocks on different days when at most one transaction is allowed. Here one 
transaction means 1 buy + 1 Sell. If it is not possible to make a profit then return 0.
Note: Stock must be bought before being sold.
Example 1:
	Input: prices[] = [7, 10, 1, 3, 6, 9, 2]
	Output: 8
	Explanation: Buy for price 1 and sell for price 9. 
Example 2:
	Input: prices[] = [7, 6, 4, 3, 1]
	Output: 0
	Explanation: Since the array is sorted in decreasing order, 0 profit can be made without making any transaction.
Example 3:
	Input: prices[] = [1, 3, 6, 9, 11]
	Output: 10
	Explanation: Since the array is sorted in increasing order, we can make maximum profit by buying at price[0] and 
	selling at price[n-1]
*/

//Time Complexity: O(n), where n is the number of days. We only loop through the array once.
//Space Complexity: O(1), We only store two variables (minPrice and maxProfit) regardless of how big the input array is.
//The Logic: You can't sell a stock before you buy it. By keeping track of the minPrice, you are essentially saying, "If I were 
//to sell today, what is the best price I could have bought it at in the past?"

/*
DRY RUN 1] 
prices = [7, 1, 5, 3, 6, 4]
minPrice = 7
maxProfit = 0

Iterations from index 1
i = 1, value = 1
	minPrice = min(7, 1) = 1
	profit = 1 - 1 = 0
	maxProfit = max(0, 0) = 0

Iterations from index 2	
i = 2, value = 5
	minPrice = min(1, 5) = 1
	profit = 5 - 1 = 4
	maxProfit = max(0, 4) = 4
	
Iterations from index 3
i = 3, value = 3
	minPrice = min(1, 3) = 1
	profit = 3 - 1 = 2
	maxProfit = max(4, 2) = 4

Iterations from index 4
i = 4, value = 6
	minPrice = min(1, 6) = 1
	profit = 6 - 1 = 5
	maxProfit = max(4, 5) = 5

Iterations from index 5
i = 5, value = 4
	minPrice = min(1, 4) = 1
	profit = 4 - 1 = 3
	maxProfit = max(5, 3) = 5


*/

package main

import (
    "fmt"
)
//[7, 1, 5, 3, 6, 4]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {

		// update minimum price
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		// calculate profit
		profit := prices[i] - minPrice

		// update max profit
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

/*
Brute Force  : Brute force checks all buy-sell pairs (O(n²)), while optimal tracks minimum price in one pass to get max profit (O(n))

1. Time Complexity: O(n^2)
This is considered inefficient for large datasets because of the nested loops.
	The Outer Loop (i): Runs n times (where n is the number of prices).
	The Inner Loop (j): Runs n-1 times, then n-2, then n-3, and so on.
	Total Operations: The total number of comparisons is roughly n(n-1)/2.
	In Big O terms: We drop the constants and lower-order terms, leaving us with O(n^2).
What this means in practice: If you have 1,000 prices, the loop runs about 500,000 times. If you have 100,000 prices, 
the loop runs about 5 billion times, which will likely cause your program to time out or run very slowly.

2) Space Complexity: O(1)
This is the "silver lining" of the brute force approach.
	Memory Usage: You are only storing a few integer variables (n, maxProfit, i, j, and profit).
	Scaling: Regardless of whether the input slice has 10 elements or 10 million elements, the amount of extra memory used by your 
	function stays exactly the same.
	In Big O terms: This is Constant Space, denoted as O(1).

for every i:
  for every j > i:
    profit = prices[j] - prices[i]
    update maxProfit

[7, 1, 5, 3, 6, 4]

i = 0 (buy = 7)
	j = 1 → profit = 1 - 7 = -6 → maxProfit = 0
	j = 2 → profit = 5 - 7 = -2 → maxProfit = 0
	j = 3 → profit = 3 - 7 = -4 → maxProfit = 0
	j = 4 → profit = 6 - 7 = -1 → maxProfit = 0
	j = 5 → profit = 4 - 7 = -3 → maxProfit = 0

i = 1 (buy = 1)
	j = 2 → profit = 5 - 1 = 4 → maxProfit = 4
	j = 3 → profit = 3 - 1 = 2 → maxProfit = 4
	j = 4 → profit = 6 - 1 = 5 → maxProfit = 5
	j = 5 → profit = 4 - 1 = 3 → maxProfit = 5

i = 2 (buy = 5)
	j = 3 → profit = 3 - 5 = -2 → maxProfit = 5
	j = 4 → profit = 6 - 5 = 1 → maxProfit = 5
	j = 5 → profit = 4 - 5 = -1 → maxProfit = 5

i = 3 (buy = 3)
	j = 4 → profit = 6 - 3 = 3 → maxProfit = 5
	j = 5 → profit = 4 - 3 = 1 → maxProfit = 5

i = 4 (buy = 6)
	j = 5 → profit = 4 - 6 = -2 → maxProfit = 5

i = 5 (buy = 4)
(no j left)
maxProfit = 5
*/
func maxProfitBruteForce(prices []int) int {
	maxProfit := 0

	for i := 0; i < len(prices); i++ { // buy day
		for j := i + 1; j < len(prices); j++ { // sell day
			profit := prices[j] - prices[i]

			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	// Define the slices
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{7, 6, 4, 3, 1}

	// Calculate results
	result1 := maxProfit(prices1)
	result2 := maxProfit(prices2)

	// Print results with comments for clarity
	fmt.Println(result1) // Output: 5
	fmt.Println(result2) // Output: 0

	fmt.Println()
	result1 = maxProfit2(prices1)
	result2 = maxProfit2(prices2)
	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Println()
	result1 = maxProfit2(prices1)
	result2 = maxProfit2(prices2)
	fmt.Println(result1)
	fmt.Println(result2)
	
	fmt.Println()
	prices := []int{4, 5, 1, 6, 7, 8, 9}
	result := maxProfitBruteForce(prices)
	fmt.Println(result)
	result = maxProfit3(prices)
	fmt.Println(result)
}
