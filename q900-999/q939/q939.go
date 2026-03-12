package q939

// 939. Minimum Area Rectangle
// 先整理一下思路:
// 1. 如何在遍历过程中判断是否存在一个矩形? 需要找到4个点, 这4个点的坐标分别是(a, b), (a, c), (d, b), (d, c)
// 2. 想到了, 能不能做两个方向的索引, 行索引和列索引, 这样 有两个点的时候, 就可以通过行索引和列索引来判断是否存在另外两个点
// 3. 最后返回最小面积, 如果没有找到矩形, 就返回0
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}
	var xIndex = make(map[int]map[int]struct{}, len(points))
	var yIndex = make(map[int]map[int]struct{}, len(points))
	for i := range points {
		if xIndex[points[i][0]] == nil {
			xIndex[points[i][0]] = make(map[int]struct{})
		}
		if yIndex[points[i][1]] == nil {
			yIndex[points[i][1]] = make(map[int]struct{})
		}
		xIndex[points[i][0]][points[i][1]] = struct{}{}
		yIndex[points[i][1]][points[i][0]] = struct{}{}
	}
	var minArea = 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] == points[j][0] || points[i][1] == points[j][1] {
				continue
			}
			if _, ok := xIndex[points[i][0]][points[j][1]]; !ok {
				continue
			}
			if _, ok := xIndex[points[j][0]][points[i][1]]; !ok {
				continue
			}
			area := abs(points[i][0]-points[j][0]) * abs(points[i][1]-points[j][1])
			if minArea == 0 || area < minArea {
				minArea = area
			}
		}
	}
	return minArea
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
