// Problem: https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	s = strings.TrimRight(s, " ")
	last := strings.Split(s, " ")
	return len(last[len(last)-1])

	//Better solution
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}