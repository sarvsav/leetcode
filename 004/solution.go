// Problem: https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
    var result []int 
    half := len(nums)/2
    for i := 0; i < half; i++ {
        result = append (result, nums[i])
        result = append (result, nums[i+half])
    }
    return result
}