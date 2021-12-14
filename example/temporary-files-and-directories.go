package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 0, 2, 4})
	check(err)
	dname, err := os.MkdirTemp("","sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)

}