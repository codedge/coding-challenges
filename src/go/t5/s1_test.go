package main

import (
	"reflect"
	"testing"
)

func TestKcomplement(t *testing.T) {
	type Input struct {
		k int
		numbers []int
	}

	type test struct {
		input Input
		want int
	}

	tests := []test {
		{
			input: Input {
				k: 6,
				numbers: []int { 1, 8, -3, 0, 1, 3, -2, 4, 5 },
			},
			want: 7,
		},
	}

	for _, tc := range tests {
		got := Kcomplement(tc.input.k, tc.input.numbers)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("input: %v, expected: %v, got: %v", tc.input, tc.want, got)
		}
	}
}