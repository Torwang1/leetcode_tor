package leetcode

// KMP substr 实现.

// KMP 核心解决的问题: 快速移动 needle, 消除无效的循环;
//
// 核心算法是: next
// 逻辑理解: 下一个待比较的字符;
//
// 时间复杂度: O(m+n)

func strStr(haystack string, needle string) int {

	// 构造 next 数组
	next := make([]int, len(needle))

	// next 计算规则:
	// - next[0] = 0. 单一元素, 没有真前缀;
	// - a[i] == a[j], next[i] = j+1. i++, j++;
	// - a[i] != a[j], j = next[j-1]. 直到: a[i] == a[j] 或者 j = 0;
	//                                否则: next[i] = 0, i++;

	next[0] = 0
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

	// 查找子串
	for i, j := 0, 0; i < len(haystack); /**/ {
		if haystack[i] == needle[j] {
			i, j = i+1, j+1
			if j == len(needle) { // 子串完全匹配, 返回结果;
				return i - j
			}
			continue
		}

		for haystack[i] != needle[j] && j != 0 {
			j = next[j-1]
		}

		if haystack[i] != needle[j] {
			i = i + 1
		}
	}

	return -1
}
