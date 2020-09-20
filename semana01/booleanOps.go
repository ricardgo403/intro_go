package main

import "fmt"

func main() {
	var myTrue bool = true
	fmt.Println(myTrue == true)
	myFalse := false
	fmt.Println("Content: ", myFalse != myTrue)
}
