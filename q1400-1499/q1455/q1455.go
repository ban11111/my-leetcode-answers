package q1455

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	split := strings.Split(sentence, " ")
	for i := range split {
		if strings.HasPrefix(split[i], searchWord) {
			return i + 1
		}
	}
	return -1
}
