// Problem: https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	var result []int
	result = append(result, nums[0])
	for i := 0; i < len(nums)-1; i++ {
		result = append(result, result[i]+nums[i+1])
	}
	return result
}
