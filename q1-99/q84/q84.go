package q84

func largestRectangleArea(heights []int) int {
	// 既然都是从左往右遍历, 按理说一次遍历应该可以整理完备, 再优化一下
	maxArea := 0
	left := make([]int, len(heights))
	stack := make([]int, 0, len(heights))
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		left[i] = i

		for len(stack) > 0 && height <= heights[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				left[i] = 0
			} else {
				left[i] = left[popIndex] // 往前溯源到最左端
			}
			area := heights[popIndex] * (i - 1 - left[popIndex] + 1) // 这个时候最右边index应该是i-1
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	farRight := stack[len(stack)-1]
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := heights[popIndex] * (farRight - left[popIndex] + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func largestRectangleAreaV2(heights []int) int {
	//fmt.Println("heights", heights)
	// 前面的代码有一个问题, 最左侧 index 没法很好记录, 需要引进一个新的数组保存最左侧index
	maxArea := 0
	left := make([]int, len(heights))
	stack := make([]int, 0, len(heights)) // 单调栈存heights下标, 从小到大入栈
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		left[i] = i
		for len(stack) > 0 && height <= heights[stack[len(stack)-1]] {
			// 出栈
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				left[i] = 0
			} else {
				left[i] = left[popIndex]
			}
		}
		stack = append(stack, i)
	}
	//fmt.Println("left", left, "remain stack", stack)

	if len(stack) > 0 {
		stack = stack[:1]
	}

	for i := 0; i < len(heights); i++ {
		// 一旦当前高度小于栈顶高度，说明栈顶高度的矩形已经结束了
		height := heights[i]
		//fmt.Println("height", height)
		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {
			// 出栈
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			area := heights[popIndex] * (i - 1 - left[popIndex] + 1) // 这个时候最右边index应该是i-1
			//fmt.Println("begin to pop", "area", area, "height", heights[popIndex], "left", left[popIndex], "i", i, "width", (i - 1 - left[popIndex] + 1))
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	farRight := stack[len(stack)-1]
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		area := heights[popIndex] * (farRight - left[popIndex] + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func largestRectangleAreaMonoStackFirstTrial(heights []int) int {
	maxArea := 0
	stack := make([]int, 0, len(heights)) // 单调栈存heights下标, 从小到大入栈
	for i, height := range heights {
		// 一旦当前高度小于栈顶高度，说明栈顶高度的矩形已经结束了
		for len(stack) > 0 && height < heights[stack[len(stack)-1]] {
			// 出栈
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 && popIndex > 0 { // 如果中途栈空了, 说明到此为止, 当前遍历位置的最小值到起始位整体的最小高度确定下来了
				area := heights[popIndex] * (popIndex + 1)
				if area > maxArea {
					maxArea = area
				}
			} else {
				// 计算面积, pop index 的高度 * 宽度 (i - pop index)
				area := heights[popIndex] * (i - popIndex)
				if area > maxArea {
					maxArea = area
				}
			}
			if height == 0 && len(heights[i+1:]) > 0 {
				area := largestRectangleArea(heights[i+1:])
				if area > maxArea {
					maxArea = area
				}
				//stack = []int{}
				//break
			}
		}
		stack = append(stack, i) // 最后一轮循环入栈最后一个高度
	}
	for len(stack) > 0 {
		popIndex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 { // 如果栈空了，说明这个高度是最小的了，宽度就是整个数组长度, 因为整个栈是递增的, pop 到最后, 肯定是最小值
			area := heights[popIndex] * len(heights)
			if area > maxArea {
				maxArea = area
			}
			break
		}
		area := heights[popIndex] * (len(heights) - popIndex)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleAreaViolent(heights []int) int {
	var maxArea, area = 0, 0
	var minHeight = 0
Loop:
	for i := 0; i < len(heights); i++ {
		height := heights[i]
		if height <= minHeight {
			continue
		}
		if height > maxArea {
			maxArea = height
		}
		for j := i + 1; j < len(heights); j++ {
			if heights[j] < height {
				height = heights[j]
				if height <= minHeight {
					continue Loop
				}
			}
			area = height * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
		minHeight = height
	}
	return maxArea
}
