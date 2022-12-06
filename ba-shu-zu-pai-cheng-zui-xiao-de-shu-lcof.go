package leetcode

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 排序规则: 拼接字符串, x + y < y + x ，则 x “小于” y
//
// https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/mian-shi-ti-45-ba-shu-zu-pai-cheng-zui-xiao-de-s-4/

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]

		xy, _ := strconv.Atoi(fmt.Sprintf("%d%d", x, y))
		yx, _ := strconv.Atoi(fmt.Sprintf("%d%d", y, x))

		if xy < yx {
			return true
		}
		return false
	})

	return strings.Trim(strings.ReplaceAll(fmt.Sprint(nums), " ", ""), "[]")
}
