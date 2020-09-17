# Stock Buy Sell

Reference: [cassidoo newsletter](https://cassidoo.co/newsletter/) 🎉 (amazing newsletter by the way)

> Given an array of numbers that represent stock prices (where each number is the price for a certain day), find 2 days when you should buy and sell your stock for the highest profit. 

Example:

```console
stockBuySell([110, 180, 260, 40, 310, 535, 695])
“buy on day 4, sell on day 7”
```

## Thoughts

Was a fun exercise, and I kept in mind that it was unnecessary to evaluate the list of days when there's only one day left. Overall, it was thought exercise on how to simplify the comparison and record of what days were optimal while iterating the loop again and again. Overall the performance is pretty good too.

```console
BenchmarkTestGetDates-12    	 6287911	       178 ns/op
```