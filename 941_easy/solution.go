// Problem: https://leetcode.com/problems/valid-mountain-array/

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[len(arr)-2] == arr[len(arr)-1] {
		return false
	}
	if arr[0] >= arr[1] {
		return false
	}
	t1, t2, flag := 0, 0, 0
	for i := 0; i < len(arr)-1; i++ {
		t1, t2 = arr[i], arr[i+1]
		switch flag {
		case 0:
			if t1 < t2 {
				continue
			} else {
				flag++
			}
		case 1:
			if t1 <= t2 {
				return false
			}
		}
	}
	if flag == 0 {
		return false
	}
	return true
}