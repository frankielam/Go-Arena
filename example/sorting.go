package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 2}
	fmt.Println(sort.IntsAreSorted(arr))
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(sort.IntsAreSorted(arr))

	str := []string{"c", "b", "d"}
	sort.Strings(str)
	fmt.Println(str)
}