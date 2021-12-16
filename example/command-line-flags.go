package main

import (
	"fmt"
	"flag"
)

func main() {
	p := fmt.Println
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	p("word:", *wordPtr)
	p("numb:", *numbPtr)
	p("fork:", *forkPtr)
	p("svar:", svar)
	p("tail:", flag.Args())
}