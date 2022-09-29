package main

import "fmt"

func isFlipedString(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    if len(s1) == 0 {
        return true
    }
    for idx1 := range s1 {
        for idx2 := range s2 {
            id := (idx1 + idx2) % len(s1)
            if s1[id] != s2[idx2] {
                break
            }
            if idx2 == len(s2) - 1 {
                return true
            }
        }
    }
    return false
}

func main() {
	fmt.Println(isFlipedString("waterbottle", "erbottlewat"))
}