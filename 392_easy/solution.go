// Problem: https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
    count := 0
    if s == "" {
        return true
    }
    for _, val := range t {
        if s[count] == byte(val) {
            count++
        } 
        if count == len(s) {
            return true
        }
    }
    return false
}