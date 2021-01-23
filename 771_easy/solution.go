// Problem: https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
    jewelFreq := make(map[rune]int)
    for _, value := range jewels {
        jewelFreq[value] += 1
    }
    count := 0
    for _,v := range stones {
        if _, ok := jewelFreq[v]; ok {
            count++
        }
    }
    return count
}

// Fast version (Help from internet)
func numJewelsInStones(J string, S string) int {
    charMap, count := make(map[rune]rune), 0
    for _, c := range J {
		charMap[c] = c
	}
    
    for _, c := range S {
        if charMap[c] == c {
            count++
        }
	}
    
    return count
}

// This is valid:
package main

import (
	"fmt"
)

func main() {
	a := make(map[string]int)
	if a["hello"] == 1 { //It won't complain that the key doesn't exist
	  fmt.Println("hi")
	}
}
