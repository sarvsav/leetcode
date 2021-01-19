// Problem: 832. Flipping an Image https://leetcode.com/problems/flipping-an-image/

// The trick here is to reverse the number, by subtracting 1 and multiplying the result by 1.
// So, 0 will give (0-1)*-1 = 1
// and 1 will give (1-1)*-1 = 0

func flipAndInvertImage(A [][]int) [][]int {
	var result [][]int
	var temp []int
	for i := 0; i < len(A); i++ {
		for j := len(A[i]) - 1; j >= 0; j-- {
			temp = append(temp, (A[i][j]-1)*-1)
		}
		result = append(result, temp)
		temp = nil
	}
	return result
}
