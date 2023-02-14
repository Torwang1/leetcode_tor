package leetcode

// 岛屿的最大面积
//
// DFS, 类似“岛屿的数量”. 差别在于 dfs 需要记录访问数量.
//

func maxAreaOfIsland(grid [][]int) int {
	var ans int

	m, n := len(grid), len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}

			var area int
			maxAreaOfIslandDFS(grid, i, j, &area)
			if area > ans {
				ans = area
			}
		}
	}

	return ans
}

func maxAreaOfIslandDFS(grid [][]int, i, j int, area *int) {
	m, n := len(grid), len(grid[0])

	if (i < 0 || i >= m) || (j < 0 || j >= n) {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0
	*area++

	for _, d := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		maxAreaOfIslandDFS(grid, i+d[0], j+d[1], area)
	}
}
