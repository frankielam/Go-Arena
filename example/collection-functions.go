package main

import (
	"fmt"
	"strings"
)

func Index(strs []string, str string) int {
	for i, w := range strs {
		if w == str {
			return i
		}
	} 
	return -1
}

func Contain(strs []string, str string) bool {
	return Index(strs, str) >= 0 
}

func Any(strs []string, f func(string) bool) bool {
	for _, w := range strs {
		if f(w) {
			return true
		}
	}
	return false
}

func All(strs []string, f func(string) bool) bool {
	for _, w := range strs {
		if !f(w) {
			return false
		}
	}
	return true
}

func Filter(strs []string, f func(string) bool) []string {
	arr := make([]string, 0)
	for  _, w := range strs {
		if f(w) {
			arr = append(arr, w)
		}
	}
	return arr
}

func Map(strs []string, f func(string) string) []string {
	mymap := make([]string, len(strs))
	for  idx, w := range strs {
		mymap[idx] = f(w)
	}
	return mymap
} 


func main() {
	str := []string{"apple", "ada",  "admin", "bob"}

	fmt.Println(Index(str, "bob"))

	fmt.Println(Contain(str, "admin"))


	fmt.Println(Any(str, func(w string) bool {
		return strings.HasPrefix(w, "a")
	}))

	fmt.Println(All(str, func(w string) bool {
		return strings.HasPrefix(w, "a")
	}))

	fmt.Println(Filter(str, func(w string) bool {
		return strings.Contains(w, "ad")
	}))

	fmt.Println(Map(str, strings.ToUpper))



}