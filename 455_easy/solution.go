// Problem: https://leetcode.com/problems/assign-cookies/
import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	tmp := 0
	for i := 0; i < len(g); i++ {
		for j := tmp; j < len(s); j++ {
			if g[i] <= s[j] {
				count++
				tmp = j + 1
				break
			}
		}
	}
	return count
}
