//Problem: https://leetcode.com/problems/greatest-common-divisor-of-strings/

func gcdOfStrings(str1 string, str2 string) string {
	lengthOfStr1 := len(str1)
	lengthOfStr2 := len(str2)
	num1, num2 := 0, 0
	if lengthOfStr1 > lengthOfStr2 {
		num1, num2 = lengthOfStr1, lengthOfStr2
	} else {
		num1, num2 = lengthOfStr2, lengthOfStr1
	}
	for remainder := 1; remainder != 0; {
		remainder = num1 % num2
		num1 = num2
		num2 = remainder
	}

	res1 := strings.Count(str1, str1[0:num1])
	if res1*num1 != lengthOfStr1 {
		return ""
	}
	res2 := strings.Count(str2, str2[0:num1])
	if res2*num1 != lengthOfStr2 {
		return ""
	}
	if str1[0:num1] != str2[0:num1] {
		return ""
	}
	return str1[0:num1]
}
