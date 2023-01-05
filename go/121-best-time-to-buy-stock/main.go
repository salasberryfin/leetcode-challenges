/*
Set a default buy price = first day and a default sell price = 0
For each price:
	1. Can I buy it cheaper? If so, update buy price
	2. Is this price-my buy price > my sell price? If so, update max sell price
*/
package main

func maxProfit(prices []int) int {
	lenPrices := len(prices)

	if lenPrices == 1 {
		return 0
	}

	buy := prices[0]
	sell := 0

	for _, price := range prices {
		if price < buy {
			buy = price
		}
		if price-buy > sell {
			sell = price - buy
		}
	}

	return sell
}
