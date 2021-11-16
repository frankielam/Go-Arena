package main

import "fmt"

type animal struct {
	name string
	hand, foot int
}

func (a animal) legs() int {
	return  a.hand + a.foot
}

func (a *animal) canRun() string {
	if a.name == "cat" {
		return "run"
	} else {
		return "unknow"
	}
}

func main() {
	cat := animal{"cat", 2, 2}
	fmt.Println(cat.legs())
	catptr := &cat
	fmt.Println(cat.canRun())
	fmt.Println(catptr.canRun())
}
