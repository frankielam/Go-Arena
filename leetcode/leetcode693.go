package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	k := n % 2
	n = n / 2
    	for {
        	if (n <= 0) {
            		break
        	}
        	if (n % 2 == k) {
            		return false
        	}
        	k = n % 2
        	n = n / 2
    	}
    	return true
}

func main() {
	rst := hasAlternatingBits(5)
	fmt.Println(rst)
	rst = hasAlternatingBits(7)
	fmt.Println(rst)
	rst = hasAlternatingBits(11)
	fmt.Println(rst)
}
