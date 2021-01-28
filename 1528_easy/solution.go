// Problem: https://leetcode.com/problems/shuffle-string/

// In beginning, I tried like this

func restoreString(s string, indices []int) string {
    stringSlice := make([]string, len(s))
    for i:=0; i< len(indices); i++ {
        stringSlice[indices[i]] = string(s[i])
    }
    return strings.Join(stringSlice[:], "")
}

// then realized, that I can a byte array, and easily convert to string using
// string

// For example:
package main

import (
	"fmt"
)

func main() {
	byteSlice := []byte{'c','o','d','i','n','g','t','h','e','r','i','g','h','t','w','a','y'}
	fmt.Println(string(byteSlice)) 
}
