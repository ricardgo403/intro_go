package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{8, 5, 1, 0, 2, 6}
	sort.Ints(s)
	fmt.Println(s)
	s1 := []int{5, 1, 0, 2, 6, 8}
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)
}
