/*
Leetcode: Design an algorithm that collects daily price quotes for some stock and returns the span of that stock's price for the current 
day. The span of the stock's price in one day is the maximum number of consecutive days (starting from that day and going 
backward) for which the stock price was less than or equal to the price of that day. For example, if the prices of the stock 
in the last four days is [7,2,1,2] and the price of the stock today is 2, then the span of today is 4 because starting from 
today, the price of the stock was less than or equal 2 for 4 consecutive days. Also, if the prices of the stock in the last 
four days is [7,34,1,2] and the price of the stock today is 8, then the span of today is 3 because starting from today, 
the price of the stock was less than or equal 8 for 3 consecutive days.
Implement the StockSpanner class:
	StockSpanner() Initializes the object of the class.
	int next(int price) Returns the span of the stock's price given that today's price is price.
Example 1:
Input ["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
		[[], [100], [80], [60], [70], [60], [75], [85]]
Output [null, 1, 1, 1, 2, 1, 4, 6]

Explanation
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // return 1
stockSpanner.next(80);  // return 1
stockSpanner.next(60);  // return 1
stockSpanner.next(70);  // return 2
stockSpanner.next(60);  // return 1
stockSpanner.next(75);  // return 4, because the last 4 prices (including today's price of 75) were less than or equal to today's price.
stockSpanner.next(85);  // return 6
 
Constraints:
	1 <= price <= 105
	At most 104 calls will be made to next.


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

func main() {
	
}