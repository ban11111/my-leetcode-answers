package q42

import "runtime/debug"

func init() {
	debug.SetMemoryLimit(1)
}

func trap(height []int) int {
	var fullArea int
	var stack = make([]int, 0, len(height))

	//fmt.Println("height", height)
	for i := 0; i < len(height); i++ {
		//fmt.Println("i", i, "stack", stack)
		h := height[i]
		area := 0
		for len(stack) > 0 && h >= height[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			//fmt.Println("popIndex", popIndex, "h", h, "height[popIndex]", height[popIndex])
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}

			width := max(i-stack[len(stack)-1]-1, i-popIndex)
			areaHeight := min(h-height[popIndex], height[stack[len(stack)-1]]-height[popIndex])
			//fmt.Println(h, popIndex, "areaHeight when len stack > 0", areaHeight, "||| last stack height", height[stack[len(stack)-1]], "width", width)
			area += width * areaHeight
			//fmt.Println("area in for loop: ", area)
		}
		//fmt.Println("area", area)
		fullArea += area

		if len(stack) > 0 && h == height[stack[len(stack)-1]] {
			continue
		}
		stack = append(stack, i)
	}

	return fullArea
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
