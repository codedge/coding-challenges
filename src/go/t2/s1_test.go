package main

import (
	"reflect"
	"testing"
)

func TestNumberToOrdinal(t *testing.T) {
	type test struct {
		input int
		want  string
	}

	tests := []test{
		{input: 1, want: "1st"},
		{input: 2, want: "2nd"},
		{input: 3, want: "3rd"},
		{input: 4, want: "4th"},
		{input: 11, want: "11th"},
		{input: 21, want: "21st"},
		{input: 22, want: "22nd"},
	}

	for _, tc := range tests {
		got := NumberToOrdinal(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}