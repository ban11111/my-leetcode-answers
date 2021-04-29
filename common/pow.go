package common

// == 这个包记录一些通用工具, 有需要直接 复制粘贴 方便

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	var result = x
	for i := 0; i < n-1; i++ {
		result *= x
	}
	return result
}
