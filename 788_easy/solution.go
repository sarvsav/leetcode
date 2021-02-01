package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(rotatedDigits(201))
}

func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if isGoodNumber(i) {
			count++
		}
	}
	return count
}

func isGoodNumber(n int) bool {
	digit := n
	remainder := 0
	final := 0
	faceValue := 1
	for digit > 0 {
		remainder = digit % 10
		switch remainder {
		case 3, 4, 7:
			return false
		case 0, 1, 8:
			final = (remainder * faceValue) + final
		case 2:
			final = (5 * faceValue) + final
		case 5:
			final = (2 * faceValue) + final
		case 6:
			final = (9 * faceValue) + final
		case 9:
			final = (6 * faceValue) + final
		}
		digit = digit / 10
		faceValue = faceValue * 10
	}

	if final == n {
		return false
	}
	fh, err := os.OpenFile("data.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	msg := fmt.Sprintln("Digit: ", n, "Final: ", final)
	_, err = fh.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	return true
}
