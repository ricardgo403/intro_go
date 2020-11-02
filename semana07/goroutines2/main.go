package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("Received: ", msg)
		// fmt.Println("Answered: ")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	channel := make(chan string)
	// var c chan string := make(chan string)
	go ping(channel)
	go pong(channel)
	go printer(channel)
	var input string
	fmt.Scanln(&input)
}
