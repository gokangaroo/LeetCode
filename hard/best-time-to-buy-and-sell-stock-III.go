package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{2, 1}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sell1 := 0
	sell2 := 0
	buy1 := math.MinInt32
	buy2 := math.MinInt32
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}
	return sell2
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
