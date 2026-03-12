package q547

import "math/big"

// isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连
// 如果 i 和 j 对换, 结果肯定也是一样的, 说明isConnected肯定是斜对称的, 如下
// [1, 0, 1]
// [0, 1, 0]
// [1, 0, 1]
// 因此, 只需要便利一半, 如上, 只需要遍历: isConnected[0][1], [0][2], [1][2]
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected) // n X n 矩阵
	// 直接用map模拟简易并查集, key为城市, value为根节点城市, 如果value一致则处于同一个集合
	// 遍历过程中顺便压缩路径, wasted too much time, try optimize....
	var set = make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := set[i]; !ok {
			set[i] = i
		}
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				root := set[i] // 因为上层已经设置过set[i], 所以这里肯定有值
				if previousRoot, ok := set[j]; ok && root != previousRoot {
					// 合并 并且 压缩路径
					for k, v := range set {
						if v == root {
							set[k] = previousRoot
						}
					}
				} else {
					set[j] = root
				}
			}
		}
	}
	var result = make(map[int]struct{})
	for _, v := range set {
		result[v] = struct{}{}
	}
	return len(result)
}

// slice 版本
func findCircleNum2(isConnected [][]int) int {
	n := len(isConnected) // n X n 矩阵
	// 直接用slice模拟简易并查集, index为城市, value为根节点城市, 如果value一致则处于同一个集合
	var set = make([]int, len(isConnected))
	for i := range set {
		set[i] = -1
	}
	for i := 0; i < n; i++ {
		if set[i] == -1 {
			set[i] = i
		}
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				root := set[i] // 因为上层已经设置过set[i], 所以这里肯定有值
				if set[j] > -1 && set[j] != set[i] {
					// 合并 并且 压缩路径
					for k, v := range set {
						if v == root {
							set[k] = set[j]
						}
					}
				} else {
					set[j] = root
				}
			}
		}
	}
	var result = make(map[int]struct{})
	for _, v := range set {
		if v < 0 {
			continue
		}
		result[v] = struct{}{}
	}
	return len(result)
}

// slice + 比特位 | 统计,  -_- ! 不行, 大数效率低下
func findCircleNum3(isConnected [][]int) int {
	n := len(isConnected) // n X n 矩阵
	// 直接用slice模拟简易并查集, index为城市, value为根节点城市, 如果value一致则处于同一个集合
	var set = make([]int, len(isConnected))
	for i := range set {
		set[i] = -1
	}
	for i := 0; i < n; i++ {
		if set[i] == -1 {
			set[i] = i
		}
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] == 1 {
				root := set[i] // 因为上层已经设置过set[i], 所以这里肯定有值
				if set[j] > -1 && set[j] != set[i] {
					// 合并 并且 压缩路径
					for k, v := range set {
						if v == root {
							set[k] = set[j]
						}
					}
				} else {
					set[j] = root
				}
			}
		}
	}
	var result = big.NewInt(0)
	for _, v := range set {
		if v < 0 {
			continue
		}
		result.Or(result, pow(2, v))
	}
	return count1s(result)
}

func count1s(num *big.Int) int {
	cnt := 0
	for i := 0; i < num.BitLen(); i++ {
		if num.Bit(i) == 1 {
			cnt++
		}
	}
	return cnt
}

func pow(x, n int) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	var result = big.NewInt(int64(x))
	for i := 0; i < n-1; i++ {
		result.Mul(result, result)
	}
	return result
}
