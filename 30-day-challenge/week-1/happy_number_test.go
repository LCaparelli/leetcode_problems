package challenge

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {
	var tests = []struct {
		number int
		want   bool
	}{
		{19, true},
		{18, false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.number)
		t.Run(testName, func(t *testing.T) {
			ans := isHappy(test.number)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}

func TestDigits(t *testing.T) {
	var tests = []struct {
		number int
		want   []int
	}{
		{19, []int{9, 1}},
		{1, []int{1}},
		{1000, []int{0, 0, 0, 1}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.number)
		t.Run(testName, func(t *testing.T) {
			ans := digits(test.number)
			if !areEqual(test.want, ans) {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
