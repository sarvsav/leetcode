// Problem: https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	var alphabets [26]int
	for i := 0; i < len(s); i++ {
		alphabets[s[i]-97]++
	}

	for i := 0; i < len(s); i++ {
		if alphabets[s[i]-97] == 1 {
			return i
		}
	}

	return -1
}
