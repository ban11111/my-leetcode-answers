package q123

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var max, min, maxSub, reverseMin, reverseMax int
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