package leetcode

import "math"

// 题目: 螺旋矩阵, 矩阵遍历
//
// 方向:
// right: ( 0,  1)
// down:  ( 1,  0)
// left:  ( 0, -1)
// up:    (-1,  0)
//
// 状态机:
// 1) 退出条件: next, next+1 都已经被访问过;

// 题目限制: -100 <= matrix[i][j] <= 100. 所以, visited 复用原始数组.
func spiralOrder(matrix [][]int) []int {
	sentry := math.MaxInt // 哨兵, 表示元素已经被 visited.

	cur := 0
	direction := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// 当前方向的下一个元素.
	next := func(i, j int) (int, int) {
		return i + direction[cur][0], j + direction[cur][1]
	}

	ans := []int{}

	i, j := 0, 0
	for {
		ans = append(ans, matrix[i][j])
		matrix[i][j] = sentry

		// next 获取下一个元素. 如果元素越界 或者 已经被访问过, 则调整方向, 重新获取.
		ti, tj := next(i, j)
		if (ti < 0 || ti >= len(matrix)) || (tj < 0 || tj >= len(matrix[0])) || matrix[ti][tj] == sentry {
			cur = (cur + 1) % len(direction)

			ti, tj = next(i, j)
			if (ti < 0 || ti >= len(matrix)) || (tj < 0 || tj >= len(matrix[0])) || matrix[ti][tj] == sentry {
				break
			}
		}

		i, j = ti, tj
	}

	return ans
}
