package main

import "math"

/*
Description:
Implement the myAtoi(string s) function, which converts a string to a 32-bit
signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

Read in and ignore any leading whitespace.
Check if the next character (if not already at the end of the string) is '-' or
'+'. Read this character in if it is either. This determines if the final
result is negative or positive respectively. Assume the result is positive if
neither is present.
Read in next the characters until the next non-digit character or the end of
the input is reached. The rest of the string is ignored.
Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no
digits were read, then the integer is 0. Change the sign as necessary (from step 2).
If the integer is out of the 32-bit signed integer range [-2^31, 2^31 - 1], then
clamp the integer so that it remains in the range. Specifically, integers less
than -2^31 should be clamped to -2^31, and integers greater than 2^31 - 1 should
be clamped to 2^31 - 1.
Return the integer as the final result.
*/

var conversionMap = map[uint8]uint8{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func myAtoi(s string) int {
	var n int64 = 0
	digits := make([]uint8, 200)
	neg := false
	started := false
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			if started {
				break
			}
			neg = true
			started = true
			continue
		} else if s[i] == '+' {
			if started {
				break
			}
			started = true
			continue
		}
		v, isNumber := conversionMap[s[i]]
		if isNumber {
			digits = append(digits, v)
			started = true
			j++
		} else if s[i] != ' ' || started {
			break
		}
	}

	firstNonZero := 0
	for i := 0; i < len(digits); i++ {
		if digits[i] != 0 {
			firstNonZero = i
			break
		} else if i == len(digits)-1 { //There is only zeros
			firstNonZero = i
		}
	}
	digits = digits[firstNonZero:]

	m := len(digits) - 1
	if m > 11 { //max length of int32
		if neg {
			return math.MinInt32
		}
		return math.MaxInt32
	} else if m < 0 {
		return 0
	}

	for i := 0; i < len(digits); i++ {
		n += int64(digits[i]) * int64(math.Pow(float64(10), float64(m)))
		if n > math.MaxInt32+1 {
			break
		}
		m--
	}

	if neg {
		n = -1 * n
	}

	if n > math.MaxInt32 {
		return math.MaxInt32
	} else if n < math.MinInt32 {
		return math.MinInt32
	}
	return int(n)
}
