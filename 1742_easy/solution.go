//Problem: https://leetcode.com/problems/maximum-number-of-balls-in-a-box/
import "strconv"

func countBalls(lowLimit int, highLimit int) int {
	boxes := make(map[int]int)
	digitSum := 0
	for i := lowLimit; i <= highLimit; i++ {
		for _, v := range []byte(strconv.Itoa(i)) {
			digitSum += int(v - 48)
		}
		boxes[digitSum]++
		digitSum = 0
	}
	max := 0
	bigBox := 0
	for i, v := range boxes {
		if max < v {
			bigBox = i
			max = v
		}
	}
	return boxes[bigBox]
}
 