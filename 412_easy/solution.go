// Problem: https://leetcode.com/problems/fizz-buzz/

import "strconv"

func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			result = append(result, "Fizz")
			continue
		}
		if i%5 == 0 {
			result = append(result, "Buzz")
			continue
		}
		result = append(result, strconv.Itoa(i))
	}
	return result
}
