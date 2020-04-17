package challenge

import (
	"fmt"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {
	var tests = []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{7, 5, 8}, 4},
		{[]int{1}, 1},
		{[]int{1, 1}, 0},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.stones)
		t.Run(testName, func(t *testing.T) {
			ans := lastStoneWeight(test.stones)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
	}
}
