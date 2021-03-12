package main

import (
	"reflect"
	"testing"
)

func TestJumpOutOfArray(t *testing.T) {
	type test struct {
		input []int
		want int
	}

	tests := []test{
		{input: []int {2, 3, -1, 1, 6}, want: 4},
		{input: []int {-1, 3, -1, 1, 6}, want: 1},
		{input: []int {1, 2, 1, 1, 1, 0}, want: -1},
	}

	for _, tc := range tests {
		got := JumpOutOfArray(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}