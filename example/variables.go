package main

import "fmt"

func main() {
	var i int = 1
	var j, k int = 2, 3
	fmt.Println(i)		//1
	fmt.Println(j, k)	//2 3
	fmt.Println(i, j, k)	//1 2 3

	var a = true
	var b bool
	fmt.Println(a, b)	//true false

	var c int
	d := 4
	fmt.Println(c, d)	//0 4
}
