package main

import "fmt"

const s string = "HELLO THERE"

func main() {
	fmt.Println(s)

	const no = 12345
	fmt.Println(no)	

	const num int = 123
	fmt.Println(num)

	fmt.Println(int64(no/num))
}
