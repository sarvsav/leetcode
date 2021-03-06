// Problem: https://leetcode.com/problems/n-th-tribonacci-number/

func tribonacci(n int) int {
	seq := []int{0, 1, 1}
	if n <= 2 {
		return seq[n]
	}

	for i := 2; i < n; i++ {
		seq[0], seq[1], seq[2] = seq[1], seq[2], seq[0]+seq[1]+seq[2]
	}
	return seq[2]
}
