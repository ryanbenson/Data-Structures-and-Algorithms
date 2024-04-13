# 121. Best Time to Buy and Sell Stock

[121. Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

```
Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

## Thoughts

This isn't too bad. There are multiple ways of doing this, one of which is the sliding window. It does in the end causing it to be longer to process.

There's another approach where you essentially loop through the list, keep updating the lowest possible item. If that entry isn't the lowest item, then you price compare and keep track of the best margin you find along the way. This ends up pretty good since it's:

> time complexity of this solution is O(n) and space complexity is O(1)
