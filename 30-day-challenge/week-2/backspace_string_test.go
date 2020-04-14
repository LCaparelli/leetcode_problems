package challenge

import (
	"fmt"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	var tests = []struct {
		s    string
		t    string
		want bool
	}{
		{"a#c", "b", false},
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a##c", "#a#c", true},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %s %s", test.s, test.t)
		t.Run(testName, func(t *testing.T) {
			ans := backspaceCompare(test.s, test.t)
			if test.want != ans {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
