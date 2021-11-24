package main

import (
	"fmt"
	"errors"
)

func f1(arg int) (int, error) {
	if arg == 35 {
		return -1, errors.New("can't work with 35")
	}
	return arg + 3, nil
}

type argError struct {
	arg int
	prob string
}

func  (a *argError) Error() string {
	return fmt.Sprintf("%d - %s", a.arg, a.prob)
}

func f2(arg int) (int, error) {
	if arg == 35 {
		return -1, &argError{arg,  "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, v := range []int{35, 17} {
		if r, e := f1(v); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)
		}
	}
	for _, v := range []int{35, 17} {
		if r, e := f2(v); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(r)
		}
	}
}
