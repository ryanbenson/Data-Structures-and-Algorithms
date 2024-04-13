# 122. Best Time to Buy and Sell Stock II

[122. Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii)

You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.

```
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
```

## Thoughts

The actual logic and approach of this one wasn't that hard. It was actually the examples that clouded your judgement if you read them over. It's pushing you to try to constantly look for all future possibilities for the most profit. So for the second example, it says to buy day 1, sell day 5, to get the profit of 4. But in reality, you should be buying the moment you see profit, and in reality, you see the same results. That drastically alters the approach, and significantly simplifies the work itself.

Once you piece that together, it's quite simple. In my opinion, this was the most misleading Leetcode so far, and not a fan of it. You can technically do it the other way, but it becomes needlessly complex in comparison.
