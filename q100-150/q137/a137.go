package q137

import "sort"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	var first = true
	for i := 0; i < len(nums); i++ {
		if first {
			first = false
		} else {
			if nums[i] != nums[i-1] {
				return nums[i-1]
			}
			i++
			first = true
		}
	}
	return nums[len(nums)-1]
}