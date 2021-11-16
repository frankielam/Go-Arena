package main

import "fmt"

func genIdx(start int) func() int {
	i := start
	return func() int {
		i++
		return i
	} 
}

func main() {
	gener := genIdx(100)
	fmt.Println(gener())
	fmt.Println(gener())
	fmt.Println(gener())

	newGener := genIdx(1000)
	fmt.Println(newGener())
	fmt.Println(newGener())
}
