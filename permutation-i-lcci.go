package leetcode

func permutation(S string) []string {
	ans := make([]string, 0)

	visited := make(map[byte]bool)

	unvisitedFunc := func() []byte {
		result := []byte{}
		for i := 0; i < len(S); i++ {
			c := S[i]
			if !visited[c] {
				result = append(result, c)
			}
		}
		return result
	}

	chars := make([]byte, 0)

	recursive(&ans, &chars, unvisitedFunc, visited)

	return ans
}

// recursive
// ans: 结果集;
// chars: 字节缓存;
func recursive(ans *[]string, chars *[]byte, unvisitedFunc func() []byte, visited map[byte]bool) {
	unvisited := unvisitedFunc()

	if len(unvisited) == 0 {
		*ans = append(*ans, string(*chars))
		return
	}

	for _, c := range unvisited {
		visited[c] = true
		*chars = append(*chars, c)
		recursive(ans, chars, unvisitedFunc, visited)
		visited[c] = false
		*chars = (*chars)[:len(*chars)-1]
	}
}
