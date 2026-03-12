package q85

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxArea = 0

	for i := range matrix {
		intersection := matrix[i]
		area := intersectionToAreaBytes(1, matrix[i], matrix[i])
		if area > maxArea {
			maxArea = area
		}
		for j := i + 1; j < len(matrix); j++ {
			// 计算 matrix[i] 和 matrix[j] 的交集
			intersection = BytesAnd(intersection, matrix[j])
			area = intersectionToAreaBytes(j-i+1, matrix[i], intersection)
			if area > maxArea {
				maxArea = area
			}
			if area == 0 {
				break
			}
		}
	}
	return maxArea
}

func BytesAnd(bytes1 []byte, bytes2 []byte) []byte {
	var result = make([]byte, len(bytes1))
	for i := range bytes1 {
		result[i] = bytes1[i] & bytes2[i]
	}
	return result
}

func intersectionToAreaBytes(rows int, startRow []byte, intersection []byte) int {
	result := BytesAnd(startRow, intersection)
	// 计算 result 中连续的 1 的最大长度
	var maxLength = 0
	var currentLength = 0
	for i := 0; i < len(result); i++ {
		if result[i] != '0' {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 0
		}
	}
	if currentLength > maxLength {
		maxLength = currentLength
	}
	return maxLength * rows
}
