package challenge

import (
	"fmt"
	"testing"
)

func TestStringShift(t *testing.T) {
	var tests = []struct {
		s      string
		shifts [][]int
		want   string
	}{
		{"abc", [][]int{{0, 1}, {1, 2}}, "cab"},
		{"abc", [][]int{{0, 1}, {1, 2}, {1, 1}}, "bca"},
		{"mecsk", [][]int{{1, 4}, {0, 5}, {0, 4}, {1, 1}, {1, 5}}, "kmecs"},
		{"yisxjwry", [][]int{{1, 8}, {1, 4}, {1, 3}, {1, 6}, {0, 6}, {1, 4}, {0, 2}, {0, 1}}, "yisxjwry"},
		{"yzeuobhwxatulpruiab", [][]int{{1, 15}, {0, 18}, {0, 12}, {0, 7}, {0, 7}, {1, 17}, {1, 15}, {0, 9}, {1, 4}, {0, 19}, {1, 16}, {0, 7}, {1, 4}, {1, 17}, {1, 19}, {1, 10}, {1, 2}, {0, 18}, {1, 15}}, "zeuobhwxatulpruiaby"},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v\t%v", test.s, test.shifts)
		t.Run(testName, func(t *testing.T) {
			ans := stringShift(test.s, test.shifts)
			if ans != test.want {
				t.Errorf("got %s, want %s", ans, test.want)
			}
		})
	}
}
