package challenge

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		numbers []int
		want    int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{0}, 0},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.numbers)
		t.Run(testName, func(t *testing.T) {
			ans := singleNumber(test.numbers)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
