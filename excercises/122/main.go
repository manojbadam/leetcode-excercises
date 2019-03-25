package main

import (
	"fmt"
)

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += (prices[i+1] - prices[i])
		}
	}
	return profit
}

//https://play.golang.org/p/YhONsicXnNS
