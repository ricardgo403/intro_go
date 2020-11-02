package main

import (
	"fmt"
	"time"
)

func f(n int) {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(n, ": ", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go f(0)

	var input string
	fmt.Scanln(&input)
}
