// Problem: https://leetcode.com/problems/calculate-money-in-leetcode-bank/

func totalMoney(n int) int {
	numberOfWeeks := n / 7
	sum := 0
	i := 0
	for i = 0; i < numberOfWeeks; i++ {
		for j := 1 + i; j <= 7+i; j++ {
			sum = sum + j
		}
		n = n - 7
	}
	tmp := i
	for k := 1; k <= n%7; k++ {
		sum += tmp + k
	}
	return sum
}
