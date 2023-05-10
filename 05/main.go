package main

// Had problems figuring this out for myself
// got solution from here:
// https://leetcode.com/problems/longest-palindromic-substring/solutions/3497051/c-java-python-javascript-detailed-explanation-easy-solution-2-approach/
// I've just converted the java code to go
func longestPalindrome(s string) string {
	n := len(s)
	startingIndex := 0
	maxLen := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if checkPosition(s, i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					startingIndex = i
				}
			}
		}
	}

	return s[startingIndex : startingIndex+maxLen]
}

func checkPosition(s string, i int, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
