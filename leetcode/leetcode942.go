package main

import (
	"fmt"
)

func diStringMatch(s string) []int {
	len := len(s)
	var ans []int = make([]int, len + 1)
	b, e := 0, len
	for i, c := range s {
		if c == 'I' {
			ans[i] = b
			b++
		} else {
			ans[i] = e
			e--
		}
		if i == len - 1 {
			ans[len] = e
		}
	}
	return ans
}

func diStringMatch2(s string) (ans []int) {
	len := len(s)
	
	b, e := 0, len
	for i, c := range s {
		if c == 'I' {
			ans = append(ans, b)
			b++
		} else {
			ans = append(ans, e)
			e--
		}
		if i == len - 1 {
			ans = append(ans, b)
		}
	}
	return 
}

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch2("III"))
	fmt.Println(diStringMatch("DDD"))
	fmt.Println(diStringMatch("IIID"))
}