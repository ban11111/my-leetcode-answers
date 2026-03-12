package q4

import "sort"

// sort, O(n * log n) + O(n/2)
func findMedianSortedArraysSimple(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2+1]) / 2
	}
	return float64(nums[len(nums)/2])
}

// double pointer, O(n/2)
func findMedianSortedArraysSimple2(nums1 []int, nums2 []int) float64 {
	i1, i2 := 0, 0
	totalLen := len(nums1) + len(nums2)
	mid := totalLen/2 + 1
	previous, current := 0, 0

	for i1+i2+2 <= mid {

	}

	if totalLen%2 == 0 {
		return float64(current+previous) / 2
	}
	return float64(current)
}

func fromArray(index int, nums []int) (int, bool) {
	if index < len(nums) {
		return nums[index], true
	}
	return 0, false
}

func canGrow(i int, nums []int) bool {
	if i < len(nums) - 1 {
		return true
	}
	return false
}

func max(i1, i2 *int, nums1, nums2 []int) int {
	v1, ok1 := fromArray(*i1, nums1)
	v2, ok2 := fromArray(*i2, nums2)
	if ok1 && ok2 {
		if v1 > v2 {
			if canGrow(*i2, nums2) {
				*i2++
			}else{
				*i1++
			}
			return v1
		}
	}
}
