//Problem: https://leetcode.com/problems/find-the-highest-altitude/

func largestAltitude(gain []int) int {
    maxAltitude := 0
    current := 0 
    for _, value := range gain {
        current = current + value
        if current > maxAltitude {
            maxAltitude = current
        }
    }
    return maxAltitude
}
