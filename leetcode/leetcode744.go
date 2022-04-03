package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
    n := len(letters)
    if (target >= letters[n-1] || target < letters[0]) {
        return letters[0]
    }
    l, r, m := 0, n-1, 0
    for l < r {
        m = l + (r - l) / 2
        if (m == l) {
            break
        }
        if (letters[m] <= target) {
            l = m
        } else {
            r = m
        }
    }
    return letters[m+1]
}

func main() {
    letters := []byte {'e', 'e', 'e', 'n', 'n'}
    var target byte = 'e'
    fmt.Printf("%c\n", nextGreatestLetter(letters, target))
}
