package a7

import "math"

func Reverse(x int) int {
	return reverse(x)
}

func reverse(x int) int {
	var result int

	var bit int
	var bits []int
	for {
		left := x / pow(10, bit)
		if left == 0 {
			bit--
			break
		}
		bits = append(bits, left%10)
		bit++
	}
	for i, b := range bits {
		result += b * pow(10, bit-i)
	}

	if overflow(result) {
		return 0
	}
	return result
}

func overflow(result int) bool {
	return result < -2147483648 || result > 2147483647
}

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

func pow10(n int) int { // 官方库效率快好多, 值得学习一下
	return int(math.Pow10(n))
}
