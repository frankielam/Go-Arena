package main

func setZeroes(matrix [][]int)  {
    m, n := len(matrix), len(matrix[0])
    var row []int = make([]int, m)
    var col []int = make([]int, n)
    for i, item := range matrix {
        for j, num := range item {
            if num == 0 {
                row[i] = 1
                col[j] = 1
            }
        }
    }
    for i, item := range matrix {
        for j := range item {
            if row[i] == 1 || col[j] == 1 {
                matrix[i][j] = 0
            }
        }
    }
}

func main() {
    d := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
    setZeroes(d)
}