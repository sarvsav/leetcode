// Problem: https://leetcode.com/problems/find-the-difference/submissions/
func findTheDifference(s string, t string) byte {
	wordFreq := make(map[rune]int)
	for _, v := range s {
		wordFreq[v]++
	}

	for _, v := range t {
		wordFreq[v]++
	}
	for i, v := range wordFreq {
		if v%2 != 0 {
			return byte(i)
		}
	}
	return byte('a')
}

//Better solution: https://leetcode.com/problems/find-the-difference/discuss/1037462/XOR-solution by _dkoval

func findTheDifference(s string, t string) byte {
	var out = t[len(t)-1]

	for i := 0; i < len(s); i++ {
		out ^= s[i] ^ t[i]
	}
	return out
}