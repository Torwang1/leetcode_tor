package leetcode

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)

	// 状态数组 && 初始化
	visited := make([][]bool, 0)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, n))
	}

	queens(n, &result, 0, &visited)

	return result
}

// queens 从上到下, 每次填充一行.
// n: 记录边界条件;
// line: 当前填充的行, [0, n);
func queens(n int, result *[][]string, line int, visited *[][]bool) {
	if line == n {
		*result = append(*result, printQueue(*visited))
		return
	}

	// 如果当前行无法放置皇后, 则是: 无效组合. 直接回溯.
	candidates := candidatePosition(n, line, *visited)
	if len(candidates) == 0 {
		return
	}

	for _, p := range candidates {
		(*visited)[line][p] = true
		queens(n, result, line+1, visited)
		(*visited)[line][p] = false
	}
}

// candidatePosition 当前行, 皇后的可选位置.
// 返回值: line 行的可用下标集合.
//
// N 皇后问题的限制:
// - 不同行; (queens for range 保证)
// - 不同列;
// - 不同对角线;
//
// - 因为从上向下填充, 对角线之需要考虑: 左上 和 右上;
func candidatePosition(n int, line int, visited [][]bool) []int {
	checkPositionOk := func(n int, line int, column int) bool {
		// 列
		for i, j := line-1, column; i >= 0; i-- {
			if visited[i][j] {
				return false
			}
		}
		// 左上对角线
		for i, j := line-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if visited[i][j] {
				return false
			}
		}
		// 右上对角线
		for i, j := line-1, column+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if visited[i][j] {
				return false
			}
		}
		return true
	}

	result := make([]int, 0, n)
	for column := 0; column < n; column++ {
		if checkPositionOk(n, line, column) {
			result = append(result, column)
		}
	}

	return result
}

// printQueue 功能函数, 输出皇后位置.
func printQueue(visited [][]bool) []string {
	result := make([]string, 0, len(visited))

	for _, line := range visited {
		bs := make([]byte, len(line))
		for i, v := range line {
			switch v {
			case true:
				bs[i] = 'Q'
			default:
				bs[i] = '.'
			}
		}
		result = append(result, string(bs))
	}

	return result
}
