package challenge

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		numbers []int
		want    int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.numbers)
		t.Run(testName, func(t *testing.T) {
			ans := maxSubArray(test.numbers)
			if ans != test.want {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
