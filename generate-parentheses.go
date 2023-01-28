package leetcode

// 相同题目: bracket-lcci.go

// 题目: 数字 n, 输出所有有效括号;
//
// 特征: 所有有效括号, 回溯(递归)输出所有可能;
//
// 有效括号的条件: '(' 出现次数 >= ')' 出现次数;

func generateParenthesisII(n int) []string {
	ans := []string{}

	chars := make([]byte, 0, 2*n)
	parenthesisII(&ans, &chars, n, n)

	return ans
}

// 回溯实现
func parenthesisII(ans *[]string, chars *[]byte, leftn int, rightn int) {
	next := nextCharsII(leftn, rightn)
	if len(next) == 0 {
		*ans = append(*ans, string(*chars))
		return
	}

	for _, c := range next {
		*chars = append(*chars, c)

		switch c {
		case '(':
			parenthesisII(ans, chars, leftn-1, rightn)
		case ')':
			parenthesisII(ans, chars, leftn, rightn-1)
		}

		*chars = (*chars)[:len(*chars)-1]
	}
}

// next 有效括号的条件, 其中 leftn rightn 表示剩余出现次数:
// - leftn == rightn == 0, next = []
// - leftn == rightn, next = ['(']
// - leftn < rightn && leftn == 0, next = [')']
// - leftn < rightn, next = ['(', ')']
func nextCharsII(leftn int, rightn int) []byte {
	result := []byte{}

	switch {
	case leftn == rightn && leftn == 0:
		// no char
	case leftn == rightn:
		result = append(result, '(')
	case leftn < rightn && leftn == 0:
		result = append(result, ')')
	case leftn < rightn:
		result = append(result, '(', ')')
	}

	return result
}
