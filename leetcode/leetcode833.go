package main

import (
    "fmt"
)

func projectionArea(grid [][]int) int {
    c := len(grid[0])

    cnt := 0
    for _, row := range grid {
        max := 0
        for  _, col := range row {
            if col > 0 {
                cnt += 1
            }
            if col > max {
                max = col
            }
        }
        cnt += max
    }
    for i := 0; i < c; i++ {
        max := 0
        for _, row := range grid {
            if row[i] > max {
                max = row[i]
            }
        }
        cnt += max
    }
    return cnt
}

func main() {
    grid := [][]int{{1, 2}, {3, 4}}
    fmt.Println(projectionArea(grid))
}
