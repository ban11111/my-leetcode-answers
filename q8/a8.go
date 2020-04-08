package q8

import (
	"math"
)

var num = map[int32]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

func myAtoi(str string) int {
	result := float64(0)
	var firstFound, positive = false, true
	var slice []int
	for _, x := range str {
		if _, ok := num[x]; !firstFound && x != ' ' && x !='-' && x != '+' && !ok {
			break
		} else if !firstFound && (x == '-' || x == '+') {
			positive = x != '-'
			firstFound = true
		} else if firstFound && !ok {
			break
		} else if !firstFound && ok {
			firstFound = true
			slice = append(slice, num[x])
		} else if firstFound && ok {
			slice = append(slice, num[x])
		}
	}
	for i, b := range slice {
		result += float64(b) * math.Pow10(len(slice)-1-i)
	}

	return join(positive, result)
}

func join(sign bool, result float64) int {
	if !sign {
		result = 0 - result
	}
	if result < -2147483648 {
		return -2147483648
	}
	if result > 2147483647 {
		return 2147483647
	}
	return int(result)
}
