package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
    for i, n := 0, len(s); i < n/2; i++ {
        if (s[i]  !=  s[n-i-1]) {
            return 2
        }
    }
    return  1
}

func main() {
    s := "ababa" 
    rst := removePalindromeSub(s)
    fmt.Println(rst) 
}
