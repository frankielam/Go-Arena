package main

import "fmt"

func main() {
	i := 0

	for i<3 {
		//fmt.Println("i is", i++)	// syntax error: unexpected ++, 
		fmt.Println("i is", i)
		i++
	}

	//for j := 0;  j < 3 {	//syntax error: unexpected {
	for j := 0;  j < 3; j++ {
		fmt.Println("j is", j)
	//	j++
	}
	
	for {
		fmt.Println("loop")
		break
	}

	for k := 0; k < 10; k++ {
		if k%2 ==0 {
			continue
		}
		fmt.Println("k is", k)
	}
}
