package leetcode

// 题目: 判断括号有效性
//
// 栈
//
// 处理:
// - 左: 入栈;
// - 右: 出栈, 判断“括号匹配”;
// - 栈剩余元素判断;
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))

	expected := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := range s {
		c := s[i]

		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != expected[c] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
