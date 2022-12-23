package leetcode

import "strings"

// removeKdigits 题目要求: 最小子序列, 即: 不能改变元素的相对位置;
//
// 相同位数的数比较: 从左向右, 第一个不相同数字的大小决定两个树的大小;
//
// 解题思路: 争取让最左侧的数字更小;
//
// 单调栈: 从栈顶到栈底, stack[i-1] <= stack[i];
func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num)-k)

	for i := range num {
		digit := num[i]

		for len(stack) > 0 && digit < stack[len(stack)-1] && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, digit)
	}

	// k: 剩余待移除的数字个数.
	ans := string(stack[:len(stack)-k])

	// 前导 0 的处理 && 空串的处理;
	ans = strings.TrimLeft(ans, "0")
	if ans == "" {
		ans = "0"
	}

	return ans
}
