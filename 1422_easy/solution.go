// Problem: https://leetcode.com/problems/maximum-score-after-splitting-a-string/

func maxScore(s string) int {
	max := 0
	sum := 0
	zeroes := 0
	ones := 0
	for i := 1; i < len(s); i++ {
		zeroes, _ = countVals(s[0:i])
		_, ones = countVals(s[i:])
		sum = zeroes + ones
		if max < sum {
			max = sum
		}
	}
	return max
}

func countVals(a string) (int, int) {
	zero := 0
	one := 0
	for _, val := range a {
		if val == 48 {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}
