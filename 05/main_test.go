package main

import "testing"

func Test_longestPalindrome(t *testing.T) {

	tests := []struct {
		arg  string
		want string
	}{
		{"a", "a"},
		{"bb", "bb"},
		{"ccd", "cc"},
		{"aba", "aba"},
		{"aaaa", "aaaa"},
		{"baaaa", "aaaa"},
		{"aaaab", "aaaa"},
		{"baaaab", "baaaab"},
		{"aaaaa", "aaaaa"},
		{"babad", "bab"},
		{"ababab", "ababa"},
		{"cbbd", "bb"},
		{"abba", "abba"},
		{"abcba", "abcba"},
		{"socorrammesubinoonibusemmarrocos", "socorrammesubinoonibusemmarrocos"},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := longestPalindrome(tt.arg); got != tt.want {
				t.Errorf("longestPalindrome() -> got %v, want %v", got, tt.want)
			}
		})
	}
}
