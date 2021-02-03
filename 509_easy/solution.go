// Problem: https://leetcode.com/problems/fibonacci-number/

func fib(n int) int {
	first := 0
	second := 1
	if n <= 0 {
		return first
	}
	if n == 1 {
		return second
	}
	prevprev := first
	prev := second
	curr := 0
	for i := 2; i <= n; i++ {
		curr = prev + prevprev
		prevprev = prev
		prev = curr
	}

	return curr

}
