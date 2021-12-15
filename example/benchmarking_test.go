package main

import (
	"testing"
)

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BenchmarkIntMax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMax(2, 3)
	}
}
