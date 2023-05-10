package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{args{"42"}, 42},
		{args{"   -42"}, -42},
		{args{"4193 with words"}, 4193},
		{args{"words and 987"}, 0},
		{args{"-91283472332"}, -2147483648},
		{args{"+1"}, 1},
		{args{"+-12"}, 0},
		{args{"-+12"}, 0},
		{args{"   +0 123"}, 0},
		{args{"20000000000000000000"}, 2147483647},
		{args{"2147483648"}, 2147483647},
		{args{"  0000000000012345678"}, 12345678},
		{args{"    0000000000000   "}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
