package main

import "fmt"

func main() {
	var input float64
	fmt.Println("Enter a number: ")
	fmt.Scan(&input)

	output := input * 2

	fmt.Println(output)
}
