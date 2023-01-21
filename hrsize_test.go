package main

import (
	"testing"
)

func TestHRSize(t *testing.T) {
	tests := []struct {
		input    HRSize
		expected string
	}{
		{42, "42"},
		{KB, "1.0K"},
		{KB + 100, "1.1K"},
		{MB, "1.0M"},
		{MB + 100*KB, "1.1M"},
		{GB, "1.0G"},
		{GB + 100*MB, "1.1G"},
		{TB, "1.0T"},
		{TB + 100*GB, "1.1T"},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if v := test.input.String(); v != test.expected {
				t.Fatalf("expected %v, have %v", test.expected, v)
			}
		})
	}
}
