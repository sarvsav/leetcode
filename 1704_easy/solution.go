// Problem: https://leetcode.com/problems/determine-if-string-halves-are-alike/
import "strings"

func halvesAreAlike(s string) bool {
	subStringFirst := strings.ToLower(s[0 : len(s)/2])
	subStringSecond := strings.ToLower(s[len(s)/2:])
	first, second := 0, 0
	for i := 0; i < len(subStringFirst); i++ {
		if subStringFirst[i] == 'a' || subStringFirst[i] == 'e' || subStringFirst[i] == 'i' || subStringFirst[i] == 'o' || subStringFirst[i] == 'u' {
			first++
		}
		if subStringSecond[i] == 'a' || subStringSecond[i] == 'e' || subStringSecond[i] == 'i' || subStringSecond[i] == 'o' || subStringSecond[i] == 'u' {
			second++
		}
	}
	return first == second
}

// Better solution
func halvesAreAlike(s string) bool {
	counter := 0
	for i := 0; i < len(s)/2; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			counter++
		}
		if s[len(s)-i-1] == 'A' || s[len(s)-i-1] == 'E' || s[len(s)-i-1] == 'I' || s[len(s)-i-1] == 'O' || s[len(s)-i-1] == 'U' || s[len(s)-i-1] == 'a' || s[len(s)-i-1] == 'e' || s[len(s)-i-1] == 'i' || s[len(s)-i-1] == 'o' || s[len(s)-i-1] == 'u' {
			counter--
		}
	}
	return counter == 0
}