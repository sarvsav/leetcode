// Problem: https://leetcode.com/problems/maximum-number-of-balloons/

func maxNumberOfBalloons(text string) int {
	runeMap := make(map[rune]int)
	for _, v := range text {
		if v == 'b' || v == 'a' || v == 'l' || v == 'n' || v == 'o' {
			runeMap[v]++
		}
	}
	i := 0
	for {

		if v, _ := runeMap['b']; v == 0 {
			return i
		}
		runeMap['b']--

		if v, _ := runeMap['a']; v == 0 {
			return i
		}
		runeMap['a']--

		if v, _ := runeMap['l']; v == 0 {
			return i
		}
		runeMap['l']--

		if v, _ := runeMap['l']; v == 0 {
			return i
		}
		runeMap['l']--

		if v, _ := runeMap['o']; v == 0 {
			return i
		}
		runeMap['o']--

		if v, _ := runeMap['o']; v == 0 {
			return i
		}
		runeMap['o']--

		if v, _ := runeMap['n']; v == 0 {
			return i
		}
		runeMap['n']--
		i++
	}
	return i
}

// Better solution: get the minimum value of b, a, l/2, o/2, and n from the map