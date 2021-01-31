// Problem: https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
    magMap := make(map[rune]int)
    for _, val := range magazine {
        magMap[val]++
    }
    
    for _, val := range ransomNote {
        if _, ok := magMap[val]; !ok {
            return false
        } 
        magMap[val]--
        if magMap[val] < 0 {
            return false
        } 
    }
    return true
}
