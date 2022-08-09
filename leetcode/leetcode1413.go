package main

import "fmt"

func minStartValue(nums []int) int {
    min, sum := 1, 0
    for _, num := range nums {
        sum += num
        if sum < min {
            min = sum
        }
    }
    if min > 0 {
        return 1
    }
    return 1 - min
}

func main() {
    nums := []int{-3, 2, -3, 4, 2}
    fmt.Println(minStartValue(nums))
}
