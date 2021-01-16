// Problem: https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	second := 0
	var result []int
	for key1, val1 := range nums {
		second = target - val1
		for key2, val2 := range nums {
			if second == val2 {
				result = []int{key2, key1}
				break
			}
		}
	}
	return result
}
