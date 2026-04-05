/*
LEETCODE] 122. Best Time to Buy and Sell Stock II
You are given an integer array prices where prices[i] is the price of a given stock on the ith day. On each day, you may 
decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can sell 
and buy the stock multiple times on the same day, ensuring you never hold more than one share of the stock.
Find and return the maximum profit you can achieve.
Example 1:
	Input: prices = [7,1,5,3,6,4]
	Output: 7
	Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
	Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
	Total profit is 4 + 3 = 7.
Example 2:
	Input: prices = [1,2,3,4,5]
	Output: 4
	Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
	Total profit is 4.
Example 3:
	Input: prices = [7,6,4,3,1]
	Output: 0
	Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.
Constraints:
	1 <= prices.length <= 3 * 104
	0 <= prices[i] <= 104
	
geeksforgeeks] Given an array prices[] representing stock prices, find the maximum total profit that can be earned by 
buying and selling the stock any number of times.
Note: We can only sell a stock which we have bought earlier and we cannot hold multiple stocks on any day. 
Examples:
	Input: prices[] = [100, 180, 260, 310, 40, 535, 695]
	Output: 865
	Explanation: Buy the stock on day 0 and sell it on day 3 = 310 - 100 = 210 and Buy the stock on day 4 and sell it on
	day 6 = 695 - 40 = 655 so the Maximum Profit  is = 210 + 655 = 865.
*/

/*
You are given stock prices for each day.
Day:    1   2   3   4   5   6
Price: [7,  1,  5,  3,  6,  4]

Rule 1: Only 1 stock at a time : You must SELL before you BUY again
Rule 2: You can buy & sell many times : Unlimited transactions allowed
Rule 3: You cannot hold multiple stocks : No buying 2 stocks together

Make maximum profit by: Buying stock & Selling stock
buy and/or sell on same day
It means:
	You can sell and then buy again immediately
	But still: At any moment → only 1 stock allowed

You want to capture every price increase

Input: [7,1,5,3,6,4]
7 → 1 (price drops)				Do nothing
1 → 5 (price increases)			Buy at 1, Sell at 5 	Profit = 4
5 → 3 (drop)					Do nothing
3 → 6 (increase)				Buy at 3, Sell at 6		Profit = 3
6 → 4 (drop)					Do nothing
Total Profit: 4 + 3 = 7


Input: [1,2,3,4,5]
(2-1) + (3-2) + (4-3) + (5-4) = 4

Input: [100,180,260,310,40,535,695]
Input: (180-100) + (260-180) + (310-260) + (535-40) + (695-535) = 865

LOGIC : 
If tomorrow price > today price → earn profit
Else → skip

Pattern : Greedy Pattern
Time Complexity: O(n) → single pass through array
Space Complexity: O(1) → no extra space used

A Greedy algorithm always makes the best choice at the current step without worrying about the future.
*/

package main 

import (
    "fmt"
)

func maxProfit(prices []int) int {
	profit:=0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			  profit = profit + (prices[i] - prices[i-1])
		}
	}
	return profit
}

func main() {
	st1:=[]int{7,1,5,3,6,4}
	fmt.Println("Maximum Profit: ", maxProfit(st1))
	st2:=[]int{1,2,3,4,5}
	fmt.Println("Maximum Profit: ", maxProfit(st2))
	st3:=[]int{7,6,4,3,1}
	fmt.Println("Maximum Profit: ", maxProfit(st3))
	st4:=[]int{100, 180, 260, 310, 40, 535, 695}
	fmt.Println("Maximum Profit: ", maxProfit(st4))
}