package main

import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main() {
	s := "sha1 this tring"	
	h := sha1.New()

	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)

	m := md5.New()
	m.Write([]byte(s))
	fmt.Printf("%x\n",m.Sum(nil))
}