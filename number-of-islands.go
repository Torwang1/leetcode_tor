package leetcode

// 题目: 岛屿的数量
//
// 约束: 上下左右, 斜角不算.
//
// 实现优化:
// - 已经被识别为‘岛屿’的位置, 标记为 0.

func numIslands(grid [][]byte) int {
	const water = '0'
	const land = '1'

	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		if (i < 0 || i >= m) || (j < 0 || j >= n) {
			return
		}
		if grid[i][j] == water {
			return
		}

		grid[i][j] = water
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	}

	var number int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == land {
				number++
				dfs(grid, r, c)
			}
		}
	}

	return number
}
