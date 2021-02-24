package main

import (
	"reflect"
	"testing"
)

func TestMaskify(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "12345", want: "12345"},
		{input: "1234-4567-8901", want: "1###-####-8901"},
		{input: "123456789", want: "1####6789"},
	}

	for _, tc := range tests {
		got := Maskify(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}