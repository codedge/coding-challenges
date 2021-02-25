package main

import (
	"reflect"
	"testing"
)

func TestCalculate(t *testing.T) {
	type test struct {
		input string
		want float64
	}

	tests := []test{
		{input: "2 5 +", want: 7},
		{input: "5 3 -", want: 2},
		{input: "4 5 *", want: 20},
		{input: "12 4 /", want: 3},
		{input: "", want: 0},
		{input: "4 7-", want: -1},
		{input: "9 8 7.5", want: 7.5},
		{input: "4 5 6", want: 6},
	}

	for _, tc := range tests {
		got, err := Calculate(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("input: %v, expected: %v, got: %v, err: %v", tc.input, tc.want, got, err)
		}
	}
}