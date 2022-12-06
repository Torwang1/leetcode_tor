package leetcode

import "math"

func waysToChange(n int) int {
	var ans int64

	changeBacktrace(&ans, math.MaxInt, n)

	return int(ans % 1000000007) // 模运算, 题目要求.
}

// changeBacktrace 回溯兑换零钱.
func changeBacktrace(ans *int64, max int, remain int) {
	if remain == 0 {
		*ans = *ans + 1
		return
	}

	nexts := nextChanges(max, remain)
	for _, chg := range nexts {
		remain = remain - chg
		changeBacktrace(ans, chg, remain)
		remain = remain + chg
	}
}

// nextChanges 所有 “小于等于” remain 的币值.
// max: 最大可使用的货币;
func nextChanges(max int, remain int) []int {

	result := make([]int, 0, 4)
	for _, chg := range []int{25, 10, 5, 1} {
		if max >= chg && remain >= chg {
			result = append(result, chg)
		}
	}

	return result
}
