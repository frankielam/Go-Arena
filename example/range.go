package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	for idx, num := range nums {
		fmt.Println("index", idx, "value", num)
	}
	
	kvs := map[string]string{"k1": "a", "k2":"b"}
	for k, v := range kvs {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k1 := range kvs {
		fmt.Println(k1)
	}

	for _, val := range kvs {
		fmt.Println(val)
	}
}
