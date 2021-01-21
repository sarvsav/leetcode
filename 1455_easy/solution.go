// Problem: https://leetcode.com/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	index := 1
	for _, word := range strings.Split(sentence, " ") {
		if len(word) >= len(searchWord) && word[0:len(searchWord)] == searchWord {
			return index
		}
		index++
	}
	return -1
}