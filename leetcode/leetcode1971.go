package main

import (
    "fmt"
)

func findCenter(edges [][]int) int {
    if edges[0][0] == edges[1][0] || edges[0][0] ==edges[1][1] {
        return edges[0][0]
    }
    return edges[0][1]
}

func main() {
    arr := [][]int {{1,2}, {2,3}, {2,4}}
    rst := findCenter(arr)
    fmt.Println(rst)
    arr = [][]int {{1,2}, {5,1}, {1,3}, {1, 4}}
    rst = findCenter(arr)
    fmt.Println(rst)
}
