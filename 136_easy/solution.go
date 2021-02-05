// Problem: https://leetcode.com/problems/single-number/

func singleNumber(nums []int) int {
    numMap := make(map[int]int)
    for _, val := range nums {
        numMap[val]++
    }
    for i,v := range numMap {
        if v == 1 {
            return i
        }
    }
    return 0
}
