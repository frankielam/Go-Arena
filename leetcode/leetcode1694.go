package main

import (
	"fmt"
	"strings"
)

func reformatNumber(number string) string {
    number = strings.Replace(number, "-", "", -1)
    number = strings.Replace(number, " ", "", -1)
    size := len(number)
    str := ""
    for i := 0; i < size;  {
        if size - i > 4 {
            str = str + number[i:i+3]
            str = str + "-"
            i = i + 3
        } else {
            if size - i == 4 {
                str = str + number[i:i+2]
                str = str + "-"
                str = str + number[i+2:]
            } else {
                str = str + number[i:]
            }
            break
        }
    }
    return str
}

func main() {
	fmt.Println(reformatNumber("1-23-45 6"))
	fmt.Println(reformatNumber("123 4-567"))
	fmt.Println(reformatNumber("123 4-5678"))
}