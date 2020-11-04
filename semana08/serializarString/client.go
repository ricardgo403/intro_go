package main

import (
	"encoding/gob"
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
	err2 := gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err2)
		return
	}
	// c.Write([]byte(msg))
	// c.Write()
}

func main() {
	go client()
	var input string
	fmt.Scanln(&input)
}
