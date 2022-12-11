package leetcode

// 相同题目: bracket-lcci.go

// 题目: 数字 n, 输出有效括号;
//
// 思路:
// (   )
// (   )
//  ...
// (   )
// 输出 n 个序列, 同时保证括号有效;
//
// 输出所有可能, 需要 回溯 算法;

func generateParenthesisII(n int) []string {
	ans := make([]string, 0)
	chars := make([]byte, 0, 2*n)

	parenthesisII(&ans, &chars, n, n)

	return ans
}

func parenthesisII(ans *[]string, chars *[]byte, left int, right int) {
	nexts := nextCharsII(left, right)
	if len(nexts) == 0 {
		*ans = append(*ans, string(*chars))
		return
	}

	for _, c := range nexts {
		*chars = append(*chars, c)
		switch c {
		case '(':
			parenthesisII(ans, chars, left-1, right)
		case ')':
			parenthesisII(ans, chars, left, right-1)
		}
		*chars = (*chars)[:len(*chars)-1]
	}
}

// left:  左括号的剩余数量;
// right: 右括号的剩余数量;
func nextCharsII(left int, right int) []byte {
	result := []byte{}

	switch {
	case left == 0 && right == 0:
		// do nothing. 边界 case (n=0) 或者 迭代结束.
	case left > 0 && right > 0:
		result = append(result, '(')
		if left < right {
			result = append(result, ')')
		}
	case left == 0 && right > 0:
		result = append(result, ')')
	default:
		panic("why here")
	}

	return result
}
