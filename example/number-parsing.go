package main

import (
	"fmt"
	"strconv"
)

func main() {
	p := fmt.Println
	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)
	u2, _ := strconv.ParseUint("-789", 0, 64)
	p(u2)

	a, _ := strconv.Atoi("125")
	p(a)

	_, e := strconv.Atoi("wat")
	p(e)
}