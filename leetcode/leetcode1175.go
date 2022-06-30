package main

import (
	"fmt"
)

func numPrimeArrangements(n int) int {
    normal, prime := 1, 0
    total := 1
    for num := 2; num <= n; num++ {
        f := true
        for i := 2; i <= num/2; i++ {
            if num % i == 0 {
                normal++
                total *= normal
                f = false
                break
            }
        }
        if f {
            prime++
            total *= prime
        }
        total %= 1000000007
        // fmt.Printf("%d , %d = %d \n", normal, prime, total)
    }
    
    return total
}

func main() {
    fmt.Println(numPrimeArrangements(10))
    fmt.Println(numPrimeArrangements(100))
}
