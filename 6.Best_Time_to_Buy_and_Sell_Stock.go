package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	currentMin := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < currentMin {
			currentMin = prices[i]
		} else if (prices[i] - currentMin) > maxProfit {
			maxProfit = prices[i] - currentMin
		}
	}
	return maxProfit
}

func main() {
	prices := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

}
