package main

import (
	"strconv"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		//{args{1}, 1},
		//{args{-1}, -1},
		//{args{12}, 21},
		//{args{123}, 321},
		//{args{-123}, -321},
		//{args{120}, 21},
		{args{1534236469}, 0},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.args.x), func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
