package leetcode

// 题目要求: 相同字母仅保存一次, 相同位数 && 最小字母序.
//
// 单调栈 + 贪心
//
func removeDuplicateLetters(s string) string {
	stack := make([]byte, 0, 26)

	remains, visited := map[byte]int{}, map[byte]bool{}

	// 统计元素出现次数. 用于判断是否可以被出栈;
	for i := 0; i < len(s); i++ {
		remains[s[i]]++
	}

	// 单调栈实现
	for i := 0; i < len(s); i++ {
		letter := s[i]

		// 如果 letter 在单调栈中, 则必然满足一下情况之一:
		// - letter 小于 栈顶;
		// - 某个元素最后一次出现, 导致局部栈被锁定. letter 恰好在这个元素之前;
		if !visited[letter] {
			for i := len(stack) - 1; i > -1 && letter < stack[i] && remains[stack[i]] > 0; i = len(stack) - 1 {
				top := stack[i]
				visited[top] = false
				stack = stack[:i]
			}

			stack = append(stack, letter)
			visited[letter] = true
		}

		remains[letter]--
	}

	return string(stack)
}
