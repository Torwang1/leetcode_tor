package leetcode

// 字符串的最大公共前后缀.
//
// KMP 的子串算法.

func longestPrefix(s string) string {
	next := make([]int, len(s))

	// next 计算规则:
	// - s[i] == s[j], next[i] = j+1. 同时移动 i 和 j;
	// - s[i] != s[j], j = next[j-1]. 直到 s[i] == s[j] 或者 j == 0;

	next[0] = 0
	for j, i := 0, 1; i < len(s); /**/ {
		if s[i] == s[j] {
			next[i] = j + 1
			i, j = i+1, j+1
			continue
		}

		for s[i] != s[j] && j != 0 {
			j = next[j-1]
		}

		if s[i] != s[j] {
			next[i] = 0
			i = i + 1
		}
	}

	return string(s[:next[len(s)-1]])
}
