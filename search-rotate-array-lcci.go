package leetcode

// 题目分析:
// - 升序数组, 经过旋转: 两段有序;
// - 重复元素, 手动剔除;
func search(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// 相同元素, 返回索引最小的. 跳过尾部相同元素.
	for arr[right] == arr[left] {
		if left < right {
			right--
		}
	}

	i := binarySearch(arr, left, right, target)
	if i < 0 {
		return -1
	}

	// 相同元素, 返回索引值最小的.
	for i > 0 && arr[i-1] == arr[i] {
		i--
	}
	return i
}

// 两段有序, 注意控制: mid
func binarySearch(arr []int, left int, right int, target int) (i int) {
	if left > right {
		return -1
	}

	mid := (right + left) / 2
	if target == arr[mid] {
		return mid
	}

	if arr[0] <= arr[mid] { // 左侧有序
		if arr[left] <= target && target < arr[mid] {
			return binarySearch(arr, left, mid-1, target)
		} else {
			return binarySearch(arr, mid+1, right, target)
		}
	} else { // 右侧有序
		if arr[right] >= target && target > arr[mid] {
			return binarySearch(arr, mid+1, right, target)
		} else {
			return binarySearch(arr, left, mid-1, target)
		}
	}
}
