package q85

import (
	"math/big"
)

func maximalRectangleBigInt(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxArea = 0
	var rows = make([]*big.Int, 0, len(matrix))
	for i := range matrix {
		rows = append(rows, rowToBigInt(matrix[i]))
	}

	for i := range rows {
		intersection := rows[i]
		area := intersectionToAreaBigInt(1, rows[i], rows[i])
		if area > maxArea {
			maxArea = area
		}
		for j := i + 1; j < len(rows); j++ {
			// 计算 rows[i] 和 rows[j] 的交集
			intersection = new(big.Int).And(intersection, rows[j])
			area = intersectionToAreaBigInt(j-i+1, rows[i], intersection)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func rowToBigInt(row []byte) *big.Int {
	var result big.Int
	x := big.NewInt(1)
	for i := uint64(0); i < uint64(len(row)); i++ {
		if row[i] == '1' {
			result.Or(&result, new(big.Int).Lsh(x, uint(i)))
		}
	}
	return &result
}

func intersectionToAreaBigInt(rows int, startRow *big.Int, intersection *big.Int) int {
	result := new(big.Int).And(startRow, intersection)
	var maxLength = 0
	var currentLength = 0
	for i := 0; i < 200; i++ {

		if result.Bit(i) != 0 {
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

func maximalRectangleUint64Version(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	var maxArea = 0

	// 先把每行的1转换成一个大整数，方便后续的位运算
	// 现用 uint64 测试一下
	var rows = make([]uint64, 0, len(matrix))
	for i := range matrix {
		rows = append(rows, rowToInt(matrix[i]))
	}

	for i := range rows {
		intersection := rows[i]
		area := intersectionToArea(1, rows[i], rows[i])
		if area > maxArea {
			maxArea = area
		}
		for j := i + 1; j < len(rows); j++ {
			// 计算 rows[i] 和 rows[j] 的交集
			intersection &= rows[j]
			area = intersectionToArea(j-i+1, rows[i], intersection)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func rowToInt(row []byte) uint64 {
	var result uint64 = 0
	for i := uint64(0); i < uint64(len(row)); i++ {
		if row[i] == '1' {
			result |= 1 << i
		}
	}
	return result
}

func intersectionToArea(rows int, startRow uint64, intersection uint64) int {
	result := startRow & intersection
	// 计算 result 中连续的 1 的最大长度
	var maxLength = 0
	var currentLength = 0
	for i := uint64(0); i < 64; i++ {
		if result&(1<<i) != 0 {
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
