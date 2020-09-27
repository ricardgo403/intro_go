package main

import "fmt"

func main() {
	var temp int
	fmt.Print("Temp: ")
	fmt.Scan(&temp)
	switch {
	case temp < 0:
		fmt.Println("Está helado")
	case temp >= 0 && temp < 12:
		fmt.Println("Está haciendo frio")
	case temp >= 12 && temp < 18:
		fmt.Println("Está a gusto")
	default:
		fmt.Println("Está haciendo calor")
	}
}
