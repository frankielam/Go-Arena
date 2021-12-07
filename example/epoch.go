package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	sec := now.Unix()
	nano := now.UnixNano()

	p(now)
	p(sec)

	millis := nano / 1000 / 1000
	p(millis)

	p(nano)

	p(time.Unix(sec,0))
	p(time.Unix(sec, 123))

	p(time.Unix(0, nano))
}