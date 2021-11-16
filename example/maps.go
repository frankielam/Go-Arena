package main

import "fmt"

func main() {
	var k1 map[string]int
	k1 = make(map[string]int)
	k1["k1"] = 1

	k2 := make(map[string]int)
	k2["k2"] = 2

	k3 := map[string]int{"k1": 1, "k2": 2}

	fmt.Println(k1)
	fmt.Println(k2)
	fmt.Println(k3)

	v, f := k3["k3"]
	fmt.Println("k3", v, f)

	fmt.Println(len(k3))
	delete(k3, "k1")
	fmt.Println(k3)
	fmt.Println(len(k3))
	
	_, exist := k3["k2"]
	fmt.Println(exist)
}
