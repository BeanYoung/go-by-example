package main

import (
	"fmt"
	"testing"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestMinIntBase(t *testing.T) {
	ans := MinInt(1, 2)
	if ans != 1 {
		t.Errorf("MinInt(1, 2) should return 1 but return %d", ans)
	}
}

func TestMinIntTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		ans  int
	}{
		{0, 1, 0},
		{1, 2, 1},
		{-1, 1, -1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d %d", test.a, test.b), func(t *testing.T) {
			ans := MinInt(test.a, test.b)
			if ans != test.ans {
				t.Errorf("Min(%d, %d) should return %d", test.a, test.b, test.ans)
			}
		})
	}
}
