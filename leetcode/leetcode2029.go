package main

import (
	"fmt"
)

func stoneGameIX(stones []int) bool {
    cnt0, cnt1, cnt2 := 0, 0, 0
    for _, val := range stones {
        if (val % 3  == 0) {
            cnt0++;
        } else if (val % 3  == 1) {
            cnt1++;
        } else {
            cnt2++;
        }
    }
    if (cnt0 % 2 == 0) {
        return cnt1 > 0 && cnt2 > 0
    }
    return cnt1 - cnt2 > 2  || cnt2 - cnt1 > 2
}

func main() {
    stones := []int{5, 1, 2, 4, 3}
    rst := stoneGameIX(stones)
    fmt.Println(rst) 
}
