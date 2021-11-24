package main

import "fmt"

type animal interface {
	foots() int
	canRun() bool
}

type cat struct  {
	head, foot int
	run bool
}

func (c cat) foots() int {
	return c.head + c.foot
}

func (c cat) canRun() bool {
	return c.run
}

func detect(a animal) {
	fmt.Println(a)
	fmt.Println(a.foots())
	fmt.Println(a.canRun())
}

func main() {
	c := cat{head: 2, foot: 2, run: true}
	detect(c)
}
