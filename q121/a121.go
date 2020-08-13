package q121

// 貌似和 q11 差不多
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var max, min, maxSub int
	max = prices[len(prices)-1]
	min = max
	for i := len(prices)-2; i > -1; i-- {
		if prices[i] > max {
			if max - min > maxSub {
				maxSub = max - min
			}
			max = prices[i]
			min = max
		} else {
			if prices[i] < min {
				min = prices[i]
			}
		}
	}
	if max - min > maxSub {
		return max - min
	}
	return maxSub
}
