package q122

// 做完前面那题, 这题就超级简单了, 只要把 maxSub 加起来就可以了
// 并不对, 理解错了 -_-!! 好吧, 重新来 ..... emm ..... 其实也差不多, 小改一下就行
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var max, min, maxSub int
	max = prices[len(prices)-1]
	min = max
	for i := len(prices) - 2; i > -1; i-- {
		if prices[i] > min {
			maxSub += max - min
			max = prices[i]
			min = max
		} else {
			if prices[i] < min {
				min = prices[i]
			}
		}
	}
	maxSub += max - min
	return maxSub
}
