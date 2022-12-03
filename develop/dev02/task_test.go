package main

import (
	"errors"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	testSample := []struct {
		input          string
		expectedOutput string
		expectedError  error
	}{
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"abcd", "abcd", nil},
		{"45", "", errors.New("invalid string: begins with a number")},
		{`qwe\4\5`, "qwe45", nil},
		{`qwe\45`, "qwe44444", nil},
		{`\5\ab8`, `5\abbbbbbbb`, nil},
		{"hey3", "heyyy", nil},
	}

	for _, sample := range testSample {
		result, err := unpackString(sample.input)
		if result != sample.expectedOutput && err != sample.expectedError {
			t.Errorf("Tested input: %s\nExpected output: %s, error: %v\n Given output: %s, error: %v\n", sample.input, sample.expectedOutput, sample.expectedError, result, err)
		}
	}
}
