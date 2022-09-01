package main

import (
    "fmt"
)

func finalPrices(prices []int) []int {
    ans := make([]int, len(prices))
    for i := 0; i < len(prices); i++ {
        ans[i] = prices[i]
        for j := i + 1; j < len(prices); j++ {
            if prices[j] <= prices[i] {
                ans[i] = prices[i] - prices[j]
                break
            }
        }
    }
    return ans
}

func main() {
    fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
}
