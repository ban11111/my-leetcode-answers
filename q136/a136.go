package q136

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return 0
}

func singleNumber2(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
