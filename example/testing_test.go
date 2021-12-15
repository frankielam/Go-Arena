package main

import (
	"fmt"
	"testing"
)

func IntMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestIntMax(t *testing.T) {
	ans := IntMax(2, -2)
	if ans != 2 {
		t.Errorf("IntMax(2, -2) = %d; want 2", ans)
	}
}

func TestIntMaxTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	} {
		{1, 2, 2},
		{3, 2, 3},
		{5, 2, 5},
		{2, 4, 4},
	}
	for _, tt :=  range tests {
		testname := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMax(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
