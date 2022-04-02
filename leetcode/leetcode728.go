package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	var list []int
	for left <= right {
		n := left
		f := true
		for n > 0 {
			if (n % 10 == 0 || left % (n % 10) != 0) {
				f = false
				break
			}
			n = n /10
		}
		if f {
			list = append(list, left)
		}
		left++
	}
	return list
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(selfDividingNumbers(47, 85))
}
