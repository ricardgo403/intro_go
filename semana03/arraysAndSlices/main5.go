package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(x), cap(x))
	slice := x[0:3] // [0:] [:5] [:]
	for i, v := range slice {
		fmt.Println(i, v)
	}
	fmt.Println(x)
}
