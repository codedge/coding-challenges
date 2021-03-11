package main

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	type test struct {
		input int
		want int
	}

	tests := []test{
		{input: 0, want: 0},
		{input: 1, want: 1},
		{input: 2, want: 1},
		{input: 3, want: 2},
		{input: 8, want: 21},
		{input: 38, want: 88169},
	}

	for _, tc := range tests {
		got, err := Fibonacci(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("input: %v, expected: %v, got: %v, err: %v", tc.input, tc.want, got, err)
		}
	}
}