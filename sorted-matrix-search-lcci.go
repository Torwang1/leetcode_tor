package leetcode

// 示例数组:
// [
//   [1,   4,  7, 11, 15],
//   [2,   5,  8, 12, 19],
//   [3,   6,  9, 16, 22],
//   [10, 13, 14, 17, 24],
//   [18, 21, 23, 26, 30]
// ]
//
// O(m+n) 算法: 从 “右上” 或 “左下” 进行

// “右上” 实现:
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	line := len(matrix)
	x, y := 0, len(matrix[0])-1

	for x < line && y >= 0 {
		switch {
		case matrix[x][y] < target:
			x++
		case matrix[x][y] == target:
			return true
		case matrix[x][y] > target:
			y--
		}
	}

	return false
}
