// https://takeuforward.org/data-structure/stock-buy-and-sell-dp-35/

package main

import "fmt"

func bestTimeToBuySellStock(nums []int) (int, []int) {
	maxProfit := 0
	buyIdx, sellIdx := 0, 0

	for i := 1; i < len(nums); i++ {
		profit := nums[i] - nums[buyIdx]

		if nums[i] < nums[buyIdx] {
			buyIdx = i
		}

		if profit > maxProfit {
			maxProfit = profit
			sellIdx = i
		}
	}

	return maxProfit, []int{buyIdx, sellIdx}
}

func main() {
	fmt.Println(bestTimeToBuySellStock([]int{7, 1, 5, 3, 6, 4}))
}
