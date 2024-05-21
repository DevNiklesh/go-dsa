// https://takeuforward.org/data-structure/stock-buy-and-sell-dp-35/

package main

import (
	"fmt"
	"math"
)

func maxProfit(nums []int) int {
	maxProfit := math.MinInt

	buy := nums[0]
	for i := 1; i < len(nums); i++ {
		profit := nums[i] - buy
		if profit > maxProfit {
			maxProfit = profit
		}

		if nums[i] < buy {
			buy = nums[i]
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
