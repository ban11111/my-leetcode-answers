package find_replications

import "sort"

// personal demand: find and mark duplicated items
// returns index of all of them
// ex: ["a", "b", "c", "b", "d", "a"] returns [0, 1, 3, 5]
func FindReplications(items []string) []int {
	// hash method
	hashMap := make(map[string]*[2]int)
	// int8[0]: 0=none, 1=first_encounter, 2=duplicated,
	// int8[1]: first_index
	result := make([]int, 0, len(items))
	for i, item := range items {
		key, ok := hashMap[item]
		if !ok {
			hashMap[item] = &[2]int{1, i}
			continue
		}
		switch key[0] {
		case 1:
			hashMap[item][0] = 2
			result = append(result, hashMap[item][1], i)
		case 2:
			result = append(result, i)
		default:
			panic("???how???")
		}
	}
	sort.Ints(result)
	return result
}