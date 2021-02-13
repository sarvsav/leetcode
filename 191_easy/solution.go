// Problem: https://leetcode.com/problems/number-of-1-bits/
import (
	"strconv"
	"strings"
)

func hammingWeight(num uint32) int {
	result := strconv.FormatInt(int64(num), 2)
	return strings.Count(result, "1")
}
