package main

import "os"

func main() {
	panic("damn")


	_, err := os.Create("/tmp/tmp_file")
	if err != nil {
		panic(err)
	}
}