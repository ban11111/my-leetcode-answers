package q9

import "github.com/ban11111/my-leetcode-answers/q7/a7"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == a7.Reverse(x)
}

// 下面这个是抄来的答案, 值得学习, 代码简洁, 还节省了一半的操作
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revertedNumber int

	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}