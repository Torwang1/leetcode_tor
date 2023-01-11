package leetcode

import "bytes"

// 题目分析: 通过在“前端”插入字符, 获得最短的回文串;
//
// 换个问法: 最长的前缀回文串.
//
// KMP 算法变种. 跳出条件: 原始字符串被耗尽;

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s // leetcode 特殊 case
	}

	sRevert := func() string {
		var buff bytes.Buffer
		for i := len(s); i > 0; i-- {
			buff.WriteByte(s[i-1])
		}
		return buff.String()
	}()

	// 变形的 KMP 算法. 以 sRevert 为原始字符串, 在其中查找 s, 直到 sRevert 消耗完.

	next := make([]int, len(s))

	// 构造子串的, 最长前后缀长度数组. 即: 下一个待比较字符;
	// - s[i] == s[j], next[i] = j + 1. i 和 j 同时移动;
	// - s[i] != s[j], j = next[j-1]. 直到: j == 0 或 s[i] == s[j];

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

	// KMP 变形查找. 注意跳出条件.
	index := 0
	for i, j := 0, 0; i < len(sRevert) && j < len(s); /**/ {
		if sRevert[i] == s[j] {
			i, j = i+1, j+1
			if i == len(sRevert) { // 原始串消耗尽, 找到回文起点.
				index = i - j
				break
			}

			continue
		}

		for sRevert[i] != s[j] && j != 0 {
			j = next[j-1]
		}

		if sRevert[i] != s[j] {
			i = i + 1
		}
	}

	return sRevert[:index] + s
}
