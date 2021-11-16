package main

import "fmt"

func main() {
	var s string = "hello " + "world"
	var x int = 3
	var y int = 4
	var f float32 = 3.0
//	var l float32 = 7.0
	var t bool = true

	fmt.Println(s + " " + s) 
	fmt.Println("x + y =", x+y)
	fmt.Println(x, " + ", y, " = ", x+y)
	fmt.Println("y / x =", y/x)
	fmt.Println("y / x =", float32(y/x))
	fmt.Println("y / f =", float32(y)/f)
	fmt.Println("boolean f = ", t)

}
