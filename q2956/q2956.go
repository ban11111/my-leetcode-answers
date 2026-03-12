package q2956

import (
	"sort"
)

func findIntersectionValuesV1(nums1 []int, nums2 []int) []int {
	var map2 = make(map[int]struct{}, len(nums2))
	for i := range nums2 {
		map2[nums2[i]] = struct{}{}
	}
	var result = make([]int, 2)
	var map1 = make(map[int]struct{}, len(nums1))
	for i := range nums1 {
		if _, ok := map2[nums1[i]]; ok {
			result[0]++
		}
		map1[nums1[i]] = struct{}{}
	}
	for i := range nums2 {
		if _, ok := map1[nums2[i]]; ok {
			result[1]++
		}
	}
	return result
}

func newInt(i int) *int {
	return &i
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var result = make([]int, 2)
	var last *int
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if last != nil {
			if i < len(nums1) && nums1[i] == *last {
				result[0]++
				i++
				continue
			}
			if j < len(nums2) && nums2[j] == *last {
				result[1]++
				j++
				continue
			}
		}
		if i >= len(nums1) || j >= len(nums2) {
			break
		}
		if nums1[i] == nums2[j] {
			last = newInt(nums1[i])
			result[0]++
			result[1]++
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}

	}
	return result
}
