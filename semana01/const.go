package main

import "fmt"

var name string = "Joel"
var global int64

func main() {
	const lastName string = "Gonzalez"
	fmt.Println(name + " " + lastName)
	fmt.Println(global)
	// lastName = " Velasco" You can't reassign a value to a const.
}
