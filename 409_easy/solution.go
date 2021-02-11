// Problem: https://leetcode.com/problems/longest-palindrome/

func longestPalindrome(s string) int {
	wordMap := make(map[rune]int)
	for _, v := range s {
		wordMap[v]++
	}

	count := 0
	i := 0
	for _, v := range wordMap {
		if v%2 == 0 {
			count = count + v
		} else {
			count = count + v - 1
			i = 1
		}
	}
	return count + i
}
