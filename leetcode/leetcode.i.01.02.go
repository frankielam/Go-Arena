package main

import (
	"fmt"
	"sort"
)

func CheckPermutation(s1 string, s2 string) bool {
	c1 := []rune(s1)
	c2 := []rune(s2)
	sort.Slice(c1, func(i int, j int) bool { return c1[i] < c1[j] })
	sort.Slice(c2, func(i int, j int) bool { return c2[i] < c2[j] })
	return string(c1) == string(c2)
}

func CheckPermutation2(s1 string, s2 string) bool {
	var a1, a2 [26]int
	for _, c := range s1 {
		a1[c - 'a']++
	}
	for _, c := range s2 {
		a2[c - 'a']++
	}
	return a1 == a2
}

// https://leetcode.cn/problems/check-permutation-lcci/
func main() {
	s1 := "abc"
	s2 := "bca"
	fmt.Println(CheckPermutation(s1, s2))
	fmt.Println(CheckPermutation2(s1, s2))
}