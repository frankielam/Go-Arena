package main

import (
    "fmt"
)

func maxLengthBetweenEqualCharacters(s string) int {
    var kv map[rune]int = make(map[rune]int)
    ans := -1
    for idx, k := range s {
        v, f := kv[k]
        if f == true {
            len := idx - v - 1
            if len > ans {
                ans = len
            }
        } else {
            kv[k] = idx
        }
    }
    return ans
}

func main() {
    fmt.Println(maxLengthBetweenEqualCharacters("aa"))
    fmt.Println(maxLengthBetweenEqualCharacters("abca"))
    fmt.Println(maxLengthBetweenEqualCharacters("cbzxy"))
}
