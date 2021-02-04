// Problem: https://leetcode.com/problems/robot-return-to-origin/

func judgeCircle(moves string) bool {
	h, v := 0, 0 //horizontal and vertical movements
	for _, val := range moves {
		switch val {
		case 85:
			v += 1
		case 68:
			v -= 1
		case 82:
			h += 1
		case 76:
			h -= 1
		}
	}
	if h == 0 && v == 0 {
		return true
	}
	return false
}

// 2 things can be improved here.
// 1. Use if, it is fast as compare to switch statements
// 2. remove if condition in last, instead do like this:
// return h == 0 && v == 0