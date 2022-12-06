package leetcode

import "sort"

// 扑克牌中的顺子
//
// 注意: 可以是多副牌, 也就是任意多大小王;
//
// 条件:
// - 非 0 数, 没有重复;
// - 非 0 数, 最大值 和 最小值 差值不超过 5;

func isStraight(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// 特殊的全 0 场景
	if nums[len(nums)-1] == 0 {
		return true
	}

	// last 探测重复牌;
	last, max := nums[len(nums)-1], nums[len(nums)-1]

	// max min 找出最大最小;
	min := nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0 && nums[i] != 0; i-- {
		v := nums[i]
		if last == v {
			return false // 重复牌
		}
		last = v
		min = v
	}

	if max-min < 5 {
		return true
	}
	return false
}
