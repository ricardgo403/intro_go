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
	channel2 := make(chan string)
	// var c chan string := make(chan string)
	go func() {
		for {
			channel <- "desde 1"
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			channel2 <- "desde 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg := <-channel:
				fmt.Println(msg)
			case msg := <-channel2:
				fmt.Println(msg)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
