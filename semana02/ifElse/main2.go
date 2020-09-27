package main

import "fmt"

func main() {
	var temp int
	fmt.Print("Temp: ")
	fmt.Scan(&temp)
	if temp < 0 {
		fmt.Println("Está helado")
	} else if temp >= 0 && temp < 12 {
		fmt.Println("Está haciendo frio")
	} else if temp >= 12 && temp < 18 {
		fmt.Println("Está a gusto")
	} else {
		fmt.Println("Está haciendo calor")
	}
}
