package q260

func singleNumber(nums []int) []int {
	var mapNum = make(map[int]struct{}, len(nums)/2+1)
	for i := range nums {
		if _, ok := mapNum[nums[i]]; ok {
			delete(mapNum, nums[i])
		} else {
			mapNum[nums[i]] = struct{}{}
		}
	}
	var result = make([]int, 2)
	var i = 0
	for k := range mapNum {
		result[i] = k
		i++
	}
	return result
}
