// Problem: https://leetcode.com/problems/maximum-product-of-three-numbers/
import "sort"

func maximumProduct(nums []int) int {

	sort.Ints(nums)

	l := len(nums)
	temp1 := nums[l-1] * nums[l-2] * nums[l-3]
	temp2 := nums[0] * nums[1] * nums[l-1]

	if temp1 > temp2 {
		return temp1
	}
	return temp2

}
