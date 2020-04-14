package challenge

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var tests = []struct {
		prices []int
		want   int
	}{
		{[]int{7,1,5,3,6,4}, 7},
		{[]int{1,2,3,4,5}, 4},
		{[]int{7,6,4,3,1}, 0},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.prices)
		t.Run(testName, func(t *testing.T) {
			ans := maxProfit(test.prices)
			if test.want != ans {
				t.Errorf("got %v, want %v", ans, test.want)
			}
		})
	}
}
