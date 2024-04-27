package main

import "fmt"

func maxProfit(prices []int, i int) (int, int) {
	n := len(prices)

	costs := make([]int, n)
	profits := make([]int, n)

	costs[0] = prices[0]
	numTrx := 0
	maxProfit := 0

	for trx := 0; trx < i; trx++ {
		for day := 1; day < n; day++ {
			costs[day] = min(costs[day-1], prices[day]-profits[day])
			profits[day] = max(profits[day-1], prices[day]-costs[day])
		}

		if maxProfit != profits[n-1] {
			numTrx++
		}

		maxProfit = profits[n-1]
	}

	return numTrx, maxProfit
}

func main() {
	pricesByDay := []int{1, 2, 6, 5, 2, 3, 6, 1, 3}
	maxTransaction := 4

	numTrx, maxProfit := maxProfit(pricesByDay, maxTransaction)
	fmt.Println("Max Profit:", maxProfit)
	fmt.Println("Number of transactions:", numTrx)
}
