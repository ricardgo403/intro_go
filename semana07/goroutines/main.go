package main

import (
	"fmt"
	"time"
)

func f(n int) {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(n, ": ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go f(0)
	go f(1)
	go f(2)
	go f(3)
	go f(4)

	var input string
	fmt.Scanln(&input)
}
