package leetcode

// 题目: 岛屿的数量
//
// 约束: 上下左右, 斜角不算.
//
// 实现优化:
// - 已经被识别为‘岛屿’的位置, 标记为 0.

func numIslands(grid [][]byte) int {
	const land = '1'

	rn, cn := len(grid), len(grid[0])

	var number int
	for r := 0; r < rn; r++ {
		for c := 0; c < cn; c++ {
			if grid[r][c] == land {
				number++
				DFMark(grid, r, c)
			}
		}
	}

	return number
}

// 深度优先遍历
func DFMark(grid [][]byte, r, c int) {
	const water = '0'
	const land = '1'

	rn, cn := len(grid), len(grid[0])

	grid[r][c] = water
	if r-1 >= 0 && grid[r-1][c] == land {
		DFMark(grid, r-1, c)
	}
	if r+1 < rn && grid[r+1][c] == land {
		DFMark(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == land {
		DFMark(grid, r, c-1)
	}
	if c+1 < cn && grid[r][c+1] == land {
		DFMark(grid, r, c+1)
	}
}
