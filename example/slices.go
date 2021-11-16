package main

import "fmt"

func main() {
	var s = make([]string, 3)
	var s2 [3]string
	fmt.Println(s)
	fmt.Println(s2)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	s = append(s, "d")
	fmt.Println(s)
	s2[0] = "0"	
	//s2 = append(s2, "d")	// first argument to append must be slice; have [1]string
	fmt.Println(s2)

	c := make([]string, len(s))
	copy(c, s)		// copy(target, source)
	fmt.Println(c)

	fmt.Println(s[0:2])	//no include s[2]
	fmt.Println(s[:2])
	fmt.Println(s[2:])

	t := []string{} //{"a", "b", "c"}	// []string NOT [3]string 
	t = append(t, "d")
	fmt.Println(t)

	twoD := make([][]int, 5)
	for i := 0; i < len(twoD); i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			twoD[i][j] = i + j
		}
		fmt.Println(twoD[i])
	}
	fmt.Println(twoD)
}
