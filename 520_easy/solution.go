// Problem: https://leetcode.com/problems/detect-capital/

func detectCapitalUse(word string) bool {
	caps, lower := 0, 0

	for i := 1; i < len(word); i++ {
		if word[i] >= 65 && word[i] <= 90 {
			caps++
			continue
		}
		lower++
	}
	if lower == len(word)-1 {
		return true
	}
	if caps == len(word)-1 && word[0] >= 65 && word[0] <= 90 {
		return true
	}
	return false
}
