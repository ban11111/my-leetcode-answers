package q403

// 2 <= stones.length <= 2000
// 0 <= stones[i] <= 231 - 1
// stones[0] == 0

// map 版, 实测 map 消耗远远高于 slice 版, leetcode map版不做优化直接超时.....
func canCross(stones []int) bool {
	dp := make([]map[int]bool, len(stones))
	for i := range dp {
		dp[i] = make(map[int]bool)
	}
	if stones[1] != 1 {
		return false
	}
	if len(stones) == 2 {
		return true
	}
	dp[1] = map[int]bool{1: true}
	for i := 2; i < len(stones); i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
		for j := i - 1; j > 0; j-- {
			k := stones[i] - stones[j]
			if k > j+1 {
				break
			}
			dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			if i == len(stones)-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}

func canCross2(stones []int) bool {
	dp := make([][]bool, len(stones))
	for i := range dp {
		dp[i] = make([]bool, len(stones))
	}
	if stones[1] != 1 {
		return false
	}
	if len(stones) == 2 {
		return true
	}
	dp[1][1] = true
	for i := 2; i < len(stones); i++ {
		if stones[i]-stones[i-1] > i {
			return false
		}
		for j := i - 1; j > 0; j-- {
			k := stones[i] - stones[j]
			//if k > j+1 {
			//	break
			//}
			if k > i {
				break
			}
			if i == len(stones)-1 && k == i {
				dp[i][k] = dp[j][k-1] || dp[j][k]
			} else {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			}
			if i == len(stones)-1 && dp[i][k] {
				return true
			}
		}
	}
	return false
}
