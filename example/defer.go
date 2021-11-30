package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("create")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writting file")
	fmt.Fprintln(f, "hello")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	err := f.Close()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/tmp.txt")
	defer closeFile(f)
	writeFile(f)
}