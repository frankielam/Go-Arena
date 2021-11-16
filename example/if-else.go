package main

import "fmt"

func main() {
	if 8%2 == 0 {
		fmt.Println("8 % 2 == 0")
	}

	n := 4
	if n%3 == 0 {
		fmt.Println(n, "% 3 == 0")
	} else {
		fmt.Println(n, "% 3 != 0")
	}

	if m := 9; m < 9 {
		fmt.Println(m, "less than 9")
	} else if m == 9 {
		fmt.Println(m, "equals 9")
	} else {
		fmt.Println(m, "large than 9")
	}
}
