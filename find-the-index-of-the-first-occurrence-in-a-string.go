package leetcode

// KMP substr 实现.
//
// 1. O(m) 实现子串的 “最长前缀长度”;
// 2. O(n) 实现长串的查找;
func strStr(haystack string, needle string) int {
	// next 记录: 最长的前、后缀长度.
	// 长度: 理解为"下一个,待匹配的字符下标".
	next := make([]int, len(needle))

	// 构造: 最长前缀长度数组;
	// 1. 如果 s[i] == s[j], next[i] = j+1. 同时移动 j 和 i;
	// 2. 如果 s[i] != s[j], j = next[j-1]. 循环直到: j = 0 || s[i] == s[j]. 移动 i;

	next[0] = 0 // 常量. 单一字符不存在真前缀. 所以, 固定为 0.
	for j, i := 0, 1; i < len(needle); /**/ {
		if needle[i] == needle[j] {
			next[i] = j + 1
			i, j = i+1, j+1
			continue
		}

		for needle[i] != needle[j] && j != 0 {
			j = next[j-1]
		}

		if needle[i] != needle[j] {
			next[i] = 0
			i++
		}
	}

	// 查找子串.
	for i, j := 0, 0; i < len(haystack) && j < len(needle); /**/ {
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
			if j == len(needle) {
				return i - j // 已完成匹配子串.
			}

			continue
		}

		for haystack[i] != needle[j] && j != 0 {
			j = next[j-1]
		}

		if haystack[i] != needle[j] {
			i++
		}
	}

	return -1
}
