package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	// b := make([]byte, 100)
	// bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Mensaje: ", msg)
		// fmt.Println("Bytes: ", bs)
	}
}

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}
