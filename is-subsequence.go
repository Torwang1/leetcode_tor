package leetcode

// 判断: s 是否为 t 的子序列
//
// 动态规划: next[i][char], 从位置 i 开始(包含), char 首次出现的下标
//
// 推理过程:
// - next[i][char] = i, t[i] = char;
// - next[i][char] = next[i+1][char], t[i] != char;
// 所以, 反向求解 next 数组;
//
// 进阶:
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，
// 你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
//
// 进阶解法:
// - t 是固定的, next 数组可以加速比较过程;
//
func isSubsequence(s string, t string) bool {

	sentry := -1 // 哨兵, 表示元素不存在.

	tlen := len(t)

	// 求解 next 数组.
	next := make([][26]int, tlen+1) // 设置边界哨兵.
	for i := 0; i < 26; i++ {
		next[tlen][i] = sentry
	}

	for i := tlen - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			char := 'a' + j
			if t[i] == byte(char) {
				next[i][j] = i
			} else {
				next[i][j] = next[i+1][j]
			}
		}
	}

	// 利用 next 信息, 判断 s 是否为 t 的子序列;
	for i, j := 0, 0; j < len(s); /**/ {
		char := s[j]

		// 字符串 t 中, 从下标 i 开始, char 首次出现的下标.
		index := next[i][int(char-'a')]

		// 哨兵, 表示: 字符出 t 中, 从下标 i 开始, 不再出现 char 字符.
		if index == sentry {
			return false
		}

		// 哨兵是合法下标, 表示: char 已经匹配, 从下一个字符继续匹配.
		i = next[i][int(char-'a')] + 1
		j = j + 1

	}

	return true
}
