package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Print
	p(rand.Intn(100), ",")	// default seed(1)
	p(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64()*5+5)
	fmt.Println(rand.Float32()*3+7)
	fmt.Println()


	s1 := rand.NewSource(time.Now().UnixNano())	
	r1 := rand.New(s1)


	for i := 0; i < 100; i++ {
		fmt.Println(r1.Intn(100))
	}

	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))


	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	
	rand.Seed(42)
	fmt.Println(r2.Intn(100))
	fmt.Println(r3.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	rand.Seed(42)
	fmt.Println(rand.Intn(100))
	
	rand.Seed(time.Now().UnixNano())

	

}