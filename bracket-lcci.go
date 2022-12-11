package leetcode

// 相同题目: generate-parentheses.go

// 算法分析:
//
// degree := “左括号的数量” - “右括号的数量”;
// - degree = 0: 只能填充 '(';
// - degree > 0: 可以填充 '(', ')';

func generateParenthesis(n int) []string {
	result := make([]string, 0)

	// 记录括号字符.
	chars := make([]byte, 0, 2*n)
	// 同步 chars, 记录 degree;
	degree := make([]int, 0, 2*n)
	// 记录 '(' 的数量;
	rN := make([]int, 0, 2*n)

	parenthesis(n, &result, &chars, &degree, &rN)

	return result
}

// parenthesis 回溯函数, 逐个字符填充.
func parenthesis(n int, result *[]string, chars *[]byte, degree *[]int, rN *[]int) {
	nextCs := nextChars(n, *degree, *rN)

	if len(nextCs) == 0 {
		*result = append(*result, string(*chars))
		return
	}

	addDegree := func(c byte) {
		var pre int
		switch {
		case len(*degree) == 0:
			pre = 0
		default:
			pre = (*degree)[len(*degree)-1]
		}

		switch {
		case c == '(':
			*degree = append(*degree, pre+1)
		case c == ')':
			*degree = append(*degree, pre-1)
		}
	}
	addRN := func(c byte) {
		var pre int
		switch {
		case len(*rN) == 0:
			pre = 0
		default:
			pre = (*rN)[len(*rN)-1]
		}

		switch {
		case c == '(':
			*rN = append(*rN, pre+1)
		case c == ')':
			*rN = append(*rN, pre)
		}
	}

	for _, c := range nextCs {
		*chars = append(*chars, c)
		addDegree(c)
		addRN(c)
		parenthesis(n, result, chars, degree, rN)
		*chars = (*chars)[:len(*chars)-1]
		*degree = (*degree)[:len(*degree)-1]
		*rN = (*rN)[:len(*rN)-1]
	}
}

// nextChars 反馈可填充的字符:
// - degree = 0, '(' ;
// - degree > 0, '(' 和 ')';
func nextChars(n int, degree []int, rN []int) []byte {
	result := make([]byte, 0, 2)

	switch {

	// - 已填充 2n 字符, 返回空;
	case len(degree) == 2*n:
		return result

	// 首次进入, degree 为空, degree = 0 的特化场景;
	case len(degree) == 0:
		result = append(result, '(')
		return result

	// 非首次进入, degree = 0;
	case degree[len(degree)-1] == 0:
		result = append(result, '(')
		return result

	// 非首次进入, degree > 0;
	case degree[len(degree)-1] > 0:
		if rN[len(rN)-1] < n {
			result = append(result, '(')
		}
		result = append(result, ')')
		return result

	default:
		panic("never here")
	}
}
