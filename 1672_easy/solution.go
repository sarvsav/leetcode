// Problem: https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
    maxWealth := 0 
    sum := 0 
    for _, val := range accounts {
        for _, money := range val {
            sum += money
        } 
        if maxWealth < sum {
            maxWealth = sum
        }
        sum = 0
    }
    return maxWealth
}
