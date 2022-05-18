package main

import (
	"fmt"
)

func findKthNumber(m int, n int, k int) int {
    left, right := 1, m * n
    for left < right {
        mid := left + (right - left) / 2
        count := mid / n * n
        for i := mid / n + 1; i <= m; i++ {
            count += mid / i
        }
        if count >= k {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func main() {
	fmt.Println(findKthNumber(9895, 28405, 100787757))
}