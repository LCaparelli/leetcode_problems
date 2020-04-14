package challenge

import (
	"fmt"
	"testing"
)

func TestCountElements(t *testing.T) {
	var tests = []struct {
		numbers []int
		want    int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 1, 3, 3, 5, 5, 7, 7}, 0},
		{[]int{1, 3, 2, 3, 5, 0}, 3},
		{[]int{1, 1, 2, 2}, 2},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.numbers)
		t.Run(testName, func(t *testing.T) {
			ans := countElements(test.numbers)
			if test.want != ans {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
