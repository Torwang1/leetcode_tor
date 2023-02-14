package leetcode

// 题目: 无重复字符的最长子串.
//
// 条件: 不含重复字符, hash 计数器.
//
// 主体算法: 滑动窗口(双指针).
//

func lengthOfLongestSubstring(s string) int {
	ans := 0

	// 计数器的数值, 只可能是 0 或 1.
	counter := map[byte]int{}

	// 左右指针
	for l, r := 0, 0; r < len(s); r++ {
		c := s[r]

		// 左边界,右移.
		for counter[c] > 0 {
			counter[s[l]] = 0
			l++
		}

		// 右边界, 加入.
		counter[c] = 1

		if length := r - l + 1; length > ans {
			ans = length
		}
	}

	return ans
}
