package q739

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)/2)
	for i := 0; i < len(temperatures); i++ {
		temp := temperatures[i]
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[popIndex] = i - popIndex
		}
		stack = append(stack, i)
	}
	return result
}

func dailyTemperaturesInPlace(temperatures []int) []int {
	// result := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)/2)
	for i := 0; i < len(temperatures); i++ {
		temp := temperatures[i]
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			popIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			temperatures[popIndex] = i - popIndex
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(stack); i++ {
		temperatures[stack[i]] = 0
	}
	return temperatures
}
