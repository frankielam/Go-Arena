package main

import "fmt"

func main() {
	var a [7]int
	fmt.Println(a)
	fmt.Println(len(a))
	a[0] = 1
	fmt.Println(a)

	var b  = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b)

	var c [2][2]int
	fmt.Println(len(c[0]))
	fmt.Println(c)

	d := [2][2]int{{0, 1}, {2, 3}}
	fmt.Println(d)
}
