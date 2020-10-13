package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(x), cap(x))
	x[4] = 100
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
	fmt.Println(x)
}
