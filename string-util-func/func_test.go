package main

import (
	"fmt"
	"testing"
)

func TestStringUtil(t *testing.T) {
	tests := []struct {
		input     string
		result    string
		operation string
		err       error
	}{
		{"hello", "HELLO", "UPPER", nil},
		{"HELLO", "hello", "LOWER", nil},
		{"hello", "Hello", "TITLE", nil},

		{"", "", "TITLE", fmt.Errorf("invalid input: input string is blank\n")},
		{"hello", "", "CAMEL", fmt.Errorf("invalid input: invalid operation %q was supplied. "+
			"valid value are [%s,%s,%s]\n", "CAMEL", UpperCase, LowerCase, TitleCase)},
	}
	for i, test := range tests {
		result, err := stringUtil(test.input, test.operation)
		// if result is not expected
		if err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("asserstion failure: test #%d failed. expected %v but got %v",
					i+1, test.err, err)
			}
		}
		if result != test.result {
			t.Errorf("asserstion failure: test #%d failed. expected %s but got %s",
				i+1, test.result, result)
		}
	}
}
