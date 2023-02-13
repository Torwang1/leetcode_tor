package leetcode

// 题目: 最小窗口子序列
//
// 基础算法: 子序列 + 动态规划
//
// 推理过程:
// 1) 设窗口起点 i, s[i] = t[0];
// 2) 在 s[i+1:] 中寻找 j, s[j] = t[1];
// 3) 在 s[j+1:] 中寻找 k, s[k] = t[2];
// ...
// 目标:
// 1) 在位置 i 时, 知道 t[j] 下一次出现的位置;
//
// next[i][26]: 在字符串 s 中, 位置 i, 下一个小写字母出现的位置.
//

func minWindow(s1 string, s2 string) string {

	sentry := len(s1) // 哨兵, sentry 是 s1 中的无效下标.

	next := [][]int{}
	for i := 0; i < len(s1); i++ {
		next = append(next, make([]int, 26))
	}

	// 计算 next 数组, 即: s1 的特征;
	// 注意: 倒序计算;
	for j := 0; j < 26; j++ {
		next[len(s1)-1][j] = sentry
	}
	for i := len(s1) - 2; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if int(s1[i+1]) == int(j+'a') {
				next[i][j] = i + 1
				continue
			}
			next[i][j] = next[i+1][j]
		}
	}

	// 利用 next, 在 s1 中快速找到下一个字母的位置;
	//
	// r 最小窗口的 右边界;
	// size 最小窗口的大小;
	// 目标值: s[r-size:r]

	var ans string

	var first int
	if s1[0] != s2[0] {
		first = next[first][s2[0]-'a']
	}

	// l 满足 “s1[l] == s2[0]” 或者 l = sentry;
	for l, i, j := first, first, 0; i < len(s1) && j < len(s2); /**/ {
		if s1[i] == s2[j] {
			i, j = i+1, j+1

			// 找到一个 window, 并记录最小窗口.
			if j == len(s2) {
				win := i - l
				if len(ans) == 0 || win < len(ans) {
					ans = string(s1[l:i])
				}
				// 移动 s 左边界, 继续比较.
				first = next[l][s2[0]-'a']
				l, i, j = first, first, 0
			}

		} else {
			i = next[i][s2[j]-'a']
		}
	}

	return ans
}
