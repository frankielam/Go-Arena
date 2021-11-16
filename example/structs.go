package main

import "fmt"

type cat struct {
	name string
	age int
}

func baby(mother string) cat {
	b := cat{"baby of " + mother, 0}
	return b
}

func main() {
	alice := cat{"alice", 2}
	fmt.Println(alice)
	bob := cat{name: "bob"}
	fmt.Println(bob)
	fmt.Println(&bob)
	p := &bob
	fmt.Println(bob.name)
	fmt.Println(p.age)
	bob.name = "bob 2"	
	fmt.Println(bob.name)

	fmt.Println(baby("tom"))
}
