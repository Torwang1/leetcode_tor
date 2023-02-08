package leetcode

// 问题: 每日温度
//
// 特征: next 高温. 单调栈
//
// 栈操作逻辑:
// 1) cur > 栈顶: 计算并弹出栈顶;
// 2) cur 下标压栈;
// 3) 跳出循环后, 针对栈内残留下标, 统一填充 0;

func dailyTemperatures(temperatures []int) []int {
	next := make([]int, len(temperatures))

	stack := make([]int, 0, len(temperatures))

	for i, cur := range temperatures {

		// 1) cur > 栈顶: 计算并弹出栈顶;
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			val := temperatures[top]
			if cur > val {
				next[top] = i - top
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		// 2) cur 下标压栈;
		stack = append(stack, i)
	}

	// 3) 跳出循环后, 针对栈内残留下标, 统一填充 0;
	for _, i := range stack {
		next[i] = 0
	}

	return next
}
