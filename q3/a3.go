package q3

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	longest, current, currentIndex := 1, 0, 0
	currentMap := make(map[int32]int)
	for i, x := range s {
		if matchedIndex, ok := currentMap[x]; ok && matchedIndex >= currentIndex {
			currentMap[x] = i
			currentIndex = matchedIndex + 1
			current = i - matchedIndex
			continue
		}
		current++
		currentMap[x] = i
		if current > longest {
			longest = current
		}
	}
	return longest
}
