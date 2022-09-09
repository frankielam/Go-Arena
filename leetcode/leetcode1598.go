package main

import "fmt"

/*
* https://leetcode.cn/problems/crawler-log-folder/
*/
func minOperations(logs []string) int {
    ans := 0
    for _, log := range logs {
        if log == "./" {
            continue
        }
        if log == "../" {
            if ans > 0 {
                ans -= 1 
            }
        } else {
            ans++
        }
    }
    return ans
}

func main() {
    logs := []string{"d1/","d2/","../","d21/","./"}
    fmt.Println(minOperations(logs))
}