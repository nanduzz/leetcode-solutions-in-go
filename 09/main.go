package main

import (
	"math"
)

// This is my solution for the follow up question :
// Follow up: Could you solve it without converting the integer to a string?
//
// Maybe a better solution would be using a Stack.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	hasEvenNumberOfDigits, numberOfDigits := hasEvenNumberOfDigitsFn(x)

	mid := numberOfDigits / 2

	digits := make([]uint8, mid)
	for i := 0; i < mid; i++ {
		lastDigit := uint8(x % 10)
		digits[i] = lastDigit
		x /= 10
	}

	if hasEvenNumberOfDigits {
		x /= 10
	}

	i := 1
	for x > 0 {
		lastDigit := uint8(x % 10)
		if lastDigit != digits[len(digits)-i] {
			return false
		}
		i++
		x /= 10
	}

	return true
}

func hasEvenNumberOfDigitsFn(x int) (bool, int) {
	numberOfDigits := int(math.Log10(float64(x))) + 1
	return numberOfDigits%2 != 0, numberOfDigits
}
