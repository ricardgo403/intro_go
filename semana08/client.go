package main

import (
	"fmt"
	"net"
)

func client() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	msg := "Hello world! :)"
	fmt.Println(msg)
	c.Write([]byte(msg))
}

func main() {
	go client()
	var input string
	fmt.Scanln(&input)
}
