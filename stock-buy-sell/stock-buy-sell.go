package stockbuysell

import "fmt"

// returns the instructions on what days to buy the stock and when to sell it
func getDates(days []int) string {
	highestProfit := 0
	highestProfitBuy := 0
	highestProfitSell := 0

	for i, day := range days {
		for k := i; k < len(days); k++ {
			profit := days[k] - day
			if profit > highestProfit {
				highestProfit = profit
				highestProfitBuy = i + 1  // storing 1 based, not array 0 based
				highestProfitSell = k + 1 // storing 1 based, not array 0 based
			}
		}
	}
	return fmt.Sprintf("buy on day %v, sell on day %v", highestProfitBuy, highestProfitSell)
}
