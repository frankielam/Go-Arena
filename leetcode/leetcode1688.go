package main

import "fmt"

func numberOfMatches(n int) int {
    cnt := 0
    for n > 1 {
        cnt += n / 2
        n = n / 2 + n % 2
    }
    return cnt
}

func main() {
    rst := numberOfMatches(14)
    fmt.Println(rst)
}
