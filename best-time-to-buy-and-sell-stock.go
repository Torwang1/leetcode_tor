package leetcode

import "math"

// 题目: 买卖股票的最佳时机. 只允许: 一次买入 & 一次卖出;
//
// 逻辑:
// - 逢低买入: 如果今天的价钱低, 今天买入;
// - 逢高卖出: 今天的价钱比最低的那天高. 今天卖出, 收益是否更大;
func maxProfit(prices []int) int {
	ans := 0

	min := math.MaxInt
	for _, cur := range prices {
		switch {
		case cur < min:
			min = cur
		default:
			if profit := cur - min; profit > ans {
				ans = profit
			}
		}
	}

	return ans
}
