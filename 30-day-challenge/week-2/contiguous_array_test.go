package challenge

import (
	"fmt"
	"testing"
)

func TestFindMaxLength(t *testing.T) {
	var tests = []struct {
		numbers []int
		want    int
	}{
		{[]int{0, 1, 0}, 2},
		{[]int{0, 1}, 2},
		{[]int{}, 0},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.numbers)
		t.Run(testName, func(t *testing.T) {
			ans := findMaxLength(test.numbers)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
