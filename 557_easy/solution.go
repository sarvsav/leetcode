// Problem: https://leetcode.com/problems/reverse-words-in-a-string-iii/
import "strings"

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	sSlice := strings.Split(s, " ")
	rev := make([]string, len(sSlice))
	for _, v := range sSlice {
		rev = append(rev, reverse(v))
	}
	return strings.TrimLeft(strings.Join(rev, " "), " ")
}

func reverse(s string) string {
	length := len(s)
	result := make([]byte, length)
	for i := 0; i <= length/2; i++ {
		result[i], result[length-i-1] = s[length-i-1], s[i]
	}
	return string(result)
}
