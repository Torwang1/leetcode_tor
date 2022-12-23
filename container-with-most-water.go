package leetcode

// 注意边界条件: height[i] == height[j]
func maxArea(height []int) int {
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for left, right := 0, len(height)-1; left < right; /**/ {
		ans = max(ans, (right-left)*min(height[left], height[right]))

		// 移动指针: 如果移动后的高度, 比已知最小值还小, 那面积一定小于当前值. 直接跳过类似下标;
		switch {
		case height[left] == height[right]:
			left++
		case height[right] < height[left]:
			for cur := height[right]; right >= left && height[right] <= cur; right-- {
			}
		case height[left] < height[right]:
			for cur := height[left]; right >= left && height[left] <= cur; left++ {
			}
		}
	}

	return ans
}
