//Problem: https://leetcode.com/problems/word-pattern/
import "strings"

func wordPattern(pattern string, s string) bool {
	wordMap := make(map[byte]string)
	sSplit := strings.Split(s, " ")
	stringMap := make(map[string]int)
	for _, v := range sSplit {
		stringMap[v]++
	}
	if len(pattern) != len(sSplit) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		if val, ok := wordMap[pattern[i]]; !ok {
			wordMap[pattern[i]] = sSplit[i]
			continue
		} else {
			if val != sSplit[i] {
				return false
			}
		}
	}
	if len(wordMap) != len(stringMap) {
		return false
	}
	return true
}
