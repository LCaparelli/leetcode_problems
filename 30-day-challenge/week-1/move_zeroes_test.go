package week_1

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var tests = []struct {
		numbers []int
		want    []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 1, 0}, []int{1, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.numbers)
		t.Run(testName, func(t *testing.T) {
			moveZeroes(test.numbers)
			if !areEqual(test.numbers, test.want) {
				t.Errorf("got %v, want %v", test.numbers, test.want)
			}
		})
	}
}
