package main

import (
    "fmt"
)

// https://leetcode.cn/problems/validate-stack-sequences/
func validateStackSequences(pushed []int, popped []int) bool {
    var slice []int
    i, j := 0, 0
    for {
        for len(slice) > 0 && slice[len(slice)-1] == popped[j] {
            slice = slice[:(len(slice)-1)]
            j++
            if j == len(popped) {
                break
            }
        }
        if i < len(pushed) {
            slice = append(slice, pushed[i])
            i++
        } else {
            break
        }
    }

    if len(slice) == 0  {
        return true
    }
    return false
}

func main() {
    pushed := []int{2,1,0}
    popped := []int{1,2,0}
    fmt.Println(validateStackSequences(pushed, popped))
}
