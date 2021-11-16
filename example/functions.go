package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func latlng(a, b float32) (float32, float32) {
	return a, b
}

func cal(nums ...int) {
	fmt.Print(nums, "\t")
	total :=0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	total := sum(2, 3)
	fmt.Println("a + b =", total)

	lat, lng := latlng(20.234, 112.123)
	fmt.Println("latitude", lat, "longitude", lng)

	nums := []int{1, 2, 3, 4}
	cal(nums...)
	cal(10, 20, 30)
}
