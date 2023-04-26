package stringCalculator

import (
	"errors"
	"testing"
)

func TestStringCalculator(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "1",
			input:    "",
			expected: 0,
		},
		{
			name:     "2",
			input:    "1",
			expected: 1,
		},
		{
			name: "3",

			input:    "3",
			expected: 3,
		},
		{
			name: "4",

			input:    "1,2",
			expected: 3,
		},
		{
			name: "5",

			input:    "3,5",
			expected: 8,
		},
		{
			name:     "6",
			input:    "1,2,3",
			expected: 6,
		},
		{
			name:     "7",
			input:    "3,5,3,9",
			expected: 20,
		},
		{
			name:     "8",
			input:    "1,2\n3",
			expected: 6,
		},
		{
			name:     "9",
			input:    "3\n5\n3,9",
			expected: 20,
		},
		{
			name:     "10",
			input:    "//;\n1;2",
			expected: 3,
		},
		{
			name:     "11",
			input:    "1000,1002,2",
			expected: 2,
		},
		{
			name:     "12",
			input:    "//[***]\\n1***2***3",
			expected: 6,
		},
	}

	for _, testCase := range testCases {
		got, _ := Add(testCase.input)
		want := testCase.expected
		if got != want {
			t.Errorf("test %q: got %d want %d", testCase.name, got, want)
		}
	}
}

func TestNegativeNumberShouldReturnError(t *testing.T) {
	_, got := Add("-1,2,-3")
	want := errors.New("negatives not allowed")
	if got == nil || want.Error() != got.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}
