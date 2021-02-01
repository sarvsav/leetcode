// Problem: https://leetcode.com/problems/bulb-switcher-iv/

func minFlips(target string) int {
	flips := 0
	sliceString := []byte(target)
	length := len(target)
	bulbSet := make([]byte, length)
	for i := 0; i < length; i++ {
		if bulbSet[i] == sliceString[i]-48 {
			continue
		}

		for j := i; j < length; j++ {
			bulbSet[j] = bulbSet[j] ^ 1
		}
		flips++
	}
	return flips
}
