package leetcode

// 最长公共前缀
//
// 搜索多个字符串的公共前缀. 可以顺序来解:
// 1) pre := prefix(str1, str2);
// 2) pre := prefix(pre, str), 其中, str 循环字符串数组;
//

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		str := strs[i]

		j := 0
		for j < len(pre) && j < len(str) {
			if pre[j] == str[j] {
				j++
				continue
			}
			break
		}

		pre = pre[0:j]
	}

	return pre
}
