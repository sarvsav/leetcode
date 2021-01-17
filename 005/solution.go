// Problem: https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool
	var max int
	for _, val := range candies {
		if val > max {
			max = val
		}
	}
	for _, val := range candies {
		if val+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
