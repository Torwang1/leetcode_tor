package leetcode

import "strings"

// 重点思路: s+s, 形成旋转后的最大字符串.

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(s+s, goal)
}
