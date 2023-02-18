package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
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
