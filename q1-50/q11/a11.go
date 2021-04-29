package q11

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var maxVolume, tmpVolume int
	for {
		if i == j {
			break
		}

		if tmpVolume = (j - i) * min(height[i], height[j]); tmpVolume > maxVolume {
			maxVolume = tmpVolume
		}

		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return maxVolume
}

func maxAreaV2(height []int) int {
	i, j := 0, len(height)-1
	var maxVolume, tmpVolume int
	for {
		if tmpVolume = (j - i) * min(height[i], height[j]); tmpVolume > maxVolume {
			maxVolume = tmpVolume
		}
		if j-i == 1 {
			break
		}
		if height[i] <= height[j] {
			current:= height[i]
			i++
			for ;height[i] < current; i++ {
				if j-i == 1 {
					break
				}
			}
		} else {
			current:= height[j]
			j--
			for ; height[j] < current; j-- {
				if j-i == 1 {
					break
				}
			}
		}
	}
	return maxVolume
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
