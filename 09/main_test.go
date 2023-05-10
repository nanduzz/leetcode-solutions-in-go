package main

import (
	"strconv"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{-1}, false},
		{args{1}, true},
		{args{0}, true},
		{args{121}, true},
		{args{1221}, true},
		{args{12321}, true},
		{args{12331}, false},
		{args{13321}, false},
		{args{-121}, false},
		{args{11}, true},
		{args{10}, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args.x), func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
